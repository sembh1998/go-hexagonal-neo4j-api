package products

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/creating"
	mooc "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform"
)

type createRequest struct {
	ID      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	BarCode string `json:"bar_code" binding:"required"`
	ImgUrl  string `json:"img_url" binding:"required"`
	Price   int    `json:"price" binding:"required"`
}

func CreateHandler(creatingProductService *creating.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := creatingProductService.CreateProduct(c, req.ID, req.Name, req.Price, req.BarCode, req.ImgUrl)
		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidProductID),
				errors.Is(err, mooc.ErrInvalidProductName),
				errors.Is(err, mooc.ErrInvalidProductPrice),
				errors.Is(err, mooc.ErrInvalidProductBarCode),
				errors.Is(err, mooc.ErrInvalidProductImgUrl):
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.Status(http.StatusCreated)
	}
}
