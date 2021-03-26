import React, {Component} from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import SeccionPrueba from "./components/SeccionPrueba";
import Tiendas from "./components/Tiendas";
import Header from "./components/Header";
import SubirArchivos from "./components/SubirArchivos";
import Productos from "./components/Productos";
import Carrito from "./components/Carrito";
import AdminPedidos from "./components/AdminPedidos";


class Router extends Component{

    render(){
        return (
            <BrowserRouter>
                <Header />
                {/* Configurar rutas y paginas*/}
                <Switch>
                    <Route exact path="/" component={Tiendas} />
                    <Route exact path="/Home" component={Tiendas} />
                    <Route exact path="/Admin" component={AdminPedidos} />
                    <Route exact path="/uploads" component={SubirArchivos} />
                    <Route exact path="/tienda" component={Productos}/>
                    <Route exact path="/Carrito" component={Carrito}/>
                </Switch>

            </BrowserRouter>
        );
    }
}

export default Router;