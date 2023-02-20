package main

import "fmt"

func printModuleToConsole(module *Module) {
	text := formatModuleOutput(module)
	fmt.Printf("%s\n", text)
}

func formatModuleOutput(module *Module) string {
	return fmt.Sprintf("%s/%s/%s - %s - %s",
		module.Namespace,
		module.Name,
		module.Provider,
		module.Version,
		module.Description)
}

func printModuleInfo(module *Module) {
	var text = formatModuleInfo(module)
	fmt.Printf("%s\n", text)
}

func formatModuleInfo(module *Module) string {
	title := fmt.Sprintf("%s/%s/%s", module.Namespace, module.Name, module.Provider)
	return fmt.Sprintf(
		"%s\n"+
			"Namespace: %s\n"+
			"Name: %s\n"+
			"Provider: %s\n"+
			"Version: %s\n"+
			"Description: %s\n",
		title, module.Namespace, module.Name, module.Provider, module.Version, module.Description)
}
