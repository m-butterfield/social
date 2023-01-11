import React from "react";
import {Route, Routes} from "react-router-dom";

const CreatePost = React.lazy(() => import("app/create_post"));
const Home = React.lazy(() => import("app/home"));
const Login = React.lazy(() => import("app/login"));
const New = React.lazy(() => import("app/new"));
const UserHome = React.lazy(() => import("app/user_home"));

const AppRoutes = () => {
  return <Routes>
    <Route
      path="/"
      element={
        <React.Suspense fallback={<>...</>}>
          <Home />
        </React.Suspense>
      }
    />
    <Route
      path="/create_post"
      element={
        <React.Suspense fallback={<>...</>}>
          <CreatePost />
        </React.Suspense>
      }
    />
    <Route
      path="/login"
      element={
        <React.Suspense fallback={<>...</>}>
          <Login />
        </React.Suspense>
      }
    />
    <Route
      path="/new"
      element={
        <React.Suspense fallback={<>...</>}>
          <New />
        </React.Suspense>
      }
    />
    <Route path="/:userName" element={
      <React.Suspense fallback={<>...</>}>
        <UserHome />
      </React.Suspense>
    } />
  </Routes>;
};

export default AppRoutes;
