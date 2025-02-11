import CarCreate from "./car/CarCreate";
import CarList from "./car/CarList";
import CarView from "./car/CarView";
import Great from "./header/Great";

function App(){
  return(
    <>
    <h1><marquee >Car Parking Management</marquee></h1>
    <Great/>
    <CarCreate/>
    <CarList/>
    <CarView/>

    </>
  );
}
export default App;