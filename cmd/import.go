/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/Zanda256/ike-go/internal/core/importers"
	"github.com/Zanda256/ike-go/internal/core/importers/wp"
	"github.com/Zanda256/ike-go/internal/core/importers/wp/stores/wpImportdb"
	"github.com/spf13/cobra"
)

var store = wpImportdb.NewStore(log, dbClient)

var imptSvc = importers.ImportService{
	WPress: wp.NewWordPressImporter(log, httpClient, store),
}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import raw data",
	Long:  `Import  data from  supported  data source`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")

	},
}

//InsertDownload(d Download) (uuid.UUID, error)

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.PersistentFlags().String("wordpress", "wp", "import from a wordpress source")
}
