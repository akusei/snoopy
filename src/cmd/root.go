package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"snoopy/protected"
)

var rootCmd = &cobra.Command{
	Use:           "snoopy",
	Long:          protected.ToolDescription,
	Version:       protected.ToolVersion,
	SilenceUsage:  true,
	SilenceErrors: true,
}

var proxy *string
var proxyPort *int16

func init() {
	proxy = rootCmd.PersistentFlags().StringP("proxy", "x", "", "Proxy to use for connections")
	proxyPort = rootCmd.PersistentFlags().Int16P("proxy-port", "p", 0, "Proxy port")
}

func addExample(cmd *cobra.Command, example string) {
	if len(cmd.Example) == 0 {
		cmd.Example = fmt.Sprintf("%s %s", rootCmd.Use, example)
	} else {
		cmd.Example = fmt.Sprintf("%s\n%s %s", cmd.Example, rootCmd.Use, example)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
