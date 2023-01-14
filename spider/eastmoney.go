package spider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/ax2/zi"
)

type EastMoneySpider struct {
	Spider
}

type FundListResult struct {
	AllNum     int64    `json:"allNum"`
	AllPages   int64    `json:"allPages"`
	AllRecords int64    `json:"allRecords"`
	BbNum      int64    `json:"bbNum"`
	EtfNum     int64    `json:"etfNum"`
	FofNum     int64    `json:"fofNum"`
	GpNum      int64    `json:"gpNum"`
	HhNum      int64    `json:"hhNum"`
	LofNum     int64    `json:"lofNum"`
	PageIndex  int64    `json:"pageIndex"`
	PageNum    int64    `json:"pageNum"`
	QdiiNum    int64    `json:"qdiiNum"`
	ZqNum      int64    `json:"zqNum"`
	ZsNum      int64    `json:"zsNum"`
	Datas      []string `json:"datas"`
}

func regJsonData(Data []byte) []byte {
	reg := regexp.MustCompile("([a-zA-Z]\\w*):")
	regStr := reg.ReplaceAllString(string(Data), `"$1":`)
	//fmt.Printf("%v\n", regStr)

	newStr := strings.Replace(regStr, `"http":`, "http:", -1)
	//fmt.Printf("%v\n", newStr)
	return []byte(newStr)
}

func (s *EastMoneySpider) Visit(url string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("referer", "fund.eastmoney.com")

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("获取原始数据错误")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("获取原始数据错误")
	}
	body = regJsonData(body)
	//fmt.Printf("body=%v, err=%v", string(body), err)

	return string(body), nil
}

func (s *EastMoneySpider) FundList(page int64) (total int64, pages int64, items []*Fund) {
	prefix := "var rankData = "
	ret, err := s.Visit(fmt.Sprintf("http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=all&rs=&gs=0&sc=1nzf&st=desc&sd=2021-12-24&ed=2022-12-24&qdii=&tabSubtype=,,,,,&pi=%d&pn=50&dx=1&v=0.4548809002246763", page))
	if err != nil || len(ret) < len(prefix)+1 {
		return
	}

	tmp := strings.Split(ret, prefix)
	if len(tmp) != 2 || len(tmp[1]) < 3 {
		return
	}

	ret2 := tmp[1][:len(tmp[1])-1]
	//fmt.Printf("ret2=%v\n", ret2)
	var result FundListResult
	err = json.Unmarshal([]byte(ret2), &result)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	//fmt.Printf("result=%+v\n", result.Datas)

	for _, s := range result.Datas {
		ff := strings.Split(s, ",")
		//fmt.Printf("ff=%v, %d\n", ff, len(ff))
		if len(ff) < 25 {
			continue
		}
		item := &Fund{
			Number:     zi.Int[int64](ff[0]),
			Name:       ff[1],
			Date:       ff[3],
			Price:      ff[4],
			Total:      ff[5],
			Daily:      zi.Float(ff[6]),
			Weekly:     zi.Float(ff[7]),
			Monthly:    zi.Float(ff[8]),
			ThreeMonth: zi.Float(ff[9]),
			SixMonth:   zi.Float(ff[10]),
			OneYear:    zi.Float(ff[11]),
			TwoYear:    zi.Float(ff[12]),
			ThreeYear:  zi.Float(ff[13]),
			FromYear:   zi.Float(ff[14]),
			FromStart:  zi.Float(ff[15]),
			Created:    ff[16],
		}
		items = append(items, item)
		fmt.Printf("%+v\n", item)
	}

	total = result.AllNum
	pages = result.AllPages

	return
}
