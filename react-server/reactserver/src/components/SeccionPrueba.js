import React, {Component} from 'react';

const Server = "http://127.0.0.1:3000";

class SeccionPrueba extends Component{

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