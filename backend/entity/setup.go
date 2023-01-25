package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		&DrugAllergy{},
		&Nurse{},
		&HistorySheet{},
	)
	db = database

	// Employee -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	dob, err := time.Parse("2006-01-02", "2535-05-30")
	employee1 := Employee{
		FirstName:            "Sarach",
		LastName:             "Yooyen",
		IdentificationNumber: "4839139280167",
		BirthDay:             dob,
		Mobile:               "0865896723",
		Address:              "237 moo. 5, Samakkhee Chai Village, Phra Bath Na Sing Sub-District, Rattanawapi District, Nong Khai Province",
		Salary:               22000,
	}
	db.Model(&Employee{}).Create(&employee1)

	// Gender -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	gender1 := Gender{Name: "ชาย/Male"}
	db.Model(&Gender{}).Create(&gender1)

	gender2 := Gender{Name: "หญิง/Female"}
	db.Model(&Gender{}).Create(&gender2)

	gender3 := Gender{Name: "อื่นๆ/Other"}
	db.Model(&Gender{}).Create(&gender3)

	// Prefix -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	prefix1 := Prefix{Name: "เด็กชาย/Master"}
	db.Model(&Prefix{}).Create(&prefix1)

	prefix2 := Prefix{Name: "เด็กหญิง/นางสาว/Miss"}
	db.Model(&Prefix{}).Create(&prefix2)

	prefix3 := Prefix{Name: "นาย/Mister"}
	db.Model(&Prefix{}).Create(&prefix3)

	prefix4 := Prefix{Name: "นาง/Mistress"}
	db.Model(&Prefix{}).Create(&prefix4)

	// Religion -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	religion1 := Religion{Name: "พุทธ/Bhuddhism"}
	db.Model(&Religion{}).Create(&religion1)

	religion2 := Religion{Name: "คริสต์/Christian"}
	db.Model(&Religion{}).Create(&religion2)

	religion3 := Religion{Name: "อิสลาม/Islam"}
	db.Model(&Religion{}).Create(&religion3)

	religion4 := Religion{Name: "ฮินดู/Hinduism"}
	db.Model(&Religion{}).Create(&religion4)

	religion5 := Religion{Name: "ซิกข์/Sikhism"}
	db.Model(&Religion{}).Create(&religion5)

	// BloodType -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	bloodtype1 := BloodType{Name: "A"}
	db.Model(&BloodType{}).Create(&bloodtype1)

	bloodtype2 := BloodType{Name: "B"}
	db.Model(&BloodType{}).Create(&bloodtype2)

	bloodtype3 := BloodType{Name: "O"}
	db.Model(&BloodType{}).Create(&bloodtype3)

	bloodtype4 := BloodType{Name: "AB"}
	db.Model(&BloodType{}).Create(&bloodtype4)

	// MaritalStatus -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	maritalstatus1 := MaritalStatus{Name: "โสด/Single"}
	db.Model(&MaritalStatus{}).Create(&maritalstatus1)

	maritalstatus2 := MaritalStatus{Name: "แต่งงาน/สมรส/Married"}
	db.Model(&MaritalStatus{}).Create(&maritalstatus2)

	maritalstatus3 := MaritalStatus{Name: "หม้าย/Widowed"}
	db.Model(&MaritalStatus{}).Create(&maritalstatus3)

	maritalstatus4 := MaritalStatus{Name: "หย่าร้าง/Divorced"}
	db.Model(&MaritalStatus{}).Create(&maritalstatus4)

	maritalstatus5 := MaritalStatus{Name: "แยกกันอยู่/Saparated"}
	db.Model(&MaritalStatus{}).Create(&maritalstatus5)

	// DrugAllergy -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	drugallergy1 := DrugAllergy{
		Name: "แพ้ยา/Drug Allergy",
	}
	db.Model(&DrugAllergy{}).Create(&drugallergy1)

	drugallergy2 := DrugAllergy{
		Name: "ไม่แพ้ยา/Not Allergic",
	}
	db.Model(&DrugAllergy{}).Create(&drugallergy2)

	// Nurse -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	dob2, err := time.Parse("2006-01-02", "2536-04-24")
	nurse1 := Nurse{
		FirstName:            "Jariyaporn",
		LastName:             "Khotsombat",
		IdentificationNumber: "1430589764258",
		BirthDay:             dob2,
		Mobile:               "0637510564",
		Address:              "184 moo. 5, Nachumchang Village, Rattanawapi Sub-District, Rattanawapi District, Nong Khai Province",
		Salary:               22000,
	}
	db.Model(&Nurse{}).Create(&nurse1)

}
