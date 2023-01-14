package fund

import (
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "fund cmd [options]",
	Short: "ziquant fund utilities",
}

func init() {
}

func GetCommands() *cobra.Command {
	return subCmd
}
