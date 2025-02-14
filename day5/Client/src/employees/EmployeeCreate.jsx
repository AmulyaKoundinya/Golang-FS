import { useState } from "react";
import PageHeader from "../header/PageHeader";
import { useNavigate } from "react-router-dom";
import axios from 'axios';

function EmployeeCreate() {
    const [employee, setEmployee] = useState({id:'', name:'', dept:'', position:''});
    const navigate = useNavigate();
    
    const txtBoxOnChange = event => {
        const updatableEmployee = {...employee};
        updatableEmployee[event.target.id] = event.target.value;
        setEmployee(updatableEmployee);
    };

    const createEmployee = async () => {
        const baseUrl = "http://localhost:8080";
        try {
            const response = await axios.post(`${baseUrl}/employees`, {...employee});
            const createdEmployee = response.data.employee;
            setEmployee(createdEmployee);
            alert(response.data.message);
            navigate('/employees/list');
        } catch(error) {
            alert('Server Error');
        }
    };

    return(
        <>
            <PageHeader/>            
            <h3><a href="/employees/list" className="btn btn-light">Go Back</a> Add Employee</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label for="name" className="form-label">Employee Name:</label>
                    <input type="text" className="form-control" id="name" 
                        placeholder="Please enter employee name"
                        value={employee.name} 
                        onChange={txtBoxOnChange}/>
                </div>
                <div className="form-group mb-3">
                    <label for="dept" className="form-label">Employee Department:</label>
                    <input type="text" className="form-control" id="dept" 
                        placeholder="Please enter employee department"
                        value={employee.dept} 
                        onChange={txtBoxOnChange}/>
                </div>
                <div className="form-group mb-3">
                    <label for="position" className="form-label">Employee Position:</label>
                    <input type="text" className="form-control" id="position" 
                        placeholder="Please enter employee position"
                        value={employee.position} 
                        onChange={txtBoxOnChange}/>
                </div>
                <button className="btn btn-primary"
                    onClick={createEmployee}>Create Employee</button>
            </div>
        </>
    );
}

export default EmployeeCreate;
