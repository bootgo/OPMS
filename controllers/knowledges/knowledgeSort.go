package knowledges

import (
	"fmt"
	"opms/controllers"
	. "opms/models/knowledges"
	"opms/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)


//文章分类管理
type ManageKnowledgesSortController struct {
	controllers.BaseController
}

func (this *ManageKnowledgesSortController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "knowledgessort-manage") {
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

	countKnowledgesSort := CountKnowledgesSort(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countKnowledgesSort)
	_, _, ks := ListKnowledgesSort(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["ks"] = ks
	this.Data["countKnowledgesSort"] = countKnowledgesSort

	this.TplName = "knowledges/knowledgessort.tpl"
}

//分类状态
type AjaxStatusKnowledgeSortController struct {
	controllers.BaseController
}

func (this *AjaxStatusKnowledgeSortController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "knowledgessort-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择分类"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 3 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeKnowledgesSortStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "分类状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "分类状态更改失败"}
	}
	this.ServeJSON()
}

//分类添加
type AddKnowledgeSortController struct {
	controllers.BaseController
}

func (this *AddKnowledgeSortController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "knowledgessort-add") {
		this.Abort("401")
	}
	this.TplName = "knowledges/knowledgessort-form.tpl"
}

func (this *AddKnowledgeSortController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "knowledgessort-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写分类名称"}
		this.ServeJSON()
		return
	}
	desc := this.GetString("desc")

	var ks KnowledgesSort
	ks.Id = utils.SnowFlakeId()
	ks.Name = name
	ks.Desc = desc
	err := AddKnowledgesSort(ks)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "分类添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "分类添加失败"}
	}
	this.ServeJSON()
}

//分类编辑
type EditKnowledgesSortController struct {
	controllers.BaseController
}

func (this *EditKnowledgesSortController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "knowledgessort-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idstr)
	ks, err := GetKnowledgesSort(int64(id))
	if err != nil {
		this.Abort("404")
	}
	this.Data["ks"] = ks
	this.TplName = "knowledges/knowledgessort-form.tpl"
}

func (this *EditKnowledgesSortController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "knowledgessort-edit") {
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
	_, err := GetKnowledgesSort(id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "分类不存在"}
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

	var ks KnowledgesSort
	ks.Name = name
	ks.Desc = desc

	err = UpdateKnowledgesSort(id, ks)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "信息修改失败"}
	}
	this.ServeJSON()
}
//分类删除
type AjaxDeleteKnowledgesSortController struct {
	controllers.BaseController
}

func (this *AjaxDeleteKnowledgesSortController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "knowledgessort-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权删除"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择分类"}
		this.ServeJSON()
		return
	}

	err := DeleteKnowledgesSort(id)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
