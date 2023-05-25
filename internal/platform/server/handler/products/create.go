package products

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/creating"
	mooc "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform"
	"github.com/sembh1998/go-hexagonal-neo4j-api/kit/command"
)

type createRequest struct {
	ID      string `json:"id" binding:"required,uuid"`
	Name    string `json:"name" binding:"required,min=3,max=100"`
	BarCode string `json:"bar_code" binding:"required,min=3,max=100"`
	ImgUrl  string `json:"img_url" binding:"required,min=3,max=100"`
	Price   int    `json:"price" binding:"required,numeric"`
}

func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := commandBus.Dispatch(c, creating.NewProductCommand(
			req.ID,
			req.Name,
			req.BarCode,
			req.ImgUrl,
			req.Price,
		))
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
