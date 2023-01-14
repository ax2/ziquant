package fund

import (
	"github.com/ax2/ziquant/spider"
	"github.com/spf13/cobra"
)

func init() {
	subCmd.AddCommand(fundListCmd)
}

var (
	fundListCmd = &cobra.Command{
		Use:   "fundlist",
		Short: "fund list",
		Run:   fundList,
	}
)

func init() {
}

func fundList(cmd *cobra.Command, args []string) {
	s := &spider.EastMoneySpider{}
	page := int64(1)
	allPages := int64(10000)
	for page < allPages {
		_, allPages, _ = s.FundList(page)
		page++
		break
	}
}
