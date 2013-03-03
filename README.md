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

The above will create a new project scaffolded for you. For example,
using the default `dot-skunk` directory, you would get this:

```
$ skunk MyProject
$ ls -lah MyProject/
total 24
drwxr-xr-x   9 mattbutcher  staff   306B Mar  2 19:15 .
drwxr-xr-x  11 mattbutcher  staff   374B Mar  2 19:15 ..
-rw-r--r--   1 mattbutcher  staff    24B Mar  2 19:15 .gitignore
-rw-r--r--   1 mattbutcher  staff   1.1K Mar  2 19:15 LICENSE.txt
-rw-r--r--   1 mattbutcher  staff    69B Mar  2 19:15 README.md
drwxr-xr-x   2 mattbutcher  staff    68B Mar  2 19:15 build
drwxr-xr-x   2 mattbutcher  staff    68B Mar  2 19:15 docs
drwxr-xr-x   2 mattbutcher  staff    68B Mar  2 19:15 lib
drwxr-xr-x   2 mattbutcher  staff    68B Mar  2 19:15 src
```

Of course, all of this is customizable.

### Customizing

* Edit your `~/.skunk/settings.json` file to your tastes.
* Create your own templates in `~/.skunk/tpl`

The template format is [documented in the Go documentation](http://golang.org/pkg/text/template/#pkg-overview)

Among other things, all of the values in `settings.json` are available
in your templates. So you can add new data to your JSON an access it in
your templates.

The default `settings.json` file might look something like this:

```json
{
  "author": "Matt Butcher",
  "email": "matt@example.com",

  "directories": ["src", "lib", "docs", "build"],

  "templates": {
    "tpl/README.tpl": "README.md",
    "tpl/MIT.txt": "LICENSE.txt",
    "tpl/gitignore.tpl": ".gitignore"
  }

}
```

Here's what these are for:

- author: A name. This is conventional (not required), and is used as a
template variable in, for example, `tpl/MIT.txt`.
- email: This is conventional, and is used as a template variable in,
for example, `tpl/MIT.txt`.
- directories: This is an array of directories to create. You can use
nested directories like this: 'foo/bar/baz' and skunk will create the
entire hierarchy for you.
- templates: This is a map of template files to destination files. Each
template will be processed, and the results will be saved at the given
directory (rooted from the project). So `tpl/README.tpl` will be run
through the template engine, and then saved as `$PROJECTNAME/README.md`.

Say you wanted to add a template variable for `subtitle`. You can
accomplish this by adding `"subtitle": "My Subtitle"` in the
`settings.json` file, and it will thereafter be available in templates
as `.subtitle`.

## License

This code is Open Source, and is licensed under the MIT license. See
LICENSE.txt.
