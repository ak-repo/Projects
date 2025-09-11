import React from "react";
import { Route, Routes } from "react-router-dom";
import PostListingPage from "./components/PostListingPage";
import AuthPage from "./components/AuthPage";

function Routing() {
  return (
    <Routes>
      <Route path="/" element={<PostListingPage />} />
      <Route path="/login" element={<AuthPage />} />
    </Routes>
  );
}

export default Routing;
