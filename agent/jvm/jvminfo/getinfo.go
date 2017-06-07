package jvminfo

import "monitor/base"

func ResGet(url,arg string) []byte{
	_,res := base.HttpGet(url+"read/java.lang:type="+arg)
	return  res
}
