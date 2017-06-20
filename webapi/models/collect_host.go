package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CollectHost struct {
	Id                           int    `orm:"column(id);auto"`
	Serverid                     int64  `orm:"column(Serverid)"`
	InfoStathostname             string `orm:"column(InfoStathostname);size(255)"`
	InfoStatuptime               int64  `orm:"column(InfoStatuptime)"`
	InfoStatbootTime             int64  `orm:"column(InfoStatbootTime)"`
	InfoStatprocs                int64  `orm:"column(InfoStatprocs)"`
	InfoStatos                   string `orm:"column(InfoStatos);size(255)"`
	InfoStatplatform             string `orm:"column(InfoStatplatform);size(255)"`
	InfoStatplatformFamily       string `orm:"column(InfoStatplatformFamily);size(255)"`
	InfoStatplatformVersion      string `orm:"column(InfoStatplatformVersion);size(255)"`
	InfoStatkernelVersion        string `orm:"column(InfoStatkernelVersion);size(255)"`
	InfoStatvirtualizationSystem string `orm:"column(InfoStatvirtualizationSystem);size(255)"`
	InfoStatvirtualizationRole   string `orm:"column(InfoStatvirtualizationRole);size(255)"`
	InfoStathostid               string `orm:"column(InfoStathostid);size(255)"`
	TimeStamp                    int64  `orm:"column(TimeStamp)"`
	CreateTime                   int64  `orm:"column(CreateTime)"`
}

func (t *CollectHost) TableName() string {
	return "collect_host"
}

func init() {
	orm.RegisterModel(new(CollectHost))
}

// AddCollectHost insert a new CollectHost into database and returns
// last inserted Id on success.
func AddCollectHost(m *CollectHost) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCollectHostById retrieves CollectHost by Id. Returns error if
// Id doesn't exist
func GetCollectHostById(id int) (v *CollectHost, err error) {
	o := orm.NewOrm()
	v = &CollectHost{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCollectHost retrieves all CollectHost matches certain condition. Returns empty list if
// no records exist
func GetAllCollectHost(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CollectHost))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CollectHost
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCollectHost updates CollectHost by Id and returns error if
// the record to be updated doesn't exist
func UpdateCollectHostById(m *CollectHost) (err error) {
	o := orm.NewOrm()
	v := CollectHost{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCollectHost deletes CollectHost by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCollectHost(id int) (err error) {
	o := orm.NewOrm()
	v := CollectHost{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CollectHost{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
