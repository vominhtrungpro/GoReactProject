import "./App.css";
import React from "react";
import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useGoogleLogin, GoogleLogin } from "@react-oauth/google";
import jwt_decode from "jwt-decode";

function App() {
  const responseMessage = (response) => {
    localStorage.setItem(
      "google-access-token",
      JSON.stringify(response.credential)
    );
    setEmail({
      email: response.credential,
    });
    handleLoginGoogle(response.credential);
  };

  const handleLoginGoogle = (token) => {
    axios
      .get(`http://localhost:9090/login-google/` + token)
      .then(function (response) {
        if (!response.data.Responsetype) {
          navigate("/register-google");
        } else {
          localStorage.setItem(
            "access-token",
            JSON.stringify(response.data.Accesstoken)
          );
          navigate("/home");
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const errorMessage = (error) => {
    console.log(error);
  };

  const navigate = useNavigate();
  const [email, setEmail] = useState({
    email: "",
  });
  useEffect(() => {
    console.log(email);
  }, [email]);
  const [data, setData] = useState({
    username: "",
    password: "",
  });
  const changeHandler = (e) => {
    setData({ ...data, [e.target.name]: e.target.value });
  };
  const handleSubmit = (e) => {
    e.preventDefault();
    axios
      .post(`http://localhost:9090/login`, data)
      .then(function (response) {
        if (!response.data.Responsetype) {
          console.log(response.data.Message);
        } else {
          localStorage.setItem(
            "access-token",
            JSON.stringify(response.data.Accesstoken)
          );
          navigate("/home");
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  return (
    <div className="App">
      <main className="form-signin w-100 m-auto">
        <form onSubmit={(e) => handleSubmit(e)}>
          <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
          <div className="form-floating">
            <input
              onChange={changeHandler}
              className="form-control"
              name="username"
              placeholder="name@example.com"
            />
            <label for="floatingInput">Email address</label>
          </div>
          <div className="form-floating">
            <input
              onChange={changeHandler}
              type="password"
              name="password"
              className="form-control"
              placeholder="Password"
            />
            <label for="floatingPassword">Password</label>
          </div>
          <button class="w-100 btn btn-lg btn-primary" type="submit">
            Login
          </button>
        </form>
        <GoogleLogin onSuccess={responseMessage} onError={errorMessage} />
      </main>
    </div>
  );
}

export default App;
