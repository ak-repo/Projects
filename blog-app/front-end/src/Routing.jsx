import React from "react";
import { Route, Routes } from "react-router-dom";
import PostListingPage from "./components/public/PostListingPage";
import AuthPage from "./components/auth/AuthPage";
import PostDetailPage from "./components/public/PostDetailPage.JSX";
import PageNotFound from "./components/404/PageNotFound";

function Routing() {
  return (
    <Routes>
      <Route path="/" element={<PostListingPage />} />
      <Route path="/postinfo" element={<PostDetailPage />} />
      <Route path="/login" element={<AuthPage />} />
      <Route path="*" element={<PageNotFound />} />
    </Routes>
  );
}

export default Routing;
