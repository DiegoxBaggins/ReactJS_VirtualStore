package EstructurasCreadas

import (
	"fmt"
	"strconv"
	"strings"
)

type Pedido struct {
	Pedidos []ListPedido
}

type ListPedido struct {
	Fecha        string  `json:"Fecha,omitempty"`
	Tienda       string  `json:"Tienda,omitempty"`
	Departamento string  `json:"Departamento,omitempty"`
	Calificacion float64 `json:"Calificacion,omitempty"`
	Productos    []ProdCod
}

type ProdCod struct {
	Codigo float64 `json:"Codigo,omitempty"`
}

func (pedido1 Pedido) ConstruirDatos(raiz *ArbolAnio) {
	listaInicial := pedido1.Pedidos
	for i := 0; i < len(listaInicial); i++ {
		objetoPedido := listaInicial[i]
		tienda := objetoPedido.Tienda
		dept := objetoPedido.Departamento
		calificacion := objetoPedido.Calificacion
		fecha := objetoPedido.Fecha
		for j := 0; j < len(objetoPedido.Productos); j++ {
			codigo := objetoPedido.Productos[j]
			raiz.IngresarPedido(codigo.Codigo, 1, tienda, calificacion, dept, fecha, raiz)
		}
	}
}

func (arbol *ArbolAnio) IngresarPedido(Codigo float64, Cantidad float64, Tienda string, Calificacion float64, Departamento string, Fecha string, raiz *ArbolAnio) {
	fecha := strings.Split(Fecha, "-")
	dia, _ := strconv.Atoi(fecha[0])
	mes, _ := strconv.Atoi(fecha[1])
	anio, _ := strconv.Atoi(fecha[2])
	anioNodo := raiz.buscarAnio(anio)
	if anioNodo == nil {
		raiz.Insertar(float64(anio))
		anioNodo = raiz.buscarAnio(anio)
	}
	mesNodo := anioNodo.meses.BuscarMes(mes)
	if mesNodo == nil {
		anioNodo.meses.InsertarMes(mes)
		mesNodo = anioNodo.meses.BuscarMes(mes)
	}
	pedidos := mesNodo.pedidos.BuscarNodoM(Departamento, dia)
	if pedidos == nil {
		mesNodo.pedidos.InsertarNodoM(dia, Departamento)
		pedidos = mesNodo.pedidos.BuscarNodoM(Departamento, dia)
	}
	pedidos.AgregarPedido(Codigo, Cantidad, Tienda, Calificacion, Departamento, Fecha)
}

type NodoAnio struct {
	anio       int
	equilibrio int
	hizq       *NodoAnio
	hder       *NodoAnio
	meses      *ListaMeses
}

func NuevoAnio(anio float64) *NodoAnio {
	return &NodoAnio{int(anio), 0, nil, nil, NewListaMes()}
}

type ArbolAnio struct {
	raiz *NodoAnio
}

func NewArAnios() *ArbolAnio {
	return &ArbolAnio{nil}
}

func (arbol *ArbolAnio) max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func (arbol *ArbolAnio) obtenerEQ(temp *NodoAnio) int {
	if temp == nil {
		return -1
	}
	return temp.equilibrio
}

func (arbol *ArbolAnio) rotacionIzquierda(temp *NodoAnio) *NodoAnio {
	aux := temp.hizq
	temp.hizq = aux.hder
	aux.hder = temp
	temp.equilibrio = arbol.max(arbol.obtenerEQ(temp.hder), arbol.obtenerEQ(temp.hizq)) + 1
	aux.equilibrio = arbol.max(arbol.obtenerEQ(temp.hizq), temp.equilibrio) + 1
	return aux
}

func (arbol *ArbolAnio) rotacionDerecha(temp *NodoAnio) *NodoAnio {
	aux := temp.hder
	temp.hder = aux.hizq
	aux.hizq = temp
	temp.equilibrio = arbol.max(arbol.obtenerEQ(temp.hder), arbol.obtenerEQ(temp.hizq)) + 1
	aux.equilibrio = arbol.max(arbol.obtenerEQ(temp.hder), temp.equilibrio) + 1
	return aux
}

func (arbol *ArbolAnio) rotacionDobleIzquierda(temp *NodoAnio) *NodoAnio {
	temp.hizq = arbol.rotacionDerecha(temp.hizq)
	return arbol.rotacionIzquierda(temp)
}

func (arbol *ArbolAnio) rotacionDobleDerecha(temp *NodoAnio) *NodoAnio {
	temp.hder = arbol.rotacionIzquierda(temp.hder)
	return arbol.rotacionDerecha(temp)
}

func (arbol *ArbolAnio) _Insertar(anio float64, raiz *NodoAnio) *NodoAnio {
	codigo := int(anio)
	if raiz == nil {
		return NuevoAnio(anio)
	} else if codigo < raiz.anio {
		raiz.hizq = arbol._Insertar(anio, raiz.hizq)
		if (arbol.obtenerEQ(raiz.hizq) - arbol.obtenerEQ(raiz.hder)) == -2 {
			if codigo < raiz.hizq.anio {
				raiz = arbol.rotacionIzquierda(raiz)
			} else {
				raiz = arbol.rotacionDobleIzquierda(raiz)
			}
		}
	} else if codigo > raiz.anio {
		raiz.hder = arbol._Insertar(anio, raiz.hder)
		if (arbol.obtenerEQ(raiz.hder) - arbol.obtenerEQ(raiz.hizq)) == 2 {
			if codigo > raiz.hder.anio {
				raiz = arbol.rotacionDerecha(raiz)
			} else {
				raiz = arbol.rotacionDobleDerecha(raiz)
			}
		}
	}
	raiz.equilibrio = arbol.max(arbol.obtenerEQ(raiz.hizq), arbol.obtenerEQ(raiz.hder)) + 1
	return raiz
}

func (arbol *ArbolAnio) Insertar(anio float64) {
	arbol.raiz = arbol._Insertar(anio, arbol.raiz)
}

func (arbol *ArbolAnio) Print() {
	arbol.inOrden(arbol.raiz)
}

func (arbol *ArbolAnio) inOrden(temp *NodoAnio) {
	if temp != nil {
		arbol.inOrden(temp.hizq)
		fmt.Println("Index: ", temp.anio)
		arbol.inOrden(temp.hder)
	}
}

func (arbol *ArbolAnio) buscarAnio(anio int) *NodoAnio {
	if arbol.raiz != nil {
		return arbol._buscarAnio(anio, arbol.raiz)
	}
	return nil
}

func (arbol *ArbolAnio) _buscarAnio(anio int, temp *NodoAnio) *NodoAnio {
	if anio == temp.anio {
		return temp
	} else {
		if anio < temp.anio {
			if temp.hizq != nil {
				return arbol._buscarAnio(anio, temp.hizq)
			}
		} else {
			if temp.hder != nil {
				return arbol._buscarAnio(anio, temp.hder)
			}
		}
	}
	return nil
}

// otros metodos del arbol
/*
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
*/

type NodoMes struct {
	codigo    int
	nombre    string
	anterior  *NodoMes
	siguiente *NodoMes
	pedidos   *CabeceraCen
}

func ConvMesNum(mes int) string {
	meses := ""
	switch mes {
	case 1:
		return "Enero"
	case 2:
		return "Febrero"
	case 3:
		return "Marzo"
	case 4:
		return "Abril"
	case 5:
		return "Mayo"
	case 6:
		return "Junio"
	case 7:
		return "Julio"
	case 8:
		return "Agosto"
	case 9:
		return "Septiembre"
	case 10:
		return "Octubre"
	case 11:
		return "Noviembre"
	case 12:
		return "Diciembre"
	}
	return meses
}

func NewMes(codigo int, nombre string) *NodoMes {
	return &NodoMes{codigo, nombre, nil, nil, NewCC()}
}

type ListaMeses struct {
	primero   *NodoMes
	ultimo    *NodoMes
	elementos int
}

func NewListaMes() *ListaMeses {
	return &ListaMeses{nil, nil, 0}
}

func (lista *ListaMeses) BuscarMes(codigo int) *NodoMes {
	auxiliar := lista.primero
	for i := 0; i < lista.elementos; i++ {
		if auxiliar.codigo == codigo {
			break
		}
		auxiliar = auxiliar.siguiente
	}
	return auxiliar
}

func (lista *ListaMeses) InsertarUltimo(mes int) {
	Mes := NewMes(mes, ConvMesNum(mes))
	lista.ultimo.siguiente = Mes
	Mes.anterior = lista.ultimo
	lista.ultimo = Mes
	lista.elementos += 1
}

func (lista *ListaMeses) InsertarPrimero(mes int) {
	Mes := NewMes(mes, ConvMesNum(mes))
	lista.primero.anterior = Mes
	Mes.siguiente = lista.primero
	lista.primero = Mes
	lista.elementos += 1
}

func (lista *ListaMeses) InsertarMedio(mes int, dir1 *NodoMes, dir2 *NodoMes) {
	Mes := NewMes(mes, ConvMesNum(mes))
	dir1.siguiente = Mes
	dir2.anterior = Mes
	Mes.siguiente = dir2
	Mes.anterior = dir1
	lista.elementos += 1
}

func (lista *ListaMeses) InsertarMes(mes int) *NodoMes {
	tamano := lista.elementos
	if tamano == 0 {
		Mes := NewMes(mes, ConvMesNum(mes))
		lista.ultimo = Mes
		lista.primero = Mes
		lista.elementos += 1
	} else {
		if tamano == 1 {
			primero := *lista.primero
			if primero.codigo > mes {
				lista.InsertarPrimero(mes)
			} else {
				lista.InsertarUltimo(mes)
			}
		} else {
			autorizacion := 0
			auxiliar := lista.primero
			for i := 1; i < lista.elementos; i++ {
				if auxiliar.codigo == mes {
					return auxiliar
				}
				if auxiliar.codigo > mes {
					if auxiliar == lista.primero {
						lista.InsertarPrimero(mes)
						autorizacion = 1
						break
					} else {
						lista.InsertarMedio(mes, auxiliar.anterior, auxiliar)
						autorizacion = 1
						break
					}
				} else {
					auxiliar = auxiliar.siguiente
				}
			}
			if autorizacion == 0 {
				if auxiliar.codigo == mes {
					return auxiliar
				} else {
					lista.InsertarUltimo(mes)
				}
			}
		}
	}
	return nil
}

type CabeceraCen struct {
	hor  *CabeceraH
	ver  *CabeceraV
	numh int
	numv int
}

func NewCC() *CabeceraCen {
	return &CabeceraCen{nil, nil, 0, 0}
}

type CabeceraH struct {
	nombre    string
	elementos int
	arriba    *CabeceraH
	abajo     *CabeceraH
	derecho   *CuerpoM
}

func NewCH(nombre string) *CabeceraH {
	return &CabeceraH{nombre, 0, nil, nil, nil}
}

type CabeceraV struct {
	dia       int
	elementos int
	izquierda *CabeceraV
	derecha   *CabeceraV
	abajo     *CuerpoM
}

func NewCV(dia int) *CabeceraV {
	return &CabeceraV{dia, 0, nil, nil, nil}
}

type CuerpoM struct {
	nombre    string
	dia       int
	arriba    *CuerpoM
	abajo     *CuerpoM
	izquierda *CuerpoM
	derecha   *CuerpoM
	pedidos   []ProductPedido
}

func BuscarCabeceraH(nombre string, cab *CabeceraH) *CabeceraH {
	if cab != nil {
		if nombre == cab.nombre {
			return cab
		} else {
			if cab.abajo != nil {
				return BuscarCabeceraH(nombre, cab.abajo)
			} else {
				return nil
			}
		}
	} else {
		return nil
	}
}

func (central *CabeceraCen) NuevaCH(nombre string) {
	if central.numh == 0 {
		cabN := NewCH(nombre)
		central.hor = cabN
	} else if central.numh == 1 {
		if central.hor.nombre < nombre {
			cabN := NewCH(nombre)
			central.hor.abajo = cabN
		} else {
			cabN := NewCH(nombre)
			cabN.abajo = central.hor
			central.hor.arriba = cabN
			central.hor = cabN
		}
	} else {
		central.InsertarCabeceraH(nombre, central.hor)
	}
	central.numh += 1
}

func (central *CabeceraCen) InsertarCabeceraH(nombre string, cab *CabeceraH) {
	if cab.nombre > nombre {
		if cab.arriba == nil {
			cabN := NewCH(nombre)
			cab.arriba = cabN
			cabN.abajo = cab
			central.hor = cabN
		} else {
			cabN := NewCH(nombre)
			cabN.arriba = cab.arriba
			cabN.abajo = cab
			cab.arriba.abajo = cabN
			cab.arriba = cabN
		}
	} else {
		if cab.abajo == nil {
			cabN := NewCH(nombre)
			cab.abajo = cabN
			cabN.arriba = cab
		} else {
			central.InsertarCabeceraH(nombre, cab.abajo)
		}
	}
}

func BuscarCabeceraV(dia int, cab *CabeceraV) *CabeceraV {
	if cab != nil {
		if dia == cab.dia {
			return cab
		} else {
			if cab.derecha != nil {
				return BuscarCabeceraV(dia, cab.derecha)
			} else {
				return nil
			}
		}
	} else {
		return nil
	}
}

func (central *CabeceraCen) NuevaCV(dia int) {
	if central.numv == 0 {
		cabN := NewCV(dia)
		central.ver = cabN
	} else if central.numv == 1 {
		if central.ver.dia < dia {
			cabN := NewCV(dia)
			central.ver.derecha = cabN
		} else {
			cabN := NewCV(dia)
			cabN.derecha = central.ver
			central.ver.izquierda = cabN
			central.ver = cabN
		}
	} else {
		central.InsertarCabeceraV(dia, central.ver)
	}
	central.numv += 1
}

func (central *CabeceraCen) InsertarCabeceraV(dia int, cab *CabeceraV) {
	if cab.dia > dia {
		if cab.izquierda == nil {
			cabN := NewCV(dia)
			cab.izquierda = cabN
			cabN.derecha = cab
			central.ver = cabN
		} else {
			cabN := NewCV(dia)
			cabN.izquierda = cab.izquierda
			cabN.derecha = cab
			cab.izquierda.derecha = cabN
			cab.izquierda = cabN
		}
	} else {
		if cab.derecha == nil {
			cabN := NewCV(dia)
			cab.derecha = cabN
			cabN.izquierda = cab
		} else {
			central.InsertarCabeceraV(dia, cab.derecha)
		}
	}
}

func (central *CabeceraCen) InsertarNodoM(dia int, nombre string) {
	cabhor := BuscarCabeceraH(nombre, central.hor)
	cabver := BuscarCabeceraV(dia, central.ver)
	if cabhor == nil {
		central.NuevaCH(nombre)
		cabhor = BuscarCabeceraH(nombre, central.hor)
	}
	if cabver == nil {
		central.NuevaCV(dia)
		cabver = BuscarCabeceraV(dia, central.ver)
	}
	aux := central.BuscarNodoM(nombre, dia)
	if aux == nil {
		arreglo1 := make([]ProductPedido, 0)
		nodoC := &CuerpoM{nombre, dia, nil, nil, nil, nil, arreglo1}
		cabhor.BuscarEspacioH(nombre, nodoC)
		cabver.BuscarEspacioV(dia, nodoC)
	}
}

func (cab1 *CabeceraH) BuscarEspacioH(nombre string, m *CuerpoM) {
	if cab1.elementos == 0 {
		cab1.derecho = m
		cab1.elementos += 1
	} else if cab1.elementos == 1 {
		if cab1.derecho.nombre < nombre {
			cab1.derecho.derecha = m
			m.izquierda = cab1.derecho
		} else {
			cab1.derecho.izquierda = m
			m.derecha = cab1.derecho
			cab1.derecho = m
		}
		cab1.elementos += 1
	} else {
		aux := cab1.derecho
		if aux.nombre > nombre {
			aux.izquierda = m
			m.derecha = aux
			cab1.derecho = m
		} else {
			cab1._BuscarEspacioH(m, aux.derecha)
		}
		cab1.elementos += 1
	}
}

func (cab1 *CabeceraH) _BuscarEspacioH(m *CuerpoM, cab *CuerpoM) {
	if cab.dia > m.dia {
		if cab.izquierda == nil {
			cab.izquierda = m
			m.derecha = cab
			cab1.derecho = m
		} else {
			m.izquierda = cab.izquierda
			m.derecha = cab
			cab.izquierda.derecha = m
			cab.izquierda = m
		}
	} else {
		if cab.derecha == nil {
			cab.derecha = m
			m.izquierda = cab
		} else {
			cab1._BuscarEspacioH(m, cab.derecha)
		}
	}
}

func (cab1 *CabeceraV) BuscarEspacioV(dia int, m *CuerpoM) {
	if cab1.elementos == 0 {
		cab1.abajo = m
		cab1.elementos += 1
	} else if cab1.elementos == 1 {
		if cab1.abajo.dia < dia {
			cab1.abajo.abajo = m
			m.arriba = cab1.abajo
		} else {
			cab1.abajo.arriba = m
			m.abajo = cab1.abajo
			cab1.abajo = m
		}
		cab1.elementos += 1
	} else {
		aux := cab1.abajo
		if aux.dia > dia {
			aux.izquierda = m
			m.abajo = aux
			cab1.abajo = m
		} else {
			cab1._BuscarEspacioV(m, aux.abajo)
		}
		cab1.elementos += 1
	}
}

func (cab1 *CabeceraV) _BuscarEspacioV(m *CuerpoM, cab *CuerpoM) {
	if cab.nombre > m.nombre {
		if cab.abajo == nil {
			cab.abajo = m
			m.arriba = cab
			cab1.abajo = m
		} else {
			m.arriba = cab.arriba
			m.abajo = cab
			cab.arriba.abajo = m
			cab.arriba = m
		}
	} else {
		if cab.abajo == nil {
			cab.abajo = m
			m.arriba = cab
		} else {
			cab1._BuscarEspacioV(m, cab.abajo)
		}
	}
}

func (central *CabeceraCen) BuscarNodoM(nombre string, dia int) *CuerpoM {
	cabhor := BuscarCabeceraH(nombre, central.hor)
	cabver := BuscarCabeceraV(dia, central.ver)
	if cabhor == nil || cabver == nil {
		return nil
	} else {
		aux := cabhor.derecho
		for i := 0; i < cabhor.elementos; i++ {
			if aux.nombre == nombre && aux.dia == dia {
				break
			}
			aux = aux.derecha
		}
		return aux
	}
}

type ProductPedido struct {
	Codigo       float64 `json:"Codigo,omitempty"`
	Cantidad     float64 `json:"Cantidad,omitempty"`
	Tienda       string  `json:"Tienda,omitempty"`
	Calificacion float64 `json:"Calificacion,omitempty"`
	Departamento string  `json:"Departamento,omitempty"`
	Fecha        string  `json:"Fecha,omitempty"`
}

func (nodo *CuerpoM) AgregarPedido(cod float64, cant float64, tienda string, cal float64, dept string, fecha string) {
	for i := 0; i < len(nodo.pedidos); i++ {
		if nodo.pedidos[i].Codigo == cod && nodo.pedidos[i].Tienda == tienda {
			nodo.pedidos[i].Cantidad += cant
			return
		}
	}
	producto := ProductPedido{cod, cant, tienda, cal, dept, fecha}
	arreglo1 := make([]ProductPedido, 1)
	arreglo1[0] = producto
	nodo.pedidos = append(nodo.pedidos, arreglo1...)
}
