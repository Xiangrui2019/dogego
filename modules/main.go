package modules

func InitAllModules() {
	InitAliyunOSSModule()
	InitLockerModule()
	InitTaskModule()
	InitAMQPModule()
}
