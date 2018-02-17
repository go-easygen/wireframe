
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}
[![PoweredBy WireFrame](PoweredBy-WireFrame.svg)](PoweredBy-WireFrame)

## {{toc 5}}

## {{.Name}} - wire-framing project for quick start

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
[Command line flag handling code auto-generation](https://github.com/go-easygen/easygen#command-line-flag-handling-code-auto-generation), especially, the [cli based command line flag handling code auto-generation](https://github.com/go-easygen/easygen#cli-based).

### Auto-generated Command line flag handling showcase using wireframe

#### $ {{exec "wireframe" | color "sh"}}

This gives full help at root level.

#### $ {{shell "wireframe put" | color "sh"}}

This gives sub-command `put` level help.

#### $ {{shell "wireframe get" | color "sh"}}

This gives sub-command `get` level help.

#### $ wireframe put -i /tmp/f


```sh
$ touch /tmp/f; wireframe put -i /tmp/f
[put]:
  &{Helper:{Help:false} Self:0xc420010240 Host:127.0.0.1 Port:8080 Daemonize:false Verbose:{value:0}}
  &{Filei:0xc4200d86c0}
  []
```

This shows getting everything from the self-config file.
Note the value of `Host`, it is read from the `wireframe_cfg.json` self-config file.

#### $ HOST=10.0.0.1 wireframe put -i /tmp/f
```sh
[put]:
  &{Helper:{Help:false} Self:0xc42008c1c0 Host:10.0.0.1 Port:8080 Daemonize:false Verbose:{value:0}}
  &{Filei:0xc4200f2660}
  []
```

This shows overriding settings from the self-config file using the environment variables. Note the value of `Host` now is taken from the environment variable, instead from the `wireframe_cfg.json` self-config file.

#### $ HOST=10.0.0.1 wireframe put -i /tmp/f -H 168.0.0.1
```sh
[put]:
  &{Helper:{Help:false} Self:0xc4200901c0 Host:168.0.0.1 Port:8080 Daemonize:false Verbose:{value:0}}
  &{Filei:0xc4200f6680}
  []
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
https://bintray.com/suntong/bin/{{.Name}}/latest, or directly under  
https://bintray.com/version/files/suntong/bin/{{.Name}}/latest  
as the result of the Continuous-Integration process.
- I.e., they are built during every git push, automatically by [travis-ci](https://travis-ci.org/), right from the source code, truely WYSIWYG.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `{{.Name}}-linux-amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `{{.Name}}`, after downloading it. 


## Debian package

Available at https://bintray.com/suntong/deb/{{.Name}},  
or directly at  https://dl.bintray.com/suntong/deb:

```
echo "deb [trusted=yes] https://dl.bintray.com/suntong/deb all main" | sudo tee /etc/apt/sources.list.d/suntong-debs.list
sudo apt-get update

sudo chmod 644 /etc/apt/sources.list.d/suntong-debs.list
apt-cache policy {{.Name}}

sudo apt-get install -y {{.Name}}
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



# Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
