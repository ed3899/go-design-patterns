package nike

import (
	"github.com/ed3899/go-patterns/abstract-factory/factory"
)

type Factory struct {
}

func (n *Factory) MakeShoe() factory.IShoe {
	shoe := new(factory.Shoe)
	shoe.SetLogo("nike")
	shoe.SetSize(14)

	return &NikeShoe{
		Shoe: *shoe,
	}
}

func (n *Factory) MakeShirt() factory.IShirt {
	shirt := new(factory.Shirt)
	shirt.SetLogo("nike")
	shirt.SetSize(14)

	return &NikeShirt{
		Shirt: *shirt,
	}
}
