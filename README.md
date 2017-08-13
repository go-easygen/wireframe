# wireframe
wire-frame construction to get the project quickly into shape

# github-repo-create - Create Repository in Github

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

# Data type def

```sh
$ jsonfiddle j2s -f yaml -i wireframe-full.yaml --name WireframeT | sed '/Wireframe\b/d; s/ `yaml:.*$//' | gofmt | tee WireframeT.go
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

$ cat wireframe-full.yaml
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

