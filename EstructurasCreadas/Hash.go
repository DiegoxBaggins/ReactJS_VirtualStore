package EstructurasCreadas

import (
	"github.com/olekukonko/tablewriter"
	"math"
	"os"
	"strconv"
	"time"
)

type Comentario struct {
	User   int
	Fecha  string
	Coment string
}

type SubComentario struct {
	Comentario *Comentario
	Sub        *SubComentario
}

type NodoHash struct {
	hash           int
	Data           *Comentario
	Subcomentarios *HashTable
}

func NewNodoHash(key int, value *Comentario) *NodoHash {
	return &NodoHash{
		hash:           key,
		Data:           value,
		Subcomentarios: nil,
	}
}

type HashTable struct {
	size      int
	carga     int
	capacidad int
	datos     []*NodoHash
}

func NewHashTable() *HashTable {
	return &HashTable{
		size:      7,
		carga:     0,
		capacidad: 60,
		datos:     make([]*NodoHash, 7),
	}
}

func nextPrime(n int) int {
	if n < 2 {
		return 2
	} else if n == 2 {
		return 3
	}
	next := n + 1
	for i := 2; i < int(next/2); i++ {
		if (next%i == 0) && (i != next) {
			return nextPrime(next)
		}
	}
	return next
}

func (tabla *HashTable) insertar(nuevo int, value *Comentario) {
	node := NewNodoHash(nuevo, value)
	pos := tabla.position(nuevo)
	tabla.datos[pos] = node
	tabla.carga++
	if ((tabla.carga * 100) / tabla.size) > tabla.capacidad {
		old := tabla.datos
		tabla.datos = make([]*NodoHash, nextPrime(tabla.size))
		tabla.size = len(tabla.datos)
		aux := 0
		for i := 0; i < len(old); i++ {
			if old[i] != nil {
				aux = tabla.position(old[i].hash)
				tabla.datos[aux] = old[i]
			}
		}
	}
}

func (tabla *HashTable) find(key int, value *Comentario) *NodoHash {
	i, p := 0, 0
	p = tabla.hashing(key)
	for !((tabla.datos[p].hash == key) && (tabla.datos[p].Data.Fecha == value.Fecha)) {
		i++
		p = tabla.closedHashing(p, i)
		if p >= tabla.size {
			p = tabla.tableLimit(p)
		}
	}
	return tabla.datos[p]
}

func (tabla *HashTable) InsertarSub(comentario *SubComentario) {
	currentTime := time.Now()
	if comentario.Sub != nil {
		tmp := tabla.find(comentario.Comentario.User, comentario.Comentario)
		if tmp.Subcomentarios == nil {
			tmp.Subcomentarios = NewHashTable()
		}
		tmp.Subcomentarios.InsertarSub(comentario.Sub)
	} else {
		comentario.Comentario.Fecha = currentTime.Format("2006-01-02 15:04:05")
		tabla.insertar(comentario.Comentario.User, comentario.Comentario)
	}
	tabla.imprimir()
}

func (tabla *HashTable) hashing(key int) int {
	return int(math.Trunc(float64(tabla.size) * ((0.2520 * float64(key)) - math.Trunc(0.2520*float64(key)))))
}

func (tabla *HashTable) closedHashing(p int, i int) int {
	return p + tabla.hashing(i*i)
}

func (tabla *HashTable) tableLimit(p int) int {
	tmp := p - tabla.size
	for tmp >= tabla.size {
		tmp = tmp - tabla.size
	}
	return tmp
}

func (tabla *HashTable) position(key int) int {
	i, p := 0, 0
	p = tabla.hashing(key)
	for tabla.datos[p] != nil {
		i++
		p = tabla.closedHashing(p, i)
		if p >= tabla.size {
			p = tabla.tableLimit(p)
		}
	}
	return p
}

func (tabla *HashTable) imprimir() {
	data := make([][]string, tabla.size)
	for i := 0; i < len(tabla.datos); i++ {
		tmp := make([]string, 2)
		aux := tabla.datos[i]
		if aux != nil {
			tmp[0] = strconv.Itoa(aux.hash)
			tmp[1] = aux.Data.Coment
		} else {
			tmp[0] = "-"
			tmp[1] = "-"
		}
		data[i] = tmp
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hash", "Valor"})
	table.SetFooter([]string{"size", strconv.Itoa(tabla.size), "Carga", strconv.Itoa(tabla.carga)})
	table.AppendBulk(data)
	table.Render()
}
