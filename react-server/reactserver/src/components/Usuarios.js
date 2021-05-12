import React, {Component} from 'react';
import axios from 'axios';
import HeaderInicio from "./HeaderInicio";
import {Redirect} from "react-router";
import HeaderAdmin from "./HeaderAdmin";
import {sha256} from "js-sha256";

const Server = "http://localhost:3000";

class Usuario extends Component{

    state = {
        redirect: null,
        imagen1: null,
        imagen2: null,
        imagen3: null,
        user: {}
    }

    dpiRef = React.createRef();
    passRef = React.createRef();
    nombreRef = React.createRef();
    correoRef = React.createRef();
    adminRef = React.createRef();
    usuarioRef = React.createRef();

    dpi2Ref = React.createRef();
    pass2Ref = React.createRef();

    ConsultarDatos = async(e) => {
        e.preventDefault();
        let dpi = parseInt(this.dpiRef.current.value);
        let pass = this.passRef.current.value;
        let nombre = this.nombreRef.current.value;
        let correo = this.correoRef.current.value;
        let cuenta = 'Usuario';
        if(this.adminRef.current.checked){
            cuenta = 'Admin'
        }
        let usuario = {
            Dpi : dpi,
            Nombre : nombre,
            Correo : correo,
            Password : pass,
            Cuenta: cuenta
        }
        await axios.post(`${Server}/newuser`, usuario).then(function (response) {
            console.log(response.data);
            alert("Usuario registrado con exito");
        });
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

    EliminarUsuario = async(e) => {
        e.preventDefault();
        let dpi = parseInt(this.dpi2Ref.current.value);
        let pass = sha256(this.pass2Ref.current.value);
        let usuario = {
            Dpi : dpi,
            Password : pass
        }
        console.log(usuario)
        await axios.post(`${Server}/borrarUsuario`, usuario).then( (response) => {
            console.log(response.data);
            alert(response.data);
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
                            <input type="password" name="password" ref={this.passRef}/>
                        </div>
                        <div className="form-group radibuttons">
                            <input type="radio" name="usuario" value="Admin" ref={this.adminRef}/> Admin
                            <input type="radio" name="usuario" value="Usuario" ref={this.usuarioRef}/> Usuario
                        </div>
                        <div className="clearfix"> </div>
                        <input type="submit" value="Enviar" className="btn btn-success"/>
                    </form>
                    <br/>
                    <h2 className="subheader"> Ingresa Datos para Elimiar </h2>
                    <form className="mid-form" onSubmit={this.EliminarUsuario}>
                        <div className="form-group">
                            <label htmlFor="dpi">dpi</label>
                            <input type="text" name="dpi" ref={this.dpi2Ref}/>
                        </div>
                        <div className="form-group">
                            <label htmlFor="password">Contrasena</label>
                            <input type="password" name="password" ref={this.pass2Ref}/>
                        </div>
                        <div className="clearfix"> </div>
                        <input type="submit" value="Eliminar" className="btn btn-success"/>
                    </form>
                    <div className="clearfix"> </div>
                    <hr/>
                    <h2>Arbol sin cifrar</h2>
                    <hr/>
                    {this.state.imagen1 &&
                    <img src={`data:image/formato;base64,${this.state.imagen1}`}  alt="Imagen"/>
                    }
                    <hr/>
                    <h2>Arbol cifrado</h2>
                    <hr/>
                    {this.state.imagen2 &&
                    <img src={`data:image/formato;base64,${this.state.imagen2}`}  alt="Imagen"/>
                    }
                    <hr/>
                    <h2>Arbol con cifrado sensible</h2>
                    <hr/>
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