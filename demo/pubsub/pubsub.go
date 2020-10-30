package pubsub

import (
	"log"
	"sync"
)

// Pubsub 发布订阅
type Pubsub struct {
	channel *sync.Map // map[string]interface{}{Channel.Name: *Channel}
}

// Topic 主题
type Topic struct {
	Title   string
	Content string
}

// Client 用户
type Client struct {
	Username string
}

func NewClient(username string) *Client {
	return &Client{Username: username}
}

func (m *Client) Broadcast(tp *Topic)  {
	log.Printf("%s 收到消息 %s - %s\n", m.Username, tp.Title, tp.Content)
}

// Channel 频道
type Channel struct {
	Name    string
	Clients *sync.Map // map[string]interface{}{Client.Username: *Client}
	Topics  chan *Topic
}

func NewChannel(name string, buf int) *Channel {
	return &Channel{Name: name, Clients: &sync.Map{}, Topics: make(chan *Topic, buf)}
}

func (ch *Channel) Publish(tp *Topic) {
	ch.Topics <- tp
}

func (ch *Channel) Subscribe(m *Client) {
	ch.Clients.Store(m.Username, m)
}

func (ch *Channel) Broadcast() {
	for {
		select {
		case tp := <-ch.Topics:
			ch.Clients.Range(func(key, value interface{}) bool {
				value.(*Client).Broadcast(tp)
				return true
			})
		}
	}
}

func NewPubsub(ch *Channel) *Pubsub {
	var pool = &sync.Map{}
	pool.Store(ch.Name, ch)
	return &Pubsub{
		pool,
	}
}

func (p *Pubsub) AddChannel(ch *Channel) {
	var pool = &sync.Map{}
	p.channel.Range(func(key, value interface{}) bool {
		pool.Store(key,value)
		return true
	})
	pool.Store(ch.Name, ch)
	p.channel = pool
}

func (p *Pubsub) Publish(channel string, tp *Topic) {
	if v,ok := p.channel.Load(channel);ok {
		var ch = v.(*Channel)
		ch.Publish(tp)
	}
}

func (p *Pubsub) Subscribe(channel string, m *Client) {
	if v,ok := p.channel.Load(channel);ok {
		var ch = v.(*Channel)
		ch.Subscribe(m)
	}
}

func (p *Pubsub) Broadcast() {
	p.channel.Range(func(key, value interface{}) bool {
		ch := value.(*Channel)
		ch.Broadcast()
		return true
	})
}
