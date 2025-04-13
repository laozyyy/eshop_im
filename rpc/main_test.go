package rpc

import (
	"context"
	"eshop_api/log"
	"testing"
)

func TestGetRandomSku(t *testing.T) {
	type args struct {
		ctx      context.Context
		pageSize int
		pageNum  int
	}
	a := args{
		ctx:      context.Background(),
		pageSize: 10,
		pageNum:  0,
	}
	got, _ := GetRandomSku(a.ctx, a.pageSize, a.pageNum)
	for _, e := range got {
		log.Infof("result: %+v", e)
	}
}

func TestMGetSku(t *testing.T) {
	type args struct {
		ctx      context.Context
		pageSize int
		pageNum  int
		tagId    string
	}
	a := args{
		ctx:      context.Background(),
		pageSize: 10,
		pageNum:  0,
		tagId:    "N_KjVQ",
	}
	got, _ := MGetSku(a.ctx, a.tagId, a.pageSize, a.pageNum)
	for _, e := range got {
		log.Infof("result: %+v", e)
	}

}

func TestOneSku(t *testing.T) {
	type args struct {
		ctx      context.Context
		pageSize int
		pageNum  int
		sku      string
	}
	a := args{
		ctx:      context.Background(),
		pageSize: 10,
		pageNum:  0,
		sku:      "155863",
	}
	got, _ := GetOneSku(a.ctx, a.sku)
	log.Infof("result: %+v", got)

}
