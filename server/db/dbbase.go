package db
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"monitor/network"
	"fmt"
)
var engine *xorm.Engine
func ConnDB() {
	dbconf := network.ServerConfig.DB
	dbconn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s",dbconf.User,dbconf.Passwd,dbconf.Host,dbconf.Protocol,dbconf.Database,dbconf.Charset)
	var err error
	engine, err = xorm.NewEngine(dbconf.Type, dbconn)
}

func GetAllowIpList() []string{

}
