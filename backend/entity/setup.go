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
  database.AutoMigrate(
    &PatientRegister{}, 
    &Employee{}, 
    &Gender{}, 
    &Prefix{}, 
    &Nationality{}, 
    &Religion{}, 
    &BloodType{}, 
    &MaritalStatus{}, 
    &SubDistrict{}, 
    &District{}, 
    &Province{}, 
    &PostCode{}, 
    &Urgency{}, 
    &Nurse{}, 
    &HistorySheet{},
  )
  db = database

  // Employee -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  dob, err := time.Parse("2006-01-02", "2535-05-30")
  employee1 := Employee{
    FirstName: "สารัช",
    LastName: "อยู่เย็น",
    IdentificationNumber: "4839139280167",
    BirthDay: dob,
    Mobile: "0865896723",
    Address: "237 หมู่ 5 บ้านสามัคคีชัย ตำบลพระบาทนาสิงห์ อำเภอรัตนวาปี จังหวัดหนองคาย",
    Salary: 22000,
  }
	db.Model(&Employee{}).Create(&employee1)

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

  // Religion -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  religion1 := Religion{Name: "พุทธ"}
  db.Model(&Religion{}).Create(&religion1)

  religion2 := Religion{Name: "คริสต์"}
  db.Model(&Religion{}).Create(&religion2)

  religion3 := Religion{Name: "อิสลาม"}
  db.Model(&Religion{}).Create(&religion3)

  religion4 := Religion{Name: "ฮินดู"}
  db.Model(&Religion{}).Create(&religion4)

  religion5 := Religion{Name: "ซิกข์"}
  db.Model(&Religion{}).Create(&religion5)

  // BloodType -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  bloodtype1 := BloodType{Name: "เอ"}
  db.Model(&BloodType{}).Create(&bloodtype1)

  bloodtype1 := BloodType{Name: "บี"}
  db.Model(&BloodType{}).Create(&bloodtype1)

  bloodtype1 := BloodType{Name: "โอ"}
  db.Model(&BloodType{}).Create(&bloodtype1)

  bloodtype1 := BloodType{Name: "เอบี"}
  db.Model(&BloodType{}).Create(&bloodtype1)

  // MaritalStatus -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  maritalstatus1 := MaritalStatus{Name: "โสด"}
  db.Model(&MaritalStatus{}).Create(&maritalstatus1)

  maritalstatus2 := MaritalStatus{Name: "สมรส"}
  db.Model(&MaritalStatus{}).Create(&maritalstatus2)

  maritalstatus3 := MaritalStatus{Name: "หม้าย"}
  db.Model(&MaritalStatus{}).Create(&maritalstatus3)

  maritalstatus4 := MaritalStatus{Name: "หย่า"}
  db.Model(&MaritalStatus{}).Create(&maritalstatus4)

  maritalstatus5 := MaritalStatus{Name: "แยกกันอยู่"}
  db.Model(&MaritalStatus{}).Create(&maritalstatus5)

  // Urgency -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  urgency1 := Urgency{Name: "สีแดง"}
  db.Model(&Urgency{}).Create(&urgency1)

  urgency2 := Urgency{Name: "สีส้ม"}
  db.Model(&Urgency{}).Create(&urgency2)

  urgency3 := Urgency{Name: "สีเหลือง"}
  db.Model(&Urgency{}).Create(&urgency3)

  urgency4 := Urgency{Name: "สีเขียว"}
  db.Model(&Urgency{}).Create(&urgency4)

  urgency5 := Urgency{Name: "สีขาว"}
  db.Model(&Urgency{}).Create(&urgency5)

  // Nurse -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  nurse1 := Nurse{
    FirstName: "จริยาภรณ์",
    LastName: "โคตรสมบัติ",
    IdentificationNumber: "1430589764258",
    BirthDay: "2536-04-24",
    Mobile: "0637510564",
    Address: "184 หมู่ 5 บ้านนาชุมช้าง ตำบลรัตนวาปี อำเภอรัตนวาปี จังหวัดหนองคาย",
    Salary: 22000,
  }
	db.Model(&Nurse{}).Create(&nurse1)

}