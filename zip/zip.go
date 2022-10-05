package zip

import (
	"archive/zip"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Ten-Tale/task-1/helpers/menuhelper"
	"github.com/dixonwille/wmenu"
)

const (
	filesDir    = "filesToArchive"
	archivesDir = "archieves"
)

func RunZIPWorker() {
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
				Title: "Create archieve",
				Handler: func(o wmenu.Opt) error {
					createMenuTemplate("CREATE_ARCHIEVE")
					return nil
				},
			},
			{
				Title: "Add files to archieve",
				Handler: func(o wmenu.Opt) error {
					createMenuTemplate("WRITE_TO_ARCHIEVE")
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
				Title: "Delete archieve",
				Handler: func(o wmenu.Opt) error {
					createMenuTemplate("DELETE_ARCHIEVE")
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

	case "CREATE_ARCHIEVE":
		menuhelper.ApplyOptionList(menu, []menuhelper.Option{
			{
				Title:   "Create archieve",
				Handler: createArchieveHandler,
			},
		})

	case "WRITE_TO_ARCHIEVE":
		var optionList []menuhelper.Option

		archieveList := listArchieves()

		if len(archieveList) == 0 {
			fmt.Println("No archieves that you can write to")
			goToMainMenuHandler()
			break
		}

		for _, file := range archieveList {
			optionList = append(optionList, menuhelper.Option{
				Title:   file,
				Handler: writeToArchieveHandler,
			})
		}

		menuhelper.ApplyOptionList(menu, optionList)

	case "LIST_FILES_TO_ARCHIEVE":
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
				Handler: func(o wmenu.Opt) error {},
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

func listArchieves() []string {
	files, err := os.ReadDir(archivesDir)

	if err != nil {
		log.Fatal(err)
	}

	var resultList []string

	for _, file := range files {
		resultList = append(resultList, file.Name())
	}

	return resultList
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

func getInputText(text string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(text)
	fmt.Println("---------------------")

	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	text = strings.Replace(text, "\r\n", "", -1)
	return text
}

func goToMainMenuHandler() {
	createMenuTemplate("MAIN")
}

func createArchieveHandler(o wmenu.Opt) error {
	archieveName := getInputText("Input archieve name")

	newArchieve, err := os.Create(fmt.Sprintf("%s/%s.zip", archivesDir, archieveName))

	if err != nil {
		log.Fatal(err)
	}

	return newArchieve.Close()
}

func writeToArchieveHandler(o wmenu.Opt) error {
	archieveWriter := chooseArchieveHandler(o.Text)

	return nil
}

func chooseArchieveHandler(filename string) *zip.Writer {
	archieve, err := os.OpenFile(fmt.Sprintf("%s/%s", archivesDir, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

	if err != nil {
		log.Fatal(err)
	}

	defer archieve.Close()

	return zip.NewWriter(archieve)
}

func chooseFileToArchieve() (string, error) {
	fileList := listFiles()

	optionList := []menuhelper.Option{}

	if len(fileList) == 0 {
		fmt.Println("No archieves that you can write to")
		goToMainMenuHandler()

		return "", errors.New("No files to write to archieve")
	}

	for _, file := range fileList {
		optionList = append(optionList, menuhelper.Option{
			Title:   file,
			Handler: writeToArchieveHandler,
		})
	}

	menuhelper.ApplyOptionList(menu, optionList)
}
