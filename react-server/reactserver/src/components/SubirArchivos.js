import React, {Component} from 'react';
import axios from 'axios';
import HeaderAdmin from "./HeaderAdmin";

const Server = "http://localhost:3000";

class SubirArchivos extends Component{

    state = {
        tiendas : [],
        selectedFile: null,
        datos : null
    }

    onFileChange = event => {
        // Update the state
        this.setState({ selectedFile: event.target.files[0] });
    };

    // On file upload (click the upload button)
    onFileUploadTienda = () => {
        const formData = new FormData();
        formData.append(
            "myFile",
            this.state.selectedFile,
            this.state.selectedFile.name
        );
        console.log(this.state.selectedFile);
        axios.post(`${Server}/cargartiendas`, formData).then(function (response) {
            console.log(response);
        });
    };

    onFileUploadInven = () => {
        const formData = new FormData();
        formData.append(
            "myFile",
            this.state.selectedFile,
            this.state.selectedFile.name
        );
        console.log(this.state.selectedFile);
        axios.post(`${Server}/cargarinventario`, formData).then(function (response) {
            console.log(response);
        });
    };

    onFileUploadPedido = () => {
        const formData = new FormData();
        formData.append(
            "myFile",
            this.state.selectedFile,
            this.state.selectedFile.name
        );
        console.log(this.state.selectedFile);
        axios.post(`${Server}/cargarpedidos`, formData).then(function (response) {
            console.log(response);
        });
    };

    onFileUploadUser = () => {
        const formData = new FormData();
        formData.append(
            "myFile",
            this.state.selectedFile,
            this.state.selectedFile.name
        );
        console.log(this.state.selectedFile);
        axios.post(`${Server}/cargarusuarios`, formData).then(function (response) {
            console.log(response);
        });
    };

    onFileUploadGrafo = () => {
        const formData = new FormData();
        formData.append(
            "myFile",
            this.state.selectedFile,
            this.state.selectedFile.name
        );
        console.log(this.state.selectedFile);
        axios.post(`${Server}/cargargrafo`, formData).then(function (response) {
            console.log(response);
        });
    };

    // File content to be displayed after
    // file upload is complete
    fileData = () => {
        if (this.state.selectedFile) {
            return (
                <div>
                    <h2>File Details:</h2>
                    <p>File Name: {this.state.selectedFile.name}</p>
                    <p>File Type: {this.state.selectedFile.type}</p>
                    <p>
                        Last Modified:{" "}
                        {this.state.selectedFile.lastModifiedDate.toDateString()}
                    </p>
                </div>
            );
        } else {
            return (
                <div>
                    <br />
                    <h4>Choose before Pressing the Upload button</h4>
                </div>
            );
        }
    };

    render(){
        return (
            <div>
                <HeaderAdmin />
                <div className="center">
                    <div id="content">
                        <h2 className="subheader"> Lectura de Archivos </h2>
                        <div>
                            <h2> Ingresar tiendas </h2>
                            <input type="file" onChange={this.onFileChange} className="btn-upload" />
                            <button onClick={this.onFileUploadTienda} className="btn-upload"> Subir </button>
                            <hr/>
                            <h2> Ingresar inventario </h2>
                            <input type="file" onChange={this.onFileChange} className="btn-upload" />
                            <button onClick={this.onFileUploadInven} className="btn-upload"> Subir </button>
                            <hr/>
                            <h2> Ingresar pedidos </h2>
                            <input type="file" onChange={this.onFileChange} className="btn-upload" />
                            <button onClick={this.onFileUploadPedido} className="btn-upload"> Subir </button>
                            <hr/>
                            <h2> Ingresar Usuarios </h2>
                            <input type="file" onChange={this.onFileChange} className="btn-upload" />
                            <button onClick={this.onFileUploadUser} className="btn-upload"> Subir </button>
                            <hr/>
                            <h2> Ingresar Grafo </h2>
                            <input type="file" onChange={this.onFileChange} className="btn-upload" />
                            <button onClick={this.onFileUploadGrafo} className="btn-upload"> Subir </button>
                            <hr/>
                        </div>
                    </div>
                    <aside id="sidebar">
                        <div id="search" className="sidebar-item">
                            {this.fileData()}
                        </div>
                    </aside>
                    <div className="clearfix"> </div>
                </div>
            </div>
        )
    }

}

export default SubirArchivos;