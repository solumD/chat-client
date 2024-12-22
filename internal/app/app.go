package app

import (
	"context"

	"github.com/solumD/chat-client/internal/config"
)

const configPath = ".env"

// App структура приложения
type App struct {
	ServiceProvider *serviceProvider
}

// NewApp возвращает объект приложения
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(_ context.Context) error {
	err := a.initConfig()
	if err != nil {
		return err
	}

	a.initServiceProvider()

	return nil
}

func (a *App) initConfig() error {
	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider() {
	a.ServiceProvider = NewServiceProvider()
}
