import { createBrowserRouter, RouterProvider } from "react-router-dom";

import App from "../App";
import Home from "../pages/Home/Home";
import Users from "../pages/Users/UserPage";
import Crawl from "../pages/Crawl/Crawl";
import RegisterGoogle from "../pages/RegisterGoogle/RegisterGoogle";
import { GoogleOAuthProvider } from "@react-oauth/google";
import React from 'react';


const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <GoogleOAuthProvider clientId="326313607776-macv5cle3gs1ehtb53rcm95je9rd3tlh.apps.googleusercontent.com">
        <React.StrictMode>
            <App />
        </React.StrictMode>
      </GoogleOAuthProvider>
    ),
  },
  {
    path: "/users",
    element: <Users />,
  },
  {
    path: "/home",
    element: <Home />,
  },
  {
    path: "/crawl",
    element: <Crawl />,
  },
  {
    path: "/register-google",
    element: <RegisterGoogle />,
  },
]);

export default router;
