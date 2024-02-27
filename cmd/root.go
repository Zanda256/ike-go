/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/Zanda256/ike-go/pkg-foundation/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg = struct {
	Web web.Config
	Db  dbsql.Config
}{}

var (
	dbClient   *dbsql.DB
	httpClient *web.ClientProvider
	log        *logger.Logger

	traceIDFunc = func(ctx context.Context) string {
		//return web.GetTraceID(ctx)
		return "not_set_up_yet"
	}

	events = logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT ******")
		},
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ike-go",
	Short: "ike-go command needs atleast one subcommand to do work",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

func initConfig() {
	// Enable automatic environment variable binding
	viper.AutomaticEnv()

	// Set default values for Web http client config
	viper.SetDefault("MAX_IDLE_CONNS", 20)
	viper.SetDefault("MAX_IDLE_CONNS_PER_HOST", 20)
	viper.SetDefault("WEB_TIMEOUT", 10)

	// Access the value of Web config
	cfg.Web.Timeout = viper.GetInt("WEB_TIMEOUT")
	cfg.Web.MaxIdleConns = viper.GetInt("MAX_IDLE_CONNS")
	cfg.Web.MaxIdleConnsPerHost = viper.GetInt("MAX_IDLE_CONNS_PER_HOST")

	cfg.Db.DisableTLS = viper.GetBool("DB_DISABLE_TLS")
	cfg.Db.User = viper.GetString("DB_USER")
	cfg.Db.Password = viper.GetString("DB_PASSWORD")
	cfg.Db.Host = viper.GetString("DB_HOST")
	cfg.Db.Name = viper.GetString("DB_NAME")
	cfg.Db.Pool.MaxConnIdleTime = viper.GetString("MAX_CONN_IDLE_TIME")
	cfg.Db.Pool.MaxConnLifetime = viper.GetString("MAX_CONN_LIFE_TIME")
	cfg.Db.Pool.MaxConnLifetimeJitter = viper.GetString("MAX_CONN_LIFE_TIME_JITTER")
	cfg.Db.Pool.MaxConns = viper.GetString("MAX_CONNS")
	cfg.Db.Pool.MinConns = viper.GetString("MIN_CONNS")
}

func init() {
	cobra.OnInitialize(initConfig)
	var err error
	dbClient, err = dbsql.Open(context.Background(), cfg.Db)
	if err != nil {
		os.Exit(1)
	}
	if err = dbsql.StatusCheck(context.Background(), dbClient); err != nil {
		os.Exit(2)
	}
	httpClient = web.NewClientProvider(cfg.Web)

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "ike-go", traceIDFunc, events)
}
