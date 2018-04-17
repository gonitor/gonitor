# Fast Server Monitor in Go
[![Build Status](https://travis-ci.org/gonitor/gonitor.svg?branch=master)](https://travis-ci.org/gonitor/gonitor)
[![codecov](https://codecov.io/gh/gonitor/gonitor/branch/master/graph/badge.svg)](https://codecov.io/gh/gonitor/gonitor)
[![Go Report Card](https://goreportcard.com/badge/github.com/gonitor/gonitor)](https://goreportcard.com/report/github.com/gonitor/gonitor)
[![GoDoc](https://godoc.org/github.com/gonitor/gonitor?status.svg)](https://godoc.org/github.com/gonitor/gonitor)

Gonitor is fast server monitor service that make monitoring servers easy and simple. Gonitor provides many cool features for monitoring server CPU, Memory, GPU, Disk, Load, Network, etc. If you like or use Gonitor, please star or share it ! 

## Installation

Get the gonitor repository

```
go get github.com/gonitor/gonitor

cd echo $GOPATH/src/github.com/gonitor/gonitor
```

And install dependencies

```
go get -u github.com/golang/dep/cmd/dep

dep ensure
```

## Running the tests

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
sh coverage.sh
```

## Build and Run

Build and run native binary

```
go build -o gonitor .

./gonitor
```
Build native binary for multiple platforms (darwin, windows and linux)

```
sh BuildMulti.sh
```

## Docker support 

Build docker image

```
docker build -t gonitor/gonitor .
```

Run docker container

```
docker run -d --name gonitor -p 9000:9000 gonitor/gonitor
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/gonitor/gonitor/tags). 

## Authors

* **Ai Nguyen** - [AiNguyenKaka](https://github.com/ainguyenkaka)

See also the list of [contributors](https://github.com/gonitor/gonitor/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

