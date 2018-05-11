package cmd

import (
	"fmt"
	"os"

	"bytes"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
)

var rootCmd = &cobra.Command{
	Use:   "yemplate <templated file>",
	Short: "yemplate is a simple CLI wrapper for the go text/template library",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var params io.Reader
		var templated io.ReadCloser

		templateFile := args[0]
		if parameterFile == "" {
			params = bytes.NewBufferString("{}")
		} else if params, err = os.Open(parameterFile); err != nil {
			return err
		}
		if templated, err = openFileOrWeb(templateFile); err != nil {
			return err
		}

		mergedMap, err := mergedParameters(params, parameters)
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
	parameters    []string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&parameterFile, "values", "f", "", "file name of the yaml file with the parameter values to insert")
	rootCmd.PersistentFlags().StringArrayVarP(&parameters, "set", "", []string{}, "parameters in the format of key=value")
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
