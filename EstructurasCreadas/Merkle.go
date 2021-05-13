package EstructurasCreadas

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type NodoMerk struct {
	Id    string
	IdIzq string
	IdDer string
	User  string
	Fecha string
	Total string
	Hoja  bool
	hizq  *NodoMerk
	hder  *NodoMerk
}

type ArbolMerk struct {
	Raiz  *NodoMerk
	Nivel int
	Datos []*NodoMerk
	Num   int
}

func NuevoNodoMerkHoja(user string, total string) *NodoMerk {
	currentTime := time.Now()
	str := user + total + currentTime.Format("2006-01-02 15:04:05.000000000")
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	return &NodoMerk{id, "", "", user, currentTime.Format("2006-01-02 15:04:05.000000000"), total, true, nil, nil}
}

func NuevoNodoMerkCuerpo(idizq string, idder string, hizq *NodoMerk, hder *NodoMerk) *NodoMerk {
	currentTime := time.Now()
	str := idizq + idder
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	return &NodoMerk{id, idizq, idder, "", currentTime.String(), "", false, hizq, hder}
}

func NuevoMerk() ArbolMerk {
	return ArbolMerk{NuevoNodoMerkHoja("", ""), 0, make([]*NodoMerk, 0), 0}
}

func (arbol *ArbolMerk) AgregarNodos(user string, total string) {
	nodo := NuevoNodoMerkHoja(user, total)
	arbol.Datos = append(arbol.Datos, nodo)
	arbol.Num += 1
}

func (arbol *ArbolMerk) CrearArbol() {
	valores := nextPowerOf2(arbol.Num)
	slicePivote := arbol.Datos
	for i := arbol.Num; i < valores; i++ {
		slicePivote = append(slicePivote, NuevoNodoMerkHoja("-1", "0"))
	}
	if valores != 0 {
		pivote := valores
		slice1 := slicePivote
		for pivote > 1 {
			pivote = pivote / 2
			slice1 = arbol.CrearNivel(slice1, pivote)
		}
		arbol.Raiz = slice1[0]
		arbol.Nivel += 1
	}
}

func (arbol *ArbolMerk) CrearNivel(slice1 []*NodoMerk, tamanio int) []*NodoMerk {
	arbol.Nivel += 1
	slice2 := make([]*NodoMerk, 0)
	j := 0
	for i := 0; i < tamanio; i++ {
		slice2 = append(slice2, NuevoNodoMerkCuerpo(slice1[j].Id, slice1[j+1].Id, slice1[j], slice1[j+1]))
		j += 2
	}
	return slice2
}

func nextPowerOf2(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func (arbol *ArbolMerk) GraficarGrafo() {
	arbol.CrearArbol()
	direct := "./react-server/reactserver/src/assets/images/grafos/merkle/"
	var graph = "digraph G{\n"
	graph += "rankdir=TB\n node[shape=box]\nconcentrate=true\n"
	str1, _ := arbol._GraficarGrafo(arbol.Raiz, 0)
	graph += str1
	graph += "\n}"
	data := []byte(graph)
	err := ioutil.WriteFile(direct+"transacciones.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", direct+"transacciones.dot").Output()
	mode := int(0777)
	err = ioutil.WriteFile(direct+"transacciones.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
}

func (arbol *ArbolMerk) _GraficarGrafo(temp *NodoMerk, nods int) (string, int) {
	grafo := ""
	pivote := nods
	str1 := ""
	if temp != nil {
		grafo += "Nodo" + strconv.Itoa(nods) + "[shape=record label=\"" + temp._GraficarGrafo() + "\"]\n"
		nods += 1
		if temp.Hoja == false {
			grafo += "Nodo" + strconv.Itoa(pivote) + "->Nodo" + strconv.Itoa(nods) + ";\n"
			str1, nods = arbol._GraficarGrafo(temp.hizq, nods)
			grafo += str1
			grafo += "Nodo" + strconv.Itoa(pivote) + "->Nodo" + strconv.Itoa(nods) + ";\n"
			str1, nods = arbol._GraficarGrafo(temp.hder, nods)
			grafo += str1
		}
	}
	return grafo, nods
}

func (nodo *NodoMerk) _GraficarGrafo() string {
	graph := ""
	if nodo.Hoja == true {
		graph += "{Id:" + nodo.Id
		graph += "| Usuario: " + nodo.User
		graph += "| Total: " + nodo.Total
		graph += "| Fecha: " + nodo.Fecha + "}"
	} else {
		graph += "{Id:" + nodo.Id
		graph += "| IdIzq: " + nodo.IdIzq
		graph += "| IdDer: " + nodo.IdDer + "}"
	}
	return graph
}

func (arbol *ArbolMerk) AgregarNodosTiendas(lista []ListaTienda) {
	for i := 0; i < len(lista); i++ {
		auxiliar := lista[i].primero
		for j := 0; j < lista[i].elementos; j++ {
			nodo := NuevoNodoMerkHoja(auxiliar.nombre, auxiliar.departamento)
			arbol.Datos = append(arbol.Datos, nodo)
			arbol.Num += 1
			auxiliar = auxiliar.siguiente
		}
	}
}

func (arbol *ArbolMerk) GraficarGrafoTiendas() {
	arbol.CrearArbol()
	direct := "./react-server/reactserver/src/assets/images/grafos/merkle/"
	var graph = "digraph G{\n"
	graph += "rankdir=TB\n node[shape=box]\nconcentrate=true\n"
	str1, _ := arbol._GraficarGrafo(arbol.Raiz, 0)
	graph += str1
	graph += "\n}"
	data := []byte(graph)
	err := ioutil.WriteFile(direct+"tiendas.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", direct+"tiendas.dot").Output()
	mode := int(0777)
	err = ioutil.WriteFile(direct+"tiendas.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
}

func (arbol *ArbolMerk) AgregarNodosProductos(lista []ListaTienda) {
	for i := 0; i < len(lista); i++ {
		auxiliar := lista[i].primero
		for j := 0; j < lista[i].elementos; j++ {
			raiz := auxiliar.productos
			raiz.AgregarNodoProductosMerkle(raiz.raiz, arbol)
			auxiliar = auxiliar.siguiente
		}
	}
}

func (arbol *ArbolProd) AgregarNodoProductosMerkle(aux *Producto, merkle *ArbolMerk) {
	nodo := NuevoNodoMerkHoja(aux.nombre, strconv.Itoa(int(aux.precio)))
	merkle.Datos = append(merkle.Datos, nodo)
	merkle.Num += 1
	if aux.hizq != nil {
		arbol.AgregarNodoProductosMerkle(aux.hizq, merkle)
	}
	if aux.hder != nil {
		arbol.AgregarNodoProductosMerkle(aux.hder, merkle)
	}
}

func (arbol *ArbolMerk) GraficarGrafoProductos() {
	arbol.CrearArbol()
	direct := "./react-server/reactserver/src/assets/images/grafos/merkle/"
	var graph = "digraph G{\n"
	graph += "rankdir=TB\n node[shape=box]\nconcentrate=true\n"
	str1, _ := arbol._GraficarGrafo(arbol.Raiz, 0)
	graph += str1
	graph += "\n}"
	data := []byte(graph)
	err := ioutil.WriteFile(direct+"productos.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", direct+"productos.dot").Output()
	mode := int(0777)
	err = ioutil.WriteFile(direct+"productos.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
}

func (arbol *ArbolMerk) AgregarNodosUsuarios(lista []Usuario) {
	for i := 0; i < len(lista); i++ {
		nodo := NuevoNodoMerkHoja(strconv.Itoa(lista[i].Dpi), lista[i].Nombre)
		arbol.Datos = append(arbol.Datos, nodo)
		arbol.Num += 1
	}
}

func (arbol *ArbolMerk) GraficarGrafoUsers() {
	arbol.CrearArbol()
	direct := "./react-server/reactserver/src/assets/images/grafos/merkle/"
	var graph = "digraph G{\n"
	graph += "rankdir=TB\n node[shape=box]\nconcentrate=true\n"
	str1, _ := arbol._GraficarGrafo(arbol.Raiz, 0)
	graph += str1
	graph += "\n}"
	data := []byte(graph)
	err := ioutil.WriteFile(direct+"usuarios.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", direct+"usuarios.dot").Output()
	mode := int(0777)
	err = ioutil.WriteFile(direct+"usuarios.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
}

type BlockChain struct {
	Indice       int
	Fecha        string
	Data         string
	Nonce        int
	PreviousHash string
	HashActual   string
}

func NuevoBlock(indice int, data string, previo string) BlockChain {
	currentTime := time.Now()
	random := rand.Intn(10000)
	str := strconv.Itoa(indice) + currentTime.Format("2006-01-02 15:04:05") + data + previo + strconv.Itoa(random)
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	return BlockChain{indice, currentTime.Format("2006-01-02 15:04:05"), data, random, previo, id}
}
