import React, {Component} from 'react';
import axios from "axios";
import {Link} from "react-router-dom";
import {Redirect} from "react-router";

const Server = "http://127.0.0.1:3000";

class Producto extends Component{

    state = {
        contador: 0,
        cantidad: 0,
        redirect: null
    };

    sumar = () => {
        if (this.state.contador !== this.state.cantidad){
            this.setState({
                contador: (this.state.contador + 1)
            });
        }
    }

    restar = () => {
        if (this.state.contador !== 0){
            this.setState({
                contador: (this.state.contador - 1)
            });
        }
    }

    AgregarCarrito = async() => {
        const {Nombre, Codigo, Descripcion, Precio, Imagen} = this.props.producto;
        console.log(this.state.contador)
        if (this.state.contador !== 0){
            let product = {
                Nombre: Nombre,
                Codigo: Codigo,
                Descripcion: Descripcion,
                Precio: Precio,
                Cantidad: this.state.contador,
                Imagen: Imagen,
                Tienda: localStorage.getItem('Nombre'),
                Calificacion: Number(localStorage.getItem('Calificacion')),
                Departamento: localStorage.getItem('Departamento')
            };
            console.log(product)
            await axios.post(`${Server}/agregarCarrito`, product).then(function (response) {
                console.log(response);
            }).finally(() => {
                alert("Agregado al carrito");
                this.setState({
                    redirect: "/Carrito"
                })
            });
        }
    }

    componentDidMount = () => {
        const {Cantidad} = this.props.producto;
        if (Cantidad === undefined ) {
            this.setState({
                    cantidad: 0
                }
            )
        }else{this.setState({
                cantidad: Cantidad
            }
        )}
    }

    render(){

        if (this.state.redirect) {
            return <Redirect to={this.state.redirect} />
        }

        const {Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen} = this.props.producto;

        return (
            <article className="article-item" id="article-template">
                <h3>{Nombre} | Codigo: {Codigo}</h3>
                <div className="image-wrap">
                    <img
                        src={Imagen}  alt={Nombre}/>
                </div>
                <h4>{Descripcion}</h4>
                <h4> Precio: Q {Precio}</h4>
                <h4>Disponibilidad: {Cantidad}</h4>
                <br/>
                <p className="p-p">
                    <input type="button" value="-" onClick={this.restar} className="btn-sum" />
                    {this.state.contador}
                    <input type="button" value="+" onClick={this.sumar} className="btn-sum" />
                    <input type="button" value="Agregar Al Carrito" onClick={this.AgregarCarrito} className="btn-upload"/>
                </p>
                <div className="clearfix"> </div>
            </article>
        )
    }

}

export default Producto;