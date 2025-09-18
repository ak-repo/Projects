import { BrowserRouter } from "react-router-dom";
import Routing from "./Routing";
import ToastMessage from "./features/toast-msg/Toast";

function App() {
  return (
    <div>
      <BrowserRouter>
        <ToastMessage />
        <Routing />
      </BrowserRouter>
    </div>
  );
}

export default App;
