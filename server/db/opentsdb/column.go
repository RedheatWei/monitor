package opentsdb

type Collect struct {
	Metric string `json:"metric"`
	TimeStamp int64 `json:"timestamp"`
	Value interface{} `json:"value"`
	Tags Tags `json:"tags"`
}
type Tags struct {
	Server string `json:"server"`
	CtimeStamp int64 `json:"ctime_stamp"`
}