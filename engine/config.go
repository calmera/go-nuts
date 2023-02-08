package engine

type Config struct {
	Nats *NatsConfig `json:"nats"`
}

type NatsConfig struct {
	Url string `json:"url"`
}
