package EstructurasCreadas

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
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

func (archivo *LecUsuarios) ConvertirArbol(arbol *ArbolB) {
	for i := 0; i < len(archivo.Usuarios); i++ {
		arbol.Insert(&archivo.Usuarios[i])
	}
}

type ArbolB struct {
	raiz *BNodo
	t    int
}

func NewArbolB(t int) *ArbolB {
	return &ArbolB{NewBNodo(t, true), t}
}

func (arbol *ArbolB) ComprobarUser(dpi int, contra string) string{
	usuario := arbol.BuscarUsuario(dpi)
	if usuario == nil {
		return ""
	}else{
		pass := fmt.Sprintf("%x", sha256.Sum256([]byte(usuario.Password)))
		if pass == contra {
			switch usuario.Cuenta {
			case "Admin": return "admin"
			case "Usuario": return "usuario"
			default: return "usuario"
			}
		} else{
			return ""
		}
	}
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

func (arbol *ArbolB) Insert(usuario *Usuario) {
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

func (arbol *ArbolB) GraficarGrafo() {
	direct := "./react-server/reactserver/src/assets/images/grafos/usuario/"
	var graph = "digraph G{\n"
	graph += "rankdir=TB\n node[shape=box]\nconcentrate=true\n"
	str1, _ := arbol._GraficarGrafo(arbol.raiz, 0)
	graph += str1
	graph += "\n}"
	data := []byte(graph)
	err := ioutil.WriteFile(direct+"usuarios.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", direct+"usuarios.dot").Output()
	mode := int(0777)
	err = ioutil.WriteFile(direct+"usuarios.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
	graph = "digraph G{\n"
	graph += "rankdir=TB\n node[shape=box]\nconcentrate=true\n"
	str1, _ = arbol._GraficarGrafoENC(arbol.raiz, 0)
	graph += str1
	graph += "\n}"
	data = []byte(graph)
	err = ioutil.WriteFile(direct+"usuariosENC.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	cmd, _ = exec.Command(path, "-Tpng", direct+"usuariosENC.dot").Output()
	mode = int(0777)
	err = ioutil.WriteFile(direct+"usuariosENC.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
	graph = "digraph G{\n"
	graph += "rankdir=TB\n node[shape=box]\nconcentrate=true\n"
	str1, _ = arbol._GraficarGrafoSEN(arbol.raiz, 0)
	graph += str1
	graph += "\n}"
	data = []byte(graph)
	err = ioutil.WriteFile(direct+"usuariosSEN.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	cmd, _ = exec.Command(path, "-Tpng", direct+"usuariosSEN.dot").Output()
	mode = int(0777)
	err = ioutil.WriteFile(direct+"usuariosSEN.png", cmd, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
}

func (arbol *ArbolB) _GraficarGrafo(temp *BNodo, nods int) (string, int) {
	grafo := ""
	pivote := nods
	str1 := ""
	if temp != nil {
		grafo += "Nodo" + strconv.Itoa(nods) + "[shape=record label=\"" + temp._GraficarGrafo() + "\"]\n"
		nods += 1
		for i := 0; i < temp.canth; i++ {
			grafo += "Nodo" + strconv.Itoa(pivote) + "->" + "Nodo" + strconv.Itoa(nods) + ";\n"
			str1, nods = arbol._GraficarGrafo(temp.hijos[i], nods)
			grafo += str1
		}
	}
	return grafo, nods
}

func (nodo *BNodo) _GraficarGrafo() string {
	graph := ""
	usuario := nodo.usuarios[0]
	for i := 0; i < nodo.llaves; i++ {
		usuario = nodo.usuarios[i]
		graph += "{" + usuario.Nombre + "|"
		graph += usuario.Password + "|"
		graph += strconv.Itoa(usuario.Dpi) + "|"
		graph += usuario.Correo + "|"
		graph += usuario.Cuenta + "}"
		if i+1 != nodo.llaves {
			graph += "|"
		}
	}
	return graph
}

func (arbol *ArbolB) _GraficarGrafoENC(temp *BNodo, nods int) (string, int) {
	grafo := ""
	pivote := nods
	str1 := ""
	if temp != nil {
		grafo += "Nodo" + strconv.Itoa(nods) + "[shape=record label=\"" + temp._GraficarGrafoENC() + "\"]\n"
		nods += 1
		for i := 0; i < temp.canth; i++ {
			grafo += "Nodo" + strconv.Itoa(pivote) + "->" + "Nodo" + strconv.Itoa(nods) + ";\n"
			str1, nods = arbol._GraficarGrafoENC(temp.hijos[i], nods)
			grafo += str1
		}
	}
	return grafo, nods
}

func (nodo *BNodo) _GraficarGrafoENC() string {
	graph := ""
	usuario := nodo.usuarios[0]
	for i := 0; i < nodo.llaves; i++ {
		usuario = nodo.usuarios[i]
		graph += "{" + fmt.Sprintf("%x", sha256.Sum256([]byte(usuario.Nombre))) + "|"
		graph += fmt.Sprintf("%x", sha256.Sum256([]byte(usuario.Password))) + "|"
		graph += fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.Itoa(usuario.Dpi)))) + "|"
		graph += fmt.Sprintf("%x", sha256.Sum256([]byte(usuario.Correo))) + "|"
		graph += fmt.Sprintf("%x", sha256.Sum256([]byte(usuario.Cuenta))) + "}"
		if i+1 != nodo.llaves {
			graph += "|"
		}
	}
	return graph
}

func (arbol *ArbolB) _GraficarGrafoSEN(temp *BNodo, nods int) (string, int) {
	grafo := ""
	pivote := nods
	str1 := ""
	if temp != nil {
		grafo += "Nodo" + strconv.Itoa(nods) + "[shape=record label=\"" + temp._GraficarGrafoSEN() + "\"]\n"
		nods += 1
		for i := 0; i < temp.canth; i++ {
			grafo += "Nodo" + strconv.Itoa(pivote) + "->" + "Nodo" + strconv.Itoa(nods) + ";\n"
			str1, nods = arbol._GraficarGrafoENC(temp.hijos[i], nods)
			grafo += str1
		}
	}
	return grafo, nods
}

func (nodo *BNodo) _GraficarGrafoSEN() string {
	graph := ""
	usuario := nodo.usuarios[0]
	for i := 0; i < nodo.llaves; i++ {
		usuario = nodo.usuarios[i]
		graph += "{" + usuario.Nombre + "|"
		graph += fmt.Sprintf("%x", sha256.Sum256([]byte(usuario.Password))) + "|"
		graph += fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.Itoa(usuario.Dpi)))) + "|"
		graph += fmt.Sprintf("%x", sha256.Sum256([]byte(usuario.Correo))) + "|"
		graph += usuario.Cuenta + "}"
		if i+1 != nodo.llaves {
			graph += "|"
		}
	}
	return graph
}
