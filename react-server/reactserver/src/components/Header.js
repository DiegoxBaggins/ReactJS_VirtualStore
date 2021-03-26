import React, {Component} from 'react';
import logo from '../assets/images/logo.png'
import { NavLink } from 'react-router-dom';

class Header extends Component{

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
                                <NavLink to="/Home" activeClassName="active">Inicio</NavLink>
                            </li>
                            <li>
                                <NavLink to="/Carrito" activeClassName="active">Carrito</NavLink>
                            </li>
                            <li>
                                <NavLink to="/Admin" activeClassName="active"> Administrador</NavLink>
                            </li>
                            <li>
                                <NavLink to="/uploads" activeClassName="active"> Uploads</NavLink>
                            </li>
                        </ul>
                    </nav>

                    <div className="clearfix"> </div>
                </div>
            </header>
        );
    }
}

export default Header;