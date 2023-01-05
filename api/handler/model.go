package handler

type OperationRequest struct {
	Entity string `json:"entity"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

type OperationResponse struct {
	Result string `json:"result"`
}
