package cmd

import (
	"github.com/ghodss/yaml"
	"io"
	"io/ioutil"
	"text/template"
)

func mergedParameters(params io.Reader, paramMap map[string]string) (map[string]interface{}, error) {
	var templateBytes []byte
	var err error
	templateMap := make(map[string]interface{})
	if templateBytes, err = ioutil.ReadAll(params); err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(templateBytes, &templateMap); err != nil {
		return nil, err
	}
	for k, v := range paramMap {
		templateMap[k] = v
	}
	return templateMap, nil
}

func doTemplate(tmplMap map[string]interface{}, target io.Reader, output io.Writer) (err error) {
	var targetBytes []byte
	var tmplTmpl *template.Template

	if targetBytes, err = ioutil.ReadAll(target); err != nil {
		return err
	}
	if tmplTmpl, err = template.New("template").Parse(string(targetBytes)); err != nil {
		return err
	}
	return tmplTmpl.Execute(output, tmplMap)
}
