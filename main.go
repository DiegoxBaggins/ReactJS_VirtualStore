package main

import (
	"Proyecto-moduls/EstructurasCreadas"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var Datos EstructurasCreadas.Data
var Vector []EstructurasCreadas.ListaTienda
var TamaVec = 0

func main() {
	//ejemplo()
	//enrutador denominado router
	router := mux.NewRouter()

	//Endpoints
	//router.HandleFunc("/getHello", HelloWorld).Methods("GET")
	router.HandleFunc("/cargartienda", CargarTienda).Methods("POST")
	router.HandleFunc("/id/{id:[0-9]+}", BusquedaPosicion).Methods("GET")
	// iniciar el servidor en el puerto 7000
	log.Fatal(http.ListenAndServe(":7000", router))
}

/*
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Esto es una peticion de tipo get")
	var matriz = Datos.TransformarDatos()
	fmt.Println(matriz)
	matriz[4].MostrarDatos()
	matriz[3].MostrarDatos()
}
*/

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

/*
func alter() {
	var clasificacion [10][26][5]EstructurasCreadas.ListaTienda
	//fmt.Println(clasificacion)
	//fmt.Println(len(clasificacion))

	for i := 0; i <27; i++ {
		for j := 0; j <10; j++ {
			for k := 0; k <5; k++ {
				fmt.Println(clasificacion[i][j][k])
			}
		}
	}

	//clasificacion[0][0][0].InsertarTienda("Genetik", "Tienda en Linea",41283319,5)
	//clasificacion[0][0][0].InsertarTienda("Genetiks", "Tienda en Linea",41283319,5)
	//clasificacion[0][0][0].InsertarTienda("Genetikss", "Tienda en Linea",41283319,5)
	//fmt.Println("lista", clasificacion[0][0][0])
	//clasificacion[0][0][0].MostrarDatos()
	clasificacion[0][0][0].InsertarTienda("Genetik", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("Dressy", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("dressy", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][2].InsertarTienda("foes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][3].InsertarTienda("foes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][4].InsertarTienda("foes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][1][1].InsertarTienda("Armani", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("Farts", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("qoes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("fies", "Tienda en Linea", 41283319, 5)
	fmt.Println("lista", clasificacion[0][0][1])
	clasificacion[0][0][1].MostrarDatos()
	LinealizarRM(10, clasificacion)
	var arreglo = LinealizarRM(10, clasificacion)
	fmt.Println(arreglo)
}
*/

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

func ejemplo() {
	type Raza struct {
		Nombre, Pais string
	}

	type Mascota struct {
		Nombre string
		Edad   int
		Raza   Raza
		Amigos []string // Arreglo de strings
	}

	// Vamos a probar...
	mascotaComoJson := []byte(`{"Nombre":"Maggie","Edad":3,"Raza":{"Nombre":"Caniche","Pais":"Francia"},"Amigos":["Bichi","Snowball","Coqueta","Cuco","Golondrino"]}`)

	// Recuerda, primero se define la variable
	var mascota Mascota

	// Y luego se manda su dirección de memoria
	err := json.Unmarshal(mascotaComoJson, &mascota)
	if err != nil {
		fmt.Printf("Error decodificando: %v\n", err)
	} else {
		// Listo. Ahora podemos imprimir
		fmt.Printf("El nombre: %s\n", mascota.Nombre)
		fmt.Printf("País de Raza: %s\n", mascota.Raza.Pais)
		fmt.Printf("Primer amigo: %v\n", mascota.Amigos[0])
	}
}
