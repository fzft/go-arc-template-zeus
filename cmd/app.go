package main

import (
	"github.com/spf13/cobra"
)

var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "Start the app server",
	Run:
}

type App struct {
}

func startApp() {

}

func init() {
	// Initialize the configuration file.
	zeus.LoadConfig()

	// Initialize the logger.
	zeus.InitLogger()

	// Initialize the database.
	initDB()

	// Initialize the messaging.
	initMessaging()
}

func initDB() {

}

func initMessaging() {

}
