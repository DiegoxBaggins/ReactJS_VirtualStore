import React, {Component} from 'react';

class Sidebar extends Component{

    render(){

        return (
            <aside id="sidebar">
                {/*
                <div id="nav-blog" className="sidebar-item">
                    <h3>Agregar JSON tiendas</h3>
                    <a href="#" className="btn btn-success">Ir</a>
                </div>
                */}
                <div id="search" className="sidebar-item">
                    <h3>Buscar tienda</h3>
                    <form>
                        <input type="text" name="search"/>
                        <input type="submit" name="submit" value="Buscar" className="btn"/>
                    </form>
                </div>
            </aside>
        );
    }
}

export default Sidebar;