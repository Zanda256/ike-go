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

// TODO: Add OpenAPI config env vars
// TIKTOKEN_CACHE_DIR
// TOGETHER_API_KEY
// OPENAI_API_KEY

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
	// Run: func(cmd *cobra.Command, args []string) {},
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
	if value == "" {
		panic(fmt.Sprintf("env var %s is required", key))
	}
	return value
}

var initConfig = func() {
	// Access the values of Web config
	cfg.Web.Timeout, _ = strconv.Atoi(getEnvValue("IKE_WEB_TIMEOUT", "10").(string))
	cfg.Web.MaxIdleConns, _ = strconv.Atoi(getEnvValue("IKE_MAX_IDLE_CONNS", "20").(string))
	cfg.Web.MaxIdleConnsPerHost, _ = strconv.Atoi(getEnvValue("IKE_MAX_IDLE_CONNS_PER_HOST", "20").(string))

	cfg.Db.DisableTLS = getEnvValue("IKE_DB_DISABLE_TLS", true).(bool)
	cfg.Db.User = mustGet("IKE_DB_USER")
	cfg.Db.Password = mustGet("IKE_DB_PASSWORD")
	cfg.Db.Host = mustGet("IKE_DB_HOST")
	cfg.Db.Name = mustGet("IKE_DB_NAME") //ike_scripts_db
	cfg.Db.Pool.MaxConnIdleTime = getEnvValue("IKE_MAX_CONN_IDLE_TIME", "20s").(string)
	cfg.Db.Pool.MaxConnLifetime = getEnvValue("IKE_MAX_CONN_LIFE_TIME", "20s").(string)
	cfg.Db.Pool.MaxConnLifetimeJitter = getEnvValue("IKE_MAX_CONN_LIFE_TIME_JITTER", "10s").(string)
	cfg.Db.Pool.MaxConns = getEnvValue("IKE_MAX_CONNS", "20").(string)
	cfg.Db.Pool.MinConns = getEnvValue("IKE_MIN_CONNS", "20").(string)
	cfg.Db.Pool.HealthCheckPeriod = getEnvValue("IKE_POOL_HEALTH_CHECK_PERIOD", "20s").(string)
}

func init() {
	//cobra.OnInitialize(initConfig)

	//var err error
	//fmt.Printf("\n%+v\n", cfg)
	//dbClient, err = dbsql.Open(context.Background(), cfg.Db)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(2)
	//}
	//if err = dbsql.StatusCheck(context.Background(), dbClient); err != nil {
	//	fmt.Printf("\ndbsql.StatusCheck: %+v\n", err.Error())
	//	os.Exit(3)
	//}
	//httpClient = web.NewClientProvider(cfg.Web)
	//
	//log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "ike-go", traceIDFunc, events)

	fmt.Printf("\nroot init running\n")
	//if dbClient == nil {
	//	fmt.Printf("\ndbClient pointer is nil\n")
	//} else {
	//	fmt.Printf("\ndbClient pointer is not nil\n")
	//}
	//if httpClient == nil {
	//	fmt.Printf("\nhttpClient pointer is nil\n")
	//} else {
	//	fmt.Printf("\nhttpClient pointer is not nil\n")
	//}
	//if log == nil {
	//	fmt.Printf("\nroot.log pointer is nil\n")
	//	//imp.Log.Info(context.Background(), "log pointer is not nil")
	//} else {
	//	fmt.Printf("\nroot.log pointer is not nil\n")
	//}
}
