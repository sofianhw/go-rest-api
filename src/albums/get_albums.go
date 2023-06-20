package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofianhw/go-rest-api/src"
)

// get list of albums:
func GetAlbums(ctx *gin.Context) {
	statusCode := http.StatusOK // code 200
	ctx.IndentedJSON(statusCode, src.Albums)
}
