package EstructurasCreadas

import (
	"fmt"
)

type Data struct {
	Datos []Alfabeto
}

type Alfabeto struct {
	Indice        string `json:"Indice,omitempty"`
	Departamentos []Tipo
}

type Tipo struct {
	Nombre  string `json:"Nombre,omitempty"`
	Tiendas []Store
}

type Store struct {
	Nombre       string  `json:"Nombre,omitempty"`
	Descripcion  string  `json:"Descripcion,omitempty"`
	Contacto     string  `json:"Contacto,omitempty"`
	Calificacion float64 `json:"Calificacion,omitempty"`
}

type StoreBus struct {
	Departamento string  `json:"Departamento,omitempty"`
	Nombre       string  `json:"Nombre,omitempty"`
	Calificacion float64 `json:"Calificacion,omitempty"`
}

func (info *Data) CalcularTamanos() ([]string, []string) {
	numInd := len(info.Datos)
	numDep := len(info.Datos[0].Departamentos)
	vecInd := make([]string, numInd)
	vecDep := make([]string, numDep)
	for indice := 0; indice < numInd; indice++ {
		vecInd[indice] = info.Datos[indice].Indice
	}
	for departamento := 0; departamento < numInd; departamento++ {
		vecDep[departamento] = info.Datos[0].Departamentos[departamento].Nombre
	}
	return vecInd, vecDep
}

func (info *Data) TransformarDatos() []ListaTienda {
	numInd := len(info.Datos)
	numDep := len(info.Datos[0].Departamentos)
	numTie := 0
	tamano := numInd * numDep * 5
	arreglo := make([]ListaTienda, tamano)
	elemento := 0
	calificacion := 0
	nombre := ""
	descripcion := ""
	contacto := ""
	for indice := 0; indice < numInd; indice++ {
		for departamento := 0; departamento < numDep; departamento++ {
			numTie = len(info.Datos[indice].Departamentos[departamento].Tiendas)
			for tienda := 0; tienda < numTie; tienda++ {
				elemento = 0
				calificacion = int(info.Datos[indice].Departamentos[departamento].Tiendas[tienda].Calificacion)
				nombre = info.Datos[indice].Departamentos[departamento].Tiendas[tienda].Nombre
				descripcion = info.Datos[indice].Departamentos[departamento].Tiendas[tienda].Descripcion
				contacto = info.Datos[indice].Departamentos[departamento].Tiendas[tienda].Contacto
				elemento = calificacion + 5*(indice+(numInd*departamento)) - 1
				arreglo[elemento].InsertarTienda(nombre, descripcion, contacto, calificacion)
			}
		}
	}
	return arreglo
}

type Tienda struct {
	nombre       string
	descripcion  string
	contacto     string
	calificacion int
	anterior     *Tienda
	siguiente    *Tienda
}

type ListaTienda struct {
	primero   *Tienda
	ultimo    *Tienda
	elementos int
}

func (lista *ListaTienda) MostrarDatos() {
	auxiliar := lista.primero
	for i := 0; i < lista.elementos; i++ {
		tienda := *auxiliar
		fmt.Println(tienda.nombre, tienda.contacto, tienda.descripcion, tienda.calificacion)
		auxiliar = tienda.siguiente
	}
	auxiliar = lista.primero
	for i := 0; i < lista.elementos; i++ {
		fmt.Println(auxiliar)
		auxiliar = auxiliar.siguiente
	}
}

func (lista *ListaTienda) VectorElementos() []Store {
	var vector = make([]Store, lista.elementos)
	auxiliar := lista.primero
	for i := 0; i < lista.elementos; i++ {
		tienda := *auxiliar
		vector[i].Nombre = tienda.nombre
		vector[i].Contacto = tienda.contacto
		vector[i].Descripcion = tienda.descripcion
		vector[i].Calificacion = float64(tienda.calificacion)
		auxiliar = tienda.siguiente
	}
	return vector
}

func (lista *ListaTienda) BuscarTienda(nombre string) (Store, int) {
	var store Store
	auxiliar := lista.primero
	for i := 0; i < lista.elementos; i++ {
		tienda := *auxiliar
		if tienda.nombre == nombre {
			store.Calificacion = float64(tienda.calificacion)
			store.Nombre = tienda.nombre
			store.Descripcion = tienda.descripcion
			store.Contacto = tienda.contacto
			return store, 1
		}
		auxiliar = tienda.siguiente
	}
	return store, 0
}

/*
func (lista *ListaTienda)InsertarTienda (nombre string, descripcion string, contacto int, calificacion int) {
	nuevaTienda := &Tienda{nombre, descripcion,contacto,calificacion, nil, nil}
	if lista.primero == nil {
		lista.ultimo = nuevaTienda
		lista.primero = nuevaTienda
		lista.elementos += 1
	} else {
		lista.ultimo.siguiente = nuevaTienda
		nuevaTienda.anterior = lista.ultimo
		lista.ultimo = nuevaTienda
		lista.elementos += 1
	}
}
*/

func (lista *ListaTienda) InsertarUltimo(nombre string, descripcion string, contacto string, calificacion int) {
	nuevaTienda := &Tienda{nombre, descripcion, contacto, calificacion, nil, nil}
	lista.ultimo.siguiente = nuevaTienda
	nuevaTienda.anterior = lista.ultimo
	lista.ultimo = nuevaTienda
	lista.elementos += 1
}

func (lista *ListaTienda) InsertarPrimero(nombre string, descripcion string, contacto string, calificacion int) {
	nuevaTienda := &Tienda{nombre, descripcion, contacto, calificacion, nil, nil}
	lista.primero.anterior = nuevaTienda
	nuevaTienda.siguiente = lista.primero
	lista.primero = nuevaTienda
	lista.elementos += 1
}

func (lista *ListaTienda) InsertarMedio(nombre string, descripcion string, contacto string, calificacion int, dir1 *Tienda, dir2 *Tienda) {
	nuevaTienda := &Tienda{nombre, descripcion, contacto, calificacion, nil, nil}
	dir1.siguiente = nuevaTienda
	dir2.anterior = nuevaTienda
	nuevaTienda.siguiente = dir2
	nuevaTienda.anterior = dir1
	lista.elementos += 1
}

func (lista *ListaTienda) InsertarTienda(nombre string, descripcion string, contacto string, calificacion int) {
	tamano := lista.elementos
	if tamano == 0 {
		nuevaTienda := &Tienda{nombre, descripcion, contacto, calificacion, nil, nil}
		lista.ultimo = nuevaTienda
		lista.primero = nuevaTienda
		lista.elementos += 1
	} else {
		if tamano == 1 {
			primero := *lista.primero
			if primero.nombre > nombre {
				lista.InsertarPrimero(nombre, descripcion, contacto, calificacion)
			} else {
				lista.InsertarUltimo(nombre, descripcion, contacto, calificacion)
			}
		} else {
			autorizacion := 0
			auxiliar := lista.primero
			for i := 1; i < lista.elementos; i++ {
				if auxiliar.nombre == nombre {
					return
				}
				if auxiliar.nombre > nombre {
					if auxiliar == lista.primero {
						lista.InsertarPrimero(nombre, descripcion, contacto, calificacion)
						autorizacion = 1
						break
					} else {
						lista.InsertarMedio(nombre, descripcion, contacto, calificacion, auxiliar.anterior, auxiliar)
						autorizacion = 1
						break
					}
				} else {
					auxiliar = auxiliar.siguiente
				}
			}
			if autorizacion == 0 {
				if auxiliar.nombre == nombre {
					return
				} else {
					lista.InsertarUltimo(nombre, descripcion, contacto, calificacion)
				}
			}
		}
	}
}
