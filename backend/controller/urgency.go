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

// POST /urgencies
func CreateUrgency(c *gin.Context) {
	var urgency entity.Urgency
	if err := c.ShouldBindJSON(&urgency); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	if err := entity.DB().Create(&urgency).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": urgency})
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

// GET /urgency/:id
func GetUrgency(c *gin.Context) {
	var urgency entity.Urgency
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM urgencies WHERE id = ?", id).Scan(&urgency).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": urgency})
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

// GET /urgencies
func ListUrgencies(c *gin.Context) {
	var urgencies []entity.Urgency
	if err := entity.DB().Raw("SELECT * FROM urgencies").Scan(&urgencies).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": urgencies})
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

// DELETE /urgencies/:id
func DeleteUrgency(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM urgencies WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "urgency not found"})
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

// PATCH /urgencies
func UpdateUrgency(c *gin.Context) {
	var urgency entity.Urgency
	if err := c.ShouldBindJSON(&urgency); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	if tx := entity.DB().Where("id = ?", urgency.ID).First(&urgency); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "urgency not found"})
		   return
	}
	if err := entity.DB().Save(&urgency).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": urgency})
}