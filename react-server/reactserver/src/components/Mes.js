import React, {Component} from 'react';
import axios from "axios";
import {Link} from "react-router-dom";
import {Redirect} from "react-router";


class Mes extends Component{

    devolverMes = () => {
        const {Nombre, Numero} = this.props.mes;
        this.props.DevolverAnio(Nombre, Numero)
    }

    render(){

        const {Nombre} = this.props.mes;

        return (
            <input type="button" value={Nombre} onClick={this.devolverMes} className="btn-mes"/>
        )
    }

}

export default Mes;