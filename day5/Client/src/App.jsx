import EmployeeList from './employees/EmployeeList';
import EmployeeCreate from './employees/EmployeeCreate';
import EmployeeView from './employees/EmployeeView';
import EmployeeEdit from './employees/EmployeeEdit';

import { BrowserRouter, Route, Routes } from 'react-router-dom';

function App() {
  return (
    <>
  
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="" element={<EmployeeList />} />
            <Route path="/employees/list" element={<EmployeeList />} />
            <Route path="/employees/create" element={<EmployeeCreate />} />
            <Route path="/employees/view/:id" element={<EmployeeView />} />
            <Route path="/employees/edit/:id" element={<EmployeeEdit />} />
          </Routes>
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;
