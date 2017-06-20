package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CollectCpuInfostat struct {
	Id                 int     `orm:"column(id);auto"`
	ServerId           int64   `orm:"column(ServerId)"`
	InfoStatcpu        int     `orm:"column(InfoStatcpu)"`
	InfoStatvendorId   string  `orm:"column(InfoStatvendorId);size(255)"`
	InfoStatfamily     string  `orm:"column(InfoStatfamily);size(255)"`
	InfoStatmodel      string  `orm:"column(InfoStatmodel);size(255)"`
	InfoStatstepping   int     `orm:"column(InfoStatstepping)"`
	InfoStatphysicalId string  `orm:"column(InfoStatphysicalId);size(255)"`
	InfoStatcoreId     string  `orm:"column(InfoStatcoreId);size(255)"`
	InfoStatcores      int     `orm:"column(InfoStatcores)"`
	InfoStatmodelName  string  `orm:"column(InfoStatmodelName);size(255)"`
	InfoStatmhz        float64 `orm:"column(InfoStatmhz)"`
	InfoStatcacheSize  int     `orm:"column(InfoStatcacheSize)"`
	InfoStatmicrocode  string  `orm:"column(InfoStatmicrocode);size(255)"`
	TimeStamp          int64   `orm:"column(TimeStamp)"`
}

func (t *CollectCpuInfostat) TableName() string {
	return "collect_cpu_infostat"
}

func init() {
	orm.RegisterModel(new(CollectCpuInfostat))
}

// AddCollectCpuInfostat insert a new CollectCpuInfostat into database and returns
// last inserted Id on success.
func AddCollectCpuInfostat(m *CollectCpuInfostat) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCollectCpuInfostatById retrieves CollectCpuInfostat by Id. Returns error if
// Id doesn't exist
func GetCollectCpuInfostatById(id int) (v *CollectCpuInfostat, err error) {
	o := orm.NewOrm()
	v = &CollectCpuInfostat{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCollectCpuInfostat retrieves all CollectCpuInfostat matches certain condition. Returns empty list if
// no records exist
func GetAllCollectCpuInfostat(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CollectCpuInfostat))
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

	var l []CollectCpuInfostat
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

// UpdateCollectCpuInfostat updates CollectCpuInfostat by Id and returns error if
// the record to be updated doesn't exist
func UpdateCollectCpuInfostatById(m *CollectCpuInfostat) (err error) {
	o := orm.NewOrm()
	v := CollectCpuInfostat{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCollectCpuInfostat deletes CollectCpuInfostat by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCollectCpuInfostat(id int) (err error) {
	o := orm.NewOrm()
	v := CollectCpuInfostat{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CollectCpuInfostat{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
