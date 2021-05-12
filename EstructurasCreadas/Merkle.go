package EstructurasCreadas

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
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
		for i := 0; i < lista[i].elementos; i++ {
			nodo := NuevoNodoMerkHoja(auxiliar.nombre, auxiliar.departamento)
			arbol.Datos = append(arbol.Datos, nodo)
			arbol.Num += 1
			auxiliar = auxiliar.siguiente
		}
	}
}
