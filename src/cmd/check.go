package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
	"snoopy/exploit"
	"strconv"
	"time"
)

var checkCmd = &cobra.Command{
	Use:   "check <TARGET> [PORT]",
	Short: "Test device for connectivity",
	Long:  "Sample long text",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  check,
}

func init() {
	addExample(checkCmd, "check 192.168.1.144 8443")
	addExample(checkCmd, "check 192.168.1.144")
	addExample(checkCmd, "check 192.168.1.144 --version")

	checkCmd.Flags().Bool("version", false, "return firmware version if vulnerable")
	rootCmd.AddCommand(checkCmd)
}

func check(cmd *cobra.Command, args []string) error {
	port := uint16(443)
	if len(args) == 2 {
		p, err := strconv.ParseUint(args[1], 10, 16)
		if err != nil {
			return err
		}
		port = uint16(p)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if proxy != nil {
		socksUrl, _ := url.Parse(fmt.Sprintf("socks5://%s:%d", *proxy, *proxyPort))
		transport.Proxy = http.ProxyURL(socksUrl)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 5,
	}

	checkVersion, _ := cmd.Flags().GetBool("version")
	version, err := exploit.CheckIfVulnerable(args[0], port, checkVersion, client)
	if err != nil {
		return err
	}

	fmt.Print("Device is vulnerable")
	if checkVersion {
		fmt.Printf(": %s", version)
	}

	fmt.Println()

	return nil
}
