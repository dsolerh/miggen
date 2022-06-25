/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dsolerh/miggen/config"
	"github.com/dsolerh/miggen/pkg/dirutils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create the configuration needed for the cli.",
	Long: `Create the configuration needed for the cli. It show a prompt
for each configuration if the flag -y is not set.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		path := filepath.Join(cwd, config.DirName, config.ConfName)
		confexist, err := dirutils.Exist(path)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		if confexist {
			reader := bufio.NewReader(os.Stdin)
			var override bool
			fmt.Println("The config file already exist in this folder.")
			for i := 0; i < 3; i++ {
				fmt.Printf("You want to override it (Y/n)?")
				data, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf("input error: %s\n", err.Error())
					continue
				}
				data = strings.ToLower(data[:len(data)-1])
				if data == "" || data == "y" || data == "yes" {
					override = true
					break
				}
				if data == "n" || data == "no" {
					break
				}
			}
			if !override {
				return
			}
		}

		defconfig, err := cmd.Flags().GetBool("yes")
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		defer f.Close()

		encoder := json.NewEncoder(f)
		encoder.SetIndent("", "\t")

		conf := config.New()
		if defconfig {
			err = encoder.Encode(conf)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().BoolP("yes", "y", false, "Use the default config")
}
