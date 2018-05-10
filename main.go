package main

import (
	"github.com/origoss/yemplate/cmd"
)

func main() {
	cmd.Execute()
	// 	flag.Parse()
	// 	var parameters, templated io.Reader
	// 	var err error
	// 	if parameters, err = os.Open(*parameterFile); err != nil {
	// 		panic(err)
	// 	}
	// 	if templated, err = os.Open(*templatedFile); err != nil {
	// 		panic(err)
	// 	}

	// 	if err := doTemplate(parameters, templated, os.Stdout); err != nil {
	// 		panic(err)
	// 	}
}
