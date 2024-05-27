package config

//type Config struct {
//	vp       *viper.Viper
//	fileName string
//	isWatch  bool
//}
//
//// NewConfig 通过文件获取读取配置的文件
//func NewConfig(fileName string, isWatch bool, run func(in fsnotify.Event)) (*Config, error) {
//	vp := viper.New()
//
//	vp.SetConfigFile(fileName)
//	if err := vp.ReadInConfig(); err != nil {
//		return nil, err
//	}
//
//	// 是否启动热加载
//	if isWatch {
//		vp.WatchConfig()
//		vp.OnConfigChange(run)
//	}
//
//	return &Config{vp: vp, fileName: fileName}, nil
//}
//
//// InitConfig 初始化配置文件信息
//func InitConfig(fileName string) (*Config, error) {
//	path, err := tools.GetPath(fileName)
//	if err != nil {
//		return nil, err
//	}
//
//	if !tools.IsExists(path) {
//		return nil, errors.New("配置文件不存在")
//	}
//
//	//// TODO 监听热加载配置信息
//	//conf, err = NewConfig(fileName, true, func(in fsnotify.Event) {
//	//	nvp, err := NewConfig()
//	//})
//
//	return NewConfig(fileName, false, nil)
//}
//
//// ReadConfig 将某个键值对信息读取到实例中
//func (c *Config) ReadConfig(k string, v interface{}) error {
//	return c.vp.UnmarshalKey(k, v)
//}
