package cmd

import (
	"fmt"
	"os"

	"bytes"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "yemplate",
	Short: "yemplate is a simple CLI wrapper for the go text/template library",
	RunE: func(cmd *cobra.Command, args []string) error {
		paramMap, err := parameterParser(parameters)
		if err != nil {
			return err
		}

		var parameters, templated io.Reader
		if parameterFile == "" {
			parameters = bytes.NewBufferString("{}")
		} else if parameters, err = os.Open(parameterFile); err != nil {
			return err
		}
		if templated, err = os.Open(templateFile); err != nil {
			return err
		}

		mergedMap, err := mergedParameters(parameters, paramMap)
		if err != nil {
			return err
		}
		return doTemplate(mergedMap, templated, os.Stdout)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	cfgFile       string
	parameterFile string
	templateFile  string
	parameters    []string
)

func parameterParser(params []string) (map[string]string, error) {
	paramMap := make(map[string]string)
	for _, param := range params {
		kvs := strings.Split(param, "=")
		if len(kvs) != 2 {
			return nil, fmt.Errorf("Invalid parameter format: %s", param)
		}
		paramMap[kvs[0]] = kvs[1]
	}
	return paramMap, nil
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&parameterFile, "parameters", "p", "", "file name of the yaml file with the parameters to insert")
	rootCmd.PersistentFlags().StringVarP(&templateFile, "template", "t", "", "file name of the template file")
	rootCmd.PersistentFlags().StringArrayVarP(&parameters, "set", "s", []string{}, "parameters in the format of key=value")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}
}
