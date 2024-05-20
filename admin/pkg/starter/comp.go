package starter

type IComp interface {
	CompName() string            // 服务名称
	GetCompItem() []CompItem     // 获取服务要进行展示在控制台的配置信息
	SetCompItem(key, val string) // 设置要进行输出的配置信息
	IsDev() bool                 // 只有在dev环境才进行输出
}

type CompItem struct {
	Key   string
	Value string
}

type comp struct {
	name      string
	isDev     bool
	compItems []CompItem
}

func NewComp(name string, isDev bool) IComp {
	c := &comp{name: name, isDev: isDev, compItems: make([]CompItem, 0)}

	RegisterComp(c)
	return c
}

func (c *comp) CompName() string {
	return c.name
}

func (c *comp) GetCompItem() []CompItem {
	return c.compItems
}

func (c *comp) SetCompItem(key, val string) {
	c.compItems = append(c.compItems, CompItem{Key: key, Value: val})
}

func (c *comp) IsDev() bool {
	return c.isDev
}
