package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type LaporanOmset struct {
	ID            uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Merchant_Name string    `gorm:"size:45;not null;unique" json:"name"`
	Omzet         float64   `gorm:"not null;" json:"omzet"`
	Created_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy     uint32    `gorm:"not null" json:"created_by"`
}

type Merchants struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint32    `gorm:"not null" json:"userId"`
	MerchantName string    `gorm:"size:45;not null;unique" json:"merchantName"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy    uint32    `gorm:"not null" json:"created_by"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy    uint32    `gorm:"not null" json:"updated_by"`
}
type Transaction struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	MerchantID uint32    `gorm:"not null" json:"merchantId"`
	OutletID   uint32    `gorm:"not null" json:"outletId"`
	BillTotal  float64   `gorm:"not null;" json:"omzet"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy  uint32    `gorm:"not null" json:"created_by"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy  uint32    `gorm:"not null" json:"updated_by"`
}

func (p *Transaction) FindLaporanByID(db *gorm.DB, pid uint64) (*Transaction, error) {
	var err error
	trx := []Transaction{}
	err = db.Debug().Model(&Transaction{}).Where("id = ?", pid).Take(&trx).Error
	if err != nil {
		return &Transaction{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.CreatedBy).Take(&p.MerchantID).Error
		if err != nil {
			return &Transaction{}, err
		}
	}
	return p, nil

}
