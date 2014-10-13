# Marktown Frontend

Web frontend for [Marktown](https://github.com/Marktown) written in [Go](http://go-lang.org) using the [Beego](http://beego.me/) web framework.

Project to learn golang.org to build a Markdown writing environment.

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

## License

See [LICENSE](LICENSE).
