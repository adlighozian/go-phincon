package logsLoginController

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adlighozian/go-belajar/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var log []models.LogsLogin
	models.DB.Find(&log)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": log})
}

func Show(c *gin.Context) {
	var log []models.LogsLogin
	id := c.Param("id")
	err := models.DB.First(&log, id).Error
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
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": log})
}

func Create(c *gin.Context) {
	var log models.LogsLogin
	if err := c.ShouldBindJSON(&log); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}
	log.LoginAt = time.Now()
	
	models.DB.Create(&log)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "data": log})
}

func Update(c *gin.Context) {
	var log models.LogsLogin
	id := c.Param("id")

	if err := c.ShouldBindJSON(&log); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}

	if models.DB.Model(&log).Where("id = ?", id).Updates(&log).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Data gagal diupdate"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {

	var log models.LogsLogin

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&log, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Data gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Data berhasil dihapus"})
}
