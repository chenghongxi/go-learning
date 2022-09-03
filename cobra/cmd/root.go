package cmd

import (
	"cmd/cobra.go/imp"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	//  "github.com/spf13/viper"
	//"demo/imp"
)

// var cfgFile string
var name string
var age int
var RootCmd = &cobra.Command{
	Use:   "demo",
	Short: "A test demo",
	Long:  `Demo is a test appcation for print things`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			cmd.Help()
			return
		}
		imp.Show(name, age)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	RootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	RootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
}
