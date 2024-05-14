package customer

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

type Customer struct {
	goclass.Class
	PyInit *module.InitPython
}

func (n Customer) New(pyInit *module.InitPython, pyModule pytypes.Module, shop *goclass.Class) (Customer, error) {
	n.PyInit = pyInit
	instance := pyclass.NewPyInstance(n.PyInit, pyModule, "Customer", shop)
	newInstance, err := instance.Create()
	if err != nil {
		return Customer{}, err
	}
	n.PyInit.FreeObject(unsafe.Pointer(newInstance))
	n.SetInstance(newInstance)

	class, err := pyclass.GetPyClass("Customer", pyModule)
	if err != nil {
		return Customer{}, err
	}
	n.PyInit.FreeObject(unsafe.Pointer(class))
	n.SetClass(class)
	n.SetPyModule(pyModule)
	return n, nil
}

func (n Customer) By_product(product *goclass.Class) error {
	_, err := pyclass.CallInstanceMethod(n.PyInit, &n.Class, "by_product", product)
	if err != nil {
		return err
	}
	return nil
}
