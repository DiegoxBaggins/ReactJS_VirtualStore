import React, {Component} from 'react';
import axios from "axios";
import ProdCar from "./ProdCar";
import HeaderUsuario from "./HeaderUsuario";

const Server = "http://127.0.0.1:3000";

class Carrito extends Component{

    state = {
        productos : [],
        total: 0
    }

    componentDidMount = async() => {
        await axios.get(`${Server}/carrito`).then( (response) => {
            console.log(response);
            if (response.data !== "No hay Articulos") {
                this.setState({
                    productos: response.data
                });
                console.log("lista no es vacia")
                console.log(response.data)
                this.setState({
                    total: 0
                });
                response.data.map((producto, i) => {
                    this.setState({
                        total: (this.state.total + (producto.Cantidad*producto.Precio))
                    });
                })
            }else{
                console.log("lista es vacia")
            }
        });
    }

    PagarCarrito = () => {
        axios.get(`${Server}/PagarCarrito`).then( (response) => {
            console.log(response);
        });
        this.setState({
            productos: [],
            total: 0
        });
        let var1 = {
            User : localStorage.getItem("DPI"),
            Total: String(this.state.total)
        }
        axios.post(`${Server}/generarTransaccion`, var1).then( (response) => {
            console.log(response);
        });
    }

    render(){

        return (
            <div>
                <HeaderUsuario />
                <div className="center">
                    <div id="content">
                        <h2 className="subheader"> Productos En el carrito</h2>
                        {!this.state.productos[0] &&
                        <div>
                            <h2> No hay productos en el Carrito </h2>
                        </div>
                        }
                        {this.state.productos[0] &&
                        this.state.productos.map((producto, i) => {
                            return (
                                <ProdCar key={i} producto={producto}/>
                            )
                        })
                        }
                    </div>
                    <aside id="sidebar">
                        <div id="search" className="sidebar-item">
                            <h2>Total de la compra:</h2>
                            <h2>Q{this.state.total}.00</h2>
                            <p className="p-p">
                                <input type="button" value="Proceder con el pedido" onClick={this.PagarCarrito} className="btn-upload"/>
                            </p>
                        </div>
                    </aside>
                    <div className="clearfix"> </div>
                </div>
            </div>
        )
    }

}

export default Carrito;