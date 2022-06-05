package main

import (
	kserver "github.com/cloudwego/kitex/server"

	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
}

type server struct {
	srv kserver.Server
}

func (s *server) Start() error {
	return s.srv.Run()
}

func (s *server) Stop(ctx context.Context) error {
	return s.srv.Stop()
}

type App struct {
	ctx         context.Context
	cancel      func()
	srvs        []Server
	sigs        []os.Signal
	stopTimeout time.Duration
}

type Option func(a *App)

func WithCancel(ctx context.Context, cancel func()) Option {
	return func(a *App) {
		a.ctx = ctx
		a.cancel = cancel
	}
}
func WithSrvs(srvs []Server) Option {
	return func(a *App) {
		a.srvs = srvs
	}
}

func WithSigs(sigs []os.Signal) Option {
	return func(a *App) {
		a.sigs = sigs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(a *App) {
		a.stopTimeout = timeout
	}
}

func (a *App) Run() error {
	// 以a的ctx作为根ctx
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}
	for _, srv := range a.srvs {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done() // Cancel函数已被调用
			// 在限定时间内Stop
			sdCtx, cancel := context.WithTimeout(a.ctx, a.stopTimeout)
			defer cancel()
			return srv.Stop(sdCtx) // Stop服务
		})
		wg.Add(1)
		eg.Go(func() error {
			defer wg.Done()
			return srv.Start() // Start服务
		})
	}
	wg.Wait()
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...) // 注册信号
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done(): // 服务因发生错误而关闭
				return ctx.Err()
			case <-c: // 接收到结束信号而关闭
				if err := a.Stop(); err != nil {
					return err
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	if a.cancel != nil {
		// 调用context的Cancel函数, 所有的Server都将关闭
		a.cancel()
	}
	return nil
}
