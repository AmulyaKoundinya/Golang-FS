import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import axios from 'axios';

function EmployeeList() {
    const [employees, setEmployees] = useState([{id: '', name: '', dept: '', position: ''}]);

    const readAllEmployees = async () => {
        try {
            const baseUrl = 'http://localhost:8080';
            const response = await axios.get(`${baseUrl}/employees`);
            const queriedEmployees = response.data;
            setEmployees(queriedEmployees);
        } catch(error) {
            alert('Server Error');
        }
    };

    const deleteEmployee = async (id) => {
        if(!window.confirm("Are you sure to delete?")) {
            return;
        }
        const baseUrl = "http://localhost:8080";
        try {
            const response = await axios.delete(`${baseUrl}/employees/${id}`);
            alert(response.data.message);
            await readAllEmployees();
        } catch(error) {
            alert('Server Error');
        }
    };

    useEffect(() => {
        readAllEmployees();
    }, []);

    return (
        <>
            <PageHeader />
            <h3>List of Employees</h3>
            <div className="container">
                <table className="table table-success table-striped">
                    <thead className="table-dark">
                        <tr>
                            <th scope="col">ID</th>
                            <th scope="col">Employee Name</th>
                            <th scope="col">Department</th>
                            <th scope="col">Position</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {(employees && employees.length > 0) ? employees.map(
                            (employee) => {
                                return (
                                    <tr key={employee.id}>
                                        <th scope="row">{employee.id}</th>
                                        <td>{employee.name}</td>
                                        <td>{employee.dept}</td>
                                        <td>{employee.position}</td>
                                        <td>
                                            <a href={`/employees/view/${employee.id}`} 
                                                className="btn btn-success">View</a>
                                            &nbsp;
                                            <a href={`/employees/edit/${employee.id}`} 
                                                className="btn btn-warning">Edit</a>
                                            &nbsp;
                                            <button  
                                                className="btn btn-danger"
                                                onClick={() => deleteEmployee(employee.id)}>
                                                Delete
                                            </button>
                                        </td>
                                    </tr>
                                );
                            }
                        ) : (
                            <tr>
                                <td colSpan="5">No Data Found</td>
                            </tr>
                        )}
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default EmployeeList;
