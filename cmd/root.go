/*
Copyright © 2022 DANIEL SOLER <dsolerh.cinter95@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/dsolerh/miggen/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "miggen",
	Short: "A migration template based generator.",
	Long: `A migration template based generator is uses a config file that is
generated with the command "init". It creates a directory ".mig" and inside
a file`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("name", "n", "", "The migration name")
	rootCmd.Flags().StringP("type", "t", config.DefaultType, "The migration type")
}
