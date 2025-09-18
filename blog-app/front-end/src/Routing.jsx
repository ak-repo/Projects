import React from "react";
import { Route, Routes } from "react-router-dom";
import PostListingPage from "./components/public/PostListingPage";
import AuthPage from "./components/auth/AuthPage";
import PostDetailPage from "./components/public/PostDetailPage.jsx";
import PageNotFound from "./components/404/PageNotFound";

import store from "./state/store/store";
import { Provider } from "react-redux";

function Routing() {
  return (
    <Provider store={store}>
      <Routes>
        <Route path="/" element={<PostListingPage />} />
        <Route path="/postinfo" element={<PostDetailPage />} />
        <Route path="/login" element={<AuthPage />} />
        <Route path="*" element={<PageNotFound />} />
      </Routes>
    </Provider>
  );
}

export default Routing;
