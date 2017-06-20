package controllers
import (
	"github.com/astaxie/beego"
	"monitor/webapplication/models"
	"fmt"
)

type JvmgetController struct {
	beego.Controller
}

func (this *JvmgetController) Get(){
	db := models.ConnDB()
	fmt.Println(db)
}