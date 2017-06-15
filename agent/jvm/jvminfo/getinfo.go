package jvminfo

import "monitor/agent/base"
//get方法获取jok信息
func ResGet(url,arg string) []byte{
	_,res := base.HttpGet(url+"read/java.lang:type="+arg)
	return  res
}
