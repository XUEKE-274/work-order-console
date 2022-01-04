package rpcReq

type DepartmentReq struct {
	Account        string  `json:"account"`
	AccountDB      string  `json:"accountDb"`
	DeptIds        []int64 `json:"deptIds"`
	ExcludeDelete  bool    `json:"excludeDelete"`
}