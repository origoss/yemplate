package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
)

var rootCmd = &cobra.Command{
	Use:   "yemplate",
	Short: "yemplate is a simple CLI wrapper for the go text/template library",
	Run: func(cmd *cobra.Command, args []string) {
		var parameters, templated io.Reader
		var err error
		if parameters, err = os.Open(parameterFile); err != nil {
			panic(err)
		}
		if templated, err = os.Open(templateFile); err != nil {
			panic(err)
		}

		if err := doTemplate(parameters, templated, os.Stdout); err != nil {
			panic(err)
		}
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
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&parameterFile, "parameters", "p", "parameters.yaml", "file name of the yaml file with the parameters to insert")
	rootCmd.PersistentFlags().StringVarP(&templateFile, "template", "t", "", "file name of the template file")
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
