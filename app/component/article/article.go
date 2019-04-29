package article

import (
	"github.com/gin-gonic/gin"
	"github.com/iphuket/iuu/app/component/article/api"
)

// Route Init
func Route(r *gin.RouterGroup) {
	api.Route(r)
}
