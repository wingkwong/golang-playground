# Setting up a Dev Environment

## Installation

1. [Go download.](https://golang.org/doc/install#download)
2. [Go install.](https://golang.org/doc/install#install)

## Setting up a Dev Environment

Add the following cmds in your rc file, in my case, that is ``~/.zshrc``.

```
# Go root
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

# a directory where your Go programs locate
export GOPATH=~/golib
export PATH=$PATH:$GOPATH/bin
```

Re-read the RC file by running

```
source <YOUR_RC_FILE>
```

Install Go extension in VS Code.

![image](https://user-images.githubusercontent.com/35857179/129442774-55069b60-3ed6-40d6-aae9-9d10c6d31441.png)


Go to ``$GOPATH`` and run the following commands

```
mkdir bin
mkdir pkg
mkdir -p src/github.com/wingkwong/golang-playground
```

Create ``Main.go``  under ``$GOPATH/src/github.com/wingkwong/golang-playground/setting-up-a-dev-environment`` and add the following line

```go
package "main"
```

VS Code would prompt some available extension. Install all of them.

![image](https://user-images.githubusercontent.com/35857179/129443290-d773193b-f422-483d-9029-59568de4d4e7.png)

Go to ``$GOPATH``

```
cd $GOPATH
```

```go
go run src/github.com/wingkwong/golang-playground/setting-up-a-dev-environment/Main.go
```

You should see 

```
Hello World
```

Or you can run 

```go
go build github.com/wingkwong/golang-playground/setting-up-a-dev-environment
```

> If you see ``go.mod file not found in current directory or any parent directory; see 'go help modules``, please run ``go env -w GO111MODULE=off`` and re-run it.

After that, you should have an executable and you can run ``./setting-up-a-dev-environment`` to see the same result.


Or you can run 

```
go install github.com/wingkwong/golang-playground/setting-up-a-dev-environment
```

and an executable will be under ``bin/``.