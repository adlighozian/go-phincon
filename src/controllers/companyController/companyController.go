package companyController

import (
	"encoding/json"
	"net/http"

	"github.com/adlighozian/go-belajar/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var company []models.Company
	models.DB.Find(&company)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": company})
}

func Show(c *gin.Context) {
	var company []models.Company
	id := c.Param("id")
	err := models.DB.First(&company, id).Error
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
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": company})
}

func Create(c *gin.Context) {

	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}

	models.DB.Create(&company)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": company})
}

func Update(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}

	if models.DB.Model(&company).Where("id = ?", id).Updates(&company).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Data gagal diupdate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {

	var company models.Company

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&company, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Data gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Data berhasil dihapus"})
}
