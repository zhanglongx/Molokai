package core

import (
	"gorm.io/gorm"
)

type HoldingStore interface {
	// Create creates a new holding and persists to database.
	// Date is in ISO8601 format and with location TZ Asia/Shanghai
	Create(Symbol Symbol, Date string, Cost float64, Reminders string) (*Holding, error)

	// DeleteBySymbol deletes record by Symbol
	DeleteBySymbol(Symbol Symbol) error

	// List returns all of records
	List() ([]*Holding, error)

	// Save persists all values
	// The Updated fields is set to current time automatically
	Save(h *Holding) error
}

type Holding struct {
	gorm.Model

	Symbol Symbol `gorm:"INDEX"`
	Date   string
	Cost   float64

	Reminders string
}

var Holdings HoldingStore

type holdings struct {
	*gorm.DB
}

func (db *holdings) Create(Symbol Symbol, Date string,
	Cost float64, Reminders string) (*Holding, error) {
	Holding := &Holding{
		Symbol:    Symbol,
		Date:      Date,
		Cost:      Cost,
		Reminders: Reminders,
	}

	return Holding, db.DB.Create(Holding).Error
}

func (db *holdings) DeleteBySymbol(Symbol Symbol) error {
	return db.Where("symbol = ?", Symbol).Delete(new(Holding)).Error
}

func (db *holdings) List() ([]*Holding, error) {
	var holdings []*Holding
	return holdings, db.Where("TRUE").Find(&holdings).Error
}

func (db *holdings) Save(h *Holding) error {
	return db.DB.Save(h).Error
}
