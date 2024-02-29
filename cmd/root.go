/*
Copyright © 2024 Sekiranda Hamza <sekirandahamza@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/Zanda256/ike-go/internal/data/dbsql"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
	"github.com/Zanda256/ike-go/pkg-foundation/web"
	"github.com/spf13/cobra"
)

var cfg = struct {
	Web web.Config
	Db  dbsql.Config
}{
	Web: web.Config{},
	Db: dbsql.Config{
		Pool: dbsql.PoolConfig{},
	},
}

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
	Short: "ike-go command needs at least one subcommand to do work",
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

func getEnvValue(key string, defaultValue any) any {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func mustGet(key string) string {
	value := os.Getenv(key)
	if value != "" {
		panic(fmt.Sprintf("env var %s is required", key))
	}
	return value
}

var initConfig = func() {
	// Access the values of Web config
	cfg.Web.Timeout, _ = strconv.Atoi(getEnvValue("WEB_TIMEOUT", "10").(string))
	cfg.Web.MaxIdleConns, _ = strconv.Atoi(getEnvValue("MAX_IDLE_CONNS", "20").(string))
	cfg.Web.MaxIdleConnsPerHost, _ = strconv.Atoi(getEnvValue("MAX_IDLE_CONNS_PER_HOST", "20").(string))

	cfg.Db.DisableTLS = getEnvValue("DB_DISABLE_TLS", true).(bool)
	cfg.Db.User = mustGet("DB_USER")
	cfg.Db.Password = mustGet("DB_PASSWORD")
	cfg.Db.Host = mustGet("DB_HOST")
	cfg.Db.Name = mustGet("DB_NAME")
	cfg.Db.Pool.MaxConnIdleTime = getEnvValue("MAX_CONN_IDLE_TIME", "20").(string)
	cfg.Db.Pool.MaxConnLifetime = getEnvValue("MAX_CONN_LIFE_TIME", "20").(string)
	cfg.Db.Pool.MaxConnLifetimeJitter = getEnvValue("MAX_CONN_LIFE_TIME_JITTER", "10").(string)
	cfg.Db.Pool.MaxConns = getEnvValue("MAX_CONNS", "20").(string)
	cfg.Db.Pool.MinConns = getEnvValue("MIN_CONNS", "20").(string)
}

func init() {
	initConfig()

	var err error
	fmt.Printf("\n%+v\n", cfg)
	dbClient, err = dbsql.Open(context.Background(), cfg.Db)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if err = dbsql.StatusCheck(context.Background(), dbClient); err != nil {
		os.Exit(3)
	}
	httpClient = web.NewClientProvider(cfg.Web)

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "ike-go", traceIDFunc, events)
}
