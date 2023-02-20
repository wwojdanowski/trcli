package main

import "fmt"

func printModuleToConsole(module *Module) {
	text := formatModuleOutput(module)
	fmt.Printf("%s\n", text)
}

func formatModuleOutput(module *Module) string {
	return fmt.Sprintf("%s/%s/%s:%s - %s",
		module.Namespace,
		module.Name,
		module.Provider,
		module.Version,
		module.Description)
}

func printModuleInfo(module *Module) {
	text := formatModuleInfo(module)
	fmt.Printf("%s\n", text)
}

func formatModuleInfo(module *Module) string {
	return fmt.Sprintf("%s (%s)\n\n%s", module.Name, module.ID, module.Description)
}
