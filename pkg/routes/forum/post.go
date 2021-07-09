package forum

import (
	"context"
	"github.com/Alisaien/example/pkg/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPost(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.GetHeader("postID"))
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// normally database interaction should be in a different package + function
	row := core.DBW.QueryRow(context.Background(), "SELECT ROW(*) FROM forum.post WHERE post_id = $1", postID)

	var post struct{} // --> replace struct{} with actual Post struct
	if err = row.Scan(&post); err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func PutPost(ctx *gin.Context) {
	// ...
}
