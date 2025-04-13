package req

type LoginReqDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterReqDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type MGetSkuReqDTO struct {
	TagID    string `json:"tag_id" binding:"required"`
	PageNum  int    `json:"page_num" binding:"required"`
	PageSize int    `json:"page_size" binding:"required"`
}

type PageReqDTO struct {
	PageNum  int `json:"page_num" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}

type AddItemRequestDTO struct {
	SkuId    string `thrift:"skuId,1,required" frugal:"1,required,string" json:"sku_id"`
	Quantity int32  `thrift:"quantity,2,required" frugal:"2,required,i32" json:"quantity"`
	Uid      string `thrift:"uid,3" frugal:"3,default,string" json:"uid"`
}
