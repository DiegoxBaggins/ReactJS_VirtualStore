package main

import (
	"Proyecto-moduls/EstructurasCreadas"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var Datos EstructurasCreadas.Data
var Vector []EstructurasCreadas.ListaTienda
var Inventario EstructurasCreadas.Invent
var TamaVec = 0
var Indices []string
var Departamentos []string

func main() {
	//enrutador denominado router
	router := mux.NewRouter()
	arbol := EstructurasCreadas.NewArbol()
	pruebaArbol(2, arbol)
	pruebaArbol(3, arbol)
	pruebaArbol(4, arbol)
	pruebaArbol(5, arbol)
	pruebaArbol(6, arbol)
	pruebaArbol(1, arbol)
	pruebaArbol(10, arbol)
	pruebaArbol(15, arbol)
	pruebaArbol(25, arbol)
	pruebaArbol(1, arbol)
	pruebaArbol(6, arbol)
	pruebaArbol(15, arbol)
	pruebaArbol(15, arbol)
	arbol.Print()

	//Endpoints
	router.HandleFunc("/cargartienda", CargarTienda).Methods("POST")
	router.HandleFunc("/id/{id:[0-9]+}", BusquedaPosicion).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	router.HandleFunc("/Eliminar", EliminarTienda).Methods("DELETE")
	router.HandleFunc("/guardar", GuardarDatos).Methods("GET")
	router.HandleFunc("/getArreglo", CrearGrafo).Methods("GET")
	router.HandleFunc("/cargarInv", CargarInven).Methods("POST")
	router.HandleFunc("/tiendas", DevolverTiendas).Methods("GET")
	// iniciar el servidor en el puerto 3000
	log.Fatal(http.ListenAndServe(":3000", router))
}

func CargarInven(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Esto es una peticion de tipo post")
	_ = json.NewDecoder(req.Body).Decode(&Inventario)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	_ = encoder.Encode("Datos cargados con exito")
	fmt.Println(Inventario)
	Inventario.SacarInventario(Vector, Indices, Departamentos)
	inventario := Vector
	fmt.Println(inventario)
	TamaVec = len(Vector)
	Indices, Departamentos = Datos.CalcularTamanos()
}

func CargarTienda(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Esto es una peticion de tipo post")
	_ = json.NewDecoder(req.Body).Decode(&Datos)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	_ = encoder.Encode("Datos cargados con exito")
	fmt.Println(Datos)
	Vector = Datos.TransformarDatos()
	fmt.Println(Vector)
	TamaVec = len(Vector)
	Indices, Departamentos = Datos.CalcularTamanos()
}

func TiendaEspecifica(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo post")
		var tienda EstructurasCreadas.StoreBus
		_ = json.NewDecoder(req.Body).Decode(&tienda)
		fmt.Println(tienda)
		first1 := tienda.Nombre[0:1]
		fmt.Println(first1)
		if tienda.Calificacion > 5 || tienda.Calificacion < 0 {
			_ = encoder.Encode("Calificacion no valida")
		} else {
			indice, dept, err1 := EncontrarIndices(tienda.Departamento, first1)
			if err1 == 1 {
				_ = encoder.Encode("El Departamento no existe")
			} else {
				elemento := int(tienda.Calificacion) + 5*(indice+(len(Indices)*dept)) - 1
				store, err := Vector[elemento].BuscarStore(tienda.Nombre)
				if err == 0 {
					_ = encoder.Encode("Tienda no existe")
				} else {
					_ = encoder.Encode(store)
				}
			}
		}
	}
}

func EliminarTienda(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo post")
		var tienda EstructurasCreadas.StoreElim
		_ = json.NewDecoder(req.Body).Decode(&tienda)
		fmt.Println(tienda)
		first1 := tienda.Nombre[0:1]
		fmt.Println(first1)
		if tienda.Calificacion > 5 || tienda.Calificacion < 0 {
			_ = encoder.Encode("Calificacion no valida")
		} else {
			indice, dept, err1 := EncontrarIndices(tienda.Departamento, first1)
			if err1 == 1 {
				_ = encoder.Encode("El Departamento no existe")
			} else {
				elemento := int(tienda.Calificacion) + 5*(indice+(len(Indices)*dept)) - 1
				store, err := Vector[elemento].BuscarStore(tienda.Nombre)
				if err == 0 {
					_ = encoder.Encode("Tienda no existe")
				} else {
					_ = encoder.Encode(store)
					Vector[elemento].EliminarTienda(tienda.Nombre)
				}
			}
		}
	}
}

func EncontrarIndices(dept string, nombre string) (int, int, int) {
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

func BusquedaPosicion(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo get")
		vars := mux.Vars(req)
		id, _ := strconv.Atoi(vars["id"])
		if id >= TamaVec {
			_ = encoder.Encode("El indice supera el tamaño de la matriz")
		} else {
			var lista = Vector[id].VectorElementos()
			_ = encoder.Encode(lista)
		}
	}
}

func GuardarDatos(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	fmt.Println(Indices)
	fmt.Println(Departamentos)
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		var Matrix EstructurasCreadas.Data
		Matrix.RegresarMatriz(Vector, Indices, Departamentos)
		fmt.Println("Esto es una peticion de tipo get")
		_ = encoder.Encode(Matrix)
		crearJson, _ := json.MarshalIndent(Matrix, "", "    ")
		err := ioutil.WriteFile("Datos.json", crearJson, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CrearGrafo(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		var graph = "digraph List {\n"
		graph += "rankdir=TB;\n"
		graph += "node [shape = record, color=black, style=filled, fillcolor=yellow];\n subgraph {\n"
		rank := "{rank = same; "
		var nodes = ""
		var pointers = ""
		for arreglo := 0; arreglo < len(Vector); arreglo++ {
			nodes += "Node" + strconv.Itoa(arreglo) + "[label=\"" + strconv.Itoa(arreglo) + "\"]\n"
			rank += "Node" + strconv.Itoa(arreglo) + "; "
			if (arreglo + 1) != len(Vector) {
				pointers += "Node" + strconv.Itoa(arreglo) + "->Node" + strconv.Itoa(arreglo+1) + ";\n"
			}
		}
		rank += "}"
		graph += rank + "\n" + nodes + "\n" + pointers
		numero := 0
		for arreglo := 0; arreglo < len(Vector); arreglo++ {
			stringLista, num := Vector[arreglo].GraficarLista(arreglo, numero)
			graph += stringLista
			numero = num
		}
		graph += "\n}\n}"
		//fmt.Println(graph)
		data := []byte(graph)
		err := ioutil.WriteFile("graph.dot", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpdf", "graph.dot").Output()
		mode := int(0777)
		err = ioutil.WriteFile("Arreglo.pdf", cmd, os.FileMode(mode))
		if err != nil {
			log.Fatal(err)
		}
		_ = encoder.Encode("Grafico con exito")
	}
}

func DevolverTiendas(w http.ResponseWriter, req *http.Request){
	setupCorsResponse(&w, req)
	arregloTiendas := make([]EstructurasCreadas.Store, 0)
	for i:= 0; i< len(Vector); i ++ {
		arregloTiendas = append(arregloTiendas, Vector[i].ReturnListStore()...)
	}
	for i:= 0; i< len(arregloTiendas); i ++ {
		fmt.Println(arregloTiendas[i].Nombre)
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo get")
		_ = encoder.Encode(arregloTiendas)
	}
}

func LinealizarRM(valor int, matriz [10][26][5]EstructurasCreadas.ListaTienda) [3000]EstructurasCreadas.ListaTienda {
	var arreglo [3000]EstructurasCreadas.ListaTienda
	for fila := 0; fila < valor; fila++ {
		for columna := 0; columna < 26; columna++ {
			for cara := 0; cara < 5; cara++ {
				elemento := 0
				elemento = cara + 5*(columna+(26*fila))
				arreglo[elemento] = matriz[fila][columna][cara]
			}
		}
	}
	for objeto := 0; objeto < len(arreglo); objeto++ {
		fmt.Println(arreglo[objeto])
	}
	return arreglo
}

func pruebaArbol(codigo float64, arbol *EstructurasCreadas.ArbolProd) {
	product := EstructurasCreadas.Product{"h", codigo, "h", 10, 10, "h"}
	arbol.Insertar(product)
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}