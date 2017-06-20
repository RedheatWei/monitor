package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CollectMem struct {
	Id                            int     `orm:"column(id);auto"`
	ServerId                      int64   `orm:"column(ServerId)"`
	VirtualMemoryStattotal        int64   `orm:"column(VirtualMemoryStattotal)"`
	VirtualMemoryStatavailable    int64   `orm:"column(VirtualMemoryStatavailable)"`
	VirtualMemoryStatused         int64   `orm:"column(VirtualMemoryStatused)"`
	VirtualMemoryStatusedPercent  float64 `orm:"column(VirtualMemoryStatusedPercent)"`
	VirtualMemoryStatfree         int64   `orm:"column(VirtualMemoryStatfree)"`
	VirtualMemoryStatactive       int64   `orm:"column(VirtualMemoryStatactive)"`
	VirtualMemoryStatinactive     int64   `orm:"column(VirtualMemoryStatinactive)"`
	VirtualMemoryStatwired        int64   `orm:"column(VirtualMemoryStatwired)"`
	VirtualMemoryStatbuffers      int64   `orm:"column(VirtualMemoryStatbuffers)"`
	VirtualMemoryStatcached       int64   `orm:"column(VirtualMemoryStatcached)"`
	VirtualMemoryStatwriteback    int64   `orm:"column(VirtualMemoryStatwriteback)"`
	VirtualMemoryStatdirty        int64   `orm:"column(VirtualMemoryStatdirty)"`
	VirtualMemoryStatwritebacktmp int64   `orm:"column(VirtualMemoryStatwritebacktmp)"`
	VirtualMemoryStatshared       int64   `orm:"column(VirtualMemoryStatshared)"`
	VirtualMemoryStatslab         int64   `orm:"column(VirtualMemoryStatslab)"`
	VirtualMemoryStatpagetables   int64   `orm:"column(VirtualMemoryStatpagetables)"`
	VirtualMemoryStatswapcached   int64   `orm:"column(VirtualMemoryStatswapcached)"`
	SwapMemoryStattotal           int64   `orm:"column(SwapMemoryStattotal)"`
	SwapMemoryStatused            int64   `orm:"column(SwapMemoryStatused)"`
	SwapMemoryStatfree            int64   `orm:"column(SwapMemoryStatfree)"`
	SwapMemoryStatusedPercent     float64 `orm:"column(SwapMemoryStatusedPercent)"`
	SwapMemoryStatsin             int64   `orm:"column(SwapMemoryStatsin)"`
	SwapMemoryStatsout            int64   `orm:"column(SwapMemoryStatsout)"`
	TimeStamp                     int64   `orm:"column(TimeStamp)"`
	CreateTime                    int64   `orm:"column(CreateTime)"`
}

func (t *CollectMem) TableName() string {
	return "collect_mem"
}

func init() {
	orm.RegisterModel(new(CollectMem))
}

// AddCollectMem insert a new CollectMem into database and returns
// last inserted Id on success.
func AddCollectMem(m *CollectMem) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCollectMemById retrieves CollectMem by Id. Returns error if
// Id doesn't exist
func GetCollectMemById(id int) (v *CollectMem, err error) {
	o := orm.NewOrm()
	v = &CollectMem{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCollectMem retrieves all CollectMem matches certain condition. Returns empty list if
// no records exist
func GetAllCollectMem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CollectMem))
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

	var l []CollectMem
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

// UpdateCollectMem updates CollectMem by Id and returns error if
// the record to be updated doesn't exist
func UpdateCollectMemById(m *CollectMem) (err error) {
	o := orm.NewOrm()
	v := CollectMem{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCollectMem deletes CollectMem by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCollectMem(id int) (err error) {
	o := orm.NewOrm()
	v := CollectMem{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CollectMem{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
