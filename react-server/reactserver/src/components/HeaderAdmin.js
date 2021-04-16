import React, {Component} from 'react';
import logo from '../assets/images/logo.png'
import { NavLink } from 'react-router-dom';

class HeaderAdmin extends Component{

    render(){

        return  (
            <header id="header">
                <div className="center">
                    { /*LOGO */}
                    <div id="logo">
                        <img src={logo} className="app-logo" alt="Logotipo"/>
                        <span id="brand"><strong> --Compras </strong>
                    </span>
                    </div>
                    { /*Menu */}
                    <nav id="menu">
                        <ul>
                            <li>
                                <NavLink to="/pedidos" activeClassName="active"> Pedidos</NavLink>
                            </li>
                            <li>
                                <NavLink to="/uploads" activeClassName="active"> Uploads</NavLink>
                            </li>
                            <li>
                                <NavLink to="/usuarios" activeClassName="active"> Usuarios</NavLink>
                            </li>
                        </ul>
                    </nav>

                    <div className="clearfix"> </div>
                </div>
            </header>
        );
    }
}

export default HeaderAdmin;