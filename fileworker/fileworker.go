package fileworker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Ten-Tale/task-1/helpers/menuhelper"
	"github.com/dixonwille/wmenu"
)

const filesDir = "files"

func RunFileWorker() {
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		err = os.Mkdir(filesDir, 0777)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := createMenuTemplate("MAIN"); err != nil {
		log.Fatal("something went wrong")
	}
}

func createMenuTemplate(key string) error {
	menu := wmenu.NewMenu("Select one of the options")
	menu.LoopOnInvalid()

	switch key {
	case "MAIN":
		menuhelper.ApplyOptionList(menu, []menuhelper.Option{
			{
				Title: "Create file",
				Handler: func(o wmenu.Opt) error {
					createMenuTemplate("CREATE_FILE")
					return nil
				},
			},
			{
				Title: "Write to file",
				Handler: func(o wmenu.Opt) error {
					createMenuTemplate("WRITE_TO_FILE")
					return nil
				},
			},
			{
				Title: "Read file",
				Handler: func(o wmenu.Opt) error {
					createMenuTemplate("READ_FILE")
					return nil
				},
			},
			{
				Title: "Delete file",
				Handler: func(o wmenu.Opt) error {
					createMenuTemplate("DELETE_FILE")
					return nil
				},
			},
			{
				Title: "Go to main menu",
				Handler: func(o wmenu.Opt) error {
					return nil
				},
			},
		})

	case "CREATE_FILE":
		menuhelper.ApplyOptionList(menu, []menuhelper.Option{
			{
				Title:   "Create file",
				Handler: createFileHandler,
			},
		})

	case "WRITE_TO_FILE":
		var optionList []menuhelper.Option

		fileList := listFiles()

		if len(fileList) == 0 {
			fmt.Println("No files that you can write to")
			goToMainMenuHandler()
			break
		}

		for _, file := range fileList {
			optionList = append(optionList, menuhelper.Option{
				Title:   file,
				Handler: writeToFileHandler,
			})
		}

		menuhelper.ApplyOptionList(menu, optionList)

	case "READ_FILE":
		var optionList []menuhelper.Option

		fileList := listFiles()

		if len(fileList) == 0 {
			fmt.Println("No files that you can read from")
			goToMainMenuHandler()
			break
		}

		for _, file := range fileList {
			optionList = append(optionList, menuhelper.Option{
				Title:   file,
				Handler: readFileHandler,
			})
		}

		menuhelper.ApplyOptionList(menu, optionList)

	case "DELETE_FILE":
		var optionList []menuhelper.Option

		fileList := listFiles()

		if len(fileList) == 0 {
			fmt.Println("No files to delete")
			goToMainMenuHandler()
			break
		}

		for _, file := range fileList {
			optionList = append(optionList, menuhelper.Option{
				Title:   file,
				Handler: deleteFileHandler,
			})
		}

		menuhelper.ApplyOptionList(menu, optionList)
	}

	if key != "MAIN" {
		menu.Option("Go to main menu", nil, false, func(o wmenu.Opt) error {
			goToMainMenuHandler()
			return nil
		})
	}

	return menu.Run()
}

func listFiles() []string {
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
	fmt.Println("Input text")
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

	if file, err := os.Create(fmt.Sprintf("%s/%s", filesDir, text)); err != nil {
		log.Fatal(err)
	} else {
		file.Close()
	}

	defer goToMainMenuHandler()

	return nil
}

func writeToFileHandler(o wmenu.Opt) error {
	text := getInputText()

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", filesDir, o.Text), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

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

	fmt.Println(string(fileContent))

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
