package neo4j

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	mooc "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

const (
	neo4jHost = "localhost"
	neo4jPort = 7687
	neo4jUser = "neo4j"
	neo4jPass = "tartamudoxd"
)

func Test_ProductRepository_Save_Success(t *testing.T) {

	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)

	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.Save(context.Background(), prod)
	assert.NoError(t, err)
}

func Test_ProductRepository_Save_RepositoryDuplicateError(t *testing.T) {
	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)
	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.Save(context.Background(), prod)
	require.NoError(t, err)

	err = productRepository.Save(context.Background(), prod)
	assert.Error(t, err)
}

func Test_ProductRepository_Save_CypherInjectionError(t *testing.T) {
	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)
	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.Save(context.Background(), prod)
	require.NoError(t, err)

	prodName = "'; MATCH (n) DETACH DELETE n; '"
	prod, err = mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	err = productRepository.Save(context.Background(), prod)
	assert.Error(t, err)
}

func Test_ProductRepository_FindByID(t *testing.T) {
	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)
	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.Save(context.Background(), prod)
	require.NoError(t, err)

	prodFound, err := productRepository.FindByID(context.Background(), prodId)
	assert.NoError(t, err)
	assert.Equal(t, prod, prodFound)
}

func Test_ProductRepository_FindByID_NotFound(t *testing.T) {
	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	prodFound, err := productRepository.FindByID(context.Background(), uuid.New().String())
	assert.Error(t, err)
	assert.Nil(t, prodFound)
}

func Test_ProductRepository_FindAll(t *testing.T) {
	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)
	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.Save(context.Background(), prod)
	require.NoError(t, err)

	prodFound, err := productRepository.FindAll(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, prodFound)
}

func Test_ProductRepository_DeleteByID(t *testing.T) {
	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)
	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.Save(context.Background(), prod)
	require.NoError(t, err)

	err = productRepository.DeleteByID(context.Background(), prodId)
	assert.NoError(t, err)

	prodFound, err := productRepository.FindByID(context.Background(), prodId)
	assert.Error(t, err)
	assert.Nil(t, prodFound)
}

func Test_ProductRepository_DeleteByID_NotFound(t *testing.T) {
	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.DeleteByID(context.Background(), uuid.New().String())
	assert.Error(t, err)
}

func Test_ProductRepository_Update(t *testing.T) {
	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)
	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.Save(context.Background(), prod)
	require.NoError(t, err)

	prodName = faker.Commerce().ProductName()
	prodBarCode = faker.Code().Ean13()
	prodImgUrl = faker.Avatar().Url("jpg", 200, 200)
	prodPrice = faker.Number().NumberInt(4)
	prod, err = mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	err = productRepository.UpdateByID(context.Background(), prod)
	assert.NoError(t, err)

	prodFound, err := productRepository.FindByID(context.Background(), prodId)
	assert.NoError(t, err)
	assert.Equal(t, prod, prodFound)
}

func Test_ProductRepository_Update_NotFound(t *testing.T) {
	prodId := uuid.New().String()
	prodName := faker.Commerce().ProductName()
	prodBarCode := faker.Code().Ean13()
	prodImgUrl := faker.Avatar().Url("jpg", 200, 200)
	prodPrice := faker.Number().NumberInt(4)
	prod, err := mooc.NewProduct(prodId, prodName, prodPrice, prodBarCode, prodImgUrl)
	require.NoError(t, err)

	neo4jDriver, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%d", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUser, neo4jPass, ""))
	require.NoError(t, err)

	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	productRepository := NewProductRepository(neo4jDriver)

	err = productRepository.UpdateByID(context.Background(), prod)
	assert.Error(t, err)
}
