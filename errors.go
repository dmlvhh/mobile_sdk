package main

import "fmt"

// APIError 表示 API 调用失败的错误
type APIError struct {
	StatusCode int
	Message    string
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: %d - %s", e.StatusCode, e.Message)
}

// NewAPIError 创建一个新的 APIError
func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}
