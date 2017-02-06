package knowledges

import (
	"fmt"
	"opms/models"
	"opms/utils"
	//"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type KnowledgesSort struct {
	Id     int64 `orm:"pk;column(sortid);"`
	Name   string
	Desc   string
	Status int
}

func (this *KnowledgesSort) TableName() string {
	return models.TableName("knowledges_sort")
}

func init() {
	orm.RegisterModel(new(KnowledgesSort))
}


func GetKnowledgesSort(id int64) (KnowledgesSort, error) {
	var ks KnowledgesSort
	var err error
	o := orm.NewOrm()

	ks = KnowledgesSort{Id: id}
	err = o.Read(&ks)

	if err == orm.ErrNoRows {
		return ks, nil
	}
	return ks, err
}

func GetKnowledgesSortName(id int64) string {
	var err error
	var name string
	err = utils.GetCache("GetKnowledgesSortName.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var ks KnowledgesSort
		o := orm.NewOrm()
		o.QueryTable(models.TableName("knowledges_sort")).Filter("sortid", id).One(&ks, "name")
		name = ks.Name
		utils.SetCache("GetKnowledgesSortName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}

func UpdateKnowledgesSort(id int64, updKs KnowledgesSort) error {
	var uks KnowledgesSort
	o := orm.NewOrm()
	uks = KnowledgesSort{Id: id}

	uks.Name = updKs.Name
	uks.Desc = updKs.Desc
	_, err := o.Update(&uks, "name", "desc")
	return err
}

func AddKnowledgesSort(updKs KnowledgesSort) error {
	o := orm.NewOrm()
	o.Using("default")
	uks := new(KnowledgesSort)

	uks.Id = updKs.Id
	uks.Name = updKs.Name
	uks.Desc = updKs.Desc
	uks.Status = 1
	_, err := o.Insert(uks)

	return err
}


func ListKnowledgeSort() (num int64, err error, ops []KnowledgesSort) {
	var sort []KnowledgesSort
	var errs error

	page := 1
	offset := 100

	start := (page - 1) * offset

	errs = utils.GetCache("ListKnowledgeSort", &sort)
	if errs != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		o := orm.NewOrm()
		o.Using("default")
		qs := o.QueryTable(models.TableName("knowledges_sort"))
		cond := orm.NewCondition()
		cond = cond.And("status", 1)
		qs = qs.SetCond(cond)

		qs.Limit(offset, start).All(&sort)
		utils.SetCache("ListKnowledgeSort", sort, cache_expire)
	}
	return num, errs, sort
}

func ListKnowledgesSort(condArr map[string]string, page int, offset int) (num int64, err error, ks []KnowledgesSort) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("knowledges_sort"))
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

	var kss []KnowledgesSort
	num, err1 := qs.Limit(offset, start).All(&kss)
	return num, err1, kss
}

//统计数量
func CountKnowledgesSort(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("knowledges_sort"))
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

//更改分类状态
func ChangeKnowledgesSortStatus(id int64, status int) error {
	o := orm.NewOrm()

	ks := KnowledgesSort{Id: id}
	err := o.Read(&ks, "sortid")
	if nil != err {
		return err
	} else {
		ks.Status = status
		_, err := o.Update(&ks)
		return err
	}
}
//删除
func DeleteKnowledgesSort(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&KnowledgesSort{Id: id})
	return err
}