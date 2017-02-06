package  customers

import (
	"fmt"
	"opms/controllers"
	. "opms/models/customers"
	"opms/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//管理客户
type ManageCustomerController struct {
	controllers.BaseController
}

func (this *ManageCustomerController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-manage") {
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

	countCustomer := CountCustomers(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countCustomer)
	_, _, customers := ListCustomers(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["customers"] = customers
	this.Data["countCustomer"] = countCustomer

	this.TplName = "customers/index.tpl"
}

type AjaxStatusCustomerController struct {
	controllers.BaseController
}

func (this *AjaxStatusCustomerController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择客户"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status > 5 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeCustomerStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "状态更改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteCustomerController struct {
	controllers.BaseController
}

func (this *AjaxDeleteCustomerController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择客户"}
		this.ServeJSON()
		return
	}

	err := DeleteCustomer(id)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

//客户添加
type AddCustomerController struct {
	controllers.BaseController
}

func (this *AddCustomerController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-add") {
		this.Abort("401")
	}
	var customer Customers
	//customer.Sex = 1
	customer.Status = 1
	this.Data["customer"] = customer
	this.TplName = "customers/form.tpl"
}

func (this *AddCustomerController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写客户名称"}
		this.ServeJSON()
		return
	}
	
	contact := this.GetString("contact")
	email := this.GetString("email")
	wechat := this.GetString("wechat")
	qq := this.GetString("qq")
	tel := this.GetString("tel")
	address := this.GetString("address")
	//remarks := this.GetString("remarks")
	phone := this.GetString("phone")
	note := this.GetString("note")
	status, _ := this.GetInt("status")

	var filepath string
	f, h, err := this.GetFile("attachment")

	if err == nil {
		defer f.Close()
		now := time.Now()
		dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
		err1 := os.MkdirAll(dir, 0755)
		if err1 != nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "目录权限不够"}
			this.ServeJSON()
			return
		}
		filename := h.Filename
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		} else {
			this.SaveToFile("attachment", dir+"/"+filename)
			filepath = strings.Replace(dir, ".", "", 1) + "/" + filename
		}
	}

	var res Customers
	res.Id = utils.SnowFlakeId()
	res.Name = name
	res.Contact = contact
	res.Email = email
	res.Wechat = wechat
	res.Qq = qq
	res.Tel = tel
	res.Address = address
	//res.Remarks = remarks
	res.Phone = phone
	res.Note = note
	res.Status = status
	res.Attachment = filepath
	err = AddCustomers(res)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	}
	this.ServeJSON()
}

//客户编辑
type EditCustomerController struct {
	controllers.BaseController
}

func (this *EditCustomerController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	customer, err := GetCustomers(int64(id))
	if err != nil {
		this.Abort("404")
	}
	this.Data["customer"] = customer
	this.TplName = "customers/form.tpl"
}

func (this *EditCustomerController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-edit") {
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
	_, err := GetCustomers(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "客户不存在"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写客户名称"}
		this.ServeJSON()
		return
	}
	contact := this.GetString("contact")
	email := this.GetString("email")
	wechat := this.GetString("wechat")
	qq := this.GetString("qq")
	tel := this.GetString("tel")
	address := this.GetString("address")
	//remarks := this.GetString("remarks")
	phone := this.GetString("phone")
	note := this.GetString("note")
	status, _ := this.GetInt("status")
	
	
	
	
	var filepath string
	f, h, err := this.GetFile("attachment")

	if err == nil {
		defer f.Close()
		now := time.Now()
		dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
		err1 := os.MkdirAll(dir, 0755)
		if err1 != nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "目录权限不够"}
			this.ServeJSON()
			return
		}
		filename := h.Filename
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
			this.ServeJSON()
			return
		} else {
			this.SaveToFile("attachment", dir+"/"+filename)
			filepath = strings.Replace(dir, ".", "", 1) + "/" + filename
		}
	}

	var res Customers
	res.Name = name
	res.Contact = contact
	res.Email = email
	res.Wechat = wechat
	res.Qq = qq
	res.Tel = tel
	res.Address = address
	//res.Remarks = remarks
	res.Phone = phone
	res.Note = note
	res.Status = status
	res.Attachment = filepath
	err = UpdateCustomers(id, res)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "信息修改失败"}
	}
	this.ServeJSON()
}
