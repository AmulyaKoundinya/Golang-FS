function CarCreate(){
    return(
        <>
        <h3><a href="cars_list.html" className="btn btn-light">Go back</a>Add Cars</h3>
    <div className="container bg-primary-subtle">
      <div className="htmlForm-group mb-3">
        <label htmlFor="number" className="htmlForm-label">Car Number:</label>
        <input type="text" className="htmlForm-control" id="number" placeholder="Please enter car number:"/>
      </div>
      <div className="htmlForm-group mb-3">
        <label htmlFor="model" className="htmlForm-label">Car Model:</label>
        <input type="text" className="htmlForm-control" id="model" placeholder="Please enter car model:"/>
      </div>
      <div className="htmlForm-group mb-3">
        <label htmlFor="type" className="htmlForm-label">Car Type:</label>
        <input type="text" className="htmlForm-control" id="type" placeholder="Please enter car type:"/>
      </div>
      <button className="btn btn-danger">Create Cars</button>
    </div>
        </>
    );
}
export default CarCreate;