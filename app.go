package main

import (
	"context"
	"fmt"
	easyenv "github.com/FriscPlusPlus/easy.env/pkg/easyenvlib"
)

// App struct
type App struct {
	ctx  context.Context
	easy *easyenv.EasyEnv
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.easy = easyenv.NewEasyEnv()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) MyTest(name string) string {
	return fmt.Sprintf("Hey %s, you are a wanama", name)
}

func (a *App) LoadDB(path string) error {
	_, err := a.easy.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetProjects() string {
	return ""
}