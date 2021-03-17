import React, {Component} from 'react';
import logo from '../assets/images/logo.png'
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
                                <a href="index.html">Inicio</a>
                            </li>
                            <li>
                                <a href="blog.html">Carrito</a>
                            </li>
                            <li>
                                <a href="formulario.html">Administrador</a>
                            </li>
                        </ul>
                    </nav>
                    { /*LIMPIAR FLOTADOS */}

                    <div className="clearfix"></div>
                </div>
            </header>
        );
    }
}

export default Header;