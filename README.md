# goapps

Small collection of apps written in Go.
Just to get started with Go and for reference.

You can use it for learning purposes, nothing really productive. Just playing ...

You need to set GOPATH.
When you build go will look in src for packages and use them.
You can set it in .bashrc.
IMPORTANT: You can set many paths, e.g.
GOPATH=/home/username/go:/home/username/otherdirectory





- Get help on some command/option:
```
    # e.g. for building an app
    go help build
```

- Compile and run
```
    go run path/to/file/with/main/function.go
```

- Just compile
```
    go build /path/to/file.go
```

- Compile for other platform
```
    GOOS=windows GOARCH=amd64 go build
```

- Format code
```
    go fmt code.go
```

- Install package
```
    go get https://path/at/Github
```

- Create documentation
```
    godoc file.go
    godoc folder
```
