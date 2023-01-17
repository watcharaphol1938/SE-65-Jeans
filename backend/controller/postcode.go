package controller

import (
	"github.com/watcharaphol1938/SE-65-Jeans/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*	POST /users
	func CreateUser(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := entity.DB().Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
*/

// POST /postcodes
func CreatePostCode(c *gin.Context) {
	var postcode entity.PostCode
	if err := c.ShouldBindJSON(&postcode); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	if err := entity.DB().Create(&postcode).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": postcode})
}

// *******************************************************************************************************

// GET /user/:id
// func GetUser(c *gin.Context) {
// 	var user entity.User
// 	id := c.Param("id")
// 	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

// GET /postcode/:id
func GetPostCode(c *gin.Context) {
	var postcode entity.PostCode
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM postcodes WHERE id = ?", id).Scan(&postcode).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": postcode})
}

// *******************************************************************************************************


// GET /users
// func ListUsers(c *gin.Context) {
// 	var users []entity.User
// 	if err := entity.DB().Raw("SELECT * FROM users").Scan(&users).Error; err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": users})
// }

// GET /postcodes
func ListPostCodes(c *gin.Context) {
	var postcodes []entity.PostCode
	if err := entity.DB().Raw("SELECT * FROM postcodes").Scan(&postcodes).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": postcodes})
}

// *******************************************************************************************************


// DELETE /users/:id
// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }

// DELETE /postcodes/:id
func DeletePostCode(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM postcodes WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "postcode not found"})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// *******************************************************************************************************

// PATCH /users
// func UpdateUser(c *gin.Context) {
// 	var user entity.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	if tx := entity.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		   return
// 	}
// 	if err := entity.DB().Save(&user).Error; err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

// PATCH /postcodes
func UpdatePostCode(c *gin.Context) {
	var postcode entity.PostCode
	if err := c.ShouldBindJSON(&postcode); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	if tx := entity.DB().Where("id = ?", postcode.ID).First(&postcode); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "postcode not found"})
		   return
	}
	if err := entity.DB().Save(&postcode).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": postcode})
}