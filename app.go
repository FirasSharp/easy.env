package main

import (
	"context"
	easyenv "github.com/FriscPlusPlus/easy.env/pkg/easyenvlib"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a *App) FilePrompt() bool {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return false
	}
	_, err = a.easy.Load(file)

	return err == nil
}

func (a *App) ListDatabases() []string {
	dbs := a.easy.GetDatabases()
	var dbNames []string
	for _, db := range dbs {
		dbNames = append(dbNames, db.Name)
	}
	return dbNames
}
