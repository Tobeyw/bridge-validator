package main

import (
	"github.com/bane-labs/bridge-validator/bane"
	"github.com/bane-labs/bridge-validator/config"
	"github.com/bane-labs/bridge-validator/neo3"
	"github.com/robfig/cron/v3"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ValidatorApp struct {
	Config config.Config
	cmd    *cobra.Command
	logger *logrus.Logger
}

func (app *ValidatorApp) initialize() {
	app.setupLogger()

	app.cmd = &cobra.Command{
		Use:   "bridge-validator",
		Short: "Bane Bridge Validator",
		Run:   app.runCmd,
	}

	cobra.OnInitialize(app.initConfig)
	viper.SetEnvPrefix("VALIDATOR")
	viper.AutomaticEnv()
	app.loadEnvFile()
}

func (app *ValidatorApp) setupLogger() {
	app.logger = logrus.New()
	app.logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	app.logger.SetLevel(logrus.InfoLevel)
}

func (app *ValidatorApp) loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		app.logger.Warn("No .env file found.")
	}
}

func (app *ValidatorApp) initConfig() {
	if file := viper.GetString("config"); file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		app.logger.Warn("Can't read config: ", err)
	}

	app.Config = config.Config{
		KeyPairFile:             getStringPointer(viper.GetString("key_pair_file")),
		MessageBrokerAddress:    getStringPointer(viper.GetString("message_broker_url")),
		NeoN3NodeURL:            getStringPointer(viper.GetString("n3_node_url")),
		BaneNodeURL:             getStringPointer(viper.GetString("bane_node_url")),
		BaneBridgeContractAddr:  getStringPointer(viper.GetString("bane_bridge_contract_addr")),
		NeoN3BridgeContractAddr: getStringPointer(viper.GetString("n3_bridge_contract_addr")),
	}
}

func getStringPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (app *ValidatorApp) runCmd(cmd *cobra.Command, args []string) {
	if !app.isConfigValid() {
		app.logger.Fatal("Not all config parameters are specified")
		os.Exit(1)
	}

	app.logger.Info("All config parameters were successfully specified.")
	app.run()
}

func (app *ValidatorApp) isConfigValid() bool {
	config := app.Config
	return config.KeyPairFile != nil &&
		config.MessageBrokerAddress != nil &&
		config.NeoN3NodeURL != nil &&
		config.BaneNodeURL != nil &&
		config.BaneBridgeContractAddr != nil &&
		config.NeoN3BridgeContractAddr != nil
}

func (app *ValidatorApp) run() {
	// This is the start of the bridge validator app.
	// Code here.
	app.logger.Info("Bridge Validator is running.")

	// neo3 server

	crontab := cron.New(cron.WithSeconds()) //精确到秒
	task := func() {
		neo3.Server(app.Config, app.logger)
	}
	spec := "*/15 * * * * ?"
	crontab.AddFunc(spec, task)
	crontab.Start()
	//select {}

	// Bane Server
	go bane.Server(app.Config, app.logger)
}

func main() {
	app := &ValidatorApp{}
	app.initialize()

	if err := app.cmd.Execute(); err != nil {
		app.logger.Fatal(err)
		os.Exit(1)
	}
}
