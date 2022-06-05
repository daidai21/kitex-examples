// Code generated by Kitex v1.9.2. DO NOT EDIT.

package marketservice

import (
	"code.byted.org/kite/kitex/byted"
	"code.byted.org/kite/kitex/server"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_market/kitex_gen/cmp/ecom/market"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler market.MarketService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, byted.InvokeSuite(serviceInfo()))

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}

// NewInvokerWithBytedConfig creates a server.Invoker with the given handler and options.
func NewInvokerWithBytedConfig(handler market.MarketService, config *byted.ServerConfig, opts ...server.Option) server.Invoker {
	var options []server.Option
	options = append(options, byted.InvokeSuiteWithConfig(serviceInfo(), config))
	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
