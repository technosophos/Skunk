package main

import (
	"cookoo"
	//"fmt"
	"os"
)

func main() {
  homedir := os.ExpandEnv("${HOME}/.skunk")
	registry, router, cxt := cookoo.Cookoo()

  cxt.Add("basedir", homedir)
  cxt.Add("PROJECT", "test")
  cxt.Add("YEAR", "2013")

	registry.
  Route("scaffold", "Scaffold a new app.").
		Does(Notice, "Testing").
		Does(LoadSettings, "settings").
			Using("file").WithDefault(homedir + "/settings.json").From("cxt:SettingsFile").
		Does(MakeDirectories, "dirs").
			Using("directories").From("cxt:directories").
    Does(RenderTemplates, "template").
      Using("templates").From("cxt:templates").
  Route("help", "Print help").
    Does(Usage, "Testing").
	Done()

	//router.HandleRequest("help", cxt, false)
	router.HandleRequest("scaffold", cxt, false)
}

