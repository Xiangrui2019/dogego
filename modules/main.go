package modules

func InitAllModules() {
	InitAliyunOSSModule()
	InitLockerModule()
	new(RedisQueue).Receive("a", func(message string) error {
		return nil
	})
}
