package creating

import (
	"context"

	mooc "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform"
)

type ProductService struct {
	productRepository mooc.ProductRepository
}

func NewProductService(prodRepo mooc.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: prodRepo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, id, name, barcode, imgUrl string, price int) error {
	product, err := mooc.NewProduct(id, name, barcode, imgUrl, price)
	if err != nil {
		return err
	}
	return s.productRepository.Save(ctx, product)
}
