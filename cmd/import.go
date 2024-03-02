/*
Copyright © 2024 Sekiranda Hamza <sekirandahamza@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/Zanda256/ike-go/internal/core/importers"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport/stores/wpImportDb"
	"github.com/spf13/cobra"
)

var svc importers.ImportService

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("import command requires atleast one argument")
		}
		err := svc.ImportWP(args)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	svc = importers.ImportService{
		WPress: wpImport.NewWordPressImporter(log, httpClient, wpImportDb.NewStore(log, dbClient)),
	}
}
