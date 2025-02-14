import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import PageHeader from "../header/PageHeader";
import axios from 'axios';

function EmployeeEdit() {
    const [employee, setEmployee] = useState({id:'', name:'', dept:'', position:''});
    const params = useParams();
    const navigate = useNavigate();

    const txtBoxOnChange = event => {
        const updatableEmployee = {...employee};
        updatableEmployee[event.target.id] = event.target.value;
        setEmployee(updatableEmployee);
    };

    const readById = async () => {
        const baseUrl = "http://localhost:8080";
        try {
            const response = await axios.get(`${baseUrl}/employees/${params.id}`);
            const queriedEmployee = response.data;
            setEmployee(queriedEmployee);
        } catch(error) {
            alert('Server Error');
        }
    };

    const updateEmployee = async () => {
        const baseUrl = "http://localhost:8080";
        try {
            const response = await axios.put(`${baseUrl}/employees/${params.id}`, {...employee});
            const updatedEmployee = response.data.employee;
            setEmployee(updatedEmployee);
            alert(response.data.message);
            navigate('/employees/list');
        } catch(error) {
            alert('Server Error');
        }
    };

    useEffect(() => {
        readById();
    }, []);

    return(
        <>
            <PageHeader/>
            <h3><a href="/employees/list" className="btn btn-light">Go Back</a>Edit Employee</h3>
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
                <button className="btn btn-warning"
                    onClick={updateEmployee}>Update Employee</button>
            </div>
        </>
    );
}

export default EmployeeEdit;
