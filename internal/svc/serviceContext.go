package svc

import (
	"exporter/internal/config"
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

const (
	EmailPassword    = "CGLHYIFLUBKQVUQM"
	localCacheExpire = time.Duration(time.Minute * 20)
	SystemListKey    = "SystemListKey"
	Keystr           = "W3WxhhoA9E9VIteCYbnhUTxDbtk2nP1Z"
	EmailListKey     = "EmailListKey"
)

type ServiceContext struct {
	Config     config.Config
	LocalCache *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		LocalCache: localCache,
	}
}
