package debugHandler

import (
	"GinDemo/controller"
	"GinDemo/errno"
	"GinDemo/logger"
	"github.com/gin-gonic/gin"
)

// @Summary Add new user to the database
// @Description path for debug
// @Tags debug
// @Accept  json
// @Produce  json
// @Router /user [post]
func Debug(c *gin.Context) {
	logger.HandlerLogger().Debug("test debug handler for log level is debug")
	controller.SendResponse(c, errno.ErrToken, nil)
}
