package entity

import (
	"time"

	"gorm.io/gorm"
)

type HistorySheet struct {
	gorm.Model
	Weight                 float32
	Height                 float32
	BMI                    float32
	Temperature            float32
	SystolicBloodPressure  uint16
	DiastolicBloodPressure uint16
	HeartRate              uint8
	RespiratoryRate        uint8
	OxygenSaturation       uint8
	DrugAllergy            string
	Symtom                 string

	PatientRegisterID *uint
	PatientRegister   PatientRegister `gorm:"references:ID"`

	NurseID *uint
	Nurse   Nurse `gorm:"references:ID"`

	UrgencyID *uint
	Urgency   Urgency `gorm:"references:ID"`
}

type Nurse struct {
	gorm.Model
	FirstName            string
	LastName             string
	IdentificationNumber string
	BirthDay             time.Time
	Mobile               string
	Address              string
	Salary               uint16
	HistorySheets        []HistorySheet `gorm:"foreignKey:NurseID"`
}

type Urgency struct {
	gorm.Model
	Name          string
	Description   string
	HistorySheets []HistorySheet `gorm:"foreignKey:UrgencyID"`
}

type PatientRegister struct {
	gorm.Model
	FirstName                          string
	LastName                           string
	IdentificationNumber               string
	Age                                uint8
	BirthDay                           time.Time
	Mobile                             string
	Email                              string
	Occupation                         string
	Address                            string
	EmergencyPersonFirstName           string
	EmergencyPersonLastName            string
	EmergencyPersonMobile              string
	EmergencyPersonOccupation          string
	EmergencyPersonRelationWithPatient string

	EmployeeID *uint
	Employee   Employee `gorm:"references:ID"`

	GenderID *uint
	Gender   Gender `gorm:"references:ID"`

	PrefixID *uint
	Prefix   Prefix `gorm:"references:ID"`

	NationalityID *uint
	Nationality   Nationality `gorm:"references:ID"`

	ReligionID *uint
	Religion   Religion `gorm:"references:ID"`

	BloodTypeID *uint
	BloodType   BloodType `gorm:"references:ID"`

	MaritalStatusID *uint
	MaritalStatus   MaritalStatus `gorm:"references:ID"`

	SubDistrictID *uint
	SubDistrict   SubDistrict `gorm:"references:ID"`

	DistrictID *uint
	District   District `gorm:"references:ID"`

	ProvinceID *uint
	Province   Province `gorm:"references:ID"`

	PostCodeID *uint
	PostCode   PostCode `gorm:"references:ID"`

	HistorySheets []HistorySheet `gorm:"foreignKey:PatientRegisterID"`
}

type Employee struct {
	gorm.Model
	FirstName            string
	LastName             string
	IdentificationNumber string
	BirthDay             time.Time
	Mobile               string
	Address              string
	Salary               uint16
	PatientRegisters     []PatientRegister `gorm:"foreignKey:EmployeeID"`
}

type Gender struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:GenderID"`
}

type Prefix struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:PrefixID"`
}

type Nationality struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:NationalityID"`
}

type Religion struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:ReligionID"`
}

type BloodType struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:BloodTypeID"`
}

type MaritalStatus struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:MaritalStatusID"`
}

type SubDistrict struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:SubDistrictID"`
}

type District struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:DistrictID"`
}

type Province struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:ProvinceID"`
}

type PostCode struct {
	gorm.Model
	Name             string
	PatientRegisters []PatientRegister `gorm:"foreignKey:PostCodeID"`
}
