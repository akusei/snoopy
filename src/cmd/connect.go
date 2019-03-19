package cmd

import (
	"github.com/spf13/cobra"
	"snoopy/terminal"
)

var connectCmd = &cobra.Command{
	Use:   "connect <IP> [PORT]",
	Short: "exploit remote device",
	Long:  "exploit a remove device and gain root access",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  connect,
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func connect(cmd *cobra.Command, args []string) error {
	return terminal.Run(args[0])
}
