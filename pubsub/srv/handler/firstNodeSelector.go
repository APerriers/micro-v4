package handler

import (
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Built in random hashed node selector
type FirstNodeSelector struct {
	Opts selector.Options
}

func (n *FirstNodeSelector) Init(Opts ...selector.Option) error {
	for _, o := range Opts {
		o(&n.Opts)
	}
	return nil
}

func (n *FirstNodeSelector) Options() selector.Options {
	return n.Opts
}

func (n *FirstNodeSelector) Select(service string, Opts ...selector.SelectOption) (selector.Next, error) {
	services, err := n.Opts.Registry.GetService(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var sopts selector.SelectOptions
	for _, opt := range Opts {
		opt(&sopts)
	}

	for _, filter := range sopts.Filters {
		services = filter(services)
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	if len(services[0].Nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	return func() (*registry.Node, error) {
		return services[0].Nodes[0], nil
	}, nil
}

func (n *FirstNodeSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (n *FirstNodeSelector) Reset(service string) {
	return
}

func (n *FirstNodeSelector) Close() error {
	return nil
}

func (n *FirstNodeSelector) String() string {
	return "first"
}
