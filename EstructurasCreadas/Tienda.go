package EstructurasCreadas

import (
	"fmt"
)

type Tienda struct {
	nombre       string
	descripcion  string
	contacto     int
	calificacion int
	anterior     *Tienda
	siguiente    *Tienda
}

type ListaTienda struct {
	primero   *Tienda
	ultimo    *Tienda
	elementos int
}

func NuevaLista() *ListaTienda {
	return &ListaTienda{nil, nil, 0}
}

func (lista *ListaTienda) MostrarDatos() {
	auxiliar := lista.primero
	for i := 0; i < lista.elementos; i++ {
		tienda := *auxiliar
		fmt.Println(tienda.nombre, tienda.contacto, tienda.descripcion, tienda.calificacion)
		auxiliar = tienda.siguiente
	}
	auxiliar = lista.primero
	for i := 0; i < lista.elementos-1; i++ {
		fmt.Println(auxiliar, auxiliar.siguiente.anterior)
		auxiliar = auxiliar.siguiente
	}
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

func (lista *ListaTienda) InsertarUltimo(nombre string, descripcion string, contacto int, calificacion int) {
	nuevaTienda := &Tienda{nombre, descripcion, contacto, calificacion, nil, nil}
	lista.ultimo.siguiente = nuevaTienda
	nuevaTienda.anterior = lista.ultimo
	lista.ultimo = nuevaTienda
	lista.elementos += 1
}

func (lista *ListaTienda) InsertarPrimero(nombre string, descripcion string, contacto int, calificacion int) {
	nuevaTienda := &Tienda{nombre, descripcion, contacto, calificacion, nil, nil}
	lista.primero.anterior = nuevaTienda
	nuevaTienda.siguiente = lista.primero
	lista.primero = nuevaTienda
	lista.elementos += 1
}

func (lista *ListaTienda) InsertarMedio(nombre string, descripcion string, contacto int, calificacion int, dir1 *Tienda, dir2 *Tienda) {
	nuevaTienda := &Tienda{nombre, descripcion, contacto, calificacion, nil, nil}
	dir1.siguiente = nuevaTienda
	dir2.anterior = nuevaTienda
	nuevaTienda.siguiente = dir2
	nuevaTienda.anterior = dir1
	lista.elementos += 1
}

func (lista *ListaTienda) InsertarTienda(nombre string, descripcion string, contacto int, calificacion int) {
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
