package main

import (
	"log"

	"github.com/nanu-c/qml-go"
	"github.com/nanu-c/ut-app-go-template/src"
)

func main() {
	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("qml/main.qml")
	if err != nil {
		return err
	}

	testvar := TestStruct{Message: "Testmessage", Number: 1}
	context := engine.Context()
	context.SetVar("testvar", &testvar)
	testvar.GetMessage()
	win := component.CreateWindow(nil)
	testvar.Root = win.Root()
	win.Show()
	win.Wait()
	return nil
}

type TestStruct struct {
	Root    qml.Object
	Message string
	Output  string
	Number  int
}

func (testvar *TestStruct) GetMessage() {
	log.Print("fetch second Message...")
	go func() {
		testvar.Number, testvar.Output = src.GetMessageFromOtherPackage(testvar.Number)
		qml.Changed(testvar, &testvar.Output)
	}()
}
