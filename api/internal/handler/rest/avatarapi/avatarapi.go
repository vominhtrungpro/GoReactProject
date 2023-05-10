package avatarapi

import (
	"GoReactProject/internal/controller/avatarservice"
	"GoReactProject/internal/model/avatarmodel"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAvatar(context *gin.Context) {
	var request avatarmodel.CreateAvatarRequest
	// var apirequest avatarmodel.CreateAvatarAPIRequest
	// err := json.NewDecoder(context.Request.Body).Decode(&apirequest)
	// if err != nil {
	// 	http.Error(context.Writer, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	userid := context.Param("id")
	id, err := strconv.Atoi(userid)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := context.Request.FormFile("file")
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	image, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	request.Image = image
	request.Userid = id

	result := avatarservice.CreateAvatar(request)
	context.IndentedJSON(http.StatusOK, result)
}

func GetAvatar(context *gin.Context) {
	userid := context.Param("id")
	id, err := strconv.Atoi(userid)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := avatarservice.GetAvatar(id)
	//t := time.Now().String()
	//fileName := "AvatarUser" + strconv.Itoa(id) + "" + t + ".jpg"
	// context.Header("Content-Disposition", "attachment; filename="+fileName+"")
	// context.Data(http.StatusOK, "application/octet-stream", result)
	context.IndentedJSON(http.StatusOK, result)
}
