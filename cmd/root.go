package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var version = "devel"

var rootCmd = &cobra.Command{
	Use:     "go-vault",
	Short:   "Vault helper",
	Long:    `TODO: changeme`,
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		Verbose(cmd)
	},
}

// Verbose Increase verbosity.
func Verbose(cmd *cobra.Command) {
	verbose, err := cmd.Flags().GetCount("verbose")
	if err != nil {
		panic(err)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})

	switch verbose {
	case 1:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case 2:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func init() {
	rootCmd.PersistentFlags().CountP("verbose", "v", "Increase verbosity")
	rootCmd.PersistentFlags().BoolP("dryrun", "n", false, "Dry run")
	rootCmd.PersistentFlags().StringP("addr", "", os.Getenv("VAULT_ADDR"), "Vault address")
}

// Execute The main function for the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
