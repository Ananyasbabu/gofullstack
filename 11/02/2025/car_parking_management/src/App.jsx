import CarCreate from "./car/CarCreate";
import CarList from "./car/CarList";
import CarView from "./car/CarView";
// import Great from "./header/Great";
import {BrowserRouter,Routes,Route} from "react-router-dom";

function App(){
  return(
    <>
      <BrowserRouter>
      <Routes>
        <Route path="" element={<CarList/>}/>
        <Route path="/car/list" element={<CarList/>}/>
        <Route path="/car/create" element={<CarCreate/>}/>
        <Route path="/car/view" element={<CarView/>}/>
        </Routes>
      </BrowserRouter>
    </>
  );
}
export default App;