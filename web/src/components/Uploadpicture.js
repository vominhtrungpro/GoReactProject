import React, { useEffect, useState } from "react";
import "./Uploadpicture.css";
import axios from "axios";
import jwt_decode from "jwt-decode";
import * as FileSaver from "file-saver";

const UploadAndDisplayImage = () => {
  const [selectedImage, setSelectedImage] = useState(null);
  const [img, setImg] = useState("https://www.prydwen.gg/static/61a1a56cc79160e400f1f0c0a5f05390/acb7c/emp-1160.png");
  const [userid, setUserId] = useState({
    Userid: 0,
  });
  useEffect(() => {
    axios.get(`http://localhost:9090/avatar/` + decoded.id).then((res) => {
    changImg(res.data);
  });
  },[])
  const changImg = (data) => {
    if(data==null)
    {
      setImg("https://www.prydwen.gg/static/61a1a56cc79160e400f1f0c0a5f05390/acb7c/emp-1160.png");
    }
    else
    {
      setImg(`data:image/jpeg;base64,${data}`);
    }
  };
  const token = JSON.parse(localStorage.getItem("access-token"));
  const decoded = jwt_decode(token);
  

  const UploadAvatar = (file) => {
    const formData = new FormData();
    formData.append("file", file);
    axios
      .post(`http://localhost:9090/avatar/` + decoded.id, formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((response) => {
        console.log(response);
      });
  };

  return (
    <div>
      <div className="image-cropper">
        <img
          className="rounded"
          alt="not found"
          width={"300px"}
          src={
            selectedImage
              ? URL.createObjectURL(selectedImage)
              : `${img}`
          }
        />
      </div>
      <br />
      <input
        type="file"
        name="myImage"
        onChange={(event) => {
          setSelectedImage(event.target.files[0]);
          UploadAvatar(event.target.files[0]);
        }}
      />
    </div>
  );
};

export default UploadAndDisplayImage;
