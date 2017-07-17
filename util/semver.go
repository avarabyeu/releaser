package util

import (
	"github.com/coreos/go-semver/semver"
	"github.com/juju/errgo/errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const snapshot semver.PreRelease = "SNAPSHOT"

//Semver represents semantic versioning file storage
type Semver struct {
	file    string
	Version *semver.Version
}

func load(file string) (*semver.Version, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Mask(err)
	}

	versionString := string(buf)
	versionString = strings.TrimSpace(versionString)

	ver, err := semver.NewVersion(versionString)

	if err != nil {
		return nil, errors.Mask(err)
	}

	return ver, nil
}

//Save saves current state to file
func (s Semver) Save() error {
	return errors.Mask(ioutil.WriteFile(s.file, []byte(s.Version.String()), 0664))
}

//Current returns current version
func (s Semver) Current() string {
	return s.Version.String()
}

//NextSnapshot bumps next snapshot version
func (s Semver) NextSnapshot(newVersion string) string {
	if "" != newVersion {
		s.Version.Set(newVersion)
	} else {
		s.Version.BumpPatch()
		s.Version.PreRelease = snapshot
	}

	return s.Version.String()
}

//Bump bumps provided new version or next patch version
func (s Semver) Bump(newVersion string) string {
	if "" != newVersion {
		s.Version.Set(newVersion)
	} else if snapshot == s.Version.PreRelease {
		s.Version.PreRelease = ""
	} else {
		s.Version.BumpPatch()

	}

	return s.Version.String()
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

//Load loads version from file
func Load(file string) *Semver {
	if !fileExists(file) {
		log.Fatalf("File %s not found", file)
	}
	ver, err := load(file)
	if nil != err {
		log.Fatalf("Cannot read file: %s", err.Error())
	}
	return &Semver{file: file, Version: ver}
}

//New creates new object
func New(file string, version *semver.Version) *Semver {
	return &Semver{file: file, Version: version}
}
