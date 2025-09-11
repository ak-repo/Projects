import { BrowserRouter } from "react-router-dom";
import PostListingPage from "./components/PostListingPage";
import Routing from "./Routing";

function App() {
  return (
    <div>
      <BrowserRouter>
        <Routing />
      </BrowserRouter>
    </div>
  );
}

export default App;
