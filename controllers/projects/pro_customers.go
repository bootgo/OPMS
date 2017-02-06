package projects

import (
	"fmt"
	"opms/controllers"
	. "opms/models/projects"
	. "opms/models/customers"
	"opms/utils"
	"strconv"
	"strings"
)

//项目成员
type CustomerProjectController struct {
	controllers.BaseController
}

func (this *CustomerProjectController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "project-customer") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	idlong := int64(id)
	project, err := GetProject(idlong)
	if err != nil {
		this.Abort("404")
	}
	this.Data["project"] = project

	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}
	offset := 500
	_, _, customers := ListProjectCustomers(idlong, page, offset)
	this.Data["customers"] = customers
	this.Data["countCustomer"] = len(customers)
	this.TplName = "projects/customer.tpl"
}

type AjaxDeleteCustomerProjectController struct {
	controllers.BaseController
}

func (this *AjaxDeleteCustomerProjectController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "project-customer-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择项目客户"}
		this.ServeJSON()
		return
	}

	err := DeleteProjectCustomers(id)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目客户删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目客户删除失败"}
	}
	this.ServeJSON()
}

type AddCustomerProjectController struct {
	controllers.BaseController
}

func (this *AddCustomerProjectController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "project-customer-add") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	project, err := GetProject(int64(id))
	if err != nil {
		this.Abort("404")
	}
	this.Data["project"] = project
	this.TplName = "projects/customer-form.tpl"
}

func (this *AddCustomerProjectController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "project-customer-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	projectid, _ := this.GetInt64("projectid")
	if projectid <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择项目"}
		this.ServeJSON()
		return
	}
	customerid, _ := this.GetInt64("customerid")
	if customerid <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写客户"}
		this.ServeJSON()
		return
	}
	realname := GetCustomerName(customerid)
	if "" == realname {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "客户不存在"}
		this.ServeJSON()
		return
	}

	checkcustomer, _ := GetProjectCustomers(customerid, projectid)
	if checkcustomer.Customerid > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "客户已经存在"}
		this.ServeJSON()
		return
	}

	var err error
	//雪花算法ID生成
	id := utils.SnowFlakeId()

	var customer ProjectsCustomers
	customer.Id = id
	customer.Customerid = customerid
	customer.Projectid = projectid

	err = AddProCustomers(customer)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目客户添加成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目客户添加失败"}
	}
	this.ServeJSON()
}
