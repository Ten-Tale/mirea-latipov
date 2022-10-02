package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dixonwille/wmenu"
)

const filesDir string = "files"

type option struct {
	title   string
	handler func(wmenu.Opt) error
}

func main() {

	if err := createMenuTemplate("MAIN"); err != nil {
		log.Fatal("something went wrong")
	}
}

func createMenuTemplate(key string) error {
	menu := wmenu.NewMenu("Select one of the options")
	menu.LoopOnInvalid()

	switch key {
	case "MAIN":
		applyOptionList(menu, []option{
			{
				"Create file",
				func(o wmenu.Opt) error {
					createMenuTemplate("CREATE_FILE")
					return nil
				},
			},
			{
				"Write to file",
				func(o wmenu.Opt) error {
					createMenuTemplate("WRITE_TO_FILE")
					return nil
				},
			},
			{
				"Read file",
				func(o wmenu.Opt) error {
					createMenuTemplate("READ_FILE")
					return nil
				},
			},
			{
				"Delete file",
				func(o wmenu.Opt) error {
					createMenuTemplate("DELETE_FILE")
					return nil
				},
			},
		})

	case "CREATE_FILE":
		applyOptionList(menu, []option{
			{
				"Create file",
				createFileHandler,
			},
			{
				"Go back",
				nil,
			},
		})

	case "WRITE_TO_FILE":
		var optionList []option

		fileList := listFiles()

		for _, file := range fileList {
			optionList = append(optionList, option{
				title:   file,
				handler: writeToFileHandler,
			})
		}

		applyOptionList(menu, optionList)

	case "READ_FILE":
		var optionList []option

		fileList := listFiles()

		for _, file := range fileList {
			optionList = append(optionList, option{
				title:   file,
				handler: readFileHandler,
			})
		}

		applyOptionList(menu, optionList)

	case "DELETE_FILE":
		var optionList []option

		fileList := listFiles()

		for _, file := range fileList {
			optionList = append(optionList, option{
				title:   file,
				handler: deleteFileHandler,
			})
		}

		applyOptionList(menu, optionList)
	}

	return menu.Run()
}

func applyOptionList(m *wmenu.Menu, ol []option) {
	for _, o := range ol {
		m.Option(o.title, nil, false, o.handler)
	}
}

func listFiles() []string {
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		err = os.Mkdir(filesDir, 0755)

		if err != nil {
			log.Fatal(err)
		}
	}

	files, err := os.ReadDir(filesDir)

	if err != nil {
		log.Fatal(err)
	}

	var resultList []string

	for _, file := range files {
		resultList = append(resultList, file.Name())
	}

	return resultList
}

func getInputText() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input file name")
	fmt.Println("---------------------")

	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	text = strings.Replace(text, "\r\n", "", -1)
	return text
}

func createFileHandler(o wmenu.Opt) error {
	text := getInputText()

	if _, err := os.Create(fmt.Sprintf("%s/%s", filesDir, text)); err != nil {
		log.Fatal(err)
	}

	defer goToMainMenuHandler()

	return nil
}

func writeToFileHandler(o wmenu.Opt) error {
	text := getInputText()

	file, err := os.Open(fmt.Sprintf("%s/%s", filesDir, o.Text))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err := file.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}

	defer goToMainMenuHandler()

	return nil
}

func readFileHandler(o wmenu.Opt) error {
	fileContent, err := os.ReadFile(fmt.Sprintf("%s/%s", filesDir, o.Text))

	if err != nil {
		log.Fatal(err)
	}

	defer goToMainMenuHandler()

	fmt.Println(fileContent)

	return nil
}

func deleteFileHandler(o wmenu.Opt) error {
	err := os.Remove(fmt.Sprintf("%s/%s", filesDir, o.Text))
	if err != nil {
		log.Fatal(err)
	}

	defer goToMainMenuHandler()

	return nil
}

func goToMainMenuHandler() {
	createMenuTemplate("MAIN")
}
