# Skunk: Build a base project

Everybody enjoys a skunkworks project now and then. But keep the
projects consistent! Skunk sets up the necessary directories for
creating a new project.

It's also a demonstation of how to write Cookoo programs in Go.

## Installation

```
$ cd $GOPATH
$ go get github.com/technosophos/skunk
$ cp -a src/github.com/technosophos/skunk/dot-skunk ~/.skunk
$ bin/skunk
```

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

Basic help is available using the `-help` flag:

```
$ skunk -help
Usage: skunk [-OPTIONS] PROJECTNAME

OPTIONS:
  -confd="/Users/mattbutcher/.skunk": Set the directory with settings.json
  -type=[]: Project type (e.g. 'go', 'php'). Separate multiple values with ','

EXAMPLES:
skunk MyProject			# Create MyProject, using defaults.
skunk -condf=skunkd MyProject	# Create MyProject using config files in ./skunkd/.
skunk -type=go,git MyProject	# Create MyProject and use the preferences for user-defined 'go' and 'git' projects.
```

#### Options Explained

- confd: The directory that holds your Skunk configuration files. By
default, Skunk looks in `~/.skunkrc`. *A skunk directory MUST have a
settings.json file*.
- type: The type of project. This allows you to add additional
directories and templates for specific kinds of projects. (See below.)


### Customizing

* Edit your `~/.skunk/settings.json` file to your tastes.
* Create your own templates in `~/.skunk/tpl`

The template format is [documented in the Go documentation](http://golang.org/pkg/text/template/#pkg-overview).
Templates can range from simple to complex, including branching logic,
constructing pipelines, and calling functions.

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

- *author*: A name. This is conventional (not required), and is used as a
template variable in, for example, `tpl/MIT.txt`.
- *email*: This is conventional, and is used as a template variable in,
for example, `tpl/MIT.txt`.
- *directories*: This is an array of directories to create. You can use
nested directories like this: 'foo/bar/baz' and skunk will create the
entire hierarchy for you.
- *templates*: This is a map of template files to destination files. Each
template will be processed, and the results will be saved at the given
directory (rooted from the project). So `tpl/README.tpl` will be run
through the template engine, and then saved as `$PROJECTNAME/README.md`.

Say you wanted to add a template variable for `subtitle`. You can
accomplish this by adding `"subtitle": "My Subtitle"` in the
`settings.json` file, and it will thereafter be available in templates
as `.subtitle`.

### Types

In addition to basic customizations, you can declare your own *type*. A
type, in Skunk, reflects *what kind (type) of project* you are creating.

In Skunk 1.0, creating types is trivially easy. You just add a type as
an object in your `settings.json` file.

Here's an example from the default `settings.json`:

```javascript
{
  "author": "Matt Butcher",
  "email": "technosophos@example.com",
	"description": "",
	"keywords": "",

  "directories": ["src", "lib", "docs", "build"],

  "templates": {
    "tpl/README.tpl": "README.md",
    "tpl/MIT.txt": "LICENSE.txt",
    "tpl/gitignore.tpl": ".gitignore"
  },


	"php": {
		"directories": ["vendor"],
		"templates": {
			"tpl/php/composer.json": "composer.json"
		}
	}
}

```

See the `php` entry? That's a declaration of the type `php`. Now, to
start a new PHP project we can invoke Skunk like this:

```
$ skunk -type php MyProject
```

And this will do the following:

- It will start with the `directories` and `tempaltes` declared at the
base of the `settings.json` file.
- It will merge the `directories` and `templates` from the `php` section
into those lists.
- It will create `MyProject` along with all of the directories and
templates in the final list.

You can set multiple types for a new project:

```
$ skunk -type php,stackato,fabric MyProject
```

This will merge the `php`, `stackato`, and `fabric` settings all into
the main settings. When it comes to overwriting, order is important. The
*last type* has the highest precendence.


## License

This code is Open Source, and is licensed under the MIT license. See
LICENSE.txt.
