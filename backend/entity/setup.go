package entity

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {
  return db
}

func SetupDatabase() {
  database, err := gorm.Open(sqlite.Open("se-65.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  database.AutoMigrate(&PatientRegister{}, &Employee{}, &Gender{}, &Prefix{}, &Nationality{}, &Religion{}, &BloodType{}, &MaritalStatus{}, &SubDistrict{}, &District{}, &Province{}, &PostCode{}, &Urgency{}, &Nurse{}, &HistorySheet{})
  db = database

  // Employee -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  
  // Gender -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  gender1 := Gender{Name: "ชาย"}
	db.Model(&Gender{}).Create(&gender1)

	gender2 := Gender{Name: "หญิง"}
	db.Model(&Gender{}).Create(&gender2)

	gender3 := Gender{Name: "อื่นๆ"}
	db.Model(&Gender{}).Create(&gender3)

  // Prefix -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  prefix1 := Prefix{Name: "เด็กชาย"}
  db.Model(&Prefix{}).Create(&prefix1)

  prefix2 := Prefix{Name: "เด็กหญิง"}
  db.Model(&Prefix{}).Create(&prefix2)

  prefix3 := Prefix{Name: "นางสาว"}
  db.Model(&Prefix{}).Create(&prefix3)

  prefix4 := Prefix{Name: "นาย"}
  db.Model(&Prefix{}).Create(&prefix4)

  prefix5 := Prefix{Name: "นาง"}
  db.Model(&Prefix{}).Create(&prefix5)

}