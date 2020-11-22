package zabbix

type Trend struct {
	Itemid   string `json:"itemid"`
	Clock    string `json:"clock"`
	Num      string `json:"num"`
	ValueMin string `json:"value_min"`
	ValueAvg string `json:"value_avg"`
	ValueMax string `json:"value_max"`
}