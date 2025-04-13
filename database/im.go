package database

import (
	"errors"
	"eshop_im/log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	MessageID  int       `gorm:"primaryKey;autoIncrement;column:message_id"`
	SenderID   string    `gorm:"column:sender_id;type:VARCHAR(36);index:idx_sender_receiver"`
	ReceiverID string    `gorm:"column:receiver_id;type:VARCHAR(36);index:idx_sender_receiver"`
	Content    string    `gorm:"column:content;type:TEXT"`
	SendTime   time.Time `gorm:"column:send_time;default:CURRENT_TIMESTAMP"`
	Status     int       `gorm:"column:status;type:TINYINT;default:0"`
}

// 修改保存消息函数
func SaveMsg(db *gorm.DB, msg string, from string, to string) (messageID int, err error) {
	db = getDBInstance(db)
	message := &Message{
		SenderID:   from,
		ReceiverID: to,
		Content:    msg,
		Status:     0,
	}

	if err := db.Create(message).Error; err != nil {
		log.Errorf("消息保存失败: %v", err)
		return 0, err
	}
	return message.MessageID, nil
}

// 新增状态更新函数
func UpdateStatus(db *gorm.DB, messageID int, status int) error {
	db = getDBInstance(db)
	result := db.Model(&Message{}).
		Where("message_id = ?", messageID).
		Update("status", status)

	if result.Error != nil {
		log.Errorf("状态更新失败: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("消息记录不存在")
	}
	return nil
}

type Receiver struct {
	Id          int
	Uid         string
	ReceiverUid string
}

func GetReceiverUid(db *gorm.DB, uid string) ([]string, error) {
	db = getDBInstance(db)
	var res []*Receiver
	err := db.Table("receiver").
		Where("uid = ?", uid).
		Find(&res).Error
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	receiver := res[0]
	receiverUids := strings.Split(receiver.ReceiverUid, ",")
	return receiverUids, nil
}

func GetOneMessage(db *gorm.DB, uid, rUid string) (*Message, error) {
	db = getDBInstance(db)
	var res []*Message
	err := db.Table("receiver").
		Where("sender_id = ?", uid).
		Where("receiver_id", rUid).
		Order("send_time DESC").
		Limit(1).
		Find(&res).Error
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func MGetMessage(db *gorm.DB, uid, rUid string, limit int) ([]*Message, error) {
	db = getDBInstance(db)
	var res []*Message
	err := db.Table("receiver").
		Where("sender_id = ?", uid).
		Where("receiver_id", rUid).
		Order("send_time ASC").
		Limit(limit).
		Find(&res).Error
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	if len(res) > 0 {
		return res, nil
	}
	return nil, nil
}
