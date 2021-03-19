import logo from './assets/images/logo.svg';
import './assets/css/App.css';

//importar componentes
import Tiendas from "./components/Tiendas";
import Header from "./components/Header";
import SidebarBuscador from "./components/SidebarBuscador";
import Footer from "./components/Footer";
import SeccionPrueba from "./components/SeccionPrueba";
import Router from "./Router";

function App() {
  return (
    <div className="App">

        <Router />
            {/*<Tiendas />*/}

        <Footer />
    </div>
  );
}

export default App;
