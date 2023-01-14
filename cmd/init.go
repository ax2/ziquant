package cmd

import (
	"github.com/ax2/ziquant/cmd/fund"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ziquant",
	Short: "some usefull quant utilities",
}

func init() {
	rootCmd.AddCommand(fund.GetCommands())
}

func RootCmd() *cobra.Command {
	return rootCmd
}
