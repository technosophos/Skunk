# Skunk: Build a base project

Everybody enjoys a skunkworks project now and then. But keep the
projects consistent! Skunk sets up the necessary directories for
creating a new project.

It's also a demonstation of how to write Cookoo programs in Go.

## Installation

* Clone this repo
* Build the program: "go build -o skunk"
* Copy `dot-skunk` to your home directory `cp -a dot-skunk ~/.skunk`

## Usage

Here's how you use it:

```
$ skunk ProjectName
```

The above will create a new project scaffolded for you.

### Customizing

* Edit your `~/.skunk/settings.json` file to your tastes.
* Create your own templates in `~/.skunk/tpl`

The template format is [documented in the Go documentation](http://golang.org/pkg/text/template/#pkg-overview)

Among other things, all of the values in `settings.json` are available
in your templates. So you can add new data to your JSON an access it in
your templates.

The End
