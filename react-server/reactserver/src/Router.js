import React, {Component} from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import SeccionPrueba from "./components/SeccionPrueba";
import Paginas from "./components/Paginas";
import Header from "./components/Header";

class Router extends Component{

    render(){
        return (
            <BrowserRouter>
                <Header />
                {/* Configurar rutas y paginas*/}
                <Switch>
                    <Route exact path="/" component={Paginas} />
                    <Route exact path="/Home" component={Paginas} />
                    <Route exact path="/ruta-prueba" component={SeccionPrueba} />

                </Switch>

            </BrowserRouter>
        );
    }
}

export default Router;