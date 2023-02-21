package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

const REGISTRY_URL = "https://registry.terraform.io/v1/modules"

func main() {
	var cmdSearch = &cobra.Command{
		Use:   "search [search in module name or description]",
		Short: "Print anything to the console",
		Long:  `A CLI application for searching terraform modules`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			processSearchCommand(args[0])
		},
	}

	var cmdInfo = &cobra.Command{
		Use:   "info [display module details]",
		Short: "Print anything to the console",
		Long:  `A CLI application for searching terraform modules`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			processInfoCommand(strings.TrimRight(args[0], "\n"))
		},
	}

	var cmdUpdate = &cobra.Command{
		Use:   "update [update modules cache]",
		Short: "Print anything to the console",
		Long:  `A CLI application for searching terraform modules`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			processUpdateCommand()
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdSearch)
	rootCmd.AddCommand(cmdInfo)
	rootCmd.AddCommand(cmdUpdate)
	rootCmd.Execute()

}

func patternMatchesModuleMetadata(pattern string, module *Module) bool {
	return strings.Contains(module.Name, pattern) || strings.Contains(module.Description, pattern)
}

func resolveModulesDir() string {
	cacheDir, _ := os.UserCacheDir()
	return fmt.Sprintf("%s/trcli", cacheDir)
}

func processInfoCommand(moduleName string) {
	mc := ModulesCache{resolveModulesDir()}
	module := mc.Find(moduleName)
	if module == nil {
		fmt.Printf("Module '%s' not found\n", moduleName)
		return
	}
	printModuleInfoToConsole(module)
}

func processUpdateCommand() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	mc := RebuildModulesCacheDir(resolveModulesDir())
	browser := RegistryBrowser{REGISTRY_URL, mc.CacheFile(),
		&SimpleModuleFetcher{REGISTRY_URL, 100},
		&SimpleModuleWriter{}}
	browser.LoopThroughModules()
	s.Stop()
}

func processSearchCommand(pattern string) {
	mc := ModulesCache{resolveModulesDir()}

	modules := loadModules(mc.CacheFile())

	for _, module := range modules {
		if patternMatchesModuleMetadata(pattern, &module) {
			printModuleToConsole(&module)
		}
	}
}
