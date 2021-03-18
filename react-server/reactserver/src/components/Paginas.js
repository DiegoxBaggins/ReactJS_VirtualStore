import React, {Component} from 'react';
import Tienda from './Tienda';

const Server = "http://localhost:3000";

class Paginas extends Component{

    state = {
        tiendas : []
    }

    constructor(props){
        super(props);
        console.log("Primer elemento")
    }

    componentDidMount = async() => {
        const response = await fetch(`${Server}/tiendas`);
        const data = await response.json();
        this.setState({
            tiendas: data
        });
        console.log(data)
    }

    render(){

        return (
            <div id="content">
                <h2 className="subheader"> Tiendas </h2>

                {/*Componente de Paginas*/}
                <div id="articles" className="Tiendas">
                {
                    this.state.tiendas.map((tienda, i) => {
                        return (
                            <Tienda key={i} tienda={tienda}/>
                        )
                    })
                }
                </div>
            </div>
        )
    }

}

export default Paginas;