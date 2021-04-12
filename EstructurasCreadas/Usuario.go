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

func (arbol *ArbolB) BuscarUsuario(k int) *Usuario {
	if arbol.raiz != nil {
		return arbol._buscarUsuario(arbol.raiz, k)
	} else {
		return nil
	}
}

func (arbol *ArbolB) _buscarUsuario(nodo1 *BNodo, k int) *Usuario {
	i := 0
	if nodo1 == nil {
		return nil
	}
	for i = 0; i < nodo1.llaves; i++ {
		if k < nodo1.usuarios[i].Dpi {
			break
		}
		if k == nodo1.usuarios[i].Dpi {
			return nodo1.usuarios[i]
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
	canth    int
}

func NewBNodo(t int, leaf bool) *BNodo {
	return &BNodo{make([]*Usuario, t), t, make([]*BNodo, t+1), 0, leaf, 0}
}

func (arbol *ArbolB) Split(nodo1 *BNodo) *BNodo {
	nuevoNodo := NewBNodo(arbol.t, nodo1.hoja)
	i := 0
	j := 0
	for i = nodo1.t / 2; i < nodo1.t; i++ {
		nuevoNodo.usuarios[j] = nodo1.usuarios[i]
		nodo1.usuarios[i] = nil
		if !nodo1.hoja {
			nuevoNodo.hijos[j+1] = nodo1.hijos[i+1]
			nodo1.hijos[i+1] = nil
			nuevoNodo.canth += 1
		}
		j += 1
		nuevoNodo.llaves += 1
	}
	if !nodo1.hoja {
		nuevoNodo.hijos[j] = nodo1.hijos[i+1]
		nodo1.hijos[i+1] = nil
		nuevoNodo.canth += 1
		nodo1.canth /= 2
	}
	nodo1.llaves = nodo1.t / 2
	return nuevoNodo
}

func (arbol *ArbolB) SplitIzq(nodo1 *BNodo) {
	nuevoNodoIzq := NewBNodo(arbol.t, nodo1.hoja)
	nuevoNodoDer := NewBNodo(arbol.t, nodo1.hoja)
	pivote := nodo1.usuarios[nodo1.t/2]
	i := 0
	j := 0
	for i = 0; i < nodo1.t/2; i++ {
		nuevoNodoIzq.usuarios[j] = nodo1.usuarios[i]
		nodo1.usuarios[i] = nil
		j += 1
		nuevoNodoIzq.llaves += 1
	}
	j = 0
	if !nodo1.hoja {
		for i = 0; i < nodo1.t/2; i++ {
			nuevoNodoIzq.hijos[j] = nodo1.hijos[i]
			nodo1.hijos[i] = nil
			j += 1
			nuevoNodoIzq.canth += 1
		}
		nuevoNodoIzq.hijos[j] = nodo1.hijos[i]
		nodo1.hijos[i] = nil
		nuevoNodoIzq.canth += 1
	}
	j = 0
	for i = nodo1.t/2 + 1; i < nodo1.t; i++ {
		nuevoNodoDer.usuarios[j] = nodo1.usuarios[i]
		nodo1.usuarios[i] = nil
		j += 1
		nuevoNodoDer.llaves += 1
	}
	j = 0
	if !nodo1.hoja {
		for i = nodo1.t/2 + 1; i < nodo1.t; i++ {
			nuevoNodoDer.hijos[j] = nodo1.hijos[i]
			nodo1.hijos[i] = nil
			j += 1
			nuevoNodoDer.canth += 1
		}
		nuevoNodoDer.hijos[j] = nodo1.hijos[i]
		nodo1.hijos[i] = nil
		nuevoNodoDer.canth += 1
	}
	nodo1.usuarios[0] = pivote
	nodo1.usuarios[nodo1.t/2] = nil
	nodo1.hijos[0] = nuevoNodoIzq
	nodo1.hijos[1] = nuevoNodoDer
	nodo1.llaves = 1
	nodo1.canth = 2
	nodo1.hoja = false
}

func (arbol *ArbolB) Insert(dpi int) {
	usuario := &Usuario{dpi, "as", "bb", "ss", "ss"}
	nuevo := arbol.Insertar(arbol.raiz, usuario)
	if nuevo != nil {
		pivote := nuevo.usuarios[0]
		for i := 0; i < nuevo.llaves-1; i++ {
			nuevo.usuarios[i] = nuevo.usuarios[i+1]
			nuevo.usuarios[i+1] = nil
		}
		nuevo.llaves -= 1
		arbol.InsertarUsuario(arbol.raiz, pivote, nuevo)
	}
	if arbol.raiz.llaves > arbol.t-1 {
		arbol.SplitIzq(arbol.raiz)
	}
}

func (arbol *ArbolB) Insertar(nodo *BNodo, usuario *Usuario) *BNodo {
	if nodo.hoja {
		arbol.InsertarUsuario(nodo, usuario, nil)
	} else {
		i := 0
		for i = 0; i < nodo.llaves; i++ {
			if usuario.Dpi < nodo.usuarios[i].Dpi {
				break
			}
		}
		nuevo := arbol.Insertar(nodo.hijos[i], usuario)
		if nuevo != nil {
			pivote := nuevo.usuarios[0]
			for i = 0; i < nuevo.llaves-1; i++ {
				nuevo.usuarios[i] = nuevo.usuarios[i+1]
				nuevo.usuarios[i+1] = nil
			}
			nuevo.llaves -= 1
			arbol.InsertarUsuario(nodo, pivote, nuevo)
		}
	}
	if nodo.llaves > nodo.t-1 && arbol.raiz != nodo {
		return arbol.Split(nodo)
	} else {
		return nil
	}
}

func (arbol *ArbolB) InsertarUsuario(nodo *BNodo, usuario *Usuario, hijoDer *BNodo) {
	k := usuario.Dpi
	pivote := 0
	for pivote = 0; pivote < nodo.llaves; pivote++ {
		if nodo.usuarios[pivote].Dpi > k {
			break
		}
	}
	i := 0
	for i = nodo.llaves - 1; i >= pivote; i-- {
		nodo.usuarios[i+1] = nodo.usuarios[i]
	}
	for i = nodo.canth - 1; i >= pivote; i-- {
		nodo.hijos[i+1] = nodo.hijos[i]
	}
	nodo.usuarios[pivote] = usuario
	nodo.llaves += 1
	if hijoDer != nil {
		nodo.hijos[pivote+1] = hijoDer
		nodo.canth += 1
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
