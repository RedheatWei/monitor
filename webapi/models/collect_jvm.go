package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CollectJvm struct {
	Id                          int    `orm:"column(id);auto"`
	ServerId                    int64  `orm:"column(ServerId)"`
	LoadedClassCount            int    `orm:"column(LoadedClassCount)"`
	UnloadedClassCount          int    `orm:"column(UnloadedClassCount)"`
	TotalLoadedClassCount       int    `orm:"column(TotalLoadedClassCount)"`
	HeapMemoryUsageInit         int64  `orm:"column(HeapMemoryUsageInit)"`
	HeapMemoryUsageCommitted    int64  `orm:"column(HeapMemoryUsageCommitted)"`
	HeapMemoryUsageMax          int64  `orm:"column(HeapMemoryUsageMax)"`
	HeapMemoryUsageUsed         int64  `orm:"column(HeapMemoryUsageUsed)"`
	NonHeapMemoryUsageInit      int64  `orm:"column(NonHeapMemoryUsageInit)"`
	NonHeapMemoryUsageCommitted int64  `orm:"column(NonHeapMemoryUsageCommitted)"`
	NonHeapMemoryUsageMax       int64  `orm:"column(NonHeapMemoryUsageMax)"`
	NonHeapMemoryUsageUsed      int64  `orm:"column(NonHeapMemoryUsageUsed)"`
	ClassPath                   string `orm:"column(ClassPath);size(255)"`
	Uptime                      int64  `orm:"column(Uptime)"`
	TotalStartedThreadCount     int    `orm:"column(TotalStartedThreadCount)"`
	PeakThreadCount             int    `orm:"column(PeakThreadCount)"`
	CurrentThreadCpuTime        int64  `orm:"column(CurrentThreadCpuTime)"`
	ThreadCount                 int    `orm:"column(ThreadCount)"`
	DaemonThreadCount           int    `orm:"column(DaemonThreadCount)"`
	AgentId                     string `orm:"column(AgentId);size(255)"`
	TimeStamp                   int64  `orm:"column(TimeStamp)"`
	CreateTime                  int64  `orm:"column(CreateTime)"`
}

func (t *CollectJvm) TableName() string {
	return "collect_jvm"
}

func init() {
	orm.RegisterModel(new(CollectJvm))
}

// AddCollectJvm insert a new CollectJvm into database and returns
// last inserted Id on success.
func AddCollectJvm(m *CollectJvm) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCollectJvmById retrieves CollectJvm by Id. Returns error if
// Id doesn't exist
func GetCollectJvmById(id int) (v *CollectJvm, err error) {
	o := orm.NewOrm()
	v = &CollectJvm{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCollectJvm retrieves all CollectJvm matches certain condition. Returns empty list if
// no records exist
func GetAllCollectJvm(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CollectJvm))
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

	var l []CollectJvm
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

// UpdateCollectJvm updates CollectJvm by Id and returns error if
// the record to be updated doesn't exist
func UpdateCollectJvmById(m *CollectJvm) (err error) {
	o := orm.NewOrm()
	v := CollectJvm{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCollectJvm deletes CollectJvm by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCollectJvm(id int) (err error) {
	o := orm.NewOrm()
	v := CollectJvm{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CollectJvm{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
