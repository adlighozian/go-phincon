package productController

import (
	"encoding/json"
	"net/http"

	"github.com/adlighozian/go-belajar/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"Status": "Success", "products": products})
}

func Show(c *gin.Context) {

	var products []models.Product
	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Error", "Message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Error", "Message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "product": products})

}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": "Tidak dapat mengupddate produk"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {

	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Error", "Message": "Tidak dapat menghapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "Data berhasil dihapus"})
}
