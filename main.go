package main

import (
	"Proyecto-moduls/EstructurasCreadas"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var Datos EstructurasCreadas.Data
var Vector []EstructurasCreadas.ListaTienda
var TamaVec = 0
var Indices []string
var Departamentos []string

func main() {
	//enrutador denominado router
	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/cargartienda", CargarTienda).Methods("POST")
	router.HandleFunc("/id/{id:[0-9]+}", BusquedaPosicion).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	router.HandleFunc("/Eliminar", EliminarTienda).Methods("DELETE")
	router.HandleFunc("/guardar", GuardarDatos).Methods("GET")
	// iniciar el servidor en el puerto 3000
	log.Fatal(http.ListenAndServe(":3000", router))
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
				store, err := Vector[elemento].BuscarTienda(tienda.Nombre)
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
				store, err := Vector[elemento].BuscarTienda(tienda.Nombre)
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
