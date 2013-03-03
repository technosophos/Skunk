package main

import (
	"cookoo"
	//"fmt"
	"os"
  "flag"
)

func main() {
  flag.Parse()
  project := flag.Arg(0)
  homedir := os.ExpandEnv("${HOME}/.skunk")
  projectdir := os.ExpandEnv("${PWD}/" + project)
	registry, router, cxt := cookoo.Cookoo()

  cxt.Add("homedir", homedir)
  cxt.Add("basedir", projectdir)
  cxt.Add("project", project)
  cxt.Add("YEAR", "2013")

	registry.
  Route("scaffold", "Scaffold a new app.").
		Does(Notice, "Testing").
		Does(LoadSettings, "settings").
			Using("file").WithDefault(homedir + "/settings.json").From("cxt:SettingsFile").
		Does(MakeDirectories, "dirs").
      Using("basedir").From("cxt:basedir").
			Using("directories").From("cxt:directories").
    Does(RenderTemplates, "template").
      Using("tpldir").From("cxt:homedir").
      Using("basedir").From("cxt:basedir").
      Using("templates").From("cxt:templates").
  Route("help", "Print help").
    Does(Usage, "Testing").
	Done()

	//router.HandleRequest("help", cxt, false)
	router.HandleRequest("scaffold", cxt, false)
}

