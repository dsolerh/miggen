/*
Copyright © 2022 DANIEL SOLER <dsolerh.cinter95@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dsolerh/miggen/config"
	"github.com/dsolerh/miggen/pkg/dirutils"
	"github.com/dsolerh/miggen/pkg/scanutil"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create the configuration needed for the cli.",
	Long: `Create the configuration needed for the cli. It show a prompt
for each configuration if the flag -y is not set.`,
	Run: InitCommand,
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().BoolP("yes", "y", false, "Use the default config")
}

func InitCommand(cmd *cobra.Command, args []string) {
	// get cwd
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// join the paths
	path := filepath.Join(cwd, config.DirName, config.ConfName)
	// check that the path exist
	confexist, err := dirutils.Exist(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	if confexist {
		fmt.Println("The config file already exist in this folder.")
		override := scanutil.ScannNChoice("You want to override it (Y/n)? ", 3, scanutil.YesOrNo)
		if !override {
			return
		}
	}

	f, err := dirutils.CreateAll(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer f.Close()

	defconfig, err := cmd.Flags().GetBool("yes")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")

	conf := config.New()
	if !defconfig {
		conf.DateFormat = scanutil.ScannNDate("Enter a date format (%s): ", 3, config.DefaultDateFormat)
		conf.Extension = scanutil.ScannNText("Enter an extension (%s): ", 3, config.DefaultExtension, scanutil.IsExt)
		conf.Separator = scanutil.ScannNText("Enter a separator (%s): ", 3, config.DefaultSeparator, scanutil.IsSep)
		conf.MigDir = scanutil.ScannNText("Enter a directory (%s): ", 3, config.DefaultMigDir, scanutil.IsDir)
	}
	err = encoder.Encode(conf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
