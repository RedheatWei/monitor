package opentsdb

type Collect struct {
	Metric string `json:"metric"`
	TimeStamp int64 `json:"timestamp"`
	Value float64 `json:"value"`
	Tags Tags `json:"tags"`
}
type Tags struct {
	Server string `json:"server"`
	CtimeStamp int64 `json:"ctime_stamp"`
	Ip string `json:"ip"`
}
type Collect_disk struct {
	Metric string `json:"metric"`
	TimeStamp int64 `json:"timestamp"`
	Value float64 `json:"value"`
	Tags Tags_disk `json:"tags"`
}
type Tags_disk struct {
	Server string `json:"server"`
	CtimeStamp int64 `json:"ctime_stamp"`
	Ip string `json:"ip"`
	UsageStatPath string `json:"usage_stat_path"`
}
type Collect_jvm struct {
	Metric string `json:"metric"`
	TimeStamp int64 `json:"timestamp"`
	Value float64 `json:"value"`
	Tags Tags_jvm `json:"tags"`
}
type Tags_jvm struct {
	Server string `json:"server"`
	CtimeStamp int64 `json:"ctime_stamp"`
	Ip string `json:"ip"`
	ClassPath string `json:"class_path"`
	AgentId string `json:"agent_id"`
	CurrentThreadCpuTime int64
	Uptime int64 `json:"uptime"`
}