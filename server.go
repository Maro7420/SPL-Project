package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

// Course struct encapsulates the data for a specific class
type Course struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Enrolled int    `json:"enrolled"`
}

// RegisterStudent is a receiver method that safely handles the enrollment logic
func (c *Course) RegisterStudent() bool {
	if c.Enrolled < c.Capacity {
		c.Enrolled++
		return true
	}
	return false
}

// RegistrationSystem acts as the centralized, thread-safe database
type RegistrationSystem struct {
	mu      sync.Mutex // Ensures thread safety across concurrent requests
	Courses map[string]*Course
}

// NewRegistrationSystem initializes the backend with sample courses
func NewRegistrationSystem() *RegistrationSystem {
	return &RegistrationSystem{
		Courses: map[string]*Course{
			"CSC230": {Code: "CSC230", Name: "Computer Architecture", Capacity: 30, Enrolled: 15},
			"NET200": {Code: "NET200", Name: "Network Engineering", Capacity: 25, Enrolled: 10},
			"MAT201": {Code: "MAT201", Name: "Advanced Calculus", Capacity: 40, Enrolled: 20},
		},
	}
}

// handleRequest processes individual client connections concurrently
func handleRequest(conn net.Conn, sys *RegistrationSystem) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	// Log when a new connection is established
	fmt.Printf("\n[+] New Connection established from: %s\n", conn.RemoteAddr().String())

	for scanner.Scan() {
		var req map[string]string
		json.Unmarshal([]byte(scanner.Text()), &req)

		sys.mu.Lock() // Lock the database to prevent race conditions
		var response []byte

		if req["action"] == "register" {
			// Extract data sent from the client GUI
			courseCode := req["course_code"]
			studentName := req["student_name"]
			studentID := req["student_id"]

			// Log the incoming request to the terminal for visibility
			fmt.Printf("[>>>] RECEIVED REQUEST: Student '%s' (ID: %s) wants to join '%s'\n", studentName, studentID, courseCode)

			course, exists := sys.Courses[courseCode]
			msg := "Error: Course not found"

			if exists && course.RegisterStudent() {
				// Format the success message containing the student's details
				msg = fmt.Sprintf("✅ Server Confirmed: Student [%s] (ID: %s) is now registered for %s", studentName, studentID, course.Code)

				// Log the successful registration
				fmt.Printf("[<<<] SUCCESS: Registered '%s'. Seats left: %d\n", studentName, course.Capacity-course.Enrolled)
			} else if exists {
				msg = "❌ Registration failed: Course is full"

				// Log the failure if the course is full
				fmt.Printf("[<<<] FAILED: Course '%s' is full.\n", courseCode)
			}

			response, _ = json.Marshal(map[string]string{"message": msg})
		}

		sys.mu.Unlock() // Unlock the database for the next user
		fmt.Fprintln(conn, string(response))
	}
}

func main() {
	sys := NewRegistrationSystem()
	listener, err := net.Listen("tcp", "127.0.0.1:65432")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	// Display startup interface in the terminal
	fmt.Println("========================================")
	fmt.Println("[LISTENING] Server is running on port 65432...")
	fmt.Println("Waiting for student requests...")
	fmt.Println("========================================")

	// Infinite loop to accept multiple incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// Create a new Goroutine for each connected student
		go handleRequest(conn, sys)
	}
}
