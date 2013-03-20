package main

import (
	"cookoo"
	"encoding/json"
	"fmt"
	"strings"
	"os"
	"path"
	"text/template"
	"flag"
)

// Print usage info.
func Usage(cxt cookoo.Context, params *cookoo.Params) interface{} {
	fmt.Fprintf(os.Stderr, "Usage: %s [-OPTIONS] PROJECTNAME\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOPTIONS:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,"\nEXAMPLES:\n")
	fmt.Fprintf(os.Stderr,"skunk MyProject\t\t\t# Create MyProject, using defaults.\n")
	fmt.Fprintf(os.Stderr,"skunk -condf=skunkd MyProject\t# Create MyProject using config files in ./skunkd/.\n")
	fmt.Fprintf(os.Stderr,"skunk -type=go,git MyProject\t# Create MyProject and use the preferences for user-defined 'go' and 'git' projects.\n")
	return true
}

// Load a settings file.
func LoadSettings(cxt cookoo.Context, params *cookoo.Params) interface{} {
	// Container for the results.
	var result map[string]interface{}

	// Open file reader
	file, ok := params.Has("file")
	if !ok {
		// Should probably return an empty map.
		return result
	}

	fmt.Println("Reading settings from ", file)

	stream, err := os.Open(file.(string))
	if err != nil {
		fmt.Println("Could not find settings file: ", err)
		return result
	}

	// Pass reader to JSON decoder
	dec := json.NewDecoder(stream)
	dec.Decode(&result)

	// Load it all into the context:
	for k, v := range result {
		cxt.Add(k, v)
	}

	// Profit
	return result
}

func MakeDirectories(cxt cookoo.Context, params *cookoo.Params) interface{} {
	basedir := params.Get("basedir", ".").(string)
	d, ok := cxt.Has("directories")
	if !ok {
		// Did nothing.
		return false
	}
	directories := d.([]interface{})

	for _, dir := range directories {
		dname := path.Join(basedir, dir.(string))
		fmt.Println("Creating ", dname)
		err := os.MkdirAll(dname, 0755)
		if err != nil {
			fmt.Println("Could not create directory: ", err)
		}
	}

	return true
}

// Merge project-specific settings into the main settings array.
func MergeProjectTypes(cxt cookoo.Context, params *cookoo.Params) interface{} {
	types, ok  := params.Has("projectTypes")
	projectTypes := types.(*templateSet)

	// No merging goin' on here.
	if !ok || projectTypes.Len() == 0 {
		return false
	}
	for _, str := range projectTypes.Templates() {
		fmt.Println(str)
	}

	return true
}

func RenderTemplates(cxt cookoo.Context, params *cookoo.Params) interface{} {
	t, ok := params.Has("templates")
	if !ok {
		return false
	}

	templates := t.(map[string]interface{})

	for k, v := range templates {
		tpath := path.Join(params.Get("tpldir", ".").(string), k)
		opath := path.Join(params.Get("basedir", ".").(string), v.(string))
		rendercopy(tpath, opath, cxt)
	}

	return true
}

func rendercopy(tpl string, destination string, cxt cookoo.Context) bool {
	t, err := template.ParseFiles(tpl)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("Rendering %s into %s...", tpl, destination)
	out, err := os.Create(destination)
	if err != nil {
		fmt.Println("Could not create ", destination)
		return false
	}
	err = t.Execute(out, cxt.AsMap())
	if err != nil {
		fmt.Println("Skipping template: ", err)
		return false
	}
	fmt.Println("done")

	return true
}

// Capture a set of templates that should be merged into the
// main array of templates.
type templateSet struct {
	templates []string
}

func newTemplateSet() *templateSet {
	t := new(templateSet)
	t.templates = make([]string, 5)
	return t
}

func (t *templateSet) Set(arg string) error {
	// Split the string
	for _, str := range strings.Split(arg, ",") {
		// Clean up string
		// append to the templateSet
		t.templates = append(t.templates, strings.TrimSpace(str))
	}
	//*t = append(*t, strings.Split(arg, ",")...)
	return nil
}

func (t *templateSet) String() string {
	//strings.Join(t, ",")
	return fmt.Sprint(t.templates)
}

func (t *templateSet) Len() int {
	return len(t.templates)
}

func (t *templateSet) Templates() []string {
	return t.templates
}
