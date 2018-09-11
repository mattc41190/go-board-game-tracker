cd ~
sudo apt-get update
sudo apt-get -y upgrade
sudo apt-get -y install git
sudo apt-get -y install golang
sudo apt-get -y install golang-go
mkdir -p go/src
mkdir -p go/pkg
mkdir -p go/bin

export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

go get github.com/mattc41190/go-board-game-tracker

cd src/github.com/mattc41190/go-board-game-tracker

go build *.go

sudo ./go-board-game-tracker
