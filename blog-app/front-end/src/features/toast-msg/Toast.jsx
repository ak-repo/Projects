import { Toaster } from "react-hot-toast";

function ToastMessage() {
  return (
    <>
      {/* your routes/components here */}
      <Toaster
        position="top-right"
        toastOptions={{
          duration: 3000,
          style: {
            background: "#333",
            color: "#fff",
          },
        }}
      />
    </>
  );
}

export default ToastMessage;

