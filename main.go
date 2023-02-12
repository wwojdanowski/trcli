package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

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

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdSearch)
	rootCmd.AddCommand(cmdInfo)
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
	printModuleInfo(module)
}

func processSearchCommand(pattern string) {
	mc := ModulesCache{resolveModulesDir()}

	if !mc.PathExists() {
		mc.BuildFullPath()
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		loopThroughModules(mc.CacheFile())
		s.Stop()

	}

	modules := loadModules(mc.CacheFile())

	for _, module := range modules {
		if patternMatchesModuleMetadata(pattern, &module) {
			printModuleToConsole(&module)
		}
	}
}
