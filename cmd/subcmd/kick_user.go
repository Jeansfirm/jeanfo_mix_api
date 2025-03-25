package subcmd

import (
	"fmt"
	session_util "jeanfo_mix/util/session"
	"os"

	"github.com/spf13/cobra"
)

func ExecKickUser(uid int, uname string) {
	fmt.Printf("now kick user with uid:%d or name:%s out\n", uid, uname)
	err := session_util.ClearUserSession(uid, uname)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func GetKickUserCmd() *cobra.Command {
	kickUserCmd := &cobra.Command{
		Use:   "kick_user",
		Short: "Kick the specified user",
		Run: func(cmd *cobra.Command, args []string) {
			uid, _ := cmd.Flags().GetInt("uid")
			uname, _ := cmd.Flags().GetString("uname")
			ExecKickUser(uid, uname)
		},
	}
	kickUserCmd.Flags().Int("uid", 0, "UserID to kick")
	kickUserCmd.Flags().StringP("uname", "n", "", "UserName to kick")

	return kickUserCmd
}
