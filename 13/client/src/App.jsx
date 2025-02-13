// import CarList from './cars/CarList'
// import CarCreate from './cars/CarCreate'
// import CarView from './cars/CarView'

import { BrowserRouter, Route, Routes } from 'react-router-dom'
import CarCreate from './cars/carCreate';
import CarEdit from './cars/carEdit';
import CarView from './cars/carView';
import CarList from './cars/carList';
// import CarEdit from './cars/CarEdit';

function App() {
  return (
    <>
  
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="" element={<CarList/>}/>
            <Route path="/cars/list" element={<CarList/>}/>
            <Route path="/cars/create" element={<CarCreate/>}/>
            <Route path="/cars/view/:id" element={<CarView/>}/>
            <Route path="/cars/edit/:id" element={<CarEdit/>}/>
          </Routes>
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;