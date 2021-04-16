import React, {Component} from 'react';
import axios from 'axios';
import HeaderInicio from "./HeaderInicio";
import {Redirect} from "react-router";
import {sha256} from "js-sha256";

const Server = "http://localhost:3000";

class InicioSesion extends Component{

    state = {
        redirect: null
    }

    dpiRef = React.createRef();
    passRef = React.createRef();

    ConsultarDatos = async(e) => {
        e.preventDefault();
        let dpi = parseInt(this.dpiRef.current.value);
        let pass = sha256(this.passRef.current.value);
        let usuario = {
            Dpi : dpi,
            Password : pass
        }
        console.log(usuario)
        await axios.post(`${Server}/ingresar`, usuario).then( (response) => {
            console.log(response.data);
            let respuesta = response.data
            if (respuesta === "F"){
                alert("Datos erroneos");
            }else{
                if (respuesta === "A"){
                    this.setState({
                        redirect: "/admin"
                    })
                }else{
                    this.setState({
                        redirect: "/Home"
                    })
                }
            }
        });
    }

    render(){

        if (this.state.redirect) {
            return <Redirect to={this.state.redirect} />
        }

        return (
            <div>
                <HeaderInicio/>
                <div className="center-2">
                    <h2 className="subheader"> Ingresa Tus Datos </h2>
                    <form className="mid-form" onSubmit={this.ConsultarDatos}>
                        <div className="form-group">
                            <label htmlFor="dpi">dpi</label>
                            <input type="text" name="dpi" ref={this.dpiRef}/>
                        </div>
                        <div className="form-group">
                            <label htmlFor="password">Contrasena</label>
                            <input type="text" name="password" ref={this.passRef}/>
                        </div>
                        <div className="clearfix"> </div>
                        <input type="submit" value="Enviar" className="btn btn-success"/>
                    </form>
                    <div className="clearfix"> </div>
                </div>
                <br/><br/><br/><br/><br/><br/><br/><br/><br/>
            </div>
        )
    }

}

export default InicioSesion;