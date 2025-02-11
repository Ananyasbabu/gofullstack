function Great() {
    return (
        <>
            <nav classname="navbar navbar-expand-lg navbar-dark bg-dark">
        <div classname="container-fluid">
          <a classname="navbar-brand" href="/car/list">Parking Management System</a>
          <button classname="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span classname="navbar-toggler-icon"></span>
          </button>
          <div classname="collapse navbar-collapse" id="navbarSupportedContent">
            <ul classname="navbar-nav me-auto mb-2 mb-lg-0">
              <li classname="nav-item">
                <a classname="nav-link " aria-current="page" href="/car/list">Cars List</a>
              </li>
              <li classname="nav-item">
                <a classname="nav-link active" href="/car/create">Add Cars</a>
              </li>
              
            </ul>
            
          </div>
        </div>
    </nav>


        </>
    );
}
export default Great;