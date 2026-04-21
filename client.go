package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func callServer(req map[string]string) string {
	conn, err := net.Dial("tcp", "127.0.0.1:65432")
	if err != nil {
		return `{"message": "Cannot connect to server."}`
	}
	defer conn.Close()

	reqBytes, _ := json.Marshal(req)
	fmt.Fprintln(conn, string(reqBytes))

	resp, _ := bufio.NewReader(conn).ReadString('\n')
	return resp
}

func main() {
	myApp := app.New()
	win := myApp.NewWindow("Student Services Registration")
	win.Resize(fyne.NewSize(400, 300))

	label := widget.NewLabel("Select a course to register:")
	
	// Dropdown menu for course selection
	courseList := widget.NewSelect([]string{"CSC230", "NET200", "MAT201"}, func(value string) {})

	registerBtn := widget.NewButton("Register Now", func() {
		if courseList.Selected == "" {
			dialog.ShowInformation("Warning", "Please select a course first.", win)
			return
		}
		
		// Network call to server
		resp := callServer(map[string]string{"action": "register", "course_code": courseList.Selected})
		
		var result map[string]string
		json.Unmarshal([]byte(resp), &result)
		
		// Display GUI Pop-up with the result
		dialog.ShowInformation("Registration Status", result["message"], win)
	})

	content := container.NewVBox(
		label,
		courseList,
		widget.NewLabel(""), // Spacer
		registerBtn,
	)

	win.SetContent(content)
	win.ShowAndRun()
}