import Greetings from "./greetings";
import Navbar from "./header/Navbar";

export default function App() {
  return (
    <>
      <Navbar/>
      <h1><marquee behavior="scroll" direction="center" scrollamount="50"> Welcome to React Ananya</marquee></h1>
      <Greetings />
    </>
  );
}
// export default App;