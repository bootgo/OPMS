package projects

import (
	"fmt"
	"opms/models"
	"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProjectsCustomers struct {
	Id        int64 `orm:"pk;"`
	Projectid int64
	Customerid    int64
	Created   int64
}

func (this *ProjectsCustomers) TableName() string {
	return models.TableName("projects_customers")
}
func init() {
	orm.RegisterModel(new(ProjectsCustomers))
}

func AddProCustomers(upd ProjectsCustomers) error {
	o := orm.NewOrm()
	pcs := new(ProjectsCustomers)

	pcs.Id = upd.Id
	pcs.Projectid = upd.Projectid
	pcs.Customerid = upd.Customerid
	pcs.Created = time.Now().Unix()
	_, err := o.Insert(pcs)
	return err
}

func ListProjectCustomers(projectId int64, page int, offset int) (num int64, err error, ops []ProjectsCustomers) {
	var pcs []ProjectsCustomers
	var errs error

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	errs = utils.GetCache("ListProjectCustomers.id."+fmt.Sprintf("%d", projectId), &pcs)
	if errs != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		o := orm.NewOrm()
		o.Using("default")
		qs := o.QueryTable(models.TableName("projects_customers"))
		cond := orm.NewCondition()
		if projectId > 0 {
			cond = cond.And("projectid", projectId)
		}
		qs = qs.SetCond(cond)

		qs.Limit(offset, start).All(&pcs)
		utils.SetCache("ListProjectCustomers.id."+fmt.Sprintf("%d", projectId), pcs, cache_expire)
	}
	return num, errs, pcs
}

func DeleteProjectCustomers(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&ProjectsCustomers{Id: id})
	return err
}

func GetProjectCustomers(customerid, projectid int64) (ProjectsCustomers, error) {
	var pcs ProjectsCustomers
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("projects_customers"))
	err := qs.Filter("customerid", customerid).Filter("projectid", projectid).One(&pcs)
	if err == orm.ErrNoRows {
		return pcs, nil
	}
	return pcs, err
}
