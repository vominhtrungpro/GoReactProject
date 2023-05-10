import React from "react";

import axios from "axios";

import "./user.css";

import "bootstrap-icons/font/bootstrap-icons.css";

import Popup from "reactjs-popup";
import "reactjs-popup/dist/index.css";
import Content from "./Content";

import "./Content.css";

export default class UserList extends React.Component {
  state = {
    users: [],
  };

  componentDidMount() {
    let config = {
      headers: {
        Authorization: `Bearer ${JSON.parse(
          localStorage.getItem("access-token")
        )}`,
      },
    };
    axios.get(`http://localhost:9090/user/all`, config).then((res) => {
      const users = res.data.User;
      this.setState({ users });
    });
  }
  render() {
    return (
      <table className="table">
        <thead className="thead-dark">
          <tr>
            <th>Id</th>
            <th>Firstname</th>
            <th>Lastname</th>
            <th>Username</th>
            <th>Email</th>
            <th>Password</th>
            
          </tr>
        </thead>
        <tbody>
          {this.state.users.map((users) => (
            <tr key={users.Id}>
              <td>{users.Id}</td>
              <td>{users.Firstname}</td>
              <td>{users.Lastname}</td>
              <td>{users.Username}</td>
              <td>{users.Email}</td>
              <td>
                {users.Password}
                <Popup modal trigger={<i className="bi bi-gear-fill float-right"></i>}>
              {(close) => <Content user={users} close={close} />}
            </Popup>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  }
}
