import React, {Component} from 'react';

class Tienda extends Component{

    render(){
        const {Nombre, Descripcion,Contacto, Calificacion, Logo} = this.props.tienda;

        return (
            <article className="article-item" id="article-template">
                <div className="image-wrap">
                    <img
                        src={Logo}  alt={Nombre}/>
                </div>
                <h2>{Nombre}</h2>
                <h4>{Descripcion}</h4>
                <h4>Calificacion: {Calificacion} Contacto: {Contacto}</h4>
                <div className="clearfix"> </div>
            </article>
        )
    }

}

export default Tienda;