package fund

import (
	"fmt"
	"sort"

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

type FundRule struct {
	List  spider.FundList
	Score func(*spider.Fund) float64
}

func Rule001(f *spider.Fund) float64 {
	return f.OneYear*0.3 + f.Monthly*0.8 + f.Weekly*0.3
}

func (n FundRule) Len() int {
	return len(n.List)
}

func (n FundRule) Less(i, j int) bool {
	return n.Score(n.List[i]) > n.Score(n.List[j])
}

func (n FundRule) Swap(i, j int) {
	n.List[i], n.List[j] = n.List[j], n.List[i]
}

func init() {
}

func fundList(cmd *cobra.Command, args []string) {
	s := &spider.EastMoneySpider{}
	page := int64(1)
	allPages := int64(10000)
	var allItems spider.FundList
	var items spider.FundList
	for page < allPages {
		_, allPages, items = s.FundList(page)
		allItems = append(allItems, items...)
		page++
		if page > 10 {
			break
		}
	}

	fr := FundRule{
		List:  allItems,
		Score: Rule001,
	}
	sort.Sort(fr)

	for _, f := range fr.List {
		fmt.Printf("[%6d]%-15.15s\t%.2f\t%.2f\t%.2f\t%.2f\t%.2f\n", f.Number, f.Name, fr.Score(f), f.Daily, f.Weekly, f.Monthly, f.OneYear)
	}
}
