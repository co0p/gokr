GOKR
====

App to show okr like numbers of a fake company. Following a mix of [uncle bob's clean-architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and [ben johnson's advice](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) on package structure.


installation
--------------

This project is using  [golang's dep](https://golang.github.io/dep/), you should use it too!

```bash
$ cd /path/to/project
$ dep ensure # installs the necessary dependencies
$ go install # builds the project and makes the 'gokr' bin available
$ gokr # start the webserver´´´
