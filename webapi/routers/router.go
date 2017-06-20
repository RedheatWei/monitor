// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"webapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/collect_cpu",
			beego.NSInclude(
				&controllers.CollectCpuController{},
			),
		),

		beego.NSNamespace("/collect_cpu_infostat",
			beego.NSInclude(
				&controllers.CollectCpuInfostatController{},
			),
		),

		beego.NSNamespace("/collect_disk",
			beego.NSInclude(
				&controllers.CollectDiskController{},
			),
		),

		beego.NSNamespace("/collect_disk_iocountersstat",
			beego.NSInclude(
				&controllers.CollectDiskIocountersstatController{},
			),
		),

		beego.NSNamespace("/collect_disk_partitionstat",
			beego.NSInclude(
				&controllers.CollectDiskPartitionstatController{},
			),
		),

		beego.NSNamespace("/collect_host",
			beego.NSInclude(
				&controllers.CollectHostController{},
			),
		),

		beego.NSNamespace("/collect_host_userstat",
			beego.NSInclude(
				&controllers.CollectHostUserstatController{},
			),
		),

		beego.NSNamespace("/collect_jvm",
			beego.NSInclude(
				&controllers.CollectJvmController{},
			),
		),

		beego.NSNamespace("/collect_load",
			beego.NSInclude(
				&controllers.CollectLoadController{},
			),
		),

		beego.NSNamespace("/collect_mem",
			beego.NSInclude(
				&controllers.CollectMemController{},
			),
		),

		beego.NSNamespace("/collect_net",
			beego.NSInclude(
				&controllers.CollectNetController{},
			),
		),

		beego.NSNamespace("/collect_net_interfacestat",
			beego.NSInclude(
				&controllers.CollectNetInterfacestatController{},
			),
		),

		beego.NSNamespace("/collect_net_interfacestat_addrs",
			beego.NSInclude(
				&controllers.CollectNetInterfacestatAddrsController{},
			),
		),

		beego.NSNamespace("/collect_net_iocountersstat",
			beego.NSInclude(
				&controllers.CollectNetIocountersstatController{},
			),
		),

		beego.NSNamespace("/server_info_base",
			beego.NSInclude(
				&controllers.ServerInfoBaseController{},
			),
		),

		beego.NSNamespace("/server_info_ip",
			beego.NSInclude(
				&controllers.ServerInfoIpController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
