// 添加在线用户查询接口
package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleOnlineUsers(ctx *gin.Context) {
	connLock.RLock()
	defer connLock.RUnlock()

	uids := make([]string, 0, len(connections))
	for uid := range connections {
		uids = append(uids, uid)
	}
	ctx.JSON(http.StatusOK, gin.H{"online_users": uids})
}
