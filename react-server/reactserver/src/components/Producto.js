import React, {Component} from 'react';

class Producto extends Component{

    render(){
        const {Nombre, Codigo, Descripcion, Precio, Cantidad, Imagen} = this.props.producto;

        return (
            <article className="article-item" id="article-template">
                <div className="image-wrap">
                    <img
                        src={Imagen}  alt={Nombre}/>
                </div>
                <h2>  Nombre: {Nombre}  Codigo: {Codigo}</h2>
                <h4>{Descripcion}</h4>
                <h4> Precio: {Precio}</h4>
                <h4> Disponibilidad: {Cantidad}</h4>
                <div className="clearfix"> </div>
            </article>
        )
    }

}

export default Producto;