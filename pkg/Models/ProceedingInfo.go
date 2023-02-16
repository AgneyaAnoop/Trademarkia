package Models

import "gorm.io/gorm"

type ProceedingInfo struct {
	ID              uint              `json:"id" gorm:"primaryKey; autoIncrement; "`
	ProceedingEntry []ProceedingEntry `json:"proceeding-entry" gorm:"foreignKey:ProceedingInfoID"`
}

type ProceedingEntry struct {
	ID                    uint               `json:"id" gorm:"primaryKey; autoIncrement; "`
	Number                int                `json:"number"`
	TypeCode              string             `json:"type-code"`
	FilingDate            string             `json:"filing-date"`
	EmployeeNumber        int                `json:"employee-number"`
	InterlocutoryAttorney string             `json:"interlocutory-attorney-name"`
	LocationCode          string             `json:"location-code"`
	DayInLocation         string             `json:"day-in-location"`
	StatusUpdateDate      string             `json:"status-update-date"`
	StatusCode            int                `json:"status-code"`
	PartyInformation      PartyInformation   `json:"party-information" gorm:"foreignKey:ProceedingEntryID"`
	ProsecutionHistory    ProsecutionHistory `json:"prosecution-history" gorm:"foreignKey:ProceedingEntryID"`
	ProceedingInfoID      uint               `json:"-"`
}

type PartyInformation struct {
	ID                uint    `json:"id" gorm:"primaryKey; autoIncrement; "`
	Party             []Party `json:"party" gorm:"foreignKey:PartyInformationID"`
	ProceedingEntryID uint    `json:"-"`
}

type Party struct {
	ID                  uint                `json:"id" gorm:"primaryKey; autoIncrement; "`
	Identifier          int                 `json:"identifier"`
	RoleCode            string              `json:"role-code"`
	Name                string              `json:"name"`
	PropertyInformation PropertyInformation `json:"property-information" gorm:"foreignKey:PartyID"`
	AddressInformation  AddressInformation  `json:"address-information" gorm:"foreignKey:PartyID"`
	PartyInformationID  uint                `json:"-"`
}

type PropertyInformation struct {
	ID       uint       `json:"id" gorm:"primaryKey; autoIncrement; "`
	Property []Property `json:"property" gorm:"foreignKey:PropertyInformationID"`
	PartyID  uint       `json:"-"`
}

type Property struct {
	ID                    uint   `json:"id" gorm:"primaryKey; autoIncrement; "`
	Identifier            int    `json:"identifier"`
	SerialNumber          int    `json:"serial-number"`
	MarkText              string `json:"mark-text"`
	PropertyInformationID uint   `json:"-"`
}

type AddressInformation struct {
	ID                uint                `json:"id" gorm:"primaryKey; autoIncrement; "`
	ProceedingAddress []ProceedingAddress `json:"proceeding-address" gorm:"foreignKey:AddressInformationID"`
	PartyID           uint                `json:"-"`
}

type ProceedingAddress struct {
	Identifier           int    `json:"identifier gorm: primary_key"`
	TypeCode             string `json:"type-code"`
	Name                 string `json:"name"`
	Address1             string `json:"address-1"`
	City                 string `json:"city"`
	State                string `json:"state"`
	Country              string `json:"country"`
	Postcode             string `json:"postcode"`
	AddressInformationID uint   `json:"-"`
}

type ProsecutionHistory struct {
	ID                uint `json:"id" gorm:"primaryKey; autoIncrement; "`
	ProceedingEntryID uint `json:"-"`
}

type ProsecutionEntry struct {
	ID          uint   `json:"id" gorm:"primaryKey; autoIncrement; "`
	Identifier  int    `json:"identifier"`
	TypeCode    string `json:"type-code"`
	Date        string `json:"date"`
	HistoryText string `json:"history-text"`
}

func MigrateProceedingInfo(db *gorm.DB) {
	db.AutoMigrate(&ProceedingInfo{})
	db.AutoMigrate(&ProceedingEntry{})
	db.AutoMigrate(&PartyInformation{})
	db.AutoMigrate(&Party{})
	db.AutoMigrate(&PropertyInformation{})
	db.AutoMigrate(&Property{})
	db.AutoMigrate(&AddressInformation{})
	db.AutoMigrate(&ProceedingAddress{})
	db.AutoMigrate(&ProsecutionHistory{})
	db.AutoMigrate(&ProsecutionEntry{})
}
