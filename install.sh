sudo apt-get install python-software-properties
sudo add-apt-repository ppa:duh/golang
sudo apt-get update
sudo apt-get install golang
sudo apt-get install git
export GOROOT=/usr/lib/go
export GOBIN=/usr/bin/go
mkdir lib
export GOPATH=$HOME/lib
mkdir lib/src
mkdir lib/src/github.com
mkdir lib/src/github.com/wenkesj
git clone https://github.com/wenkesj/go-server.git lib/src/github.com/wenkesj/go-server
go get github.com/wenkesj/go-server
