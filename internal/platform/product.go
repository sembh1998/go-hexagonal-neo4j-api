package platform

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidProductID = errors.New("invalid product id")

type ProductID string

func NewProductID(id string) (ProductID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("%w:%s", ErrInvalidProductID, id)
	}
	return ProductID(id), nil
}

var ErrInvalidProductName = errors.New("invalid product name")

type ProductName string

func NewProductName(name string) (ProductName, error) {
	if len(name) == 0 {
		return "", fmt.Errorf("%w: name is required", ErrInvalidProductName)
	}
	return ProductName(name), nil
}

var ErrInvalidProductPrice = errors.New("invalid product price")

type ProductPrice int

func NewProductPrice(price int) (ProductPrice, error) {
	if price < 0 {
		return 0, fmt.Errorf("%w: price must be greater than 0", ErrInvalidProductPrice)
	}
	return ProductPrice(price), nil
}

var ErrInvalidProductBarCode = errors.New("invalid product bar code")

type ProductBarCode string

func NewProductBarCode(barCode string) (ProductBarCode, error) {
	if len(barCode) == 0 {
		return "", fmt.Errorf("%w: barCode is required", ErrInvalidProductBarCode)
	}
	return ProductBarCode(barCode), nil
}

var ErrInvalidProductImgUrl = errors.New("invalid product img url")

type ProductImgUrl string

func NewProductImgUrl(imgUrl string) (ProductImgUrl, error) {
	if len(imgUrl) == 0 {
		return "", fmt.Errorf("%w: imgUrl is required", ErrInvalidProductImgUrl)
	}
	return ProductImgUrl(imgUrl), nil
}

type Product struct {
	ID      ProductID
	Name    ProductName
	Price   ProductPrice
	BarCode ProductBarCode
	ImgUrl  ProductImgUrl
}

func NewProduct(id, name, barCode, imgUrl string, price int) (*Product, error) {
	Id, err := NewProductID(id)
	if err != nil {
		return nil, err
	}

	Name, err := NewProductName(name)
	if err != nil {
		return nil, err
	}

	Price, err := NewProductPrice(price)
	if err != nil {
		return nil, err
	}

	BarCode, err := NewProductBarCode(barCode)
	if err != nil {
		return nil, err
	}

	ImgUrl, err := NewProductImgUrl(imgUrl)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:      Id,
		Name:    Name,
		Price:   Price,
		BarCode: BarCode,
		ImgUrl:  ImgUrl,
	}, nil
}

// equals returns true if the product is equal to another product.
func (p *Product) Equals(other *Product) bool {
	return p.ID == other.ID &&
		p.Name == other.Name &&
		p.Price == other.Price &&
		p.BarCode == other.BarCode &&
		p.ImgUrl == other.ImgUrl
}

type ProductRepository interface {
	// Save persists the product.
	Save(ctx context.Context, product *Product) error
	// FindByID returns the product with the given ID.
	FindByID(ctx context.Context, id string) (*Product, error)
	// FindAll returns all products.
	FindAll(ctx context.Context) ([]*Product, error)
	// DeleteByID deletes the product with the given ID.
	DeleteByID(ctx context.Context, id string) error
	// UpdateByID updates the product with the given ID.
	UpdateByID(ctx context.Context, product *Product) error
}

// go:generate mockery --case snake --outpkg storagemocks --output ./storage/storagemocks --name ProductRepository
