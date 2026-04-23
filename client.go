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

// callServer establishes a TCP connection, sends JSON, and waits for a response
func callServer(req map[string]string) string {
	conn, err := net.Dial("tcp", "127.0.0.1:65432")
	if err != nil {
		return `{"message": "❌ Cannot connect to server. Is it running?"}`
	}
	defer conn.Close()

	// Serialize and transmit the request
	reqBytes, _ := json.Marshal(req)
	fmt.Fprintln(conn, string(reqBytes))

	// Read the incoming response from the server
	resp, _ := bufio.NewReader(conn).ReadString('\n')
	return resp
}

func main() {
	myApp := app.New()
	win := myApp.NewWindow("Student Services Registration")
	win.Resize(fyne.NewSize(450, 400))

	// Create input fields for student data
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter Full Name...")

	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("Enter Student ID...")

	// Dropdown menu for course selection
	courseList := widget.NewSelect([]string{"CSC230", "NET200", "MAT201"}, func(value string) {})

	// Button trigger logic
	registerBtn := widget.NewButton("Register & Send Request", func() {

		// Client-Side Validation: Ensure the student filled in all required fields
		if nameEntry.Text == "" || idEntry.Text == "" || courseList.Selected == "" {
			dialog.ShowInformation("Missing Info", "⚠️ Please fill in your Name, ID, and select a course.", win)
			return
		}

		// Prepare the request payload with the input data
		requestData := map[string]string{
			"action":       "register",
			"course_code":  courseList.Selected,
			"student_name": nameEntry.Text,
			"student_id":   idEntry.Text,
		}

		// Send the request and receive the JSON response
		resp := callServer(requestData)

		var result map[string]string
		json.Unmarshal([]byte(resp), &result)

		// Display the confirmation message received from the server
		dialog.ShowInformation("Server Response", result["message"], win)
	})

	// Design the vertical UI layout
	content := container.NewVBox(
		widget.NewLabel("Student Name:"),
		nameEntry,
		widget.NewLabel("Student ID:"),
		idEntry,
		widget.NewLabel("Select Course:"),
		courseList,
		widget.NewLabel(""), // Empty space for better UI spacing
		registerBtn,
	)

	win.SetContent(content)
	win.ShowAndRun()
}
