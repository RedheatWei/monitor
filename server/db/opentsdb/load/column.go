package load

type Collect_load struct {
	Metric string `json:"metric"`
	TimeStamp int64 `json:"timestamp"`
	Value string `json:"value"`
	Tags Tags `json:"tags"`
}
type Tags struct {
	AvgStatload1      float64 `json:"AvgStatload1"`
	AvgStatload5      float64 `json:"AvgStatload5"`
	AvgStatload15      float64 `json:"AvgStatload15"`
	MiscStatprocsRunning      int `json:"MiscStatprocsRunning"`
	MiscStatprocsBlocked      int `json:"MiscStatprocsBlocked"`
	MiscStatctxt      int `json:"MiscStatctxt"`
	TimeStamp      int64 `json:"TimeStamp"`
}
