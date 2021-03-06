package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/avarabyeu/releaser/util"
	"github.com/juju/errgo/errors"
	"github.com/spf13/cobra"
	"gopkg.in/cheggaaa/pb.v1"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	refreshRate  = time.Millisecond * 10
	btVersionURL = "https://api.bintray.com/packages/{{.org}}/{{.repo}}/{{.package}}/versions"
)

var btVersionTemplate *template.Template

func init() {
	btVersionTemplate, _ = template.New("btVersionTMPL").Parse(btVersionURL)
}

var bintrayCommand = &cobra.Command{
	Use:   "bintray",
	Short: "Uploads to bintray",
	Long:  `Uploads artifacts to bintray`,
	Run: func(cmd *cobra.Command, args []string) {
		err := uploadToBintray(cmd)
		if nil != err {
			log.Fatalf("Cannot upload to bintray. %s", err)
		}
	},
}

func uploadToBintray(cmd *cobra.Command) error {
	wd, err := os.Getwd()
	if nil != err {
		log.Printf("Cannot get work dir! %s", err.Error())
		return err
	}

	err = checkData(config.Bintray)
	if nil != err {
		log.Printf("BintrayConf data isn't provided: %s", err.Error())
		return err
	}

	artifactsFolder := path.Join(wd, cmd.Flag("artifactsFolder").Value.String())
	files, _ := ioutil.ReadDir(artifactsFolder)

	if len(files) > 0 {
		er := createVersion(GetSemver(cmd).Version.String(), config.Bintray)
		if nil != er {
			return er
		}

	}

	// update bars
	//pool, err := pb.StartPool()
	if err != nil {
		return err
	}

	//wg := new(sync.WaitGroup)
	for _, f := range files {
		//wg.Add(1)
		log.Print(f.Name())
		uploadFile(GetSemver(cmd).Current(), path.Join(artifactsFolder, f.Name()), config.Bintray)
	}

	//wg.Wait()
	//pool.Stop()
	return nil
}

func createVersion(version string, bd *BintrayConf) error {
	details := map[string]string{"name": version, "desc": "New version"}
	rqBody, err := json.Marshal(details)
	if nil != err {
		return err
	}

	url := new(bytes.Buffer)
	btVersionTemplate.Execute(url, map[string]string{"org": bd.Org, "repo": bd.Repo, "package": bd.Pack})

	log.Println(url.String())
	rq, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(rqBody))
	if err != nil {
		return err
	}
	rq.Header.Set("Authorization", bd.getAuth())
	rq.Header.Set("Content-Type", "application/json")

	rs, err := http.DefaultClient.Do(rq)
	if err != nil {
		return err
	}
	defer rs.Body.Close()

	if rs.StatusCode/100 != 2 {
		body, _ := ioutil.ReadAll(rs.Body)
		fmt.Println("Error:", string(body))
		if 409 == rs.StatusCode {
			//already exists. That's OK
			return nil
		}
		return errors.Newf("Cannot create new version. Status code: %d", rs.StatusCode)
	}

	return nil
}

func uploadFile(version string, filePath string, bt *BintrayConf) {
	//defer wg.Done()
	var err error
	var f *os.File
	var fi os.FileInfo

	var bar *pb.ProgressBar

	if f, err = os.Open(filePath); err != nil {
		log.Fatal(err)
	}
	if fi, err = f.Stat(); err != nil {
		log.Fatal(err)
	}
	bar = pb.New64(fi.Size()).Prefix(util.SubstrAfterLast(f.Name(), "/")).SetUnits(pb.U_BYTES).SetRefreshRate(refreshRate)
	//pool.Add(bar)

	r, w := io.Pipe()
	go func() {
		defer w.Close()
		defer f.Close()

		rqw := io.MultiWriter(w, bar)
		if _, err = io.Copy(rqw, f); err != nil {
			log.Fatal(err)
		}

	}()

	rq, err := http.NewRequest("PUT", fmt.Sprintf("https://api.bintray.com/content/%s/%s/%s/%s", bt.Org, bt.Repo, version, fi.Name()), r)
	rq.Header.Set("Content-Type", "application/octet-stream")
	rq.Header.Set("X-Bintray-Package", bt.Pack)
	rq.Header.Set("X-Bintray-Version", version)
	rq.Header.Set("X-Bintray-Override", "1")
	rq.Header.Set("X-Bintray-Publish", "1")
	rq.Header.Set("Authorization", bt.getAuth())

	resp, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	rs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(rs))
		log.Println(resp.StatusCode)
	}
}

func checkData(btData *BintrayConf) error {
	var errs string

	if "" == btData.Org {
		errs += "BintrayConf organization not set\n"
	}

	if "" == btData.User {
		errs += "BintrayConf user not set\n"
	}

	if "" == btData.Token {
		errs += "BintrayConf token not set\n"
	}

	if "" == btData.Pack {
		errs += "BintrayConf package not set\n"
	}

	if "" == btData.Repo {
		errs += "BintrayConf repo not set\n"
	}

	if "" != errs {
		log.Fatalf("Validation failed: \n%s", errs)
		return errors.New(errs)
	}

	return nil

}
