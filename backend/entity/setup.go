package entity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		//&PostCode{},
		&DrugAllergy{},
		&Nurse{},
		&HistorySheet{},
	)
	db = database

	// Employee -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	dob, err := time.Parse("2006-01-02", "2535-05-30")
	if err != nil {
		panic(err)
	}
	employee1 := Employee{
		Name:                 "Sarach Yooyen",
		IdentificationNumber: "4839139280167",
		BirthDay:             dob,
		Mobile:               "0865896723",
		Email:                "SarachTH6@gmail.com",
		Password:             "SYBGPathum06",
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
	bloodtype1 := BloodType{Name: "เอ/A"}
	db.Model(&BloodType{}).Create(&bloodtype1)

	bloodtype2 := BloodType{Name: "บี/B"}
	db.Model(&BloodType{}).Create(&bloodtype2)

	bloodtype3 := BloodType{Name: "โอ/O"}
	db.Model(&BloodType{}).Create(&bloodtype3)

	bloodtype4 := BloodType{Name: "เอบี/AB"}
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
	if err != nil {
		panic(err)
	}
	nurse1 := Nurse{
		Name:                 "Jariyaporn Khotsombat",
		IdentificationNumber: "1430589764258",
		BirthDay:             dob2,
		Mobile:               "0637510564",
		Email:                "Jariyaporn24@gmail.com",
		Password:             "JKNurse24",
		Salary:               22000,
	}
	db.Model(&Nurse{}).Create(&nurse1)

	GetSubDistrictList(db)
	GetDistrictList(db)
	GetProvinceList(db)
	GetNationalityList(db)

}

func GetProvinceList(db *gorm.DB) {
	resp, err := http.Get("https://raw.githubusercontent.com/kongvut/thai-province-data/master/api_province.json")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// println(string(body))
	var provinceResp []struct {
		ID     uint   `json:"id"`
		NameTH string `json:"name_th"`
		NameEN string `json:"name_en"`
	}
	json.Unmarshal(body, &provinceResp)

	var newProvince []Province
	for _, province := range provinceResp {
		newProvince = append(newProvince, Province{
			Model: gorm.Model{ID: province.ID},
			Name:  fmt.Sprintf("%s/%s", province.NameTH, province.NameEN),
		})
	}
	db.Model(&Province{}).Create(&newProvince)
}

func GetDistrictList(db *gorm.DB) {
	resp, err := http.Get("https://raw.githubusercontent.com/kongvut/thai-province-data/master/api_amphure.json")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// println(string(body))
	var districtResp []struct {
		ID         uint   `json:"id"`
		NameTH     string `json:"name_th"`
		NameEN     string `json:"name_en"`
		ProvinceID uint   `json:"province_id"`
	}
	json.Unmarshal(body, &districtResp)

	var newDistrict []District
	for _, district := range districtResp {
		newDistrict = append(newDistrict, District{
			Model:      gorm.Model{ID: district.ID},
			Name:       fmt.Sprintf("%s/%s", district.NameTH, district.NameEN),
			ProvinceID: district.ProvinceID,
		})
	}
	db.Model(&District{}).Create(&newDistrict)
}

func GetSubDistrictList(db *gorm.DB) {
	resp, err := http.Get("https://raw.githubusercontent.com/kongvut/thai-province-data/master/api_tambon.json")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// println(string(body))
	var subDistrictResp []struct {
		ID         uint   `json:"id"`
		NameTH     string `json:"name_th"`
		NameEN     string `json:"name_en"`
		DistrictID uint   `json:"amphure_id"`
		PostCode   uint   `json:"zip_code"`
	}
	json.Unmarshal(body, &subDistrictResp)

	// var newSubDistrict []SubDistrict
	for _, subDistrict := range subDistrictResp {
		newSubDistrict := SubDistrict{
			Model:      gorm.Model{ID: subDistrict.ID},
			Name:       fmt.Sprintf("%s/%s", subDistrict.NameTH, subDistrict.NameEN),
			DistrictID: subDistrict.DistrictID,
			PostCode:   subDistrict.PostCode,
		}
		db.Model(&SubDistrict{}).Create(&newSubDistrict)
	}

}

func GetNationalityList(db *gorm.DB) {
	resp, err := http.Get("https://raw.githubusercontent.com/ponlawat-w/country-list-th/master/country-list-th.json")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// println(string(body))
	var nationalityResp []struct {
		ID     uint   `json:"id"`
		NameTH string `json:"name"`
		NameEN string `json:"enName"`
	}
	json.Unmarshal(body, &nationalityResp)

	var newNationality []Nationality
	for _, nationality := range nationalityResp {
		newNationality = append(newNationality, Nationality{
			Model: gorm.Model{ID: nationality.ID},
			Name:  fmt.Sprintf("%s/%s", nationality.NameTH, nationality.NameEN),
		})
	}
	db.Model(&Nationality{}).Create(&newNationality)
}
