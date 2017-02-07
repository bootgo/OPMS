package projects

import (
	"fmt"
	"opms/models"
	"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProjectsRoles struct {
	Id     int64 `orm:"pk;column(projectsrolesid);"`
	Name   string
	Desc   string
	Status int
}

func (this *ProjectsRoles) TableName() string {
	return models.TableName("projects_roles")
}
func init() {
	orm.RegisterModel(new(ProjectsRoles))
}

func GetProjectsRoles(id int64) (ProjectsRoles, error) {
	var projectsroles ProjectsRoles
	var err error
	o := orm.NewOrm()

	projectsroles = ProjectsRoles{Id: id}
	err = o.Read(&projectsroles)

	if err == orm.ErrNoRows {
		return projectsroles, nil
	}
	return projectsroles, err
}

func GetProjectsRolesName(id int64) string {
	var err error
	var name string
	err = utils.GetCache("GetProjectsRolesName.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var projectsroles ProjectsRoles
		o := orm.NewOrm()
		o.QueryTable(models.TableName("projects_roles")).Filter("projectsrolesid", id).One(&projectsroles, "name")
		name = projectsroles.Name
		utils.SetCache("GetProjectsRolesName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}

func UpdateProjectsRoles(id int64, updPr ProjectsRoles) error {
	var pr ProjectsRoles
	o := orm.NewOrm()
	pr = ProjectsRoles{Id: id}

	pr.Name = updPr.Name
	pr.Desc = updPr.Desc
	_, err := o.Update(&pr, "name", "desc")
	return err
}

func AddProjectsRoles(updPr ProjectsRoles) error {
	o := orm.NewOrm()
	o.Using("default")
	pr := new(ProjectsRoles)

	pr.Id = updPr.Id
	pr.Name = updPr.Name
	pr.Desc = updPr.Desc
	pr.Status = 1
	_, err := o.Insert(pr)

	return err
}

func ListProjectsRoles(condArr map[string]string, page int, offset int) (num int64, err error, dep []ProjectsRoles) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("projects_roles"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	var prs []ProjectsRoles
	num, err1 := qs.Limit(offset, start).All(&prs)
	return num, err1, prs
}

//统计数量
func CountProjectsRoles(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("projects_roles"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

//更改角色状态
func ChangeProjectsRolestatus(id int64, status int) error {
	o := orm.NewOrm()

	pr := ProjectsRoles{Id: id}
	err := o.Read(&pr, "projectsrolesid")
	if nil != err {
		return err
	} else {
		pr.Status = status
		_, err := o.Update(&pr)
		return err
	}
}
