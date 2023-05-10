import "./RegisterGoogleForm.css";
import jwt_decode from "jwt-decode";
import React from "react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function RegisterGoogleForm() {
  function Logout() {
    localStorage.removeItem("access-token");
    localStorage.removeItem("google-access-token");
    navigate("/");
  }
  const token = JSON.parse(localStorage.getItem("google-access-token"));
  const navigate = useNavigate();
  const decoded = jwt_decode(token);
  const [data, setData] = useState({
    firstname: "",
    lastname: "",
    username: "",
    email: decoded.email,
    password: "",
  });
  const changeHandler = (e) => {
    setData({ ...data, [e.target.name]: e.target.value });
  };
  const handleSubmit = (e) => {
    e.preventDefault();
    axios
      .post(`http://localhost:9090/user`, data)
      .then(function (response) {
        if (!response.data.Resulttype) {
          console.log(response.data.Message);
        } else {
          navigate("/");
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  return (
    <div>
      <div className="container h-100">
        <div className="row d-flex justify-content-center align-items-center h-100">
          <div className="col-lg-12 col-xl-11">
            <div className="card text-black">
              <div className="card-body p-md-5">
                <div className="row justify-content-center">
                  <div className="col-md-10 col-lg-6 col-xl-5 order-2 order-lg-1">
                    <p className="text-center h1 fw-bold mb-5 mx-1 mx-md-4 mt-4">
                      Sign up
                    </p>
                    <form
                      className="mx-1 mx-md-4"
                      onSubmit={(e) => handleSubmit(e)}
                    >
                      <div className="d-flex flex-row align-items-center mb-4">
                        <i className="fas fa-user fa-lg me-3 fa-fw"></i>
                        <div className="form-outline flex-fill mb-0">
                          <input
                            onChange={changeHandler}
                            className="form-control"
                            name="firstname"
                            placeholder=""
                          />
                          <label className="form-label" for="form3Example1c">
                            First Name
                          </label>
                        </div>
                      </div>

                      <div className="d-flex flex-row align-items-center mb-4">
                        <i className="fas fa-envelope fa-lg me-3 fa-fw"></i>
                        <div className="form-outline flex-fill mb-0">
                          <input
                            onChange={changeHandler}
                            className="form-control"
                            name="lastname"
                            placeholder=""
                          />
                          <label className="form-label" for="form3Example3c">
                            Last Name
                          </label>
                        </div>
                      </div>

                      <div className="d-flex flex-row align-items-center mb-4">
                        <i className="fas fa-lock fa-lg me-3 fa-fw"></i>
                        <div className="form-outline flex-fill mb-0">
                          <input
                            onChange={changeHandler}
                            className="form-control"
                            name="username"
                            placeholder=""
                          />
                          <label className="form-label" for="form3Example4c">
                            Username
                          </label>
                        </div>
                      </div>
                      <div className="d-flex flex-row align-items-center mb-4">
                        <i className="fas fa-key fa-lg me-3 fa-fw"></i>
                        <div className="form-outline flex-fill mb-0">
                          <input
                            onChange={changeHandler}
                            className="form-control"
                            name="password"
                            type="password"
                            placeholder=""
                          />
                          <label className="form-label" for="form3Example4cd">
                            Password
                          </label>
                        </div>
                      </div>
                      <div className="d-flex justify-content-center mx-4 mb-3 mb-lg-4">
                        <button
                          type="submit"
                          className="btn btn-primary btn-lg"
                        >
                          Register
                        </button>
                        <button
                          className="btn btn-outline-success ml-5"
                          onClick={Logout}
                        >
                          Logout
                        </button>
                      </div>
                    </form>
                  </div>
                  <div className="col-md-10 col-lg-6 col-xl-7 d-flex align-items-center order-1 order-lg-2">
                    <img
                      src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-registration/draw1.webp"
                      className="img-fluid"
                      alt="Sample image"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default RegisterGoogleForm;
