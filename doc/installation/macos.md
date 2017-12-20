# Install Golang on macos

```
brew update
brew install golang
```

edit ~/.zshrc (or .bashrc)
```
export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

```
source ~/.zshrc
```

create `src` `pkg` `bin`, these three folder in the $GOPATH

```
mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin
```


# test

create hello folder and hello.go inside
```
package main
import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

```
go run hello.go
```

build binary

```
go install hello 
```

then hello.go will put into somewhere in the $GOPATH/bin
and it can be executed directly

```
hello
```

http://sourabhbajaj.com/mac-setup/Go/README.html