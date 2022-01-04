package rpcResp

import "time"

type BaseRpcResponse struct {
	Msg string `json:"msg"`
	Code int32 `json:"code"`
	Success bool `json:"success"`
	TraceId string `json:"traceId"`
	RequestId string `json:"requestId"`
	Time time.Time `json:"time"`
}