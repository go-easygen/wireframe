
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{toc 5}}

## {{.Name}} - wire-framing project for quick start

wire-frame construction to get the project quickly into shape

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


# Download/Install

## Download binaries

- The latest binary executables are available under  
https://bintray.com/suntong/bin/{{.Name}}#files/{{.Name}}  
as the result of the Continuous-Integration process.
- I.e., they are built right from the source code during every git commit automatically by [travis-ci](https://travis-ci.org/).
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `{{.Name}}_linux_VER_amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `{{.Name}}`, after downloading it. 


## Debian package

Available at https://dl.bintray.com/suntong/deb.

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


# Credits



# Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

