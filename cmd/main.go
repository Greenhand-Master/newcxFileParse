package main

import (
	logo "fileparse/res"
	"log"

	"go.uber.org/zap"
)

func main() {
	zap.S().Info(logo.Logo)
	zap.S().Info("hello")
	//_logger.InitLogger(conf.Log.ErrorPath, conf.Log.InfoPath, conf.Log.DebugPath, conf.Log.Format, conf.Log.MaxSaveTime, conf.Log.RotationTime)
	zap.S().Info("[INIT] Initialized Configuration")
	zap.S().Info("[INIT] Initialized Log")
	log.Println("code is running here!")
	//This is
}
