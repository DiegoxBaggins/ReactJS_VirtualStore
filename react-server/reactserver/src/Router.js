import React, {Component} from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import Tiendas from "./components/Tiendas";
import SubirArchivos from "./components/SubirArchivos";
import Productos from "./components/Productos";
import Carrito from "./components/Carrito";
import AdminPedidos from "./components/AdminPedidos";
import InicioSesion from "./components/InicioSesion";
import Usuario from "./components/Usuarios";
import Merkles from "./components/Merkles";



class Router extends Component{

    render(){
        return (
            <BrowserRouter>
                {/* Configurar rutas y paginas*/}
                <Switch>
                    <Route exact path="/" component={InicioSesion} />
                    <Route exact path="/Home" component={Tiendas} />
                    <Route exact path="/pedidos" component={AdminPedidos} />
                    <Route exact path="/uploads" component={SubirArchivos} />
                    <Route exact path="/admin" component={SubirArchivos} />
                    <Route exact path="/usuarios" component={Usuario}/>
                    <Route exact path="/tienda" component={Productos}/>
                    <Route exact path="/Carrito" component={Carrito}/>
                    <Route exact path="/Merkle" component={Merkles}/>
                </Switch>

            </BrowserRouter>
        );
    }
}

export default Router;