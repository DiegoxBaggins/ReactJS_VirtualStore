import React, {Component} from 'react';
import axios from 'axios';
import Producto from "./Producto";
import {ListItem} from "semantic-ui-react";
import Anio from "./Anio";

const Server = "http://localhost:3000";

class SubirArchivos extends Component{

    state = {
        anios : [],
        selectedFile: null,
        datos : null,
        arbol : null,
        matriz : null
    }

    DevolverMatriz = (anio, mes, mesnum) => {
        this.setState(
            {
                matriz : null
            }
        )
        const direccion = anio + "/" + mesnum + "/" + mes
        console.log(direccion)
        axios.get(`${Server}/matriz/${direccion}`).then( (response) => {
            console.log(response);
            if (response.data !== "No hay Pedidos") {
                this.setState({
                    matriz: response.data
                });
                console.log("lista no es vacia")
                console.log(response.data)
            }else{
                console.log("lista es vacia")
            }
        });
    }

    componentDidMount = () => {
        axios.get(`${Server}/aniosmeses`).then( (response) => {
            if (response.data !== "No hay Pedidos") {
                this.setState({
                    anios: response.data
                });
                console.log("lista no es vacia")
                console.log(response.data)
            }else{
                console.log("lista es vacia")
            }
        });
        axios.get(`${Server}/aniosmesesimg`).then( (response) => {
            console.log(response);
            if (response.data !== "No hay pedidos") {
                this.setState({
                    arbol: response.data
                });
                console.log("lista no es vacia")
                console.log(response.data)
            }else{
                console.log("lista es vacia")
            }
        });
    }

    render(){
        return (
            <div className="center-2">
                <div id="content-1">
                    <h2 className="subheader"> Pedidos </h2>
                    {this.state.arbol &&
                    <img src={`data:image/formato;base64,${this.state.arbol}`}  alt="Imagen"/>
                    }
                    <div>
                        <h2> Matriz </h2>
                        {this.state.matriz &&
                        <img src={`data:image/formato;base64,${this.state.matriz}`}  alt="Imagen"/>
                        }
                    </div>
                </div>
                <aside id="sidebar-1">
                    <div id="search" className="sidebar-item">
                        {this.state.anios[0] &&
                            this.state.anios.map((anio, i) => {
                                return(
                                    <Anio key={i} anio={anio} DevolverMatriz={this.DevolverMatriz}/>
                                )
                            })
                        }
                    </div>
                </aside>
                <div className="clearfix"> </div>
            </div>
        )
    }

}

export default SubirArchivos;