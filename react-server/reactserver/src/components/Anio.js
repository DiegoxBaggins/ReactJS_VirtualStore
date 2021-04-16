import React, {Component} from 'react';
import Mes from "./Mes";


class Anio extends Component{

    DevolverAnio = (mes, mesnum) => {
        const {Nombre} = this.props.anio;
        this.props.DevolverMatriz(Nombre, mes, mesnum)
    }

    render(){

        const {Nombre, Meses} = this.props.anio;

        return (
            <div>
                <h4>{Nombre}</h4>
                <div className="siderbar-div">
                    {Meses.map((mes, i) => {
                        return(
                            <Mes key={i} mes={mes} DevolverAnio={this.DevolverAnio}/>
                        )
                    })
                    }
                </div>
            </div>
        )
    }

}

export default Anio;