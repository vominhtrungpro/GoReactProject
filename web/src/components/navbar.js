import React from "react";
import {useNavigate} from 'react-router-dom'

function Navbar() {
  const navigate = useNavigate()
  function Logout() {
    localStorage.removeItem('access-token');
    localStorage.removeItem('google-access-token');
    navigate('/')
  }

  return (
    <div>
      <header>
        <nav className="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
          <div className="container-fluid">
            <a className="navbar-brand" href="Home">
              Home
            </a>
            <button
              className="navbar-toggler"
              type="button"
              data-bs-toggle="collapse"
              data-bs-target="#navbarCollapse"
              aria-controls="navbarCollapse"
              aria-expanded="false"
              aria-label="Toggle navigation"
            >
              <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarCollapse">
              <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item">
                  <a className="nav-link" href="/users">
                    User
                  </a>
                </li>
                <li className="nav-item">
                  <a className="nav-link" href="/crawl">
                    Crawl
                  </a>
                </li>
              </ul>
              <form className="d-flex" role="search">
                <input
                  className="form-control me-2"
                  type="search"
                  placeholder="Search"
                  aria-label="Search"
                />
                <button className="btn btn-outline-success" type="submit">
                  Search
                </button>
              </form>
              <button className="btn btn-outline-success ml-5" onClick={Logout} >
                  Logout
                </button>
            </div>
          </div>
        </nav>
      </header>
    </div>
  );
}

export default Navbar;
