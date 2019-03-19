package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var downloadCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "download <remote> <local>",
		Short: "download a remote file",
		Long:  "download a remote file to the local destination",
		Args:  cobra.ExactArgs(2),
	},
	Run: download,
}

func init() {
	remoteShell.AddCommand(downloadCmd)
}

func download(sh *shell.Shell, args []string) error {
	fmt.Println("look at you, sneakin' a file, how cute")
	return nil
}
