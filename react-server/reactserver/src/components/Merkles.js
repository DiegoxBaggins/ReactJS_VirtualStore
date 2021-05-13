import React, {Component} from 'react';
import axios from 'axios';
import HeaderInicio from "./HeaderInicio";
import {Redirect} from "react-router";
import HeaderAdmin from "./HeaderAdmin";
import {sha256} from "js-sha256";

const Server = "http://localhost:3000";

class Merkles extends Component{

    state = {
        redirect: null,
        imagen1: null,
        imagen2: null,
        imagen3: null,
        imagen4: null,
    }

    componentDidMount = async() => {
        await axios.get(`${Server}/arbolesMerkle`).then( (response) => {
            console.log(response.data)
            if (response.data !== "No hay Usuarios") {
                let imagenes = response.data
                this.setState({
                    imagen1: imagenes.Imagen1,
                    imagen2: imagenes.Imagen2,
                    imagen3: imagenes.Imagen3,
                    imagen4: imagenes.Imagen4
                });
                console.log("lista no es vacia")
            }else{
                console.log("lista es vacia")
            }
        });
    }

    render(){

        return (
            <div>
                <HeaderAdmin />
                <div className="center-2">
                    <h2 className="subheader"> Arboles de Merkle </h2>
                    <h2>Transacciones</h2>
                    <hr/>
                    {this.state.imagen1 &&
                    <img src={`data:image/formato;base64,${this.state.imagen1}`}  alt="Imagen"/>
                    }
                    <hr/>
                    <h2>Tiendas</h2>
                    <hr/>
                    {this.state.imagen2 &&
                    <img src={`data:image/formato;base64,${this.state.imagen2}`}  alt="Imagen"/>
                    }
                    <hr/>
                    <h2>Productos</h2>
                    <hr/>
                    {this.state.imagen3 &&
                    <img src={`data:image/formato;base64,${this.state.imagen3}`}  alt="Imagen"/>
                    }
                    <hr/>
                    <h2>Usuarios</h2>
                    <hr/>
                    {this.state.imagen4 &&
                    <img src={`data:image/formato;base64,${this.state.imagen4}`}  alt="Imagen"/>
                    }
                    <div className="clearfix"> </div>
                </div>
            </div>
        )
    }

}

export default Merkles;