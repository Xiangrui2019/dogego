package serializer

// Response 团队基础序列化器
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total interface{} `json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}
