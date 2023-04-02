package initialize

import "go.uber.org/zap"

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stderr",
		//"stdout",
		"./mx_shop_log.log",
	}
	logger, _ := cfg.Build()
	defer logger.Sync()
	// 替换线程安全的的全局logger，后面都可以使用zap.S()调用
	zap.ReplaceGlobals(logger)
}
