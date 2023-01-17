package main

import (
  "github.com/watcharaphol1938/SE-65-Jeans/controller"
  "github.com/watcharaphol1938/SE-65-Jeans/entity"
  "github.com/gin-gonic/gin"
)

func main() {
  entity.SetupDatabase()
  r := gin.Default()
  r.Use(CORSMiddleware())

  // PatientRegister Routes
  r.GET("/patientregisters", controller.ListPatientRegisters)
  r.GET("/patientregister/:id", controller.GetPatientRegister)
  r.POST("/patientregisters", controller.CreatePatientRegister)
  r.PATCH("/patientregisters", controller.UpdatePatientRegister)
  r.DELETE("/patientregisters/:id", controller.DeletePatientRegister)

  // Employee Routes
  r.GET("/employees", controller.ListEmployees)
  r.GET("/employee/:id", controller.GetEmployee)
  r.POST("/employees", controller.CreateEmployee)
  r.PATCH("/employees", controller.UpdateEmployee)
  r.DELETE("/employees/:id", controller.DeleteEmployee)

  // Gender Routes
  r.GET("/genders", controller.ListGenders)
  r.GET("/gender/:id", controller.GetGender)
  r.POST("/genders", controller.CreateGender)
  r.PATCH("/genders", controller.UpdateGender)
  r.DELETE("/genders/:id", controller.DeleteGender)

  // Prefix Routes
  r.GET("/prefixes", controller.ListPrefixes)
  r.GET("/prefix/:id", controller.GetPrefix)
  r.POST("/prefixes", controller.CreatePrefix)
  r.PATCH("/prefixes", controller.UpdatePrefix)
  r.DELETE("/prefixes/:id", controller.DeletePrefix)

  // Nationality Routes
  r.GET("/nationalities", controller.ListNationalities)
  r.GET("/nationality/:id", controller.GetNationality)
  r.POST("/nationalities", controller.CreateNationality)
  r.PATCH("/nationalities", controller.UpdateNationality)
  r.DELETE("/nationalities/:id", controller.DeleteNationality)

  // Religion Routes
  r.GET("/religions", controller.ListReligions)
  r.GET("/religion/:id", controller.GetReligion)
  r.POST("/religions", controller.CreateReligion)
  r.PATCH("/religions", controller.UpdateReligion)
  r.DELETE("/religions/:id", controller.DeleteReligion)

  // BloodType Routes
  r.GET("/bloodtypes", controller.ListBloodTypes)
  r.GET("/bloodtype/:id", controller.GetBloodType)
  r.POST("/bloodtypes", controller.CreateBloodType)
  r.PATCH("/bloodtypes", controller.UpdateBloodType)
  r.DELETE("/bloodtypes/:id", controller.DeleteBloodType)

  // MaritalStatus Routes
  r.GET("/maritalstatuses", controller.ListMaritalStatuses)
  r.GET("/maritalstatus/:id", controller.GetMaritalStatus)
  r.POST("/maritalstatuses", controller.CreateMaritalStatus)
  r.PATCH("/maritalststuses", controller.UpdateMaritalStatus)
  r.DELETE("/maritalstatuses/:id", controller.DeleteMaritalStautus)

  // SubDistrict Routes
  r.GET("/subdistricts", controller.ListSubDistricts)
  r.GET("/subdistrict/:id", controller.GetSubDistrict)
  r.POST("/subdistricts", controller.CreateSubDistrict)
  r.PATCH("/subdistricts", controller.UpdateSubDistrict)
  r.DELETE("/subdistricts/:id", controller.DeleteSubDistrict)

  // District Routes
  r.GET("/districts", controller.ListDistricts)
  r.GET("/district/:id", controller.GetDistrict)
  r.POST("/districts", controller.CreateDistrict)
  r.PATCH("/districts", controller.UpdateDistrict)
  r.DELETE("/districts/:id", controller.DeleteDistrict)

  // Province Routes
  r.GET("/provinces", controller.ListProvinces)
  r.GET("/province/:id", controller.GetProvince)
  r.POST("/provinces", controller.CreateProvince)
  r.PATCH("/provinces", controller.UpdateProvince)
  r.DELETE("/provinces/:id", controller.DeleteProvince)

  // PostCode Routes
  r.GET("/postcodes", controller.ListPostCodes)
  r.GET("/postcode/:id", controller.GetPostCode)
  r.POST("/postcodes", controller.CreatePostCode)
  r.PATCH("/postcodes", controller.UpdatePostCode)
  r.DELETE("/postcodes/:id", controller.DeletePostCode)

  // Run the server
  r.Run()
}
func CORSMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
        if c.Request.Method == "OPTIONS" {
          c.AbortWithStatus(204)
          return
        }
        c.Next()
  }
}