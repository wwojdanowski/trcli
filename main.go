package main

import (
	"encoding/json"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const REGISTRY_URL = "https://registry.terraform.io/v1/modules"

func loopThroughModules(modulesPath string) {
	offset := 0
	limit := 100

	shouldContinue := true

	for i := 0; shouldContinue; i++ {
		response, err := http.Get(fmt.Sprintf("%s?limit=%d&offset=%d", REGISTRY_URL, limit, offset))
		if err != nil {
			log.Fatalln(err)
		}
		defer response.Body.Close()
		responseBody, _ := ioutil.ReadAll(response.Body)
		modulesInfo := ModulesInfo{}
		_ = json.Unmarshal([]byte(responseBody), &modulesInfo)
		//printAllModules(&modulesInfo)
		storeAllModules(modulesPath, &modulesInfo)
		shouldContinue = modulesInfo.Meta.NextOffset > 0
		offset = modulesInfo.Meta.NextOffset

	}
}

func printAllModules(modulesInfo *ModulesInfo) {
	for i := 0; i < len(modulesInfo.Modules); i++ {
		fmt.Println(modulesInfo.Modules[i].ID, modulesInfo.Modules[i].Name)
	}
}

func storeAllModules(modulesFile string, modulesInfo *ModulesInfo) {
	for i := 0; i < len(modulesInfo.Modules); i++ {
		storeModuleData(&modulesInfo.Modules[i], modulesFile)
	}
}

func storeModuleData(module *Module, fileName string) {
	data, _ := json.Marshal(module)

	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s\n", string(data))); err != nil {
		log.Println(err)
	}
}

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
	printModuleDetail(module)
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

	fmt.Println()
	for _, module := range modules {
		if patternMatchesModuleMetadata(pattern, &module) {
			fmt.Printf("%s - %s\n", module.Name, module.Description)
		}
	}
}

func filterModules() {

}
