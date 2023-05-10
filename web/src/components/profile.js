import axios from "axios";
import jwt_decode from "jwt-decode";
import React, { useState, useEffect } from "react";
import UploadAndDisplayImage from "./Uploadpicture";

function Profile() {
  const token = JSON.parse(localStorage.getItem("access-token"));

  const decoded = jwt_decode(token);

  const [user, setResponseData] = useState({
    Id: "",
    Firstname: "",
    Lastname: "",
    Username: "",
    Email: "",
    Password: "",
  });

  useEffect(() => {
    const config = {
      headers: {
        Authorization: `Bearer ${JSON.parse(
          localStorage.getItem("access-token")
        )}`,
      },
    };
    axios
      .get(`http://localhost:9090/user/` + decoded.id, config)
      .then((res) => {
        setResponseData(res.data.User);
      });
  }, []);
  return (
    <div className="container rounded bg-white mt-5 mb-5">
      <div className="row">
        <div className="col">
          <div className="d-flex flex-column align-items-center text-center p-3 py-5">
            <UploadAndDisplayImage />
            <span className="font-weight-bold">
              {user.Firstname} {user.Lastname}
            </span>
            <span className="text-black-50">{user.Email}</span>
            <span> </span>
          </div>
        </div>
        <div className="col">
          <div className="p-3 py-5">
            <div className="d-flex justify-content-between align-items-center mb-3">
              <h4 className="text-right">Profile Settings</h4>
            </div>
            <div className="row mt-2">
              <div className="col-md-6">
                <label className="labels">First name</label>
                <input
                  type="text"
                  className="form-control"
                  placeholder="First name"
                  value={user.Firstname}
                />
              </div>
              <div className="col-md-6">
                <label className="labels">Last name</label>
                <input
                  type="text"
                  className="form-control"
                  placeholder="Last name"
                  value={user.Lastname}
                />
              </div>
            </div>
            <div className="row mt-3">
              <div className="col-md-12">
                <label className="labels">Email</label>
                <input
                  type="text"
                  className="form-control"
                  placeholder="Email"
                  value={user.Email}
                />
              </div>
            </div>
            <div className="mt-5 text-center">
              <button className="btn btn-primary profile-button" type="button">
                Save
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Profile;
