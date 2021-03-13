package EstructurasCreadas

import (
	"fmt"
	"go/types"
	"strconv"
)

type Product struct {
	Nombre      string  `json:"Nombre,omitempty"`
	Codigo      float64 `json:"Codigo,omitempty"`
	Descripcion string  `json:"Descripcion,omitempty"`
	Precio      float64 `json:"Precio,omitempty"`
	Cantidad    float64 `json:"Cantidad,omitempty"`
	Imagen      string  `json:"Imagen,omitempty"`
}

func (product *Product) ConvertirTienda() *Producto {
	nombre := product.Nombre
	codigo := int(product.Codigo)
	descripcion := product.Descripcion
	precio := product.Precio
	cantidad := int(product.Cantidad)
	imagen := product.Imagen
	return &Producto{nombre, codigo, descripcion, precio, cantidad, imagen, 0, nil, nil}
}

type Producto struct {
	nombre      string
	codigo      int
	descripcion string
	precio      float64
	cantidad    int
	imagen      string
	equilibrio  int
	hizq        *Producto
	hder        *Producto
}

type ArbolProd struct {
	raiz Producto
}

func NewArbol() *ArbolProd {
	return &ArbolProd{nil}
}

func (producto *Producto) buscarProd(codigo int) *Producto {
	if producto.codigo == codigo {
		return producto
	} else if producto.codigo < codigo {
		return producto.hizq.buscarProd(codigo)
	} else if producto.codigo > codigo {
		return producto.hder.buscarProd(codigo)
	}
	return nil
}

func (producto *Producto) obtenerEq() int {
	if producto == nil {
		return -1
	} else {
		return producto.equilibrio
	}
}

func (producto *Producto) RotarIzq() *Producto {
	auxiliar := producto.hder
}
