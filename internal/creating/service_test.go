package creating

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	mooc "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform"
	storagemocks "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func Test_ProductService_CreateProduct_RepositoryError(t *testing.T) {
	prodID := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodPrice := faker.Number().NumberInt(4)
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("png", 520, 520)

	product, err := mooc.NewProduct(prodID, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	productRepository := new(storagemocks.ProductRepository)
	productRepository.On("Save", mock.Anything, product).Return(errors.New("something went wrong"))

	productService := NewProductService(productRepository)

	err = productService.CreateProduct(context.Background(), prodID, prodName, prodPrice, prodBarCode, prodImgUrl)

	productRepository.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_ProductService_CreateProduct_Success(t *testing.T) {
	prodID := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodPrice := faker.Number().NumberInt(4)
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("png", 520, 520)

	product, err := mooc.NewProduct(prodID, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	productRepository := new(storagemocks.ProductRepository)
	productRepository.On("Save", mock.Anything, product).Return(nil)

	productService := NewProductService(productRepository)

	err = productService.CreateProduct(context.Background(), prodID, prodName, prodPrice, prodBarCode, prodImgUrl)

	productRepository.AssertExpectations(t)
	assert.NoError(t, err)
}
