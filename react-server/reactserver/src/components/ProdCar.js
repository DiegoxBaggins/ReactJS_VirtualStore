import React, {Component} from 'react';
import axios from "axios";
import {Link} from "react-router-dom";
import {Redirect} from "react-router";

const Server = "http://127.0.0.1:3000";

class ProdCar extends Component{

    state = {
        redirect: null
    }

    EliminarCarrito = async() => {
        const {Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda, Calificacion, Departamento} = this.props.producto;
        let product = {
            Nombre: Nombre,
            Codigo: Codigo,
            Descripcion: Descripcion,
            Precio: Precio,
            Cantidad: Cantidad,
            Imagen: Imagen,
            Tienda: Tienda,
            Calificacion: Calificacion,
            Departamento: Departamento
        };
        console.log(product)
        await axios.post(`${Server}/eliminarCarrito`, product).then(function (response) {
            console.log(response);
        }).finally(() => {
            alert("Tienda borrada del carrito");
            this.setState({
                redirect: "/Carrito"
            })
        });
    }

    render(){

        if (this.state.redirect) {
            return <Redirect to={this.state.redirect} />
        }

        const {Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen, Tienda} = this.props.producto;

        return (
            <article className="article-item" id="article-template">
                <h3>{Nombre} | Codigo: {Codigo}</h3>
                <h5>Tienda: {Tienda}</h5>
                <div className="image-wrap">
                    <img
                        src={Imagen}  alt={Nombre}/>
                </div>
                <h4>{Descripcion}</h4>
                <h4> Precio: Q {Precio}</h4>
                <h4>Cantidad: {Cantidad}</h4>
                <h4>Subtotal: {Precio*Cantidad}</h4>
                <br/>
                <p className="p-p">
                    <input type="button" value="Eliminar del Carrito" onClick={this.EliminarCarrito} className="btn-upload"/>
                </p>
                <div className="clearfix"> </div>
            </article>
        )
    }

}

export default ProdCar;