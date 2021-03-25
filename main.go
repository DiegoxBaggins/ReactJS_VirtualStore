package main

import (
	"Proyecto-moduls/EstructurasCreadas"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Datos EstructurasCreadas.Data
var Vector []EstructurasCreadas.ListaTienda
var Inventario EstructurasCreadas.Invent
var TamaVec = 0
var Indices []string
var Departamentos []string
var Carrito []EstructurasCreadas.ProductCarr
var JsonPedidos EstructurasCreadas.Pedido
var Pedidos EstructurasCreadas.ArbolAnio

func main() {
	pruebaMatriz()
	//enrutador denominado router
	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/cargartienda", CargarTienda).Methods("POST")
	router.HandleFunc("/id/{id:[0-9]+}", BusquedaPosicion).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	router.HandleFunc("/Eliminar", EliminarTienda).Methods("DELETE")
	router.HandleFunc("/guardar", GuardarDatos).Methods("GET")
	router.HandleFunc("/getArreglo", CrearGrafo).Methods("GET")
	router.HandleFunc("/cargarInv", CargarInven).Methods("POST")

	router.HandleFunc("/tiendas", DevolverTiendasAPI).Methods("GET")
	router.HandleFunc("/cargartiendas", CargarTiendasAPI).Methods("POST")
	router.HandleFunc("/cargarinventario", CargarInvenAPI).Methods("POST")
	router.HandleFunc("/productos/{dept}/{cal}/{nom}", DevolverInventarioAPI).Methods("GET")
	router.HandleFunc("/agregarCarrito", AgregarAlCarritoAPI).Methods("POST")
	router.HandleFunc("/carrito", ReturnCarritoAPI).Methods("GET")
	router.HandleFunc("/eliminarCarrito", EliminarDelCarritoAPI).Methods("POST")
	router.HandleFunc("/PagarCarrito", PagarCarritoAPI).Methods("GET")
	router.HandleFunc("/cargarpedidos", CargarPedidosAPI).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	// iniciar el servidor en el puerto 3000
	log.Fatal(http.ListenAndServe(":3000", handler))
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

func CargarInven(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo post")
		_ = json.NewDecoder(req.Body).Decode(&Inventario)
		_ = encoder.Encode("Datos cargados con exito")
		fmt.Println(Inventario)
		Inventario.SacarInventario(Vector, Indices, Departamentos)
		inventario := Vector
		fmt.Println(inventario)
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

func pruebaMatriz() {
	matriz := EstructurasCreadas.NewCC()
	matriz.InsertarNodoM(5, "C")
	matriz.InsertarNodoM(1, "C")
	matriz.InsertarNodoM(2, "C")
	matriz.InsertarNodoM(4, "A")
	matriz.InsertarNodoM(3, "B")
	matriz.InsertarNodoM(11, "J")
	matriz.InsertarNodoM(2, "B")
	hol := matriz
	fmt.Println(hol)
}

/*
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
*/

// metodos de la api rest

func DevolverTiendasAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	arregloTiendas := make([]EstructurasCreadas.StoreFront, 0)
	for i := 0; i < len(Vector); i++ {
		arregloTiendas = append(arregloTiendas, Vector[i].ReturnListStore()...)
	}
	for i := 0; i < len(arregloTiendas); i++ {
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

func CargarTiendasAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	fmt.Println("Esto es una peticion de tipo post")
	var buf bytes.Buffer
	file, header, err := req.FormFile("myFile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	io.Copy(&buf, file)
	contents := buf.String()
	fmt.Println(contents)
	_ = json.Unmarshal(buf.Bytes(), &Datos)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	_ = encoder.Encode("Datos cargados con exito")
	fmt.Println(Datos)
	Vector = Datos.TransformarDatos()
	fmt.Println(Vector)
	TamaVec = len(Vector)
	Indices, Departamentos = Datos.CalcularTamanos()
}

func CargarInvenAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	fmt.Println("Esto es una peticion de tipo post")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo post")
		var buf bytes.Buffer
		file, header, err := req.FormFile("myFile")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		fmt.Printf("File name %s\n", name[0])
		io.Copy(&buf, file)
		contents := buf.String()
		fmt.Println(contents)
		_ = json.Unmarshal(buf.Bytes(), &Inventario)
		_ = encoder.Encode("Datos cargados con exito")
		fmt.Println(Inventario)
		Inventario.SacarInventario(Vector, Indices, Departamentos)
	}
}

func DevolverInventarioAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	fmt.Println("Esto es una peticion de tipo post")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	vars := mux.Vars(req)
	nombre, _ := vars["nom"]
	calificacion, _ := strconv.Atoi(vars["cal"])
	departamento, _ := vars["dept"]
	first1 := nombre[0:1]
	fmt.Println(first1)
	indice, dept, _ := EncontrarIndices(departamento, first1)
	elemento := int(calificacion) + 5*(indice+(len(Indices)*dept)) - 1
	store, err := Vector[elemento].BuscarTienda(nombre)
	if err == 0 {
		_ = encoder.Encode("Tienda no existe")
	} else {
		if store.ReturnRaiz() == nil {
			_ = encoder.Encode("No hay productos")
		} else {
			arreglo := store.ReturnInventario()
			_ = encoder.Encode(arreglo)
		}
	}
}

func AgregarAlCarritoAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	arregloProd := make([]EstructurasCreadas.ProductCarr, 1)
	var producto EstructurasCreadas.ProductCarr
	fmt.Println(req.Body)
	_ = json.NewDecoder(req.Body).Decode(&producto)
	arregloProd[0] = producto
	fmt.Println(producto)
	Carrito = append(Carrito, arregloProd...)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo get")
		_ = encoder.Encode("Agregado")
	}
}

func ReturnCarritoAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	var carr = Carrito
	fmt.Println(carr)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if len(Carrito) == 0 {
		_ = encoder.Encode("No hay Articulos")
	} else {
		fmt.Println("Esto es una peticion de tipo get")
		_ = encoder.Encode(Carrito)
	}
	fmt.Println(Carrito)
}

func EliminarDelCarritoAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	var producto EstructurasCreadas.ProductCarr
	_ = json.NewDecoder(req.Body).Decode(&producto)
	fmt.Println(producto)
	espacio := 0
	for espacio = 0; espacio < len(Carrito); espacio++ {
		if Carrito[espacio].Codigo == producto.Codigo && Carrito[espacio].Tienda == producto.Tienda && Carrito[espacio].Departamento == producto.Departamento && Carrito[espacio].Calificacion == producto.Calificacion {
			break
		}
	}
	Carrito = append(Carrito[:espacio], Carrito[espacio+1:]...)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo post")
		_ = encoder.Encode("Eliminado")
	}
}

func PagarCarritoAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	espacio := 0
	for espacio = 0; espacio < len(Carrito); espacio++ {
		nombre := Carrito[espacio].Tienda
		calificacion := Carrito[espacio].Calificacion
		departamento := Carrito[espacio].Departamento
		first1 := nombre[0:1]
		indice, dept, _ := EncontrarIndices(departamento, first1)
		elemento := int(calificacion) + 5*(indice+(len(Indices)*dept)) - 1
		store, _ := Vector[elemento].BuscarTienda(nombre)
		store.RestarInventario(int(Carrito[espacio].Codigo), int(Carrito[espacio].Cantidad))
	}
	Carrito = nil
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo get")
		_ = encoder.Encode("Pagado")
	}
}

func CargarPedidosAPI(w http.ResponseWriter, req *http.Request) {
	//setupCorsResponse(&w, req)
	fmt.Println("Esto es una peticion de tipo post")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if TamaVec == 0 {
		_ = encoder.Encode("Los datos no han sido ingresados")
	} else {
		fmt.Println("Esto es una peticion de tipo post")
		var buf bytes.Buffer
		file, header, err := req.FormFile("myFile")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		fmt.Printf("File name %s\n", name[0])
		io.Copy(&buf, file)
		contents := buf.String()
		fmt.Println(contents)
		_ = json.Unmarshal(buf.Bytes(), &JsonPedidos)
		_ = encoder.Encode("Datos cargados con exito")
		fmt.Println(Pedidos)
		JsonPedidos.ConstruirDatos(&Pedidos)
		inventario := JsonPedidos
		invetario1 := Pedidos
		fmt.Println(inventario)
		fmt.Println(invetario1)
	}
}
