package etc

import (
	"github.com/Ten-Tale/task-1/filejson"
	"github.com/Ten-Tale/task-1/filesystem"
	"github.com/Ten-Tale/task-1/fileworker"
	"github.com/Ten-Tale/task-1/filexml"
	"github.com/Ten-Tale/task-1/helpers/menuhelper"
	"github.com/dixonwille/wmenu"
)

func CreateMainMenu() {
	menu := wmenu.NewMenu("Select one of the options")
	menu.LoopOnInvalid()

	menuhelper.ApplyOptionList(menu, []menuhelper.Option{
		{
			Title: "Show disk info",
			Handler: func(o wmenu.Opt) error {
				filesystem.ShowDeviceInfo()

				CreateMainMenu()

				return nil
			},
		},
		{
			Title: "File worker",
			Handler: func(o wmenu.Opt) error {
				fileworker.RunFileWorker()

				CreateMainMenu()

				return nil
			},
		},
		{
			Title: "JSON worker",
			Handler: func(o wmenu.Opt) error {
				filejson.RunJSONFileWorker()
				CreateMainMenu()

				return nil
			},
		},
		{
			Title: "XML worker",
			Handler: func(o wmenu.Opt) error {
				filexml.RunXMLFileWorker()
				CreateMainMenu()

				return nil
			},
		},
	})

	menu.Run()
}
