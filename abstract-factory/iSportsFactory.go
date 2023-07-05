package abstractfactory

import (
	"fmt"
	"github.com/ed3899/go-patterns/abstract-factory/factory"
	"github.com/ed3899/go-patterns/abstract-factory/factory/adidas"
)

// Reference
// https://refactoring.guru/design-patterns/abstract-factory

type ISportsFactory interface {
	MakeShoe() factory.IShoe
	MakeShirt() factory.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
			return &adidas.Adidas{}, nil
	}

	if brand == "nike" {
			return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}