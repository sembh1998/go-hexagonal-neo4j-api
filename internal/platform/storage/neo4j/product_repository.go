package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	mooc "github.com/sembh1998/go-hexagonal-api/internal/platform"
)

type ProductRepository struct {
	Conn neo4j.DriverWithContext
}

func NewProductRepository(conn neo4j.DriverWithContext) mooc.ProductRepository {
	return &ProductRepository{
		Conn: conn,
	}
}

// Save implements the ProductRepository interface.
func (r *ProductRepository) Save(ctx context.Context, product *mooc.Product) error {
	session := r.Conn.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	// Define the product to be saved
	prod := graphProduct{
		ID:      string(product.ID),
		Name:    string(product.Name),
		Price:   int(product.Price),
		BarCode: string(product.BarCode),
		ImgUrl:  string(product.ImgUrl),
	}

	// Define the cypher query to save the product
	cypher := `
		CREATE (p:Product {id: $id, name: $name, price: $price, bar_code: $bar_code, img_url: $img_url})
		RETURN p AS product
	`
	// convert the product to a map with json tags
	mapparams := structToMap(prod)

	// Execute the cypher query
	_, err := session.Run(ctx, cypher, mapparams)
	if err != nil {
		return err
	}

	return nil
}

// FindByID implements the ProductRepository interface.
func (r *ProductRepository) FindByID(ctx context.Context, id string) (*mooc.Product, error) {
	panic("implement me")
}

// FindAll implements the ProductRepository interface.
func (r *ProductRepository) FindAll(ctx context.Context) ([]*mooc.Product, error) {
	panic("implement me")
}

// DeleteByID implements the ProductRepository interface.
func (r *ProductRepository) DeleteByID(ctx context.Context, id string) error {
	panic("implement me")
}

// UpdateByID implements the ProductRepository interface.
func (r *ProductRepository) UpdateByID(ctx context.Context, id string, product *mooc.Product) error {
	panic("implement me")
}
