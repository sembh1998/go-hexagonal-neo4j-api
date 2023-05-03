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

func (s *ProductService) CreateProduct(ctx context.Context, id, name string, price int, barcode, imgUrl string) error {
	product, err := mooc.NewProduct(id, name, price, barcode, imgUrl)
	if err != nil {
		return err
	}
	return s.productRepository.Save(ctx, product)
}
