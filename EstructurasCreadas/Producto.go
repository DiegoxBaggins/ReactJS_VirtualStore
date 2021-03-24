package EstructurasCreadas

import "fmt"

type Invent struct {
	Inventarios []Inventario
}

type Inventario struct {
	Tienda       string  `json:"Tienda,omitempty"`
	Departamento string  `json:"Departamento,omitempty"`
	Calificacion float64 `json:"Calificacion,omitempty"`
	Productos    []Product
}

type Product struct {
	Nombre      string  `json:"Nombre,omitempty"`
	Codigo      float64 `json:"Codigo,omitempty"`
	Descripcion string  `json:"Descripcion,omitempty"`
	Precio      float64 `json:"Precio,omitempty"`
	Cantidad    float64 `json:"Cantidad,omitempty"`
	Imagen      string  `json:"Imagen,omitempty"`
}

func (product *Product) ConvertirProduct() *Producto {
	nombre := product.Nombre
	codigo := int(product.Codigo)
	descripcion := product.Descripcion
	precio := product.Precio
	cantidad := int(product.Cantidad)
	imagen := product.Imagen
	return &Producto{nombre, codigo, descripcion, precio, cantidad, imagen, 0, nil, nil}
}

func (inventario *Invent) SacarInventario(Vector []ListaTienda, Indices []string, Departamentos []string) {
	for i := 0; i < len(inventario.Inventarios); i++ {
		tienda := BuscarEspacio(Vector, Indices, Departamentos, inventario.Inventarios[i].Departamento, inventario.Inventarios[i].Tienda[0:1], int(inventario.Inventarios[i].Calificacion), inventario.Inventarios[i].Tienda)
		for j := 0; j < len(inventario.Inventarios[i].Productos); j++ {
			tienda.productos.Insertar(inventario.Inventarios[i].Productos[j])
		}
		//tienda.productos.Print()
	}
}

func BuscarEspacio(Vector []ListaTienda, Indices []string, Departamentos []string, departamento string, first1 string, calificacion int, nombre string) *Tienda {
	indice, dept, err1 := EncontrarIndices(Indices, Departamentos, departamento, first1)
	if err1 == 1 {
		fmt.Println("El Departamento no existe")
	} else {
		elemento := calificacion + 5*(indice+(len(Indices)*dept)) - 1
		store, err := Vector[elemento].BuscarTienda(nombre)
		if err == 0 {
			fmt.Println("Tienda no existe")
		} else {
			return store
		}
	}
	return nil
}

func EncontrarIndices(Indices []string, Departamentos []string, dept string, nombre string) (int, int, int) {
	indice := 0
	departamento := 0
	err := 0
	for indice = 0; indice < len(Indices); indice++ {
		if nombre == Indices[indice] {
			break
		}
	}
	for departamento = 0; departamento < len(Departamentos); departamento++ {
		if dept == Departamentos[departamento] {
			err = 2
			return indice, departamento, err
		}
	}
	err = 1
	return indice, departamento, err
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

func (producto *Producto) ConvertirProducto() Product {
	return Product{producto.nombre, float64(producto.codigo), producto.descripcion, producto.precio, float64(producto.cantidad), producto.imagen}
}

func NuevoProducto(nombre string, codigo float64, descripcion string, precio float64, cantidad float64, imagen string) *Producto {
	return &Producto{nombre, int(codigo), descripcion, precio, int(cantidad), imagen, 0, nil, nil}
}

func (producto *Producto) buscarProd(codigo int) *Producto {
	if producto.codigo == codigo {
		return producto
	} else if producto.codigo > codigo {
		return producto.hizq.buscarProd(codigo)
	} else if producto.codigo < codigo {
		return producto.hder.buscarProd(codigo)
	}
	return nil
}

type ArbolProd struct {
	raiz *Producto
}

func NewArbol() *ArbolProd {
	return &ArbolProd{nil}
}

func (arbol *ArbolProd) max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func (arbol *ArbolProd) obtenerEQ(temp *Producto) int {
	if temp == nil {
		return -1
	}
	return temp.equilibrio
}

func (arbol *ArbolProd) rotacionIzquierda(temp *Producto) *Producto {
	aux := temp.hizq
	temp.hizq = aux.hder
	aux.hder = temp
	temp.equilibrio = arbol.max(arbol.obtenerEQ(temp.hder), arbol.obtenerEQ(temp.hizq)) + 1
	aux.equilibrio = arbol.max(arbol.obtenerEQ(temp.hizq), temp.equilibrio) + 1
	return aux
}

func (arbol *ArbolProd) rotacionDerecha(temp *Producto) *Producto {
	aux := temp.hder
	temp.hder = aux.hizq
	aux.hizq = temp
	temp.equilibrio = arbol.max(arbol.obtenerEQ(temp.hder), arbol.obtenerEQ(temp.hizq)) + 1
	aux.equilibrio = arbol.max(arbol.obtenerEQ(temp.hder), temp.equilibrio) + 1
	return aux
}

func (arbol *ArbolProd) rotacionDobleIzquierda(temp *Producto) *Producto {
	temp.hizq = arbol.rotacionDerecha(temp.hizq)
	return arbol.rotacionIzquierda(temp)
}

func (arbol *ArbolProd) rotacionDobleDerecha(temp *Producto) *Producto {
	temp.hder = arbol.rotacionIzquierda(temp.hder)
	return arbol.rotacionDerecha(temp)
}

func (arbol *ArbolProd) _Insertar(product Product, raiz *Producto) *Producto {
	codigo := int(product.Codigo)
	if raiz == nil {
		return product.ConvertirProduct()
	} else if codigo < raiz.codigo {
		raiz.hizq = arbol._Insertar(product, raiz.hizq)
		if (arbol.obtenerEQ(raiz.hizq) - arbol.obtenerEQ(raiz.hder)) == -2 {
			if codigo < raiz.hizq.codigo {
				raiz = arbol.rotacionIzquierda(raiz)
			} else {
				raiz = arbol.rotacionDobleIzquierda(raiz)
			}
		}
	} else if codigo > raiz.codigo {
		raiz.hder = arbol._Insertar(product, raiz.hder)
		if (arbol.obtenerEQ(raiz.hder) - arbol.obtenerEQ(raiz.hizq)) == 2 {
			if codigo > raiz.hder.codigo {
				raiz = arbol.rotacionDerecha(raiz)
			} else {
				raiz = arbol.rotacionDobleDerecha(raiz)
			}
		}
	} else {
		raiz.cantidad += int(product.Cantidad)
	}
	raiz.equilibrio = arbol.max(arbol.obtenerEQ(raiz.hizq), arbol.obtenerEQ(raiz.hder)) + 1
	return raiz
}

func (arbol *ArbolProd) Insertar(product Product) {
	arbol.raiz = arbol._Insertar(product, arbol.raiz)
}

func (arbol *ArbolProd) Print() {
	arbol.inOrden(arbol.raiz)
}

func (arbol *ArbolProd) inOrden(temp *Producto) {
	if temp != nil {
		arbol.inOrden(temp.hizq)
		fmt.Println("Index: ", temp.codigo, temp.equilibrio, temp.cantidad)
		arbol.inOrden(temp.hder)
	}
}


// otros metodos del arbol

func (arbol *ArbolProd) DevolverListaProducts(producto *Producto) []Product {
	arreglo := make([]Product, 1)
	if producto.hizq != nil {
		arreglo = append(arreglo, arbol.DevolverListaProducts(producto.hizq)...)
	}
	product := producto.ConvertirProducto()
	arreglo[0] = product
	if producto.hder != nil {
		arreglo = append(arreglo, arbol.DevolverListaProducts(producto.hder)...)
	}
	return arreglo
}

func (arbol *ArbolProd) RestarInven(producto *Producto, codigo int, cantidad int) {
	if producto.codigo == codigo {
		producto.cantidad -= cantidad
	}else {
		if codigo < producto.codigo {
			arbol.RestarInven(producto.hizq, codigo, cantidad)
		}else{
			arbol.RestarInven(producto.hder, codigo, cantidad)
		}
	}
}

type ProductCarr struct {
	Nombre      	string  `json:"Nombre,omitempty"`
	Codigo      	float64 `json:"Codigo,omitempty"`
	Descripcion 	string  `json:"Descripcion,omitempty"`
	Precio      	float64 `json:"Precio,omitempty"`
	Cantidad    	float64 `json:"Cantidad,omitempty"`
	Imagen      	string  `json:"Imagen,omitempty"`
	Tienda      	string  `json:"Tienda,omitempty"`
	Calificacion 	float64 `json:"Calificacion,omitempty"`
	Departamento 	string  `json:"Departamento,omitempty"`
}

func NuevoProdCarr(nombre string, codigo float64, descripcion string, precio float64, cantidad float64, imagen string, tienda string, cal float64, dept string) *ProductCarr {
	return &ProductCarr{nombre, codigo, descripcion, precio, cantidad, imagen, tienda, cal, dept}
}














