
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

# Go wire-frame building

The detailed explanation on the [Go wire-frame building can be found here]( https://github.com/go-easygen/wireframe/wiki/Go-project-wire-frame-building).

## Command line flag handling code auto-generation

Refer to

[Command line flag handling code auto-generation](https://github.com/go-easygen/wireframe/wiki/Command-line-flag-handling-code-auto-generation#auto-gen)



# Download/Install

_(The following is the base template for all `wireframe` based projects)_

## Download binaries

- The latest binary executables are available right under the github release page  
https://github.com/suntong/{{.Name}}/releases  
as the result of the Continuous-Integration process.
- I.e., they are built during every git tagged push, automatically by [GitHub Actions](https://github.com/features/actions), right from the source code, truely WYSIWYG.
- The `.deb`, `.rpm` and `.apk` packages are readily available, as well as the executables for other Linux and Windows as well.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `{{.Name}}_ver_linux_amd64.tar.gz` file.
- Unzip it and put the executable somewhere in the PATH, after downloading it. 


## Install Source

To install the source code instead:

```
go get github.com/go-easygen/wireframe
```


# Similar Projects

At the beginning, all the following similar projects have been attempted before rolling out on my own solution instead. The listed url points to the limitations they had, _by the time this project was created_.

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
