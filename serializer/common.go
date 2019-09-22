package serializer

import "time"

// Response 团队基础序列化器
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Error     string      `json:"error"`
	TimeStamp int64       `json:"timestamp"`
}

func (response Response) Result() *Response {
	response.TimeStamp = time.Now().Unix()

	return &response
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

func (response TrackedErrorResponse) Result() *TrackedErrorResponse {
	response.TimeStamp = time.Now().Unix()

	return &response
}
