package request

type TemplateAddRequest struct {
	Name string `json:"name"` // 用户名称
}

type NewTemplateAddRequest struct {
	TemplateName string `json:"templateName"`
	WorkflowName string `json:"workflowName"`
	Fields []*FieldRequest `json:"fields"`
	Nodes []*NodeRequest `json:"nodes"`
}

type FieldRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type NodeRequest struct {
	Name string `json:"name"`
}