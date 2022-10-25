package vcdcmd

import (
	"fmt"

	executor "github.com/juanfont/gitlab-machine"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var shellVcdCmd = &cobra.Command{
	Use:   "shell cmd",
	Short: "Opens a shell with the specified command",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing parameters")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		vcdDriver, err := getVcdDriver()
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating vcd driver")
		}
		e, _ := executor.NewExecutor(vcdDriver)
		err = e.Shell(args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating executor")
		}
	},
}
