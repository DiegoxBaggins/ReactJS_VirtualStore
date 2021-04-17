package EstructurasCreadas

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type Grafo struct {
	Nodos                []NodoGrafo
	PosicionInicialRobot string `json:"PosicionInicialRobot,omitempty"`
	Entrega              string `json:"Entrega,omitempty"`
	Matriz               MatrizGrafo
}

type NodoGrafo struct {
	Nombre  string `json:"Nombre,omitempty"`
	Enlaces []EnlaceGrafo
}

type EnlaceGrafo struct {
	Nombre    string  `json:"Nombre,omitempty"`
	Distancia float64 `json:"Distancia,omitempty"`
}

type MatrizGrafo struct {
}

func (grafo *Grafo) GraficarGrafo() {
	direct := "./react-server/reactserver/src/assets/images/grafos/paquete/"
	var graph = "graph G{\n"
	graph += "edge [weight=1000]\n"
	graph += grafo._Graficar()
	graph += "\n}"
	data := []byte(graph)
	err := ioutil.WriteFile(direct+"paquete.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", direct+"paquete.dot").Output()
	mode := 0777
	err = ioutil.WriteFile(direct+"paquete.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
}

func (grafo *Grafo) _Graficar() string {
	graph := ""
	num := 0
	for i := 0; i < len(grafo.Nodos); i++ {
		graph += "Nodo" + strconv.Itoa(i) + "[label=\"" + grafo.Nodos[i].Nombre
		graph += "\""
		if grafo.Nodos[i].Nombre == grafo.PosicionInicialRobot {
			graph += " color=\"green\""
		} else if grafo.Nodos[i].Nombre == grafo.Entrega {
			graph += " color=\"red\""
		}
		graph += "]\n"
	}
	for i := 0; i < len(grafo.Nodos); i++ {
		for j := 0; j < len(grafo.Nodos[i].Enlaces); j++ {
			num = int(grafo.Nodos[i].Enlaces[j].Distancia)
			graph += "Nodo" + strconv.Itoa(i) + "--Nodo" + grafo._BuscarNodo(grafo.Nodos[i].Enlaces[j].Nombre) + "[label=\"" + strconv.Itoa(num) + "\"]\n"
		}
	}
	return graph
}

func (grafo *Grafo) _BuscarNodo(nombre string) string{
	for i := 0; i < len(grafo.Nodos); i++ {
		if nombre == grafo.Nodos[i].Nombre {
			return strconv.Itoa(i)
		}
	}
	return ""
}