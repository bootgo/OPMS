package  venders

import (
	"fmt"
	"opms/controllers"
	. "opms/models/venders"
	"opms/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//管理供应商
type ManageVenderController struct {
	controllers.BaseController
}

func (this *ManageVenderController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "vender-manage") {
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

	countVender := CountVenders(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countVender)
	_, _, venders := ListVenders(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["venders"] = venders
	this.Data["countVender"] = countVender

	this.TplName = "venders/index.tpl"
}

type AjaxStatusVenderController struct {
	controllers.BaseController
}

func (this *AjaxStatusVenderController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "vender-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择供应商"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status > 5 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeVenderStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "状态更改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteVenderController struct {
	controllers.BaseController
}

func (this *AjaxDeleteVenderController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "vender-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择供应商"}
		this.ServeJSON()
		return
	}

	err := DeleteVender(id)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

//供应商添加
type AddVenderController struct {
	controllers.BaseController
}

func (this *AddVenderController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "vender-add") {
		this.Abort("401")
	}
	var vender Venders
	//vender.Sex = 1
	vender.Status = 1
	this.Data["vender"] = vender
	this.TplName = "venders/form.tpl"
}

func (this *AddVenderController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "vender-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写供应商名称"}
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

	var res Venders
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
	err = AddVenders(res)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	}
	this.ServeJSON()
}

//供应商编辑
type EditVenderController struct {
	controllers.BaseController
}

func (this *EditVenderController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "vender-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	vender, err := GetVenders(int64(id))
	if err != nil {
		this.Abort("404")
	}
	this.Data["vender"] = vender
	this.TplName = "venders/form.tpl"
}

func (this *EditVenderController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "vender-edit") {
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
	_, err := GetVenders(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "供应商不存在"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写供应商名称"}
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

	var res Venders
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
	err = UpdateVenders(id, res)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "信息修改失败"}
	}
	this.ServeJSON()
}
