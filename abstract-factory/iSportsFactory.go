package abstractfactory

import (
	"fmt"

	"github.com/ed3899/go-patterns/abstract-factory/factory"
	"github.com/ed3899/go-patterns/abstract-factory/factory/adidas"
	"github.com/ed3899/go-patterns/abstract-factory/factory/nike"
)

// Reference
// https://refactoring.guru/design-patterns/abstract-factory

type ISportsFactory interface {
	MakeShoe() factory.IShoe
	MakeShirt() factory.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas.Factory{}, nil
	}

	if brand == "nike" {
		return &nike.Factory{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
