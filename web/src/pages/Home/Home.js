import React from "react";
import Navbar from "../../components/navbar";
import Profile from "../../components/profile";

function Home() {
  return (
    <div className="App">
      <Navbar />
      <div style={{ marginTop: "50px" }}>
        <Profile />
      </div>
    </div>
  );
}

export default Home;
