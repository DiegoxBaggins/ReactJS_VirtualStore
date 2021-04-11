package EstructurasCreadas

import (
	"fmt"
)

type LecUsuarios struct {
	Usuarios []Usuario
}
type Usuario struct {
	Dpi      int    `json:"Dpi,omitempty"`
	Nombre   string `json:"Nombre,omitempty"`
	Correo   string `json:"Correo,omitempty"`
	Password string `json:"Password,omitempty"`
	Cuenta   string `json:"Cuenta,omitempty"`
}

type ArbolB struct {
	raiz *BNodo
	t    int
}

func NewArbolB(t int) *ArbolB {
	return &ArbolB{NewBNodo(t, true), t}
}

func (arbol *ArbolB) traverse() {
	if arbol.raiz != nil {
		arbol.raiz.traverse()
	}
	fmt.Println()
}

func (arbol *ArbolB) BuscarUsuario(k int) *BNodo {
	if arbol.raiz != nil {
		return arbol._buscarUsuario(arbol.raiz, k)
	} else {
		return nil
	}
}

func (arbol *ArbolB) _buscarUsuario(nodo1 *BNodo, k int) *BNodo {
	i := 0
	if nodo1 == nil {
		return nil
	}
	for i = 0; i < nodo1.llaves; i++ {
		if k < nodo1.usuarios[i].Dpi {
			break
		}
		if k == nodo1.usuarios[1].Dpi {
			return nodo1
		}
	}
	if nodo1.hoja {
		return nil
	} else {
		return arbol._buscarUsuario(nodo1.hijos[i], k)
	}
}

type BNodo struct {
	usuarios []*Usuario
	t        int
	hijos    []*BNodo
	llaves   int
	hoja     bool
}

func NewBNodo(t int, leaf bool) *BNodo {
	return &BNodo{make([]*Usuario, t-1), t, make([]*BNodo, t), 0, leaf}
}

func (arbol *ArbolB) Split(nodo1 *BNodo, pos int, nodo2 *BNodo) {
	nodo3 := NewBNodo(arbol.t, nodo2.hoja)
	for j := 0; j < arbol.t-1; j++ {
		nodo3.usuarios[j] = nodo2.usuarios[j+arbol.t]
	}
	if !nodo2.hoja {
		for j := 0; j < arbol.t; j++ {
			nodo3.hijos[j] = nodo2.hijos[j+arbol.t]
		}
	}
	nodo3.llaves = arbol.t - 1
	for j := nodo1.llaves; j >= pos+1; j-- {
		nodo1.hijos[j+1] = nodo1.hijos[j]
	}
	nodo1.hijos[pos+1] = nodo3
	for j := nodo1.llaves - 1; j >= pos; j-- {
		nodo1.usuarios[j+1] = nodo1.usuarios[j]
	}
	nodo1.usuarios[pos] = nodo2.usuarios[arbol.t-1]
	nodo1.llaves = nodo1.llaves + 1
}

func (arbol *ArbolB) Insertar(dpi int) {
	usuario := &Usuario{dpi, "as", "bb", "ss", "ss"}
	raiz := arbol.raiz
	if raiz.llaves == arbol.t-1 {
		s := NewBNodo(arbol.t, false)
		raiz = s
		s.hijos[0] = raiz
		arbol.Split(s, 0, raiz)
		arbol.InsertarUsuario(s, usuario)
	} else {
		arbol.InsertarUsuario(raiz, usuario)
	}
}

func (arbol *ArbolB) InsertarUsuario(nodo *BNodo, usuario *Usuario) {
	k := usuario.Dpi
	if nodo.hoja {
		i := 0
		for i = nodo.llaves - 1; i >= 0 && k < nodo.usuarios[i].Dpi; i-- {
			nodo.usuarios[i+1] = nodo.usuarios[i]
		}
		nodo.usuarios[i+1] = usuario
		nodo.llaves += 1
	} else {
		i := 0
		for i = nodo.llaves - 1; i >= 0 && k < nodo.usuarios[i].Dpi; i-- {
		}
		i++
		tmp := nodo.hijos[i]
		if tmp.llaves == arbol.t-1 {
			arbol.Split(nodo, i, tmp)
			if k > nodo.usuarios[i].Dpi {
				i++
			}
		}
		arbol.InsertarUsuario(nodo.hijos[i], usuario)
	}
}

func (nodo *BNodo) buscarDpi(k int) int {
	for i := 0; i < nodo.llaves; i++ {
		if nodo.usuarios[i].Dpi == k {
			return i
		}
	}
	return -1
}

func (nodo *BNodo) traverse() {
	i := 0
	for i = 0; i < nodo.llaves; i++ {
		if nodo.hoja == false {
			nodo.hijos[i].traverse()
		}
		fmt.Print(nodo.usuarios[i].Dpi, ", ")
	}
	if nodo.hoja == false {
		nodo.hijos[i].traverse()
	}
}
