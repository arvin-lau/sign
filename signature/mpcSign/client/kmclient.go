package client

// 定义一个结构体对象
var mpcSignClient *MpcSignClient

// InitMpcSignClient 初始化服务器
func InitMpcSignClient(server []string) error {
	mpcSignClient = NewMpcSignClient(server)
	return nil
}

// GetKmClient 获取初始化值
func GetKmClient() *MpcSignClient {
	return mpcSignClient
}
