package mysqldb
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"fmt"
	"monitor/server/base"
)
var engine *xorm.Engine

var ServerConfig base.ServerConfig
func init()  {
	ServerConfig = base.GetConfig()
}

func ConnDB() *xorm.Engine{
	dbconf := ServerConfig.DB
	dbconn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s",dbconf.User,dbconf.Passwd,dbconf.Protocol,dbconf.Host,dbconf.Database,dbconf.Charset)
	var err error
	engine, err = xorm.NewEngine(dbconf.Type, dbconn)
	if err!=nil{
		fmt.Println(err)
	}
	return engine
}

func GetAllowIpList() []map[string]interface{}{
	db := ConnDB()
	//iplist := new(Server_info_ip)
	sql := "SELECT *,server_info_ip.serverid serverid,server_info_base.server server FROM server_info_ip,server_info_base WHERE server_info_base.id=server_info_ip.serverid"
	results, err := db.Sql(sql).Query().List()
	if err != nil{
		fmt.Println(err)
	}
	return results
}
