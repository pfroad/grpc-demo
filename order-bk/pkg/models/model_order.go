package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Order
type Order struct {
	
	Id   uint     `gorm:"column:id" form:"id" json:"id" comment:"" sql:"bigint(20) unsigned,PRI"`
	ParkNo   string     `gorm:"column:park_no" form:"park_no" json:"park_no" comment:"" sql:"char(10),MUL"`
	ParkName   string     `gorm:"column:park_name" form:"park_name" json:"park_name" comment:"" sql:"varchar(255)"`
	SpaceId   int     `gorm:"column:space_id" form:"space_id" json:"space_id" comment:"" sql:"bigint(20)"`
	SpaceNo   string     `gorm:"column:space_no" form:"space_no" json:"space_no" comment:"" sql:"varchar(32)"`
	UserId   int     `gorm:"column:user_id" form:"user_id" json:"user_id" comment:"" sql:"bigint(20)"`
	Mobile   string     `gorm:"column:mobile" form:"mobile" json:"mobile" comment:"" sql:"varchar(16)"`
	CarNo   string     `gorm:"column:car_no" form:"car_no" json:"car_no" comment:"" sql:"varchar(16)"`
	Status   int     `gorm:"column:status" form:"status" json:"status" comment:"" sql:"tinyint(3)"`
	Type   int     `gorm:"column:type" form:"type" json:"type" comment:"1-共享,2-临停,4-月保,8-日租" sql:"tinyint(3)"`
	PaymentType   int     `gorm:"column:payment_type" form:"payment_type" json:"payment_type" comment:"1-余额,2-金币,4-微信,8-支付宝,16-建行, other" sql:"int(10)"`
	StartTime   *time.Time     `gorm:"column:start_time" form:"start_time" json:"start_time,omitempty" comment:"" sql:"timestamp"`
	EndTime   *time.Time     `gorm:"column:end_time" form:"end_time" json:"end_time,omitempty" comment:"" sql:"timestamp"`
	TotalFee   int     `gorm:"column:total_fee" form:"total_fee" json:"total_fee" comment:"" sql:"int(10)"`
	DiscountFee   int     `gorm:"column:discount_fee" form:"discount_fee" json:"discount_fee" comment:"" sql:"int(10)"`
	OtherFee   int     `gorm:"column:other_fee" form:"other_fee" json:"other_fee" comment:"" sql:"int(10)"`
	IsDisabled   int     `gorm:"column:is_disabled" form:"is_disabled" json:"is_disabled" comment:"" sql:"tinyint(3)"`
	IsDeleted   int     `gorm:"column:is_deleted" form:"is_deleted" json:"is_deleted" comment:"" sql:"tinyint(3)"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
}
//TableName
func (m *Order) TableName() string {
	return "order"
}
//One
func (m *Order) One() (one *Order, err error) {
	one = &Order{}
	err = crudOne(m, one)
	return
}
//All
func (m *Order) All(q *PaginationQuery) (list *[]Order, total uint, err error) {
	list = &[]Order{}
	total, err = crudAll(m, q, list)
	return
}
//Update
func (m *Order) Update() (err error) {
	where := Order{Id: m.Id}
	m.Id = 0
	
	return crudUpdate(m, where)
}
//Create
func (m *Order) Create() (err error) {
	m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete
func (m *Order) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
