package neo4j

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	mooc "github.com/sembh1998/go-hexagonal-neo4j-api/internal/platform"
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
	cypher := `CREATE (p:` + graphProductNodeLabel + ` {id: $id, name: $name, price: $price, bar_code: $bar_code, img_url: $img_url})`
	// convert the product to a map with json tags
	mapparams := structToMap(prod)

	// Execute the cypher query
	_, err := session.Run(ctx, cypher, mapparams)

	if err != nil {
		return fmt.Errorf("error saving product: %w", err)
	}

	return nil
}

// FindByID implements the ProductRepository interface.
func (r *ProductRepository) FindByID(ctx context.Context, id string) (*mooc.Product, error) {

	session := r.Conn.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	// Define the cypher query to find the product
	cypher := `MATCH (p:` + graphProductNodeLabel + ` {id: $id}) RETURN p.id, p.name, p.price, p.bar_code, p.img_url`
	// Execute the cypher query
	result, err := session.Run(ctx, cypher, map[string]interface{}{"id": id})

	if err != nil {
		return nil, fmt.Errorf("error finding product: %w", err)
	}

	// Get the first record
	record, err := result.Single(ctx)

	if err != nil {
		return nil, fmt.Errorf("error finding product: %w", err)
	}

	// Get the product from the record and map it to a Product struct
	product := &mooc.Product{
		ID:      mooc.ProductID(record.Values[0].(string)),
		Name:    mooc.ProductName(record.Values[1].(string)),
		Price:   mooc.ProductPrice(int(record.Values[2].(int64))),
		BarCode: mooc.ProductBarCode(record.Values[3].(string)),
		ImgUrl:  mooc.ProductImgUrl(record.Values[4].(string)),
	}

	return product, nil

}

// FindAll implements the ProductRepository interface.
func (r *ProductRepository) FindAll(ctx context.Context) ([]*mooc.Product, error) {
	session := r.Conn.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	// Define the cypher query to find all products
	cypher := `MATCH (p:` + graphProductNodeLabel + `) RETURN p.id, p.name, p.price, p.bar_code, p.img_url`
	// Execute the cypher query
	result, err := session.Run(ctx, cypher, nil)

	if err != nil {
		return nil, fmt.Errorf("error finding products: %w", err)
	}

	// Get all records
	records, err := result.Collect(ctx)

	if err != nil {
		return nil, fmt.Errorf("error finding products: %w", err)
	}

	// Create a slice of products
	products := make([]*mooc.Product, len(records))

	// Iterate over the records and map them to a Product struct
	for i, record := range records {
		products[i] = &mooc.Product{
			ID:      mooc.ProductID(record.Values[0].(string)),
			Name:    mooc.ProductName(record.Values[1].(string)),
			Price:   mooc.ProductPrice(int(record.Values[2].(int64))),
			BarCode: mooc.ProductBarCode(record.Values[3].(string)),
			ImgUrl:  mooc.ProductImgUrl(record.Values[4].(string)),
		}
	}

	return products, nil

}

// DeleteByID implements the ProductRepository interface.
func (r *ProductRepository) DeleteByID(ctx context.Context, id string) error {
	session := r.Conn.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	// Define the cypher query to delete the product and return the number of deleted nodes
	cypher := `MATCH (p:` + graphProductNodeLabel + ` {id: $id}) DETACH DELETE p RETURN count(p)`
	// Execute the cypher query
	result, err := session.Run(ctx, cypher, map[string]interface{}{"id": id})

	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	// Get the first record
	record, err := result.Single(ctx)

	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	// Get the number of deleted nodes
	deletedNodes := record.Values[0].(int64)

	// If the number of deleted nodes is 0, the product was not found
	if deletedNodes == 0 {
		return fmt.Errorf("error deleting product: product not found")
	}

	return nil
}

// UpdateByID implements the ProductRepository interface.
func (r *ProductRepository) UpdateByID(ctx context.Context, product *mooc.Product) error {
	session := r.Conn.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(context.Background())

	// Define the product to be updated
	prod := graphProduct{
		ID:      string(product.ID),
		Name:    string(product.Name),
		Price:   int(product.Price),
		BarCode: string(product.BarCode),
		ImgUrl:  string(product.ImgUrl),
	}

	// Define the cypher query to update the product and return the number of afected nodes
	cypher := `MATCH (p:` + graphProductNodeLabel + ` {id: $id}) SET p.name = $name, p.price = $price, p.bar_code = $bar_code, p.img_url = $img_url  RETURN count(p)`
	// convert the product to a map with json tags
	mapparams := structToMap(prod)

	// Execute the cypher query
	result, err := session.Run(ctx, cypher, mapparams)

	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	// Get the first record
	record, err := result.Single(ctx)

	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	// Get the number of afected nodes
	afectedNodes := record.Values[0].(int64)

	// If the number of afected nodes is 0, the product was not found
	if afectedNodes == 0 {
		return fmt.Errorf("error updating product: product not found")
	}

	return nil
}
