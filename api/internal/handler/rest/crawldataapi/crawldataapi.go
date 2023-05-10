package crawldataapi

import (
	"GoReactProject/internal/controller/crawldataservice"
	"GoReactProject/internal/model/crawldatamodel"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrawlData(context *gin.Context) {
	var request crawldatamodel.CrawlDataRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := crawldataservice.CollyCrawl(request)
	context.IndentedJSON(http.StatusOK, result)
}
