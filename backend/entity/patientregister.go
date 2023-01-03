package entity

import (
  "time"
  "gorm.io/gorm"
)

type PatientRegister struct {
  gorm.Model
  FirstName        string
  LastName         string
  Email            string
  Age              uint8
  BirthDay         time.Time

  EmployeeID  *uint
  Employee  Employee  `gorm:"references:ID"`
  
  GenderID  *uint
  Gender  Gender  `gorm:"references:ID"`
  
  PrefixID  *uint
  Prefix  Prefix  `gorm:"references:ID"`
}

type Employee struct{
  gorm.Model
  FirstName string
  LastName  string
  IdentificationNumber  string
  BirthDay  time.Time
  Mobile  string
  Address string
  Salary  uint8
  PatientRegisters  []PatientRegister `gorm:"foreignKey:EmployeeID"`
}

type Gender struct{
  gorm.Model
  Name  string
  PatientRegisters  []PatientRegister `gorm:"foreignKey:GenderID"`
}

type Prefix struct{
  gorm.Model
  Name  string
  PatientRegisters  []PatientRegister `gorm:"foreignKey:PrefixID"`
}