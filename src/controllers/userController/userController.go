package userController

import (
	"encoding/json"
	"net/http"

	"github.com/adlighozian/go-belajar/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var user []models.User
	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": user})
}

func Show(c *gin.Context) {
	var user []models.User
	id := c.Param("id")
	err := models.DB.First(&user, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Error", "message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": user})
}

func Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Data gagal diupdate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {

	var user models.User

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Data gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Data berhasil dihapus"})

}
