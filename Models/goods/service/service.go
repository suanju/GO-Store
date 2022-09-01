package service

type Service struct {
	ID   uint   `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name" `
}

type ServicesList struct {
	ID        uint    `json:"id" gorm:"id"`
	GoodsId   uint    `json:"goods_id" gorm:"goods_id"`
	ServiceId uint    `json:"service_id" gorm:"service_id"`
	Service   Service `json:"service" gorm:"foreignKey:ServiceId"`
}

type ServicesLists []ServicesList

func (Service) TableName() string {
	return "ml_goods_service"
}

func (ServicesList) TableName() string {
	return "ml_goods_service_list"
}
