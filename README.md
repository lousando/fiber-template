# Go Fiber Template

Comes pre-loaded with [Bulma](https://bulma.io/) and [HTMX](https://htmx.org/)

Checklist 📋

- Change the module name in `go.mod` 
- Change the site title in `views/index.tmpl`


## Getting Started 🚀

```bash
# create your .env file
cp .env.example .env

# (if not already installed) install the air utility globally
go install github.com/cosmtrek/air@latest

# Now you can just run air to start the server 
air
```

#### If you get: "command not found: air" or "No such file or directory"

The solution to this is to add the following variables to your
shell profile.

```shell
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export PATH=$PATH:$(go env GOPATH)/bin
```