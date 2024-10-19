package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yann0917/fs-dl/utils"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "查看当前登录的账号账号",
	Long:  `使用 fs-dl user 查看当前登录的账号信息`,
	Args:  cobra.ExactArgs(0),

	RunE: func(cmd *cobra.Command, args []string) error {
		return user()
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}

func user() (err error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "手机号", "姓名", "账户", "积分", "是否绑定微信", "邮箱", "注册时间"})
	table.SetAutoFormatHeaders(true)
	table.SetAutoWrapText(false)

	user, err := Instance.GetUserInfo()
	if err != nil {
		return
	}
	table.Append([]string{
		utils.Int2String(user.Uid),
		user.Mobile, user.Username,
		utils.Float2String(user.AccountBalance),
		utils.Int2String(user.Point),
		utils.Bool2String(user.BindWeChat),
		user.Email,
		utils.UnixMilli2String(user.FirstLoginTime),
	})

	table.Render()
	return
}
