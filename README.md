# serve-go
serve static website

Inspired by npm serve, https://www.npmjs.com/package/serve

However sometimes we don't have internet to download npm packages, golang application provided a way to do offline installation


### build project
```shell
go build .

# build for windows in linux
env GOOS=windows GOARCH=amd64 go build
```

### usage
```shell

serve-go ./

serve-go ./ -port 8080

#proxy sample configuration
serve-go ./ -port 8080 -proxy "/api/,https://api.randomuser.me,root"

# root refer to browse without context

# multiple proxy configuration
serve-go ./ -port 8080 -proxy "/api/,https://api.randomuser.me" -proxy "/api/,https://api.randomuser.me"

#help
serve-go -h

```

### Installation
Download your respect of binary from https://github.com/yhjhoo/serve-go/releases and put into your environment path


### Build from local
```shell
go build
```
