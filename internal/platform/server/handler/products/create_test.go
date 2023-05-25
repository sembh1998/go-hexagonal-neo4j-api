package products

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/creating"
	"github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/bus/inmemory"
	storagemocks "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	productRepository := new(storagemocks.ProductRepository)
	productRepository.On("Save", mock.Anything, mock.AnythingOfType("*platform.Product")).Return(nil).Once()
	productService := creating.NewProductService(productRepository)

	var commandBus = inmemory.NewCommandBus()

	createProductCommandHandler := creating.NewCreateProductCommandHandler(productService)
	commandBus.Register(creating.ProductCreateCommandType, createProductCommandHandler)

	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.POST("/products", CreateHandler(commandBus))

	t.Run("should return 400 when the request body is invalid", func(t *testing.T) {
		// Arrange
		createRequest := createRequest{
			ID:      "123",
			Name:    "Sal",
			BarCode: "123456789",
			ImgUrl:  "https://www.google.com",
			Price:   180,
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return 201 when the product is created", func(t *testing.T) {
		// Arrange
		createRequest := createRequest{
			ID:      "95637397-56cc-47b1-bd7e-1b24ee376ab3",
			Name:    "Sal",
			BarCode: "123456789",
			ImgUrl:  "https://www.google.com",
			Price:   180,
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

}
