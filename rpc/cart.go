package rpc

import (
	"eshop_api/kitex_gen/eshop/cart"
	"eshop_api/kitex_gen/eshop/cart/cartservice"
	"eshop_api/log"
	"eshop_api/model/req"
	"eshop_api/model/resp"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
)

var cartService cartservice.Client

func init() {
	var err error
	cartService, err = cartservice.NewClient("hello", client.WithHostPorts("117.72.72.114:20002"))
	//cartService, err = cartservice.NewClient("hello", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Errorf("error: %v", err)
	}
}

func AddItem(ctx *gin.Context, req req.AddItemRequestDTO) (*resp.AddItemRespDTO, error) {
	r := &cart.AddItemRequest{
		SkuId:    req.SkuId,
		Quantity: req.Quantity,
		Uid:      req.Uid,
	}
	item, err := cartService.AddItem(ctx, r)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	ret := &resp.AddItemRespDTO{
		Code: int(item.Code),
	}
	if item.ErrStr != nil {
		ret.Info = *item.ErrStr
	} else {
		ret.Info = "success"
	}
	return ret, nil
}

func GetList(ctx *gin.Context, request cart.PageRequest) ([]*cart.CartItem, error) {
	r := &cart.PageRequest{
		PageSize: request.PageSize,
		PageNum:  request.PageNum,
		Uid:      request.Uid,
	}
	list, err := cartService.GetList(ctx, r)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	return list.Items, nil
}
