import React, {Component} from 'react';

class SeccionPrueba extends Component{

    contador = 0;

    state = {
        contador: 0
    };

    sumar = () => {
        this.setState({
                contador: (this.state.contador + 1)
            });
    }

    restar = () => {
        this.setState({
                contador: (this.state.contador - 1)
            });
    }

    render() {
        return(
            <section id="content">
                <h2 className="subheader">Últimos artículos</h2>
                <p>
                    Contador: {this.state.contador}
                </p>
                <p>
                    <input type="button" value="Sumar" onClick={this.sumar}/>
                    <input type="button" value="Restar" onClick={this.restar}/>
                </p>

                <section className="componentes" >
                </section>
            </section>
        );
    }
}

export default SeccionPrueba;