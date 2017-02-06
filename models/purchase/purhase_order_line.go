package purchase

import (
	"pms/models/base"
	"pms/models/partner"
	"pms/models/product"
)

type PurchaseOrderLine struct {
	base.Base
	Name                string                  `orm:"unique"`  //采购订单明细号
	Product             *product.ProductProduct `orm:"rel(fk)"` //产品
	ProductName         string                  //产品名称
	SupplierProductName string                  //供应商产品名称
	Supplier            *partner.Partner        `orm:"rel(fk)"` //供应商
}
