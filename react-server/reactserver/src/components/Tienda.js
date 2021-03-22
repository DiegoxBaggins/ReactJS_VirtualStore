import React, {Component} from 'react';
import {Link} from "react-router-dom";

class Tienda extends Component{

    state = {
        tienda : {}
    }

    GuardarDatos = () => {
        //console.log(this.state.tienda)
        localStorage.setItem('Nombre',this.state.tienda.Nombre);
        localStorage.setItem('Calificacion',this.state.tienda.Calificacion);
        localStorage.setItem('Departamento',this.state.tienda.Departamento)
    }

    componentDidMount = () => {
        this.setState({
            tienda : this.props.tienda
        })
    }

    render(){

        const {Nombre, Descripcion,Contacto, Calificacion, Logo} = this.props.tienda;

        return (
            <article className="article-item" id="article-template">
                <div className="image-wrap">
                    <img
                        src={Logo}  alt={Nombre}/>
                </div>
                <h2>{Nombre}</h2>
                <h4>{Descripcion}</h4>
                <h4>Calificacion: {Calificacion} Contacto: {Contacto}</h4>
                <Link to={'/tienda'}>
                    <button onClick={this.GuardarDatos} className="btn-upload" >Redirect</button>
                </Link>
                <div className="clearfix"> </div>
            </article>
        )
    }

}

export default Tienda;