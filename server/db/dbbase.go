package db
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"fmt"
	"monitor/base"
)
var engine *xorm.Engine

var ServerConfig base.ServerConfig
func init()  {
	ServerConfig = base.GetConfig()
}

func ConnDB() *xorm.Engine{
	dbconf := ServerConfig.DB
	dbconn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s",dbconf.User,dbconf.Passwd,dbconf.Host,dbconf.Protocol,dbconf.Database,dbconf.Charset)
	var err error
	engine, err = xorm.NewEngine(dbconf.Type, dbconn)
	if err!=nil{
		fmt.Printf(err)
	}
	return engine
}

func GetAllowIpList() []string{

}
