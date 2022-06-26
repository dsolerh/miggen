/*
Copyright © 2022 DANIEL SOLER <dsolerh.cinter95@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dsolerh/miggen/config"
	"github.com/dsolerh/miggen/pkg/dirutils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "miggen",
	Short: "A migration template based generator.",
	Long: `A migration template based generator is uses a config file that is
generated with the command "init". It creates a directory ".mig" and inside
a file`,
	Run: MigCommand,
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

func MigCommand(cmd *cobra.Command, args []string) {
	// get cwd
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	path := filepath.Join(cwd, config.DirName, config.ConfName)
	// check that the path exist
	confexist, err := dirutils.Exist(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	conf := config.New()

	if confexist {
		f, err := os.Open(path)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		defer f.Close()

		decoder := json.NewDecoder(f)
		err = decoder.Decode(conf)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	if name == "" {
		fmt.Println("must specified a name for the migration")
	}
	typ, err := cmd.Flags().GetString("type")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	migName := time.Now().Format(conf.DateFormat) + conf.Separator + name + conf.Extension
	f, err := dirutils.CreateAll(filepath.Join(cwd, conf.MigDir, migName))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(conf.Types[typ])
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
