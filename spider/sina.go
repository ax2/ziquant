package spider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SinaSpider struct {
	Spider
}

func (s *SinaSpider) Visit(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("获取原始数据错误")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("获取原始数据错误")
	}
	//fmt.Printf("body=%v, err=%v", string(body), err)

	return string(body), nil
}

func (s *SinaSpider) FundList(page int64) (total int64) {
	xsrv := "IO.XSRV2.CallbackList['hLfu5s99aaIUp7D4']"
	ret, err := s.Visit(fmt.Sprintf("http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/%s/NetValueReturn_Service.NetValueReturnOpen?page=%d&num=40&sort=form_year&asc=0&ccode=&type2=0&type3=", xsrv, page))
	if err != nil || len(ret) < len(xsrv)+1 {
		return
	}

	tmp := strings.Split(ret, xsrv)
	if len(tmp) != 2 || len(tmp[1]) < 3 {
		return
	}

	ret2 := tmp[1][1 : len(tmp[1])-2]
	var result interface{}
	err = json.Unmarshal([]byte(ret2), &result)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	fmt.Printf("%v\n", result)

	return
}
