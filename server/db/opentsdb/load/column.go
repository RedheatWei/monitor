package load

type Collect_load struct {
	Metric string `json:"Metric"`
	TimeStamp int64 `json:"TimeStamp"`
	Value string `json:"Value"`
	//Tags Tags `json:"Tags"`
	Tags interface{} `json:"Tags"`
}
type Tags struct {
	ServerId      int64  `json:"ServerId"`
	AvgStatload1      float64 `json:"AvgStatload1"`
	AvgStatload5      float64 `json:"AvgStatload5"`
	AvgStatload15      float64 `json:"AvgStatload15"`
	MiscStatprocsRunning      int `json:"MiscStatprocsRunning"`
	MiscStatprocsBlocked      int `json:"MiscStatprocsBlocked"`
	MiscStatctxt      int `json:"MiscStatctxt"`
	TimeStamp      int64 `json:"TimeStamp"`
}
