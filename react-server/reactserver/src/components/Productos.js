import React, {Component} from 'react';
import SidebarBuscador from "./SidebarBuscador";
import Producto from "./Producto";
import axios from "axios";

const Server = "http://127.0.0.1:3000";

class Productos extends Component{

    state = {
        tienda : {},
        productos : []
    }

    componentDidMount = () => {
        this.EstablecerDatos();
        const direccion = localStorage.getItem('Departamento') + "/" + localStorage.getItem('Calificacion') + "/" + localStorage.getItem('Nombre')
        axios.post(`${Server}/productos/${direccion}`).then( (response) => {
            console.log(response);
            if (response.data !== "No hay productos") {
                this.setState({
                    productos: response.data
                });
                console.log("lista no es vacia")
                console.log(response.data)
            }else{
                console.log("lista es vacia")
            }
        });
    }

    EstablecerDatos = async() => {
        const Nombre = localStorage.getItem('Nombre');
        const Calificacion = localStorage.getItem('Calificacion');
        const Departamento = localStorage.getItem('Departamento');
        console.log(Nombre);
        await this.setState({
                           tienda : {
                               Nombre: Nombre,
                               Calificacion: Calificacion,
                               Departamento: Departamento,
                           }
                       },() => (console.log(this.state.tienda.Nombre, this.state.tienda.Calificacion, this.state.tienda.Departamento)));
    }

    render(){

        return (
            <div className="center">
                <div id="content">
                    <h2 className="subheader">Tienda: {this.state.tienda.Nombre} </h2>
                    <h2 className="subheader"> Productos </h2>
                    {!this.state.productos[0] &&
                    <div>
                        <h2> La tienda no tiene productos </h2>
                    </div>
                    }
                    {this.state.productos[0] &&
                    this.state.productos.map((producto, i) => {
                        return (
                            <Producto key={i} producto={producto}/>
                        )
                    })
                    }
                </div>
                <SidebarBuscador />
                <div className="clearfix"> </div>
            </div>
        )
    }

}

export default Productos;