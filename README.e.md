
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)

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

## github-create-repo - Create Repository in Github

```sh
 # Create an organization Github Repository 
$ easygen -tf github-create-repo.tmpl wireframe_proj.yaml
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/orgs/repos -d '{"name":"wireframe", "description": "wire-frame construction to get the project quickly into shape", "auto_init": true, "license_template": "mit", "gitignore_template": "Go"}'

 # Create a normal user Github Repository 
sed 's/^  Vendor: go-easygen/  User: suntong/' wireframe_proj.yaml > /tmp/wireframe_proj.yaml

$ easygen -tf github-create-repo.tmpl /tmp/wireframe_proj.yaml
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

First of all, this is what auto-generated help looks like:

#### $ {{exec "wireframe" | color "sh"}}

This gives full help at root level.

There are also two sub-commands, which are:

#### $ {{shell "wireframe put" | color "sh"}}

The above gives sub-command `put` level help, and the next is for `get`:

#### $ {{shell "wireframe get" | color "sh"}}

The above gives sub-command `get` level help.

Before we see how it runs, let's take a look at how to define and get all the above. Here is the single source of CLI definition for all above:

<a name="cli.yaml"/>

#### {{cat "wireframe_cli.yaml" | color "yaml"}}


The above `yaml` definition is all it takes to get a wire-framed Go code to start with.

We don't need to jump into the generate code itself now, just take a look what we will get out of the box first:

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

Basically, the above functionalities are what we can get out of the box from the above [single source of CLI definition file](#cli.yaml) automatically, without a single line of customization code.

## github-create-release - Create Release in Github

```
GITHUB_TOKEN=...

GITHUB_TAG=1.0.0
GITHUB_RELEASE_TEXT="Release v$GITHUB_TAG"

git push

$ easygen -tf ../../go-easygen/wireframe/github-create-release.tmpl wireframe_proj.yaml
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/repos/go-easygen/wireframe/releases -d '{"tag_name":"'$GITHUB_TAG'", "name": "wireframe-'$GITHUB_TAG'", "body": "'"$GITHUB_RELEASE_TEXT"'"}'

```

The copy and do the `curl` command on the command line. E.g.:

With `curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/repos/go-easygen/wireframe/releases -d '{"tag_name":"'$GITHUB_TAG'", "name": "wireframe-'$GITHUB_TAG'", "body": "'"$GITHUB_RELEASE_TEXT"'"}'` for `wireframe`:

```json
{
  "url": "https://api.github.com/repos/go-easygen/wireframe/releases/14826407",
  "assets_url": "https://api.github.com/repos/go-easygen/wireframe/releases/14826407/assets",
  "upload_url": "https://uploads.github.com/repos/go-easygen/wireframe/releases/14826407/assets{?name,label}",
  "html_url": "https://github.com/go-easygen/wireframe/releases/tag/1.0.0",
  "id": 14826407,
  "node_id": "MDc6UmVsZWFzZTE0ODI2NDA3",
  "tag_name": "1.0.0",
  "target_commitish": "master",
  "name": "wireframe-1.0.0",
  "draft": false,
  "author": {
    "login": "suntong",
    "id": 422244,
    "node_id": "MDQ6VXNlcjQyMjI0NA==",
    "avatar_url": "https://avatars1.githubusercontent.com/u/422244?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/suntong",
    "html_url": "https://github.com/suntong",
    "followers_url": "https://api.github.com/users/suntong/followers",
    "following_url": "https://api.github.com/users/suntong/following{/other_user}",
    "gists_url": "https://api.github.com/users/suntong/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/suntong/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/suntong/subscriptions",
    "organizations_url": "https://api.github.com/users/suntong/orgs",
    "repos_url": "https://api.github.com/users/suntong/repos",
    "events_url": "https://api.github.com/users/suntong/events{/privacy}",
    "received_events_url": "https://api.github.com/users/suntong/received_events",
    "type": "User",
    "site_admin": false
  },
  "prerelease": false,
  "created_at": "2019-01-07T03:51:46Z",
  "published_at": "2019-01-07T04:15:30Z",
  "assets": [

  ],
  "tarball_url": "https://api.github.com/repos/go-easygen/wireframe/tarball/1.0.0",
  "zipball_url": "https://api.github.com/repos/go-easygen/wireframe/zipball/1.0.0",
  "body": "Release v1.0.0"
}

```

For `ffcvt`:

```
$ easygen -tf ../../go-easygen/wireframe/github-create-release.tmpl ffcvt_proj.yaml 
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/repos/suntong/ffcvt/releases -d '{"tag_name":"'$GITHUB_TAG'", "name": "ffcvt-'$GITHUB_TAG'", "body": "'"$GITHUB_RELEASE_TEXT"'"}'

GITHUB_TAG=1.3.2
GITHUB_RELEASE_TEXT="Add subtitle streams copy support"
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/repos/suntong/ffcvt/releases -d '{"tag_name":"'$GITHUB_TAG'", "name": "ffcvt-'$GITHUB_TAG'", "body": "'"$GITHUB_RELEASE_TEXT"'"}'

```

will get:

```json
{
  "url": "https://api.github.com/repos/suntong/ffcvt/releases/14826435",
  "assets_url": "https://api.github.com/repos/suntong/ffcvt/releases/14826435/assets",
  "upload_url": "https://uploads.github.com/repos/suntong/ffcvt/releases/14826435/assets{?name,label}",
  "html_url": "https://github.com/suntong/ffcvt/releases/tag/1.3.2",
  "id": 14826435,
  "node_id": "MDc6UmVsZWFzZTE0ODI2NDM1",
  "tag_name": "1.3.2",
  "target_commitish": "master",
  "name": "ffcvt-1.3.2",
  "draft": false,
  "author": {
    "login": "suntong",
    "id": 422244,
    "node_id": "MDQ6VXNlcjQyMjI0NA==",
    "avatar_url": "https://avatars1.githubusercontent.com/u/422244?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/suntong",
    "html_url": "https://github.com/suntong",
    "followers_url": "https://api.github.com/users/suntong/followers",
    "following_url": "https://api.github.com/users/suntong/following{/other_user}",
    "gists_url": "https://api.github.com/users/suntong/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/suntong/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/suntong/subscriptions",
    "organizations_url": "https://api.github.com/users/suntong/orgs",
    "repos_url": "https://api.github.com/users/suntong/repos",
    "events_url": "https://api.github.com/users/suntong/events{/privacy}",
    "received_events_url": "https://api.github.com/users/suntong/received_events",
    "type": "User",
    "site_admin": false
  },
  "prerelease": false,
  "created_at": "2019-01-07T02:56:43Z",
  "published_at": "2019-01-07T04:20:21Z",
  "assets": [

  ],
  "tarball_url": "https://api.github.com/repos/suntong/ffcvt/tarball/1.3.2",
  "zipball_url": "https://api.github.com/repos/suntong/ffcvt/zipball/1.3.2",
  "body": "Add subtitle streams copy support"
}

```



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
