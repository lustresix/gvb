// Package cmd /*
package cmd

import (
	"gbv2/config"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbv2",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
	SilenceUsage: true,
	// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gbv2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Read the configuration file")
	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
