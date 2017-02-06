package venders

import (
	//"fmt"
	"opms/models"
	//"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Venders struct {
	Id         int64 `orm:"pk;column(venderid);"`
	//Realname   string
	//Sex        int
	//Birth      int64
	//Edu        int
	//Work       int
	Name		string
	Contact		string
	Email       string
	Wechat     string
	Qq          string
	Phone       string
	Tel         string
	Address     string
	Note        string
	//Remarks		string
	Attachment string
	Created    int64
	Status     int

}

func (this *Venders) TableName() string {
	return models.TableName("venders")
}
func init() {
	orm.RegisterModel(new(Venders))
}

func GetVenders(id int64) (Venders, error) {
	var vender Venders
	var err error
	o := orm.NewOrm()

	vender = Venders{Id: id}
	err = o.Read(&vender)

	if err == orm.ErrNoRows {
		return vender, nil
	}
	return vender, err
}

func UpdateVenders(id int64, upd Venders) error {
	o := orm.NewOrm()
	vds := Venders{Id: id}
	vds.Name = upd.Name
	vds.Contact = upd.Contact
	vds.Email = upd.Email
	vds.Wechat = upd.Wechat
	vds.Qq = upd.Qq
	vds.Attachment = upd.Attachment
	vds.Status = upd.Status
	vds.Note = upd.Note
	vds.Phone = upd.Phone
	vds.Tel = upd.Tel
	vds.Address = upd.Address
	//vds.Remarks = upd.Remarks

	if upd.Attachment != "" {
		vds.Attachment = upd.Attachment
		_, err := o.Update(&vds, "name", "contact", "email", "wechat", "qq", "note", "phone", "tel", "address","status","attachment")
		return err
	} else {
		_, err := o.Update(&vds, "name", "contact", "email", "wechat", "qq", "note", "phone", "tel", "address","status")
		return err
	}
}

func AddVenders(upd Venders) error {
	o := orm.NewOrm()
	o.Using("default")
	vds := new(Venders)

	vds.Id = upd.Id
	vds.Name = upd.Name
	vds.Contact = upd.Contact
	vds.Email = upd.Email
	vds.Wechat = upd.Wechat
	vds.Qq = upd.Qq
	vds.Tel = upd.Tel
	vds.Address = upd.Address
	//vds.Remarks = upd.Remarks
	vds.Attachment = upd.Attachment
	vds.Created = time.Now().Unix()
	vds.Status = upd.Status
	vds.Note = upd.Note
	vds.Phone = upd.Phone
	_, err := o.Insert(vds)
	return err
}

func ListVenders(condArr map[string]string, page int, offset int) (num int64, err error, vds []Venders) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("venders"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("realname__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	//if condArr["sex"] != "" {
	//	cond = cond.And("sex", condArr["sex"])
	//}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	var vdses []Venders
	qs = qs.OrderBy("-venderid")
	num, err1 := qs.Limit(offset, start).All(&vdses)
	return num, err1, vdses
}

//统计数量
func CountVenders(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("venders"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("realname__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

func ChangeVenderStatus(id int64, status int) error {
	o := orm.NewOrm()

	res := Venders{Id: id}
	err := o.Read(&res, "venderid")
	if nil != err {
		return err
	} else {
		res.Status = status
		_, err := o.Update(&res)
		return err
	}
}

func DeleteVender(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Venders{Id: id})
	return err
}
