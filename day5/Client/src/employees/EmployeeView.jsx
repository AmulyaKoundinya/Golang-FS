import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import PageHeader from "../header/PageHeader";
import axios from 'axios';

function EmployeeView() {
    const [employee, setEmployee] = useState({id: '', name: '', dept: '', position: ''});
    const params = useParams();

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

    useEffect(() => {
        readById();
    }, []);

    return (
        <>
            <PageHeader />
            <h3><a href="/employees/list" className="btn btn-light">Go Back</a>View Employee</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="name" className="form-label">Employee Name:</label>
                    <div className="form-control" id="name">{employee.name}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="dept" className="form-label">Department:</label>
                    <div className="form-control" id="dept">{employee.dept}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="position" className="form-label">Position:</label>
                    <div className="form-control" id="position">{employee.position}</div>
                </div>
            </div>
        </>
    );
}

export default EmployeeView;
