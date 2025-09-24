package drmodel

import "time"

type SaleReport struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement;"`
	Date      string    `json:"date" gorm:"type:char(10);comment:'Y-m-d'"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (SaleReport) TableName() string {
	return "dr_sale_reports"
}
