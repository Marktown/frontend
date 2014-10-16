# Marktown Frontend

Web frontend for [Marktown](https://github.com/Marktown) written in [Go](http://go-lang.org) using the [Beego](http://beego.me/) web framework.

Project to learn golang.org to build a Markdown writing environment.

[![Build Status](https://travis-ci.org/Marktown/frontend.svg?branch=master)](https://travis-ci.org/Marktown/frontend)

## Development Requirements

* [Go](http://golang.org/)
* [Ruby](http://ruby-lang.org) for [Sass](http://sass-lang.com/)
* [Node.js](http://nodejs.org/) + [npm](https://github.com/npm/npm) for asset management via [Bower](http://bower.io)
* [Livereload](http://livereload.com/)

## Development Environment

1. Setup dependencies

        $ make prepare

2. Run the Beego frontend on [http://localhost:8080](http://localhost:8080)

        $ make run

3. Open the web browser

        $ open http://localhost:8080

### Advanced

* [Livereload](http://livereload.com/) and recompile assets

        $ make watch

## General TODOs

* better name for FileStore "base"
* implement base.FileStore.ReadDirTree
* documentation: my proposal for function-docu-style:

        // Create a new user 
        // Param name : string
        // Param age  : int
        // Return User
        func CreateUser (name string, age int) (user User)

## License

See [LICENSE](LICENSE).
