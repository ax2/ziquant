package spider

type Spider interface {
	FundList(int64) (int64, int64)
	Visit(url string)
}

type Fund struct {
	Number     int64   `json:"symbol"`      // 基金编号
	Name       string  `json:"name"`        // 名称
	Date       string  `json:"date"`        // 日期
	Price      string  `json:"dwjz"`        // 单位净值
	Total      string  `json:"ljjz"`        // 累计净值
	Daily      float64 `json:"daily"`       // 单日收益
	Weekly     float64 `json:"weekly"`      // 周收益
	Monthly    float64 `json:"monthly"`     // 月收益
	ThreeMonth float64 `json:"three_month"` // 最近三个月收益
	SixMonth   float64 `json:"six_month"`   // 最近半年收益
	OneYear    float64 `json:"one_year"`    // 最近一年收益
	TwoYear    float64 `json:"two_year"`    // 最近两年收益
	ThreeYear  float64 `json:"three_year"`  // 最近三年收益
	FromYear   float64 `json:"form_year"`   // 今年以来收益
	FromStart  float64 `json:"form_start"`  // 成立以来收益

	Created  string  `json:"clrq"`       // 成立日期
	Company  string  `json:"jjglr_code"` // 基金公司
	Manager  string  `json:"jjjl"`       // 基金经理
	Jzrq     string  `json:"jzrq"`       //
	PerNav   string  `json:"per_nav"`    //
	Sname    string  `json:"sname"`      // 别名
	TotalNav string  `json:"total_nav"`  //
	Zjzfe    float64 `json:"zjzfe"`      //
	Zmjgm    string  `json:"zmjgm"`      //
}

type FundList []*Fund
