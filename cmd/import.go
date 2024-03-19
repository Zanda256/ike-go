/*
Copyright © 2024 Sekiranda Hamza <sekirandahamza@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/Zanda256/ike-go/pkg-foundation/web"
	"os"

	"github.com/Zanda256/ike-go/internal/core/importers"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport"
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport/stores/wpImportDb"
	"github.com/spf13/cobra"
)

var svc *importers.ImportService

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a download from a source",
	Long: `import command takes a url from which you want to import content.

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	//Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("import command requires atleast one argument")
		}
		if svc.WPress == nil {
			fmt.Printf("\nsvc.WPress is nil\n")
			return fmt.Errorf("svc.WPress is nil")
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
	initConfig()
	var err error
	fmt.Printf("\n%+v\n", cfg)
	dbClient, err = dbsql.Open(context.Background(), cfg.Db)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if err = dbsql.StatusCheck(context.Background(), dbClient); err != nil {
		fmt.Printf("\ndbsql.StatusCheck: %+v\n", err.Error())
		os.Exit(3)
	}
	httpClient = web.NewClientProvider(cfg.Web)

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "ike-go", traceIDFunc, events)
	imp := wpImport.NewWordPressImporter(log, httpClient, wpImportDb.NewStore(log, dbClient))
	//if imp.Storage != nil {
	//	fmt.Printf("\nstorage pointer is not nil\n")
	//}
	//if imp.WebClient == nil {
	//	fmt.Printf("\nWebClient pointer is nil\n")
	//}
	//if imp.Log == nil {
	//	fmt.Printf("\nlog pointer is nil\n")
	//	//imp.Log.Info(context.Background(), "log pointer is not nil")
	//}
	//if dbClient == nil {
	//	fmt.Printf("\ndbClient pointer is nil\n")
	//}
	//if httpClient == nil {
	//	fmt.Printf("\nhttpClient pointer is nil\n")
	//}
	//if log == nil {
	//	fmt.Printf("\nroot.log pointer is nil\n")
	//	//imp.Log.Info(context.Background(), "log pointer is not nil")
	//}
	fmt.Printf("\nimport init running\n")
	svc = &importers.ImportService{
		WPress: imp,
	}
}
