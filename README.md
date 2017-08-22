# Releaser


[![Build Status](https://travis-ci.org/avarabyeu/releaser.svg?branch=master)](https://travis-ci.org/avarabyeu/releaser)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/avarabyeu/releaser/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/avarabyeu/releaser)](https://goreportcard.com/report/github.com/avarabyeu/releaser)
[![Code Coverage](https://codecov.io/gh/avarabyeu/releaser/branch/master/graph/badge.svg)](https://codecov.io/gh/avarabyeu/releaser)

```sh
Release assistant built by avarabyeu

Usage:
  releaser [flags]
  releaser [command]

Available Commands:
  bintray     Uploads to bintray
  bump        Bump new version number
  exec        Executes bunch of commands
  help        Help about any command
  init        Creates version file
  release     Release new version
  replace     Replaces placeholders in files
  show        Print the current version number
  tag         Tags new version in git

Flags:
      --artifactsFolder string   Folder with artifacts to upload (default "release")
      --bintray.org string       BintrayConf organization
      --bintray.pack string      BintrayConf package
      --bintray.repo string      BintrayConf repository
      --bintray.token string     BintrayConf token
      --bintray.user string      BintrayConf user name
  -f, --file string              Version file to store current version (default "VERSION")
  -h, --help                     help for releaser
      --replace string           Replaces placeholders in files (default "release")
  -v, --version string           Version to be used

Use "releaser [command] --help" for more information about a command.
```
