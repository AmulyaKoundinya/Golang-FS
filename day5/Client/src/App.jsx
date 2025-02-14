import CarList from './employees/EmployeeList'
import CarCreate from './employees/EmployeeCreate'
import CarView from './employees/CarView'

import { BrowserRouter, Route, Routes } from 'react-router-dom'
import CarEdit from './employees/EmployeeEdit';

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
