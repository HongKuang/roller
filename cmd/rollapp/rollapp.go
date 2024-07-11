package rollapp

import (
	"github.com/spf13/cobra"

	initrollapp "github.com/dymensionxyz/roller/cmd/rollapp/init"
	"github.com/dymensionxyz/roller/cmd/rollapp/status"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollapp [command]",
		Short: "Commands to initialize and run a RollApp",
	}

	cmd.AddCommand(initrollapp.Cmd)
	cmd.AddCommand(status.Cmd())

	return cmd
}