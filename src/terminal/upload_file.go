package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var uploadCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "upload <local_src> <remote_dst>",
		Short: "upload a file to the remote host",
		Long:  "upload a file to the remote host",
		Args:  cobra.ExactArgs(2),
	},
	Run: upload,
}

func init() {
	remoteShell.AddCommand(uploadCmd)
}

func upload(sh *shell.Shell, cmd *shell.Command, args []string) error {
	fmt.Println("uploading all the datas")
	return nil
}
