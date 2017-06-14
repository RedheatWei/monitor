package base

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"monitor/conf"
)
type JvmInfo struct {
	//class
	LoadedClassCount      int32 `json:"LoadedClassCount"`
	UnloadedClassCount    int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
	//memory
	HeapMemoryUsageInit      int64 `json:"HeapMemoryUsageInit"`
	HeapMemoryUsageCommitted int64 `json:"HeapMemoryUsageCommitted"`
	HeapMemoryUsageMax       int64 `json:"HeapMemoryUsageMax"`
	HeapMemoryUsageUsed      int64 `json:"HeapMemoryUsageUsed"`

	NonHeapMemoryUsageInit      int64 `json:"NonHeapMemoryUsageInit"`
	NonHeapMemoryUsageCommitted int64 `json:"NonHeapMemoryUsageCommitted"`
	NonHeapMemoryUsageMax       int64 `json:"NonHeapMemoryUsageMax"`
	NonHeapMemoryUsageUsed      int64 `json:"NonHeapMemoryUsageUsed"`
	//rumtime
	ClassPath string `json:"ClassPath"`
	Uptime    int64 `json:"Uptime"`
	//thread
	TotalStartedThreadCount int32 `json:"TotalStartedThreadCount"`
	PeakThreadCount         int32 `json:"PeakThreadCount"`
	CurrentThreadCpuTime    int64 `json:"CurrentThreadCpuTime"`
	ThreadCount             int32 `json:"ThreadCount"`
	DaemonThreadCount       int32 `json:"DaemonThreadCount"`
	//base
	AgentId   string `json:"agentId"`
	TimeStamp int64 `json:"TimeStamp"`
}
type SysMemInfo struct {
	//memory
	VirtualMemoryStatTotal        uint64 `json:"VirtualMemoryStatTotal"`
	VirtualMemoryStatAvailable    uint64 `json:"VirtualMemoryStatAvailable"`
	VirtualMemoryStatUsed         uint64 `json:"VirtualMemoryStatUsed"`
	VirtualMemoryStatUsedPercent  float64 `json:"VirtualMemoryStatUsedPercent"`
	VirtualMemoryStatFree         uint64 `json:"VirtualMemoryStatFree"`
	VirtualMemoryStatBuffers      uint64 `json:"VirtualMemoryStatBuffers"`
	VirtualMemoryStatCached       uint64 `json:"VirtualMemoryStatCached"`
	VirtualMemoryStatWriteback    uint64 `json:"VirtualMemoryStatWriteback"`
	VirtualMemoryStatDirty        uint64 `json:"VirtualMemoryStatDirty"`
	VirtualMemoryStatWritebackTmp uint64 `json:"VirtualMemoryStatWritebackTmp"`
	VirtualMemoryStatShared       uint64 `json:"VirtualMemoryStatShared"`
	VirtualMemoryStatSlab         uint64 `json:"VirtualMemoryStatSlab"`
	VirtualMemoryStatPageTables   uint64 `json:"VirtualMemoryStatPageTables"`
	VirtualMemoryStatSwapCached   uint64 `json:"VirtualMemoryStatSwapCached"`
	//swap
	SwapMemoryStatTotal       uint64 `json:"SwapMemoryStatTotal"`
	SwapMemoryStatUsed        uint64 `json:"SwapMemoryStatUsed"`
	SwapMemoryStatFree        uint64 `json:"SwapMemoryStatFree"`
	SwapMemoryStatUsedPercent float64 `json:"SwapMemoryStatUsedPercent"`
	SwapMemoryStatSin         uint64 `json:"SwapMemoryStatSin"`
	SwapMemoryStatSout        uint64 `json:"SwapMemoryStatSout"`
}
type SysCpuInfo struct {
	InfoStatCPU        int32 `json:"InfoStatCPU"`
	InfoStatStepping   int32 `json:"InfoStatStepping"`
	InfoStatCores      int32 `json:"InfoStatCores"`
	InfoStatMhz        float64 `json:"InfoStatMhz"`
	InfoStatCacheSize  int32 `json:"InfoStatCacheSize"`
	InfoStatVendorID   string `json:"InfoStatVendorID"`
	InfoStatFamily     string `json:"InfoStatFamily"`
	InfoStatModel      string `json:"InfoStatModel"`
	InfoStatPhysicalID string `json:"InfoStatPhysicalID"`
	InfoStatCoreID     string `json:"InfoStatCoreID"`
	InfoStatModelName  string `json:"InfoStatModelName"`
	InfoStatMicrocode  string `json:"InfoStatMicrocode"`
	InfoStatFlags      []string `json:"InfoStatFlags"`

	TimesStatCPU       float64 `json:"TimesStatCPU"`
	TimesStatUser      float64 `json:"TimesStatUser"`
	TimesStatSystem    float64 `json:"TimesStatSystem"`
	TimesStatIdle      float64 `json:"TimesStatIdle"`
	TimesStatNice      float64 `json:"TimesStatNice"`
	TimesStatIowait    float64 `json:"TimesStatIowait"`
	TimesStatIrq       float64 `json:"TimesStatIrq"`
	TimesStatSoftirq   float64 `json:"TimesStatSoftirq"`
	TimesStatSteal     float64 `json:"TimesStatSteal"`
	TimesStatGuest     float64 `json:"TimesStatGuest"`
	TimesStatGuestNice float64 `json:"TimesStatGuestNice"`
	TimesStatStolen    float64 `json:"TimesStatStolen"`
}
type SysDiskInfo struct {
	IOCountersStatReadCount        uint64 `json:"IOCountersStatReadCount"`
	IOCountersStatMergedReadCount  uint64 `json:"IOCountersStatMergedReadCount"`
	IOCountersStatWriteCount       uint64 `json:"IOCountersStatWriteCount"`
	IOCountersStatMergedWriteCount uint64 `json:"IOCountersStatMergedWriteCount"`
	IOCountersStatReadBytes        uint64 `json:"IOCountersStatReadBytes"`
	IOCountersStatWriteBytes       uint64 `json:"IOCountersStatWriteBytes"`
	IOCountersStatReadTime         uint64 `json:"IOCountersStatReadTime"`
	IOCountersStatWriteTime        uint64 `json:"IOCountersStatWriteTime"`
	IOCountersStatIopsInProgress   uint64 `json:"IOCountersStatIopsInProgress"`
	IOCountersStatIoTime           uint64 `json:"IOCountersStatIoTime"`
	IOCountersStatWeightedIO       uint64 `json:"IOCountersStatWeightedIO"`
	IOCountersStatName             string `json:"IOCountersStatName"`
	IOCountersStatSerialNumber     string `json:"IOCountersStatSerialNumber"`

	PartitionStatDevice     string `json:"PartitionStatDevice"`
	PartitionStatMountpoint string `json:"PartitionStatMountpoint"`
	PartitionStatFstype     string `json:"PartitionStatFstype"`
	PartitionStatOpts       string `json:"PartitionStatOpts"`

	UsageStatPath              string `json:"UsageStatPath"`
	UsageStatFstype            string `json:"UsageStatFstype"`
	UsageStatTotal             uint64 `json:"UsageStatTotal"`
	UsageStatFree              uint64 `json:"UsageStatFree"`
	UsageStatUsed              uint64 `json:"UsageStatUsed"`
	UsageStatUsedPercent       float64 `json:"UsageStatUsedPercent"`
	UsageStatInodesTotal       uint64 `json:"UsageStatInodesTotal"`
	UsageStatInodesUsed        uint64 `json:"UsageStatInodesUsed"`
	UsageStatInodesFree        uint64 `json:"UsageStatInodesFree"`
	UsageStatInodesUsedPercent float64 `json:"UsageStatInodesUsedPercent"`
}
type SysHostInfo struct {
	InfoStatHostname             string `json:"InfoStatHostname"`
	InfoStatUptime               uint64 `json:"InfoStatUptime"`
	InfoStatBootTime             uint64 `json:"InfoStatBootTime"`
	InfoStatProcs                uint64 `json:"InfoStatProcs"`
	InfoStatOS                   string `json:"InfoStatOS"`
	InfoStatPlatform             string `json:"InfoStatPlatform"`
	InfoStatPlatformFamily       string `json:"InfoStatPlatformFamily"`
	InfoStatPlatformVersion      string `json:"InfoStatPlatformVersion"`
	InfoStatKernelVersion        string `json:"InfoStatKernelVersion"`
	InfoStatVirtualizationSystem string `json:"InfoStatVirtualizationSystem"`
	InfoStatVirtualizationRole   string `json:"InfoStatVirtualizationRole"`
	InfoStatHostID               string `json:"InfoStatHostID"`

	LSBID          string `json:""`
	LSBRelease     string `json:""`
	LSBCodename    string `json:""`
	LSBDescription string `json:""`

	TemperatureStatSensorKey   string `json:"TemperatureStatSensorKey"`
	TemperatureStatTemperature float64 `json:"TemperatureStatTemperature"`

	UserStatUser     string `json:"UserStatUser"`
	UserStatTerminal string `json:"UserStatTerminal"`
	UserStatHost     string `json:"UserStatHost"`
	UserStatStarted  int `json:"UserStatStarted"`
}
type SysLoadInfo struct {
	AvgStatLoad1  float64 `json:"AvgStatLoad1"`
	AvgStatLoad5  float64 `json:"AvgStatLoad5"`
	AvgStatLoad15 float64 `json:"AvgStatLoad15"`

	MiscStatProcsRunning int `json:"MiscStatProcsRunning"`
	MiscStatProcsBlocked int `json:"MiscStatProcsBlocked"`
	MiscStatCtxt         int `json:"MiscStatCtxt"`
}
type SysNetInfo struct {
	ConnectionStatFd     uint32 `json:"ConnectionStatFd"`
	ConnectionStatFamily uint32 `json:"ConnectionStatFamily"`
	ConnectionStatType   uint32 `json:"ConnectionStatType"`
	ConnectionStatLaddr  Addr  `json:"ConnectionStatLaddr"`
	ConnectionStatRaddr  Addr `json:"ConnectionStatRaddr"`
	ConnectionStatStatus string `json:"ConnectionStatStatus"`
	ConnectionStatUids   []int32 `json:"ConnectionStatUids"`
	ConnectionStatPid    int32 `json:"ConnectionStatPid"`

	IOCountersStatName        string `json:"IOCountersStatName"`
	IOCountersStatBytesSent   uint64 `json:"IOCountersStatBytesSent"`
	IOCountersStatBytesRecv   uint64 `json:"IOCountersStatBytesRecv"`
	IOCountersStatPacketsSent uint64 `json:"IOCountersStatPacketsSent"`
	IOCountersStatPacketsRecv uint64 `json:"IOCountersStatPacketsRecv"`
	IOCountersStatErrin       uint64 `json:"IOCountersStatErrin"`
	IOCountersStatErrout      uint64 `json:"IOCountersStatErrout"`
	IOCountersStatDropin      uint64 `json:"IOCountersStatDropin"`
	IOCountersStatDropout     uint64 `json:"IOCountersStatDropout"`
	IOCountersStatFifoin      uint64 `json:"IOCountersStatFifoin"`
	IOCountersStatFifoout     uint64 `json:"IOCountersStatFifoout"`

	InterfaceStatMTU          int `json:""`
	InterfaceStatName         string `json:""`
	InterfaceStatHardwareAddr string `json:""`
	InterfaceStatFlags        []string `json:""`
	InterfaceStatAddrs        []InterfaceAddr `json:""`

	ProtoCountersStatProtocol string `json:"ProtoCountersStatProtocol"`
	ProtoCountersStatStats    map[string]int64 `json:"ProtoCountersStatStats"`
}
type SysProcessInfo struct {
	IOCountersStatReadCount  uint64 `json:"IOCountersStatReadCount"`
	IOCountersStatWriteCount uint64 `json:"IOCountersStatWriteCount"`
	IOCountersStatReadBytes  uint64 `json:"IOCountersStatReadBytes"`
	IOCountersStatWriteBytes uint64 `json:"IOCountersStatWriteBytes"`

	MemoryInfoExStatRSS    uint64 `json:"MemoryInfoExStatRSS"`
	MemoryInfoExStatVMS    uint64 `json:"MemoryInfoExStatVMS"`
	MemoryInfoExStatShared uint64 `json:"MemoryInfoExStatShared"`
	MemoryInfoExStatText   uint64 `json:"MemoryInfoExStatText"`
	MemoryInfoExStatLib    uint64 `json:"MemoryInfoExStatLib"`
	MemoryInfoExStatData   uint64 `json:"MemoryInfoExStatData"`
	MemoryInfoExStatDirty  uint64 `json:"MemoryInfoExStatDirty"`

	MemoryInfoStatRSS  uint64 `json:"MemoryInfoStatRSS"`
	MemoryInfoStatVMS  uint64 `json:"MemoryInfoStatVMS"`
	MemoryInfoStatSwap uint64 `json:"MemoryInfoStatSwap"`

	MemoryMapsStatPath         string `json:"MemoryMapsStatPath"`
	MemoryMapsStatRss          uint64 `json:"MemoryMapsStatRss"`
	MemoryMapsStatSize         uint64 `json:"MemoryMapsStatSize"`
	MemoryMapsStatPss          uint64 `json:"MemoryMapsStatPss"`
	MemoryMapsStatSharedClean  uint64 `json:"MemoryMapsStatSharedClean"`
	MemoryMapsStatSharedDirty  uint64 `json:"MemoryMapsStatSharedDirty"`
	MemoryMapsStatPrivateClean uint64 `json:"MemoryMapsStatPrivateClean"`
	MemoryMapsStatPrivateDirty uint64 `json:"MemoryMapsStatPrivateDirty"`
	MemoryMapsStatReferenced   uint64 `json:"MemoryMapsStatReferenced"`
	MemoryMapsStatAnonymous    uint64 `json:"MemoryMapsStatAnonymous"`
	MemoryMapsStatSwap         uint64 `json:"MemoryMapsStatSwap"`

	NumCtxSwitchesStatVoluntary   int64 `json:"NumCtxSwitchesStatVoluntary"`
	NumCtxSwitchesStatInvoluntary int64 `json:"NumCtxSwitchesStatInvoluntary"`

	OpenFilesStatPath string `json:"OpenFilesStatPath"`
	OpenFilesStatFd   uint64 `json:"OpenFilesStatFd"`

	ProcessPid int32 `json:"ProcessPid"`

	RlimitStatResource int32 `json:"RlimitStatResource"`
	RlimitStatSoft     int32 `json:"RlimitStatSoft"`
	RlimitStatHard     int32 `json:"RlimitStatHard"`
}
type Addr struct {
	IP   string `json:"ip"`
	Port uint32 `json:"port"`
}
type InterfaceAddr struct {
	Addr string `json:"addr"`
}
type Senddata struct {
	Type           string `json:"type"`
	JvmInfo        JvmInfo `json:"JvmInfo"`
	SysMemInfo     SysMemInfo `json:"SysMemInfo"`
	SysCpuInfo     SysCpuInfo `json:"SysCpuInfo"`
	SysDiskInfo    SysDiskInfo `json:"SysDiskInfo"`
	SysHostInfo    SysHostInfo `json:"SysHostInfo"`
	SysLoadInfo    SysLoadInfo `json:"SysLoadInfo"`
	SysNetInfo     SysNetInfo `json:"SysNetInfo"`
	SysProcessInfo SysProcessInfo `json:"SysProcessInfo"`
}

//get方法
func HttpGet(url string) (n int,res []byte){
	response,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 1,[]byte{}
	}else{
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		return 0,body
	}
}
//读取配置文件
func ReadAgentConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/agent.conf")
	return config.Read(mod,par)
}
//读取服务端配置文件
func ReadServerConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/server.conf")
	return config.Read(mod,par)
}
