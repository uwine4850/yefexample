package main

import (
	"fmt"
	"github.com/uwine4850/yefgo/pyclass/module"
	"github.com/uwine4850/yefgotest/gen/customer"
	"github.com/uwine4850/yefgotest/gen/shop"
)

func main() {
	init := module.InitPython{}
	init.Initialize()
	defer init.Finalize()

	shopMod, err := init.GetPyModule("proj.shop")
	if err != nil {
		panic(err)
	}
	customerMod, err := init.GetPyModule("proj.customer")
	if err != nil {
		panic(err)
	}
	shopClass, err := shop.Shop{}.New(&init, shopMod)
	if err != nil {
		panic(err)
	}
	productClass, err := shop.Product{}.New(&init, shopMod, "prod1")
	if err != nil {
		panic(err)
	}
	customerClass, err := customer.Customer{}.New(&init, customerMod, &shopClass.Class)
	if err != nil {
		panic(err)
	}
	err = shopClass.Add_product(&productClass.Class)
	if err != nil {
		panic(err)
	}
	fmt.Println(shopClass.Get_products())
	err = customerClass.By_product(&productClass.Class)
	if err != nil {
		panic(err)
	}
	fmt.Println(shopClass.Get_products())
}
