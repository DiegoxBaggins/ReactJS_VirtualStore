import React, {Component} from 'react';

class SidebarBuscador extends Component{

    render(){

        return (
            <aside id="sidebar">

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

export default SidebarBuscador;