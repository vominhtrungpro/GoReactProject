import "./CrawlContent.css";
import { useEffect, useState } from "react";
import axios from "axios";

function CrawlContent() {
  const [data, setData] = useState({
    url: "",
  });
  const [url, setUrl] = useState([]);
  const changeHandler = (e) => {
    setData({ ...data, [e.target.name]: e.target.value });
  };
  const handleSubmit = (e) => {
    e.preventDefault();
    axios
      .post(`http://localhost:9090/crawl`, data)
      .then(function (response) {
        setUrl(response.data.Url);
        console.log(response.data.Url);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  return (
    <div className="text-center">
      <main className="form-signin w-10 m-auto">
        <form onSubmit={(e) => handleSubmit(e)}>
          <h1 className="h3 mb-3 fw-normal">Enter URL</h1>
          <div className="form-floating">
            <input
              onChange={changeHandler}
              className="form-control"
              name="url"
              placeholder="url"
            />
            <label for="floatingInput">URL</label>
          </div>
          <button class="w-100 btn btn-lg btn-primary" type="submit">
            Submit
          </button>
        </form>
      </main>
      {url.map((item) => (
          <img
          src={`https:${item}`}
        />
      ))}
    </div>
  );
}

export default CrawlContent;
