package model

import "time"

// User 表
type User struct {
	ID         int64     `json:"-"`        // 自增的唯一标识 ID
	UID        string    `json:"uid"`      // 用户唯一标识
	Name       string    `json:"name"`     // 用户姓名
	Phone      string    `json:"phone"`    // 用户电话号码
	Email      string    `json:"email"`    // 用户电子邮箱
	Password   string    `json:"password"` // 用户密码
	Role       int       `json:"role"`     // 用户权限-1为买家，2为卖家，3为管理员
	CreateTime time.Time `json:"-"`        // 创建时间
	UpdateTime time.Time `json:"-"`        // 更新时间
	IsDeleted  bool      `json:"-"`        // 是否删除标记，0 表示未删除，1 表示已删除
}

// Score 分数表
type Score struct {
	ID         int64     `json:"-"`     // 自增的唯一标识 ID
	UID        string    `json:"uid"`   // 用户唯一标识
	Score      float64   `json:"score"` // 分数，保留三位小数
	CreateTime time.Time `json:"-"`     // 创建时间
	UpdateTime time.Time `json:"-"`     // 更新时间
	IsDeleted  bool      `json:"-"`     // 是否删除标记，0 表示未删除，1 表示已删除
}
