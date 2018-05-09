package main

import (
	"flag"
	"github.com/ghodss/yaml"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

var parameterFile = flag.String("p", "parameters.yaml", "parameter file in yaml format")
var templatedFile = flag.String("t", "", "templated file")

func doTemplate(tmpl, target io.Reader, output io.Writer) (err error) {
	templateMap := map[string]interface{}{}
	var templateBytes []byte
	var targetBytes []byte
	var tmplTmpl *template.Template

	if templateBytes, err = ioutil.ReadAll(tmpl); err != nil {
		return err
	}
	if targetBytes, err = ioutil.ReadAll(target); err != nil {
		return err
	}
	if err = yaml.Unmarshal(templateBytes, &templateMap); err != nil {
		return err
	}
	if tmplTmpl, err = template.New("template").Parse(string(targetBytes)); err != nil {
		return err
	}
	return tmplTmpl.Execute(output, templateMap)
}

func main() {
	flag.Parse()
	var parameters, templated io.Reader
	var err error
	if parameters, err = os.Open(*parameterFile); err != nil {
		panic(err)
	}
	if templated, err = os.Open(*templatedFile); err != nil {
		panic(err)
	}

	if err := doTemplate(parameters, templated, os.Stdout); err != nil {
		panic(err)
	}
}
