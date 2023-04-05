package bot

type BotConfig struct {
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
	}
}

func (b *Bot) OnChat(handler func(*Bot, []byte)) {
	b.config.onChat = handler
}

func (c *BotConfig) OnChat(handler func(*Bot, []byte)) {
	c.onChat = handler
}

func (b *Bot) OnWhisper(handler func(*Bot, []byte)) {
	b.config.onWhisper = handler
}

func (c *BotConfig) OnWhisper(handler func(*Bot, []byte)) {
	c.onWhisper = handler
}

func (b *Bot) OnUserJoin(handler func(*Bot, []byte)) {
	b.config.onUserJoin = handler
}

func (c *BotConfig) OnUserJoin(handler func(*Bot, []byte)) {
	c.onUserJoin = handler
}

func (b *Bot) OnUserLeave(handler func(*Bot, []byte)) {
	b.config.onUserLeave = handler
}

func (c *BotConfig) OnUserLeave(handler func(*Bot, []byte)) {
	c.onUserLeave = handler
}
