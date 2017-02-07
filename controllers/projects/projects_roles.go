package projects

import (
	"fmt"
	"opms/controllers"
	. "opms/models/projects"
	"opms/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//项目角色管理
type ManageProjectsRolesController struct {
	controllers.BaseController
}

func (this *ManageProjectsRolesController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "projectsroles-manage") {
		this.Abort("401")
	}
	page, err := this.GetInt("p")
	status := this.GetString("status")
	keywords := this.GetString("keywords")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)
	condArr["status"] = status
	condArr["keywords"] = keywords

	countProjectsRoles := CountProjectsRoles(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countProjectsRoles)
	_, _, projectsroles := ListProjectsRoles(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["projectsroles"] = projectsroles
	this.Data["countProjectsRoles"] = countProjectsRoles

	this.TplName = "projects/projectroles.tpl"
}

//部门状态
type AjaxStatusProjectsRolesController struct {
	controllers.BaseController
}

func (this *AjaxStatusProjectsRolesController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "projectsroles-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择项目角色"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 3 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeProjectsRolestatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目角色状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目角色状态更改失败"}
	}
	this.ServeJSON()
}

//项目角色添加
type AddProjectsRolesController struct {
	controllers.BaseController
}

func (this *AddProjectsRolesController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "projectsroles-add") {
		this.Abort("401")
	}
	this.TplName = "projects/projectroles-form.tpl"
}

func (this *AddProjectsRolesController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "projectsroles-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写名称"}
		this.ServeJSON()
		return
	}
	desc := this.GetString("desc")

	var pr ProjectsRoles
	pr.Id = utils.SnowFlakeId()
	pr.Name = name
	pr.Desc = desc
	err := AddProjectsRoles(pr)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目角色添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目角色添加失败"}
	}
	this.ServeJSON()
}

//部门编辑
type EditProjectsRolesController struct {
	controllers.BaseController
}

func (this *EditProjectsRolesController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "projectsroles-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	pr, err := GetProjectsRoles(int64(id))
	if err != nil {
		this.Abort("404")
	}
	this.Data["pr"] = pr
	this.TplName = "projects/projectroles-form.tpl"
}

func (this *EditProjectsRolesController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "projectsroles-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "参数出错"}
		this.ServeJSON()
		return
	}
	_, err := GetProjectsRoles(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目角色不存在"}
		this.ServeJSON()
		return
	}

	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写名称"}
		this.ServeJSON()
		return
	}
	desc := this.GetString("desc")

	var pr ProjectsRoles
	pr.Name = name
	pr.Desc = desc

	err = UpdateProjectsRoles(id, pr)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "信息修改失败"}
	}
	this.ServeJSON()
}
