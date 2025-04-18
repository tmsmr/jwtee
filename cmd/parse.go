package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tmsmr/jwtee/internal/pkg/log"
	"github.com/tmsmr/jwtee/internal/pkg/stdin"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		val, err := stdin.Read()
		if err != nil {
			log.Error("Failed to read input", err)
		}
		log.Debug(val)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
