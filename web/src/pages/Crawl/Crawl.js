import React from "react";
import Navbar from "../../components/navbar";
import CrawlContent from "../../components/CrawlContent";

function Crawl() {
  return (
    <div className="App">
      <Navbar />
      <div style={{ marginTop: "50px" }}>
        <CrawlContent />
      </div>
    </div>
  );
}

export default Crawl;
