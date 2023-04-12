package products

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/sembh1998/go-hexagonal-api/internal/platform"
)

type createRequest struct {
	ID      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	BarCode string `json:"bar_code" binding:"required"`
	ImgUrl  string `json:"img_url" binding:"required"`
	Price   int    `json:"price" binding:"required"`
}

func CreateHandler(productRepository mooc.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product, err := mooc.NewProduct(req.ID, req.Name, req.Price, req.BarCode, req.ImgUrl)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := productRepository.Save(c, product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "OK",
		})
	}
}
