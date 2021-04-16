import React, {Component} from 'react';
import Tienda from './Tienda';
import SidebarBuscador from "./SidebarBuscador";
import HeaderUsuario from "./HeaderUsuario";

const Server = "http://localhost:3000";

class Tiendas extends Component{

    state = {
        tiendas : [],
        datos : null
    }

    componentDidMount = async() => {
        const response = await fetch(`${Server}/tiendas`);
        const data = await response.json();
        if (data!=="Los datos no han sido ingresados") {
            this.setState({
                tiendas: data
            });
            console.log("lista no es vacia")
            console.log(data)
            this.state.datos = "";
        }else{
            console.log("lista es vacia")
        }
    }

    render(){
        return (
                <div>
                <HeaderUsuario />
                <div className="center">
                    <div id="content">
                        <h2 className="subheader"> Tiendas </h2>

                        {!this.state.tiendas[0] &&
                        <div>
                            <h2> No se han ingresado tiendas </h2>
                        </div>
                        }
                        {/*Componente de Tiendas*/}
                        {this.state.tiendas[0] &&
                        this.state.tiendas.map((tienda, i) => {
                            return (
                                <Tienda key={i} tienda={tienda}/>
                            )
                        })
                        }
                    </div>
                    <SidebarBuscador />
                    <div className="clearfix"> </div>
                </div>
            </div>
        )
    }

}

export default Tiendas;