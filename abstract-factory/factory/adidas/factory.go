package adidas

import (
	"github.com/ed3899/go-patterns/abstract-factory/factory"
)

type Factory struct {
}

func (a *Factory) MakeShoe() factory.IShoe {
	shoe := new(factory.Shoe)
	shoe.SetLogo("adidas")
	shoe.SetSize(14)

	return &AdidasShoe{
		Shoe: *shoe,
	}
}

func (a *Factory) MakeShirt() factory.IShirt {
	shirt := new(factory.Shirt)
	shirt.SetLogo("adidas")
	shirt.SetSize(14)
	return &AdidasShirt{
		Shirt: *shirt,
	}
}
