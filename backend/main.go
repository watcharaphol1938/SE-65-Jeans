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