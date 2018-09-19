
# wireframe

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-easygen/wireframe?status.svg)](http://godoc.org/github.com/go-easygen/wireframe)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-easygen/wireframe)](https://goreportcard.com/report/github.com/go-easygen/wireframe)
[![travis Status](https://travis-ci.org/go-easygen/wireframe.svg?branch=master)](https://travis-ci.org/go-easygen/wireframe)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)

## TOC
- [wireframe - wire-framing project for quick start](#wireframe---wire-framing-project-for-quick-start)
- [wire-frame building](#wire-frame-building)
  - [github-repo-create - Create Repository in Github](#github-repo-create---create-repository-in-github)
  - [gitlab-repo-create - Create Repository in Gitlab](#gitlab-repo-create---create-repository-in-gitlab)
  - [Data type def](#data-type-def)

## wireframe - wire-framing project for quick start

*wire-frame* provides wire-framing for Go cli based projects, from start to finish.

It is a tiny Go cli code that demonstrates how to quickly get a Go based command line program started, and deployed.

It illustrates

- what basic info to prepare and how to utilize this info for all the following tasks
- how to create the github repository with it from command line
- how to handle command line parameters using code gen
- how to use Continuous-Integration [travis-ci](https://travis-ci.org/) to build and release binary executables (of all OS platforms) every time you do a `git commit`
- how to package the final tool as debian package and upload to your PPA on bintray, so that people can easily install and get your updates

Check out the executables and package building log here:
https://travis-ci.org/go-easygen/wireframe/builds/265785563

# wire-frame building

## github-repo-create - Create Repository in Github

```sh
 # Create an organization Github Repository 
$ easygen -tf github-repo-create.tmpl wireframe_proj.yaml 
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/orgs/repos -d '{"name":"wireframe", "description": "wire-frame construction to get the project quickly into shape", "auto_init": true, "license_template": "mit", "gitignore_template": "Go"}'

 # Create a normal user Github Repository 
sed 's/^  Vendor: go-easygen/  User: suntong/' wireframe_proj.yaml > /tmp/wireframe_proj.yaml

$ easygen -tf github-repo-create.tmpl /tmp/wireframe_proj.yaml 
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/user/repos -d '{"name":"wireframe", "description": "wire-frame construction to get the project quickly into shape", "auto_init": true, "license_template": "mit", "gitignore_template": "Go"}'

```

The GitHub Token is for accessing [GitHub API](https://developer.github.com/v3) to create repository or deploy the artefacts to GitHub etc. You can create one [here](https://github.com/settings/tokens/new).

## gitlab-repo-create - Create Repository in Gitlab

```sh
ghrn=wireframe
ghrd='wire-frame construction to get the project quickly into shape'
ghun=go-easygen
GITLAB_TOKEN=xxxx

namespace_id=`curl -s --header "PRIVATE-TOKEN: $GITLAB_TOKEN" "https://gitlab.com/api/v3/namespaces" | jq --arg name $ghun '.[] | select(.name==$name) | .id'`

curl -H "Content-Type:application/json" https://gitlab.com/api/v3/projects?private_token=$GITLAB_TOKEN -d "{ \"name\": \"$ghrn\", \"description\": \"$ghrd\", \"namespace_id\": $namespace_id"',"only_allow_merge_if_build_succeeds":true,"only_allow_merge_if_all_discussions_are_resolved":true}'
```


## Data type def

```sh
$ jsonfiddle j2s -f yaml -i wireframe_full.yaml --name WireframeT | sed '/Wireframe\b/d; s/ `yaml:.*$//' | gofmt | tee WireframeT.go
package main

type WireframeT struct {
        Author  string
        Desc    string
        Lang    string
        License string
        Proj    string
        User    string
        Vendor  string
}

$ cat wireframe_full.yaml
Wireframe:
  Proj: wireframe
  Desc: wire-frame construction to get the project quickly into shape
  Lang: Go
  User: <empty>
  Vendor: go-easygen
  Author: Tong Sun <suntong@cpan.org>
  License: MIT
```

The `jsonfiddle` is the JSON Fiddling tool that makes it easy to look at the JSON data from different aspects, which is [available here](https://github.com/go-jsonfile/jsonfiddle).

## Command line flag handling code auto-generation

Refer to

- [Command line flag handling code auto-generation](https://github.com/go-easygen/easygen#command-line-flag-handling-code-auto-generation), especially,
- the [cli based command line flag handling code auto-generation](https://github.com/go-easygen/easygen#cli-based).
- for the four different `UsageStyles` that can be used to control the `usage()` output, check out the [UsageStyle of package cli](https://github.com/go-easygen/wireframe/wiki/UsageStyle-of-package-cli) wiki.

### Auto-generated Command line flag handling showcase using wireframe

#### $ wireframe
```sh
wire framing
Version 0.1.0 built on 2018-09-18

Tool to showcase wire-framing command line app fast prototype

Options:

  -h, --help        display help information 
  -c, --config      config file [=$__EXEC_FILENAME.json]

  -H, --host        host addr [=$HOST]
  -p, --port        listening port 

      --long        Now can use the \n to arrange parameters in groups
			Just like what is showing here, even with extreme long usage text that can spread across multiple lines [=$Demo]

  -D, --daemonize   daemonize the service 
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity) 


Commands:

  put   Upload into service
  get   Get from the service
```

This gives full help at root level.

#### $ wireframe put
```sh
Upload into service

Usage:
  wireframe put -i /tmp/f

Options:

  -h, --help        display help information 
  -c, --config      config file [=$__EXEC_FILENAME.json]

  -H, --host        host addr [=$HOST]
  -p, --port        listening port 

      --long        Now can use the \n to arrange parameters in groups
			Just like what is showing here, even with extreme long usage text that can spread across multiple lines [=$Demo]

  -D, --daemonize   daemonize the service 
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity) 

  -i, --input      *The file to upload from (mandatory)
```

This gives sub-command `put` level help.

#### $ wireframe get
```sh
Get from the service

Usage:
  wireframe get -o /tmp/f some more args

Options:

  -h, --help        display help information 
  -c, --config      config file [=$__EXEC_FILENAME.json]

  -H, --host        host addr [=$HOST]
  -p, --port        listening port 

      --long        Now can use the \n to arrange parameters in groups
			Just like what is showing here, even with extreme long usage text that can spread across multiple lines [=$Demo]

  -D, --daemonize   daemonize the service 
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity) 

  -i, --input      *The file to upload from (mandatory) 
  -o, --output      The output file (default: some file)
```

This gives sub-command `get` level help.

#### $ wireframe put -i /tmp/f


```sh
$ touch /tmp/f; wireframe put -i /tmp/f
[put]:
  &{Helper:{Help:false} Self:0xc420010240 Host:127.0.0.1 Port:8080 Daemonize:false Verbose:{value:0}}
  &{Filei:0xc4200d86c0}
  []
wireframe v 0.1.0. Upload into service
Copyright (C) 2018, Myself <me@mine.org>
```

This shows getting everything from the self-config file.
Note the value of `Host`, it is read from the `wireframe_cfg.json` self-config file.

#### $ HOST=10.0.0.1 wireframe put -i /tmp/f
```sh
[put]:
  &{Helper:{Help:false} Self:0xc42008c1c0 Host:10.0.0.1 Port:8080 Daemonize:false Verbose:{value:0}}
  &{Filei:0xc4200f2660}
  []
wireframe v 0.1.0. Upload into service
Copyright (C) 2018, Myself <me@mine.org>
```

This shows overriding settings from the self-config file using the environment variables. Note the value of `Host` now is taken from the environment variable, instead from the `wireframe_cfg.json` self-config file.

#### $ HOST=10.0.0.1 wireframe put -i /tmp/f -H 168.0.0.1
```sh
[put]:
  &{Helper:{Help:false} Self:0xc4200901c0 Host:168.0.0.1 Port:8080 Daemonize:false Verbose:{value:0}}
  &{Filei:0xc4200f6680}
  []
wireframe v 0.1.0. Upload into service
Copyright (C) 2018, Myself <me@mine.org>
```

This shows overriding settings on the command line. Note the value of `Host` now is taken from the command line. So the priority of setting the `Host` value is, from higher priority to lower:

- command line
- environment variable
- self-config file

Three different levels.

#### $ wireframe get -o /tmp/f some more args

```sh
$ HOST=10.0.0.1 wireframe get -o /tmp/f some more args
[get]:
  &{Helper:{Help:false} Self:0xc420090180 Host:10.0.0.1 Port:8080 Daemonize:false Verbose:{value:0}}
  &{Fileo:0xc4200f8680}
  [some more args]
wireframe v 0.1.0. Get from the service
Copyright (C) 2018, Myself <me@mine.org>
```

This just shows how to make use of the extra arguments passed from the command line. Note the setting is a bit different between `put` and `get` regarding what is mandatory on the command line. I.e., for `get`, there much be some extra command line arguments.

## Binary releases

``` sh
gpkg=$(basename $(pwd))

export BINTRAY_USER=suntong BINTRAY_REPO_BIN=bin
easygen -tf bintray-bin ${gpkg}_proj | tee bintray-bin.json

export BINTRAY_REPO_DEB=deb PKG1ST=`expr $gpkg : '^\(.\)'`
easygen -tf bintray-pkg ${gpkg}_proj | tee bintray-pkg.json

export PKG_MAINT="My Name <myid@myorg.com>"
easygen -tf travis ${gpkg}_proj | tee .travis.yml
```

Then,

- inspect the generated `bintray-bin.json`, `bintray-pkg.json`, and `.travis.yml`. 
- try to do a `git push` and manually fix any remaining issues. 

The `easygen` is the universal code/text generator, which is [available here](https://github.com/go-easygen/easygen).

The above steps assume that the user and the `BINTRAY_REPO_BIN` and `BINTRAY_REPO_DEB` repos are already exist on [bintray.com](https://bintray.com/). Check out the [Hosting Debian Packages on Bintray](https://blog.bintray.com/2014/12/16/hosting-debian-packages-on-bintray-rocks/) for details, and [Bintray Debian Repository Creation and Upload file using API](https://stackoverflow.com/questions/45516482/bintray-debian-repository-creation-and-upload-file-using-api/45519360#45519360) for the condensed and practical version. All you need to do before doing the above steps are,

- Create an Bintray account (free).
- Obtain `BINTRAY_API_KEY` from its web site.
- Create two repos for binary executables (`BINTRAY_REPO_BIN`) and debian packages (`BINTRAY_REPO_DEB`).

That's it. The above steps should take care of the rest.


# Download/Install

## Download binaries

- The latest binary executables are available under  
https://bintray.com/suntong/bin/wireframe/latest, or directly under  
https://bintray.com/version/files/suntong/bin/wireframe/latest  
as the result of the Continuous-Integration process.
- I.e., they are built during every git push, automatically by [travis-ci](https://travis-ci.org/), right from the source code, truely WYSIWYG.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `wireframe-linux-amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `wireframe`, after downloading it. 


## Debian package

Available at https://bintray.com/suntong/deb/wireframe,  
or directly at  https://dl.bintray.com/suntong/deb:

```
echo "deb [trusted=yes] https://dl.bintray.com/suntong/deb all main" | sudo tee /etc/apt/sources.list.d/suntong-debs.list
sudo apt-get update

sudo chmod 644 /etc/apt/sources.list.d/suntong-debs.list
apt-cache policy wireframe

sudo apt-get install -y wireframe
```


## Install Source

To install the source code instead:

```
go get github.com/go-easygen/wireframe
```


# Similar Projects

All the following similar projects have been attempted before rolling out on my own solution instead. The listed url points to the limitations they have by the time this project was created.

- [**goreleaser**](https://github.com/goreleaser/goreleaser/issues/15#issuecomment-321949280)
- [**go-github-release**](https://github.com/mh-cbon/go-github-release/issues/6), and also see the issues [here](https://github.com/mh-cbon/go-github-release/issues/10#issuecomment-307646985), [here](https://github.com/mh-cbon/go-github-release/issues/18) and [here](https://github.com/mh-cbon/go-github-release/issues/20).


# Credits


# Promoting WireFrame

Please help promoting WireFrame by using one of the following badges:

```
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)
```

[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)


# Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
