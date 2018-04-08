package main

import (
	"os"

	"github.com/therecipe/qt/quickcontrols2"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	// Create application
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	// material, imagine, universe
	quickcontrols2.QQuickStyle_SetStyle("material")

	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	// Load the main qml file
	engine.Load(core.NewQUrl3("qrc:qml/main.qml", 0))

	// Execute app
	gui.QGuiApplication_Exec()
}
