package cmd

import (
	"github.com/ghodss/yaml"
	"io"
	"io/ioutil"
	"text/template"
)

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
