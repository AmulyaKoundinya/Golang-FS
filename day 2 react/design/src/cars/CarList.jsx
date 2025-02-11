function CarList(){
    return(
        <>
        <h4>Cars List</h4>
      <div className="container">
      <table className="table table-striped">
        <thead className="table-dark">
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Number</th>
            <th scope="col">Model</th>
            <th scope="col">className</th>
            <th scope="col"></th>

          </tr>
        </thead>
        <tbody>
          <tr>
            <th scope="row">car1</th>
            <td>KA09abc1234</td>
            <td>Zen</td>
            <td>Smaller</td>
            <td><a className="btn btn-success" href="car_view.html?car1">view</a></td>
          </tr>
          <tr>
            <th scope="row">car2</th>
            <td>KA13xyz5678</td>
            <td>Beatle</td>
            <td>Hatchback</td>
            <td><a className="btn btn-success" href="car_view.html?car2">view</a></td>

          </tr>
          
        </tbody>
      </table>
    </div>
        </>
    );
}
export default CarList;