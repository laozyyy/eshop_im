package handler

import (
	"encoding/json"
	"eshop_im/database"
	"eshop_im/log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域，根据实际安全需求调整
	},
}

// HandleWebSocket 处理WebSocket连接
// 在文件顶部添加连接池和锁
var (
	connections = make(map[string]*websocket.Conn)
	connLock    sync.RWMutex
)

type MsgEntity struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

const MessageAlreadySent int = 1

func HandleWebSocket(ctx *gin.Context) {
	// 升级HTTP连接到WebSocket
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Errorf("WebSocket upgrade failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket连接建立失败"})
		return
	}
	defer conn.Close()

	// 用户身份验证（示例使用uid参数）
	uid := ctx.Query("uid")
	if uid == "" {
		log.Error("缺少用户身份标识")
		return
	}

	// 用户身份验证后注册连接
	connLock.Lock()
	connections[uid] = conn
	connLock.Unlock()

	// 添加连接关闭处理
	defer func() {
		connLock.Lock()
		delete(connections, uid)
		connLock.Unlock()
	}()

	// 修改消息处理循环
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorf("读取消息失败: %v", err)
			break
		}

		// 解析消息格式：{"to": "target_uid", "content": "message"}
		var msg MsgEntity
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Errorf("消息解析失败: %v", err)
			continue
		}

		// 落库
		msgId, err := database.SaveMsg(nil, msg.Content, uid, msg.To)
		if err != nil {
			log.Errorf("error: %d", err)
			return
		}

		// 查找目标连接
		connLock.RLock()
		targetConn, exists := connections[msg.To]
		connLock.RUnlock()
		if !exists {
			log.Infof("用户uid: %s 不在线", msg.To)
			//resp, _ := json.Marshal(map[string]string{"error": "用户不在线"})
			//conn.WriteMessage(websocket.TextMessage, resp)
			continue
		}

		// 转发消息
		if err := targetConn.WriteMessage(websocket.TextMessage, []byte(msg.Content)); err != nil {
			log.Errorf("消息转发失败: %v", err)
			continue
		}

		err = database.UpdateStatus(nil, msgId, MessageAlreadySent)
		if err != nil {
			log.Errorf("error: %d", err)
			return
		}
	}
}

type MGetResponse struct {
	Info      string           `json:"info"`
	Receivers []ReceiverEntity `json:"receivers"`
}
type ReceiverEntity struct {
	Uid          string    `json:"uid"`
	Name         string    `json:"name"`
	LastMessage  string    `json:"last_message"`
	LastSendTime time.Time `json:"last_send_time"`
}

func HandleMgetReceiver(ctx *gin.Context) {
	// 用户身份验证（示例使用uid参数）
	uid := ctx.Query("uid")
	if uid == "" {
		log.Error("缺少用户身份标识")
		return
	}
	receiverUids, _ := database.GetReceiverUid(nil, uid)
	receivers := make([]ReceiverEntity, 0)
	for _, rUid := range receiverUids {
		message, err := database.GetOneMessage(nil, uid, rUid)
		var tmp ReceiverEntity
		if message == nil {
			tmp = ReceiverEntity{
				Uid:          rUid,
				Name:         "test",
				LastMessage:  "none",
				LastSendTime: time.Now(),
			}
		} else {
			tmp = ReceiverEntity{
				Uid:          rUid,
				Name:         "test",
				LastMessage:  message.Content,
				LastSendTime: message.SendTime,
			}
		}
		if err != nil {
			log.Errorf("error: %d", err)
			return
		}
		receivers = append(receivers, tmp)
	}
	ctx.JSON(http.StatusOK, receivers)
}

type GetOneHistoryReq struct {
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
}
type GetOneHistoryResp struct {
	Info     string              `json:"info"`
	Messages []*database.Message `json:"messages"`
}

func HandleOneHistory(ctx *gin.Context) {
	var req GetOneHistoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Errorf("error: %d", err)
		return
	}
	messages, _ := database.MGetMessage(nil, req.SenderId, req.ReceiverId, 10)
	resp := GetOneHistoryResp{
		Info:     "success",
		Messages: messages,
	}
	ctx.JSON(http.StatusOK, resp)

}
