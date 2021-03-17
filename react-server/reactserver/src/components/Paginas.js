import React, {Component} from 'react';
import Tienda from './Tienda';

class Paginas extends Component{

    state = {
        tiendas: [
            {
                Nombre: "Amaya, Rolón y Chavarría Asociados",
                Descripcion: "Illum accusantium voluptate voluptatem in corrupti dolorem velit et.",
                Contacto: "976191834",
                Calificacion: 5,
                Logo:"https://economipedia.com/wp-content/uploads/2015/10/apple-300x300.png"
            },
            {
                Nombre: "Aguayo Cepeda S.A.",
                Descripcion: "Numquam ea est error inventore et porro veritatis.",
                Contacto: "949 586 354",
                Calificacion: 4,
                Logo:"https://i.pinimg.com/originals/3d/0f/0e/3d0f0e8f600627fde858f6c6e668e999.gif"
            }
        ],
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