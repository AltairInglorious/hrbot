package bot

type BotConfig struct {
	onInit      func(*Bot)
	onChat      func(*Bot, []byte)
	onWhisper   func(*Bot, []byte)
	onUserJoin  func(*Bot, []byte)
	onUserLeave func(*Bot, []byte)
}

func NewBotConfig() *BotConfig {
	return &BotConfig{
		onChat:      func(*Bot, []byte) {},
		onWhisper:   func(*Bot, []byte) {},
		onUserJoin:  func(*Bot, []byte) {},
		onUserLeave: func(*Bot, []byte) {},
		onInit:      func(*Bot) {},
	}
}

func (c *BotConfig) OnChat(handler func(*Bot, []byte)) {
	c.onChat = handler
}

func (c *BotConfig) OnWhisper(handler func(*Bot, []byte)) {
	c.onWhisper = handler
}

func (c *BotConfig) OnUserJoin(handler func(*Bot, []byte)) {
	c.onUserJoin = handler
}

func (c *BotConfig) OnUserLeave(handler func(*Bot, []byte)) {
	c.onUserLeave = handler
}

func (c *BotConfig) OnInit(handler func(*Bot)) {
	c.onInit = handler
}
