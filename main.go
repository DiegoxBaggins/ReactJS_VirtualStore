package main

import (
	"Proyecto-moduls/EstructurasCreadas"
	"fmt"
)

func main() {
	var clasificacion [27][10][5]EstructurasCreadas.ListaTienda
	//fmt.Println(clasificacion)
	//fmt.Println(len(clasificacion))

	/*for i := 0; i <27; i++ {
		for j := 0; j <10; j++ {
			for k := 0; k <5; k++ {
				fmt.Println(clasificacion[i][j][k])
			}
		}
	}*/

	//clasificacion[0][0][0].InsertarTienda("Genetik", "Tienda en Linea",41283319,5)
	//clasificacion[0][0][0].InsertarTienda("Genetiks", "Tienda en Linea",41283319,5)
	//clasificacion[0][0][0].InsertarTienda("Genetikss", "Tienda en Linea",41283319,5)
	//fmt.Println("lista", clasificacion[0][0][0])
	//clasificacion[0][0][0].MostrarDatos()
	clasificacion[0][0][1].InsertarTienda("Genetik", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("Dressy", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("dressy", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("foes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("foes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("foes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("Armani", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("Farts", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("qoes", "Tienda en Linea", 41283319, 5)
	clasificacion[0][0][1].InsertarTienda("fies", "Tienda en Linea", 41283319, 5)
	fmt.Println("lista", clasificacion[0][0][1])
	clasificacion[0][0][1].MostrarDatos()

}
