import React from "react";

import "./Content.css";

export default ({ user, close }) => (
  <div>
    <a className="close" onClick={close}>
      ×
    </a>
    <div className="header"></div>
    <div className="content">
      <form>
        <div className="form-group">
          <label>Id</label>
          <input type="text" className="form-control" value={user.Id} />
        </div>
        <div className="form-group">
          <label>Firstname</label>
          <input type="text" className="form-control" value={user.Firstname} />
        </div>
        <div className="form-group">
          <label>Lastname</label>
          <input type="text" className="form-control" value={user.Lastname} />
        </div>
        <div className="form-group">
          <label>Username</label>
          <input type="text" className="form-control" value={user.Username} />
        </div>
        <div className="form-group">
          <label>Email</label>
          <input type="text" className="form-control" value={user.Email} />
        </div>
        <div className="form-group">
          <label>Password</label>
          <input type="text" className="form-control" value={user.Password} />
        </div>
        <button type="submit" className="btn btn-primary">Save</button>
      </form>
    </div>
  </div>
);
