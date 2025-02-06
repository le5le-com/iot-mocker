package utils

import "context"

// Ctx 全局上下文
var Ctx context.Context

// CtxCancel 程序退出，同时关闭goroutine和清理资源
var CtxCancel context.CancelFunc

func init() {
	// 创建一个全局的上下文，方便在任何地方使用，关闭goroutine
	Ctx, CtxCancel = context.WithCancel(context.Background())
}
