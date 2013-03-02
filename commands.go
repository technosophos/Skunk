package main

import (
	"cookoo"
	"fmt"
	"encoding/json"
	"os"
	"text/template"
)

func Usage(cxt cookoo.Context, params *cookoo.Params) interface{} {
	fmt.Println("Usage: skunk PROJECTNAME")
	return true
}

func Notice(cxt cookoo.Context, params *cookoo.Params) interface{} {
	fmt.Println("Got Here")
	return true
}

func LoadSettings(cxt cookoo.Context, params *cookoo.Params) interface{} {
	// Container for the results.
	var result map[string]interface{}

	// Open file reader
	file, ok := params.Has("file")
	if !ok {
		// Should probably return an empty map.
		return result
	}

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
  d, ok := cxt.Has("directories")
  if !ok {
    // Did nothing.
    return false
  }
	directories := d.([]interface{})

  for _, dir := range directories {
    fmt.Println("Directory: ", dir.(string))
    os.MkdirAll(dir.(string), 0755)
  }

  return true
}

func RenderTemplates(cxt cookoo.Context, params *cookoo.Params) interface{} {
	t, ok := cxt.Has("templates")
	if !ok {
		return false
	}

	templates := t.(map[string]interface{})

	for k, v := range templates {
		tpath := cxt.Get("basedir").(string) + "/" + k
		rendercopy(tpath, v.(string), cxt)
	}

	return true
}

func rendercopy(tpl string, destination string, cxt cookoo.Context) bool {
	t, err := template.ParseFiles(tpl)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Yay")
	out, err := os.Create(destination)
	if err != nil {
		fmt.Println("Could not create ", destination)
		return false
	}
	err = t.Execute(out, cxt.AsMap())
	if err != nil {
		fmt.Println("Skipping template: ", err)
	}

	return true
}
