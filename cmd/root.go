package cmd

import (
	"errors"
	"github.com/spf13/pflag"
	"github.com/tmsmr/jwtee/internal/pkg/log"
	"os"

	"github.com/spf13/cobra"
)

var (
	debug bool
)

func fatal(msg string, err error) {
	if err == nil {
		return
	}
	log.Error(msg, "err", err)
	os.Exit(1)
}

var rootCmd = &cobra.Command{
	Use:   "jwtee",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		parseCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.ParseFlags(os.Args); !errors.Is(err, pflag.ErrHelp) {
		fatal("Failed to parse flags", err)
	}
	log.EnableDebug(debug)
	fatal("Failed to execute rootCmd command", rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
}
