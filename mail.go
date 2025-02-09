package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func printLogo() {
	fmt.Println(Blue + `
___________.__.__           ________                            .__                     
\_   _____/|__|  |   ____   \_____  \_______  _________    ____ |__|_______ ___________ 
 |    __)  |  |  | _/ __ \   /   |   \_  __ \/ ___\__  \  /    \|  \___   // __ \_  __ \
 |     \   |  |  |_\  ___/  /    |    \  | \/ /_/  > __ \|   |  \  |/    /\  ___/|  | \/
 \___  /   |__|____/\___  > \_______  /__|  \___  (____  /___|  /__/_____ \\___  >__|   
     \/                 \/          \/     /_____/     \/     \/         \/    \/      
` + Reset)
	fmt.Println(Yellow + "  Organize your files 🚀" + Reset)
	fmt.Println("")
}

type Config struct {
	Extensions map[string]string `json:"extensions"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf(Red+"Config issue: %v"+Reset, err)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf(Red+"Error decoding config: %v"+Reset, err)
	}

	return &config, nil
}

func main() {
	printLogo() 

	configFile := flag.String("config", "config.json", "Path to config file")
	dir := flag.String("path", ".", "Path to organize")
	dryRun := flag.Bool("dry-run", false, "Perform a dry run")
	flag.Parse()

	config, err := loadConfig(*configFile)
	if err != nil {
		fmt.Println(Red + "❌ Error loading config:" + Reset, err)
		return
	}

	if config == nil {
		fmt.Println(Red + "❌ Error: Config is nil" + Reset)
		return
	}

	processDirectories(*dir, *dryRun, config.Extensions)
}

func processDirectories(dirname string, dryRun bool, allowedExtensions map[string]string) {
	f, err := os.Open(dirname)
	if err != nil {
		fmt.Println(Red + "❌ Error opening directory:" + Reset, err)
		return
	}
	defer f.Close()

	files, err := f.ReadDir(0)
	if err != nil {
		fmt.Println(Red + "❌ Error reading directory:" + Reset, err)
		return
	}

	var changes int
	for _, v := range files {
		if v.IsDir() || strings.HasPrefix(v.Name(), ".") {
			continue
		}

		ext := filepath.Ext(v.Name())
		folderName, ok := allowedExtensions[ext]
		if !ok || folderName == "" {
			continue
		}

		targetDir := filepath.Join(dirname, folderName)
		os.MkdirAll(targetDir, 0755)

		srcPath := filepath.Join(dirname, v.Name())
		destPath := filepath.Join(targetDir, v.Name())

		fmt.Printf(Green+"✔ Moving %s ➜ %s"+Reset+"\n", srcPath, destPath)

		if !dryRun {
			err := os.Rename(srcPath, destPath)
			if err != nil {
				fmt.Println(Red + "❌ Error moving file:" + Reset, err)
			} else {
				changes++
			}
		}
	}

	if changes > 0 {
		fmt.Printf(Yellow+"✨ Files Moved: %d"+Reset+"\n", changes)
	} else {
		fmt.Println(Blue + "ℹ️ No changes" + Reset)
	}
}
