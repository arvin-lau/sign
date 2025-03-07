package client

// 定义一个结构体对象
var kmClient *KmClient

// InitKmClient 初始化服务器
func InitKmClient(server []string, token string) error {
	kmClient = NewKmClient(server, token)
	return nil
}

// GetKmClient 获取初始化值
func GetKmClient() *KmClient {
	return kmClient
}
