package main

import (
	//"github.com/masterminds/cookoo/src/cookoo"
	"cookoo"
	"flag"
	"os"
	"path"
	"time"
	"fmt"
)

func main() {
	homedir := os.ExpandEnv("${HOME}/.skunk")
	var sets templateSet
	flag.StringVar(&homedir, "confd", homedir, "Set the directory with settings.json")
	flag.Var(&sets, "type", "Project type (e.g. 'go', 'php'). Separate multiple values with ','")

	registry, router, cxt := cookoo.Cookoo()
	registry.
		Route("scaffold", "Scaffold a new app.").
		Does(LoadSettings, "settings").
			Using("file").WithDefault(path.Join(homedir, "settings.json")).
			From("cxt:SettingsFile").
		Does(MergeProjectTypes, "types").
			Using("projectTypes").From("cxt:templateSets").
			Using("directories").From("cxt:directories").
			Using("templates").From("cxt:templates").
		Does(MakeDirectories, "dirs").
			Using("basedir").From("cxt:basedir").
			Using("directories").From("cxt:directories").
		Does(RenderTemplates, "template").
			Using("tpldir").From("cxt:homedir").
			Using("basedir").From("cxt:basedir").
			Using("templates").From("cxt:templates").

		Route("help", "Print help").
		Does(Usage, "HelpText").
		Done()

	// I <3 Closures
	flag.Usage = func() {
		router.HandleRequest("help", cxt, false)
		return
	}
	flag.Parse()

	project := flag.Arg(0)
	projectdir := os.ExpandEnv("${PWD}/" + project)
	cxt.Add("homedir", homedir)
	cxt.Add("basedir", projectdir)
	cxt.Add("project", project)
	cxt.Add("templateSets", &sets)
	cxt.Add("now", time.Now())
	cxt.Add("SettingsFile", path.Join(homedir, "settings.json"))

	// No arg[0] is an error.
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: No project name specified\n\n")
		router.HandleRequest("help", cxt, false)
		return
	}
	//router.HandleRequest("help", cxt, false)
	router.HandleRequest("scaffold", cxt, false)
}

