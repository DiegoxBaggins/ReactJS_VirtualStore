import logo from './assets/images/logo.svg';
import './assets/css/App.css';

//importar componentes
import Paginas from "./components/Paginas";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import Footer from "./components/Footer";
import SeccionPrueba from "./components/SeccionPrueba";
import Router from "./Router";

function App() {
  return (
    <div className="App">



        <div className="center">
            <Router />
            {/*<Paginas />*/}

            <Sidebar />
            <div className="clearfix"> </div>
        </div>  {/* end div center */}

        <Footer />
    </div>
  );
}

export default App;