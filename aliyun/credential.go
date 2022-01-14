package aliyun

import (
	`github.com/aliyun/alibaba-cloud-sdk-go/services/sts`
)

type Credentials struct {
	// 日志存储信息
	SLS SLS `json:"sls"`
	// 证书相关
	Credentials sts.Credentials `json:"credentials"`
}
