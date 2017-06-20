package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CollectNetIocountersstat struct {
	Id                        int    `orm:"column(id);auto"`
	ServerId                  int64  `orm:"column(ServerId)"`
	IOCountersStatname        string `orm:"column(IOCountersStatname);size(255)"`
	IOCountersStatbytesSent   int64  `orm:"column(IOCountersStatbytesSent)"`
	IOCountersStatbytesRecv   int64  `orm:"column(IOCountersStatbytesRecv)"`
	IOCountersStatpacketsSent int64  `orm:"column(IOCountersStatpacketsSent)"`
	IOCountersStatpacketsRecv int64  `orm:"column(IOCountersStatpacketsRecv)"`
	IOCountersStaterrin       int64  `orm:"column(IOCountersStaterrin)"`
	IOCountersStaterrout      int64  `orm:"column(IOCountersStaterrout)"`
	IOCountersStatdropin      int64  `orm:"column(IOCountersStatdropin)"`
	IOCountersStatdropout     int64  `orm:"column(IOCountersStatdropout)"`
	IOCountersStatfifoin      int64  `orm:"column(IOCountersStatfifoin)"`
	IOCountersStatfifoout     int64  `orm:"column(IOCountersStatfifoout)"`
	TimeStamp                 int64  `orm:"column(TimeStamp)"`
}

func (t *CollectNetIocountersstat) TableName() string {
	return "collect_net_iocountersstat"
}

func init() {
	orm.RegisterModel(new(CollectNetIocountersstat))
}

// AddCollectNetIocountersstat insert a new CollectNetIocountersstat into database and returns
// last inserted Id on success.
func AddCollectNetIocountersstat(m *CollectNetIocountersstat) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCollectNetIocountersstatById retrieves CollectNetIocountersstat by Id. Returns error if
// Id doesn't exist
func GetCollectNetIocountersstatById(id int) (v *CollectNetIocountersstat, err error) {
	o := orm.NewOrm()
	v = &CollectNetIocountersstat{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCollectNetIocountersstat retrieves all CollectNetIocountersstat matches certain condition. Returns empty list if
// no records exist
func GetAllCollectNetIocountersstat(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CollectNetIocountersstat))
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

	var l []CollectNetIocountersstat
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

// UpdateCollectNetIocountersstat updates CollectNetIocountersstat by Id and returns error if
// the record to be updated doesn't exist
func UpdateCollectNetIocountersstatById(m *CollectNetIocountersstat) (err error) {
	o := orm.NewOrm()
	v := CollectNetIocountersstat{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCollectNetIocountersstat deletes CollectNetIocountersstat by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCollectNetIocountersstat(id int) (err error) {
	o := orm.NewOrm()
	v := CollectNetIocountersstat{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CollectNetIocountersstat{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
