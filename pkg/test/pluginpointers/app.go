package pluginpointers

import (
	"sync"
)

type App struct {
	plugins map[string]Plugin
}

func NewApp() *App {
	return &App{
		plugins: make(map[string]Plugin),
	}
}
func (a *App) RegisterPlugin(plugin Plugin) {
	a.plugins[plugin.GetName()] = plugin

}
func (a *App) Run() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go a.RunHandler(&wg)
	}

	wg.Wait()
}
func (a *App) RunHandler(wg *sync.WaitGroup) {
	defer wg.Done()

	steps := make(map[string]StepHandle)

	for _, plugin := range a.plugins {
		handler := plugin.CreateHandler(1)
		for key, handle := range handler.StepHandles() {
			steps[key] = handle
		}
	}

	handle, ok := steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["xaxaxa"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
	handle, ok = steps["I add HTTP header :name with value :value"]
	if ok {
		args := StepArgs{}
		handle(args)
	}
}
