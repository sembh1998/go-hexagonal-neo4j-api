package creating

import (
	"context"
	"errors"

	"github.com/sembh1998/go-hexagonal-neo4j-api/kit/command"
)

const ProductCreateCommandType command.Type = "command.creating.product"

// ProductCommand is the command dispatched to create a product.
type ProductCommand struct {
	ID      string
	Name    string
	BarCode string
	ImgUrl  string
	Price   int
}

func NewProductCommand(id, name, barCode, imgUrl string, price int) command.Command {
	return &ProductCommand{
		ID:      id,
		Name:    name,
		BarCode: barCode,
		ImgUrl:  imgUrl,
		Price:   price,
	}
}

func (p *ProductCommand) Type() command.Type {
	return ProductCreateCommandType
}

// ProductCommandHandler is the handler that creates the product.
type ProductCommandHandler struct {
	service *ProductService
}

func NewCreateProductCommandHandler(service *ProductService) command.Handler {
	return &ProductCommandHandler{
		service: service,
	}
}

func (h *ProductCommandHandler) Handle(ctx context.Context, command command.Command) error {
	p, ok := command.(*ProductCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateProduct(ctx, p.ID, p.Name, p.BarCode, p.ImgUrl, p.Price)
}
