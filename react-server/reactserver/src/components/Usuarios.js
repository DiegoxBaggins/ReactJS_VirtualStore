import React, {Component} from 'react';
import axios from 'axios';
import HeaderInicio from "./HeaderInicio";
import {Redirect} from "react-router";
import HeaderAdmin from "./HeaderAdmin";

const Server = "http://localhost:3000";

class Usuario extends Component{

    state = {
        redirect: null,
        imagen1: null,
        imagen2: null,
        imagen3: null
    }

    dpiRef = React.createRef();
    passRef = React.createRef();
    nombreRef = React.createRef();
    correoRef = React.createRef();
    adminRef = React.createRef();
    usuarioRef = React.createRef();

    ConsultarDatos = async(e) => {
        e.preventDefault();
        let dpi = this.dpiRef.current.value;
        let pass = this.dpiRef.current.value;
        let usuario = {
            dpi : dpi,
            password : pass
        }
        await axios.post(`${Server}/ingresar`, usuario).then(function (response) {
            console.log(response.data);
        });
    }

    componentDidMount = async() => {
        await axios.get(`${Server}/arbolusuarios1`).then( (response) => {
            console.log(response.data)
            if (response.data !== "No hay Usuarios") {
                this.setState({
                    imagen1: response.data
                });
                console.log("lista no es vacia")
            }else{
                console.log("lista es vacia")
            }
        });
        await axios.get(`${Server}/arbolusuarios2`).then( (response) => {
            console.log(response.data)
            if (response.data !== "No hay Usuarios") {
                this.setState({
                    imagen2: response.data
                });
                console.log("lista no es vacia")
            }else{
                console.log("lista es vacia")
            }
        });
        await axios.get(`${Server}/arbolusuarios3`).then( (response) => {
            console.log(response.data)
            if (response.data !== "No hay Usuarios") {
                this.setState({
                    imagen3: response.data
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
                        <h2 className="subheader"> Ingresa Tus Datos </h2>
                        <form className="mid-form" onSubmit={this.ConsultarDatos}>
                            <div className="form-group">
                                <label htmlFor="dpi">dpi</label>
                                <input type="text" name="dpi" ref={this.dpiRef}/>
                            </div>
                            <div className="form-group">
                                <label htmlFor="nombre">Nombre</label>
                                <input type="text" name="nombre" ref={this.nombreRef}/>
                            </div>
                            <div className="form-group">
                                <label htmlFor="correo">Correo</label>
                                <input type="text" name="correo" ref={this.correoRef}/>
                            </div>
                            <div className="form-group">
                                <label htmlFor="password">Contrasena</label>
                                <input type="text" name="password" ref={this.passRef}/>
                            </div>
                            <div className="form-group radibuttons">
                                <input type="radio" name="usuario" value="admin" ref={this.adminRef}/> Admin
                                <input type="radio" name="usuario" value="usuario" ref={this.usuarioRef}/> Usuario
                            </div>
                            <div className="clearfix"> </div>
                            <input type="submit" value="Enviar" className="btn btn-success"/>
                        </form>
                    <div className="clearfix"> </div>
                    {this.state.imagen1 &&
                    <img src={`data:image/formato;base64,${this.state.imagen1}`}  alt="Imagen"/>
                    }
                    {this.state.imagen2 &&
                    <img src={`data:image/formato;base64,${this.state.imagen2}`}  alt="Imagen"/>
                    }
                    {this.state.imagen3 &&
                    <img src={`data:image/formato;base64,${this.state.imagen3}`}  alt="Imagen"/>
                    }
                    <div className="clearfix"> </div>
                </div>
            </div>
        )
    }

}

export default Usuario;