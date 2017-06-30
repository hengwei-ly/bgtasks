package app

import (
	"bgtasks/app/libs"
	"bgtasks/app/models"
	"bgtasks/app/routes"

	"github.com/revel/revel"
	"github.com/three-plus-three/forms"

	_ "github.com/three-plus-three/modules/bind"
	"github.com/three-plus-three/modules/toolbox"
	"github.com/three-plus-three/modules/web_ext"
)

var Lifecycle *libs.Lifecycle

func init() {
	web_ext.Init("bgtasks", "例子", "",
		func(data *web_ext.Lifecycle) error {
			//if err := models.DropTables(data.ModelEngine); err != nil {
			//	return err
			//}
			if err := models.InitTables(data.ModelEngine); err != nil {
				return err
			}

			data.Variables["userLevel"] = []forms.InputChoice{{Value: "1", Label: "high"},
				{Value: "2", Label: "modium"},
				{Value: "3", Label: "low"}}
			revel.TemplateFuncs["userLevel_format"] = func(level string) string {
				switch level {
				case "1":
					return "high"
				case "2":
					return "modium"
				case "3":
					return "low"
				}
				return level
			}

			Lifecycle = &libs.Lifecycle{
				Lifecycle: data,
				DB:        models.DB{Engine: data.ModelEngine},
				DataDB:    models.DB{Engine: data.DataEngine},
			}
			return nil
		},
		func(data *web_ext.Lifecycle) ([]toolbox.Menu, error) {
			return []toolbox.Menu{
				{Title: "主页", Name: "Home", URL: routes.Home.Index()},
				{Title: "用户", URL: "#", Children: []toolbox.Menu{}}}, nil
		})

}
