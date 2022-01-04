package rpcResp

type DepartmentRpcResp struct {
	BaseRpcResponse
	Data *[]DepartmentResp `json:"data"`

}

type DepartmentResp struct {
	Id        int64  `json:"_id"`
	Name      string `json:"name"`
	ParentId  int64  `json:"parentid"`
}
