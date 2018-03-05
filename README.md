# Ubuntu Touch Python CMake App Template

A template to generate apps for Ubuntu Touch that are GO 1.6 and QML apps using [clickable](http://clickable.bhdouglass.com/en/latest/) to compile it.

## Get Started

To get started with this template, simply run `clickable init` and choose the
'Go/QML App' option.

or

go get -d github.com/nanu-c/ut-app-go-template

## Dependencies

This uses the github.com/nanu-c/qml-go fork which supports go 1.6 and up.
The gopkg.in/qml.v1 can be replaced without code changes by the github.com/nanu-c/qml-go repro.

## Usage
* install go in the host system. It's not necessary but makes live easier.
* go to the template `cd $GOPATH/github.com/nanu-c/ut-app-go-template`
* get Dependencies `go get -d ./...`
* Set your [gopath](https://github.com/golang/go/wiki/GOPATH) in `clickable.json`
* Run `clickable` to compile and install your new app!
* Get in touch with the [go<->qml api documentation](https://godoc.org/gopkg.in/qml.v1)

## Quirks

* main.go has to be in the main folder
* `ObjectName` can't be found in `Components{}`

## License

Copyright (C) 2018 [Aaron Kimmig](http://nanu-c.org)

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License version 3, as published
by the Free Software Foundation.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranties of MERCHANTABILITY, SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
