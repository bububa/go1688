# 1688.com电商采购API

详情参见1688电商采购[API文档](https://open.1688.com/solution/solutionDetail.htm?solutionKey=1559112028404#apiAndMessageList)

```
package main

import (
	"log"
  
	"github.com/bububa/go1688"
	"github.com/bububa/go1688/requests/alibaba/product"
)

const (
	AppKey           = "xxx"
	AppSecret        = "xxx"
	productId uint64 = 594279676823
)

func main() {
	client := go1688.NewClient(AppKey, AppSecret, nil)
	getProductInfo(client, productId)
}

func getProductInfo(client *go1688.Client, productId uint64) {
	req := &product.CpsMediaProductInfoRequest{
		OfferId: productId,
	}
	info, bizInfos, err := product.CpsMediaProductInfo(client, req, "")
	if err != nil {
		log.Fatalln(err)
		return
	}
  log.Printf("info: %+v, biz: %+v", info, bizInfos)
}
```
