package main

import (
	//"github.com/masterminds/cookoo/src/cookoo"
	"cookoo"
	"fmt"
	"flag"
	"os"
	"path"
	"time"
	"strings"
)

func main() {
	homedir := os.ExpandEnv("${HOME}/.skunk")
	var sets templateSet
	flag.StringVar(&homedir, "confd", homedir, "Set the directory with settings.json")
	flag.Var(&sets, "type", "Project type (e.g. 'go', 'php'). Separate multiple values with ','")
	flag.Parse()

	project := flag.Arg(0)
	projectdir := os.ExpandEnv("${PWD}/" + project)
	registry, router, cxt := cookoo.Cookoo()

	cxt.Add("homedir", homedir)
	cxt.Add("basedir", projectdir)
	cxt.Add("project", project)
	cxt.Add("now", time.Now())

	registry.
		Route("scaffold", "Scaffold a new app.").
		Does(LoadSettings, "settings").
		Using("file").WithDefault(path.Join(homedir, "settings.json")).From("cxt:SettingsFile").
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

	//router.HandleRequest("help", cxt, false)
	router.HandleRequest("scaffold", cxt, false)
}

type templateSet []string

func (t *templateSet) Set(arg string) error {
	// Split the string
	/*for _, str := range strings.Split(value, ",") {
		// Clean up string
		// append to the templateSet
		*t = append(*t, str)
	}*/
	*t = append(*t, strings.Split(arg, ",")...)
	return nil
}

func (t *templateSet) String() string {
	//strings.Join(t, ",")
	return fmt.Sprint(*t)
}
