/*
Copyright © 2022 DANIEL SOLER <dsolerh.cinter95@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "miggen",
	Short: "A migration template based generator.",
	Long: `A migration template based generator is uses a config file that is
generated with the command "init". It creates a directory ".miggen" and inside
a file ".config.{json|yaml}". The format for storage is json for default
For example:
	miggen init 			# prompt various config option
	miggen init -y 		# create the config with default options

For generating the migration file use the command with 
the flags -type and -name. The type is optional and defaults to "blank"
wich creates an empty file
For example:
	miggen -name=<migration-name> -type=<custom-type>
	miggen -name=<migration-name> 		# with default type "blank"

Custom types can be added to the config file .miggen/.config.{json|yaml}`,
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
