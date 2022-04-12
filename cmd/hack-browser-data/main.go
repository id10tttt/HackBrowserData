package main

import (
	"fmt"
	"os"

	"hack-browser-data/internal/browser"
	"hack-browser-data/internal/log"
	"hack-browser-data/internal/outputter"
)

var (
	browserName       string
	exportDir         string
	outputFormat      string
	verbose           bool
	compress          bool
	customProfilePath string
)

func main() {
	Execute()
}

func Execute() {
	browsers := browser.PickBrowser("firefox")
	log.InitLog("debug")
	filetype := "json"
	dir := "results"
	// dir := "result"
	output := outputter.New(filetype)
	var file *os.File
	fmt.Println("browser number:", len(browsers))
	for _, b := range browsers {
		browsingData, err := b.GetBrowsingData()
		if err != nil {
			panic(err)
		}
		for _, source := range browsingData.Sources {
			filename := fmt.Sprintf("%s_%s.%s", b.Name(), source.Name(), filetype)
			file, err = output.CreateFile(dir, filename)
			err = output.Write(source, file)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

// func Execute() {
// 	app := &cli.App{
// 		Name:  "hack-browser-data",
// 		Usage: "Export passwords/cookies/history/bookmarks from browser",
// 		UsageText: "[hack-browser-data -b chrome -f json -dir results -cc]\n 	Get all browingdata(password/cookie/history/bookmark) from chrome",
// 		Version: "0.3.7",
// 		Flags: []cli.Flag{
// 			&cli.BoolFlag{Name: "verbose", Aliases: []string{"vv"}, Destination: &verbose, Value: false, Usage: "verbose"},
// 			&cli.BoolFlag{Name: "compress", Aliases: []string{"cc"}, Destination: &compress, Value: false, Usage: "compress result to zip"},
// 			&cli.StringFlag{Name: "browser", Aliases: []string{"b"}, Destination: &browserName, Value: "all", Usage: "available browsers: all|" + strings.Join(core.ListBrowser(), "|")},
// 			&cli.StringFlag{Name: "results-dir", Aliases: []string{"dir"}, Destination: &exportDir, Value: "results", Usage: "export dir"},
// 			&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Destination: &outputFormat, Value: "csv", Usage: "format, csv|json|console"},
// 			&cli.StringFlag{Name: "profile-dir-path", Aliases: []string{"p"}, Destination: &customProfilePath, Value: "", Usage: "custom profile dir path, get with chrome://version"},
// 		},
// 		HideHelpCommand: true,
// 		Action: func(c *cli.Context) error {
// 			var (
// 				browsers []core.Browser
// 				err      error
// 			)
// 			if verbose {
// 				log.InitLog("debug")
// 			} else {
// 				log.InitLog("error")
// 			}
// 			if customProfilePath != "" {
// 				browsers, err = core.PickCustomBrowser(browserName, customProfilePath, customKeyPath)
// 				if err != nil {
// 					log.Error(err)
// 				}
// 			} else {
// 				// default select all browsers
// 				browsers, err = core.PickBrowser(browserName)
// 				if err != nil {
// 					log.Error(err)
// 				}
// 			}
// 			err = utils.MakeDir(exportDir)
// 			if err != nil {
// 				log.Error(err)
// 			}
// 			for _, browser := range browsers {
// 				err := browser.InitSecretKey()
// 				if err != nil {
// 					log.Error(err)
// 				}
// 				// default select all items
// 				// you can get single item with browser.GetItem(itemName)
// 				items, err := browser.GetAllItems()
// 				if err != nil {
// 					log.Error(err)
// 				}
// 				name := browser.Name()
// 				key := browser.GetSecretKey()
// 				for _, item := range items {
// 					err := item.CopyDB()
// 					if err != nil {
// 						log.Error(err)
// 					}
// 					switch browser.(type) {
// 					case *core.Chromium:
// 						err := item.ChromeParse(key)
// 						if err != nil {
// 							log.Error(err)
// 						}
// 					case *core.Firefox:
// 						err := item.FirefoxParse()
// 						if err != nil {
// 							log.Error(err)
// 						}
// 					}
// 					err = item.Release()
// 					if err != nil {
// 						log.Error(err)
// 					}
// 					err = item.OutPut(outputFormat, name, exportDir)
// 					if err != nil {
// 						log.Error(err)
// 					}
// 				}
// 			}
// 			if compress {
// 				err = utils.Compress(exportDir)
// 				if err != nil {
// 					log.Error(err)
// 				}
// 			}
// 			return nil
// 		},
// 	}
// 	err := app.Run(os.Args)
// 	if err != nil {
// 		log.Error(err)
// 	}
// }