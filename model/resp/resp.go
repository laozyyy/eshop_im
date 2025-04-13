package resp

import (
	"eshop_im/kitex_gen/eshop/cart"
	"eshop_im/model"
)

type LoginRespDTO struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Uid  string `json:"uid,omitempty"`
}

type RegisterRespDTO struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Uid  string `json:"uid,omitempty"`
}

type UserRespDTO struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	User *model.User `json:"user"`
}

type Tag struct {
	// 一级分类id
	ID string `json:"id"`
	// 一级分类名字
	Name string `json:"name"`
	// 一级分类图片
	Picture string `json:"picture"`
}

type TagRespDTO struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Result []Tag  `json:"result"`
}

type GetSkuRespDTO struct {
	Code int             `json:"code"`
	Info string          `json:"info"`
	Data *model.GoodsSku `json:"data,omitempty"`
}

type PageRespDTO struct {
	Code int               `json:"code"`
	Info string            `json:"info"`
	Data []*model.GoodsSku `json:"data,omitempty"`
}
type AddItemRespDTO struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type CartListRespDTO struct {
	Code int              `json:"code"`
	Info string           `json:"info"`
	Data []*cart.CartItem `json:"data,omitempty"`
}
