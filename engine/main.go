package engine

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
	"sync"
)

func NewEngine(config Config) (*Engine, error) {
	return &Engine{
		config: config,
		wg:     &sync.WaitGroup{},
	}, nil
}

type Engine struct {
	config Config
	nc     *nats.Conn
	js     nats.JetStreamContext
	wg     *sync.WaitGroup
}

func (e *Engine) Start() error {
	// -- connect to the embedded nats server
	nc, err := nats.Connect(e.config.Nats.Url)
	if err != nil {
		e.Close()
		return err
	}
	e.nc = nc

	js, err := nc.JetStream()
	if err != nil {
		e.Close()
		return fmt.Errorf("unable to initialize jetstream: %w", err)
	}
	e.js = js

	e.wg.Add(1)
	log.Info().Msg("Node Ready!")

	return nil
}

func (e *Engine) Wait() {
	e.wg.Wait()
}

func (e *Engine) Close() {
	// -- close the client
	e.nc.Close()

	// -- wait for all modules to be closed
	e.wg.Done()
}
