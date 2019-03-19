package terminal

import (
	"fmt"
	"github.com/spf13/cobra"
	"snoopy/shell"
)

var openDirCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "open-dir",
		Short: "manage the root open directory",
		Long:  "Enable, disable or check the status of the root open directory",
	},
}

var openDirEnableCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "enable",
		Short: "enable the root open directory",
	},
	Run: openEnable,
}

var openDirDisableCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "disable",
		Short: "disable the root open directory",
	},
	Run: openDisable,
}

var openDirStatusCmd = &shell.Command{
	Command: &cobra.Command{
		Use:   "status",
		Short: "check the root open directory status",
	},
	Run: openStatus,
}

func init() {
	openDirCmd.AddCommand(remoteShell, openDirEnableCmd)
	openDirCmd.AddCommand(remoteShell, openDirDisableCmd)
	openDirCmd.AddCommand(remoteShell, openDirStatusCmd)
	remoteShell.AddCommand(openDirCmd)
}

func openEnable(sh *shell.Shell, args []string) error {
	fmt.Println("this will let you browse all files in a interwebz browser")
	return nil
}

func openDisable(sh *shell.Shell, args []string) error {
	fmt.Println("no soup for you!")
	return nil
}

func openStatus(sh *shell.Shell, args []string) error {
	fmt.Println("ummm... I don't know what the status is")
	return nil
}
