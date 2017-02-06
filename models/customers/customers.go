package customers

import (
	"fmt"
	"opms/models"
	"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Customers struct {
	Id         int64 `orm:"pk;column(customerid);"`
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

func (this *Customers) TableName() string {
	return models.TableName("customers")
}
func init() {
	orm.RegisterModel(new(Customers))
}

func GetCustomers(id int64) (Customers, error) {
	var customer Customers
	var err error
	o := orm.NewOrm()

	customer = Customers{Id: id}
	err = o.Read(&customer)

	if err == orm.ErrNoRows {
		return customer, nil
	}
	return customer, err
}

func UpdateCustomers(id int64, upd Customers) error {
	o := orm.NewOrm()
	res := Customers{Id: id}
	res.Name = upd.Name
	res.Contact = upd.Contact
	res.Email = upd.Email
	res.Wechat = upd.Wechat
	res.Qq = upd.Qq
	res.Attachment = upd.Attachment
	res.Status = upd.Status
	res.Note = upd.Note
	res.Phone = upd.Phone
	res.Tel = upd.Tel
	res.Address = upd.Address
	//res.Remarks = upd.Remarks

	if upd.Attachment != "" {
		res.Attachment = upd.Attachment
		_, err := o.Update(&res, "name", "contact", "email", "wechat", "qq", "note", "phone", "tel", "address","status","attachment")
		return err
	} else {
		_, err := o.Update(&res, "name", "contact", "email", "wechat", "qq", "note", "phone", "tel", "address","status")
		return err
	}
}

func AddCustomers(upd Customers) error {
	o := orm.NewOrm()
	o.Using("default")
	res := new(Customers)

	res.Id = upd.Id
	res.Name = upd.Name
	res.Contact = upd.Contact
	res.Email = upd.Email
	res.Wechat = upd.Wechat
	res.Qq = upd.Qq
	res.Tel = upd.Tel
	res.Address = upd.Address
	//res.Remarks = upd.Remarks
	res.Attachment = upd.Attachment
	res.Created = time.Now().Unix()
	res.Status = upd.Status
	res.Note = upd.Note
	res.Phone = upd.Phone
	_, err := o.Insert(res)
	return err
}

func ListCustomers(condArr map[string]string, page int, offset int) (num int64, err error, res []Customers) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("customers"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("realname__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	if condArr["sex"] != "" {
		cond = cond.And("sex", condArr["sex"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	var reses []Customers
	qs = qs.OrderBy("-customerid")
	num, err1 := qs.Limit(offset, start).All(&reses)
	return num, err1, reses
}

//统计数量
func CountCustomers(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("customers"))
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

func ChangeCustomerStatus(id int64, status int) error {
	o := orm.NewOrm()

	res := Customers{Id: id}
	err := o.Read(&res, "customerid")
	if nil != err {
		return err
	} else {
		res.Status = status
		_, err := o.Update(&res)
		return err
	}
}

func DeleteCustomer(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Customers{Id: id})
	return err
}

func GetCustomerName(id int64) string {
	var err error
	var name string

	err = utils.GetCache("GetCustomerName.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var customer Customers
		o := orm.NewOrm()
		o.QueryTable(models.TableName("customers")).Filter("customerid", id).One(&customer, "name")
		name = customer.Name
		utils.SetCache("GetCustomerName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}
