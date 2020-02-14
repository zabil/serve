`serve` is a command line tool to quickly start a web server to 
view static content.

### Usage

Serve the current directory

```
$ serve
Serving at http://localhost:3000
```

Serve specific directory

```
$ serve -d build/site
Serving at http://localhost:3000
```

Serve on a specific port

```
$ serve -p 1234
Serving at http://localhost:1234
```

### Install

```
$ go get github.com/zabil/serve
```
