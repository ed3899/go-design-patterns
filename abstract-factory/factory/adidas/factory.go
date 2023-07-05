package adidas

import (
	"github.com/ed3899/go-patterns/abstract-factory/factory"
)

type Adidas struct {
}

func (a *Adidas) MakeShoe() factory.IShoe {
    shoe := new(factory.Shoe)
    shoe.SetLogo("adidas")
    shoe.SetSize(14)

    return &AdidasShoe{
        Shoe: *shoe,
    }
}

func (a *Adidas) MakeShirt() factory.IShirt {
    shirt := new(factory.Shirt)
    shirt.SetLogo("adidas")
    shirt.SetSize(14)
    return &AdidasShirt{
        Shirt: *shirt,
    }
}