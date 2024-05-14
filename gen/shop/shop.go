package shop

/*
#include <Python.h>
#cgo pkg-config: python3
*/
import "C"
import (
	"github.com/uwine4850/yefgo/goclass"
	"github.com/uwine4850/yefgo/pyclass"
	"github.com/uwine4850/yefgo/pyclass/module"
	"github.com/uwine4850/yefgo/pytypes"
	"unsafe"
)

type Product struct {
	goclass.Class
	PyInit *module.InitPython
}

func (n Product) New(pyInit *module.InitPython, pyModule pytypes.Module, name string) (Product, error) {
	n.PyInit = pyInit
	instance := pyclass.NewPyInstance(n.PyInit, pyModule, "Product", name)
	newInstance, err := instance.Create()
	if err != nil {
		return Product{}, err
	}
	n.PyInit.FreeObject(unsafe.Pointer(newInstance))
	n.SetInstance(newInstance)

	class, err := pyclass.GetPyClass("Product", pyModule)
	if err != nil {
		return Product{}, err
	}
	n.PyInit.FreeObject(unsafe.Pointer(class))
	n.SetClass(class)
	n.SetPyModule(pyModule)
	return n, nil
}

type Shop struct {
	goclass.Class
	PyInit *module.InitPython
}

func (n Shop) New(pyInit *module.InitPython, pyModule pytypes.Module) (Shop, error) {
	n.PyInit = pyInit
	instance := pyclass.NewPyInstance(n.PyInit, pyModule, "Shop")
	newInstance, err := instance.Create()
	if err != nil {
		return Shop{}, err
	}
	n.PyInit.FreeObject(unsafe.Pointer(newInstance))
	n.SetInstance(newInstance)

	class, err := pyclass.GetPyClass("Shop", pyModule)
	if err != nil {
		return Shop{}, err
	}
	n.PyInit.FreeObject(unsafe.Pointer(class))
	n.SetClass(class)
	n.SetPyModule(pyModule)
	return n, nil
}

func (n Shop) Add_product(product *goclass.Class) error {
	_, err := pyclass.CallInstanceMethod(n.PyInit, &n.Class, "add_product", product)
	if err != nil {
		return err
	}
	return nil
}

func (n Shop) Delete_products(product *goclass.Class) error {
	_, err := pyclass.CallInstanceMethod(n.PyInit, &n.Class, "delete_products", product)
	if err != nil {
		return err
	}
	return nil
}

func (n Shop) Get_products() (*[]Product, error) {
	res, err := pyclass.CallInstanceMethod(n.PyInit, &n.Class, "get_products")
	if err != nil {
		return nil, err
	}
	var output []Product
	err = pyclass.MethodOutput(n.PyInit, res, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
