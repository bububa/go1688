# 1688.com电商采购API

详情参见1688电商采购[API文档](https://open.1688.com/solution/solutionDetail.htm?solutionKey=1559112028404#apiAndMessageList)
[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/go1688.svg)](https://pkg.go.dev/github.com/bububa/go1688)
[![Go](https://github.com/bububa/go1688/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/go1688/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/go1688/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/go1688/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/go1688.svg)](https://github.com/bububa/go1688)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/go1688)](https://goreportcard.com/report/github.com/bububa/go1688)
[![GitHub license](https://img.shields.io/github/license/bububa/go1688.svg)](https://github.com/bububa/go1688/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/go1688.svg)](https://GitHub.com/bububa/go1688/releases/)


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
