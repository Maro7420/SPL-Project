package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

// --- Object-Oriented Data Models ---
type Course struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Enrolled int    `json:"enrolled"`
}

// Encapsulation: Method to handle business logic safely
func (c *Course) RegisterStudent() bool {
	if c.Enrolled < c.Capacity {
		c.Enrolled++
		return true
	}
	return false
}

type RegistrationSystem struct {
	mu      sync.Mutex // Ensures Thread Safety
	Courses map[string]*Course
}

func NewRegistrationSystem() *RegistrationSystem {
	return &RegistrationSystem{
		Courses: map[string]*Course{
			"CSC230": {Code: "CSC230", Name: "Computer Architecture & Assembly", Capacity: 30, Enrolled: 15},
			"NET200": {Code: "NET200", Name: "Network Engineering & Security", Capacity: 25, Enrolled: 10},
			"MAT201": {Code: "MAT201", Name: "Advanced Calculus", Capacity: 40, Enrolled: 20},
		},
	}
}

// --- Networking Layer ---
func handleRequest(conn net.Conn, sys *RegistrationSystem) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		var req map[string]string
		json.Unmarshal([]byte(scanner.Text()), &req)

		sys.mu.Lock() // Locking to prevent race conditions during registration
		var response []byte
		
		if req["action"] == "get_courses" {
			list := []Course{}
			for _, c := range sys.Courses {
				list = append(list, *c)
			}
			response, _ = json.Marshal(list)
			
		} else if req["action"] == "register" {
			course, exists := sys.Courses[req["course_code"]]
			msg := "Error: Course not found"
			
			if exists && course.RegisterStudent() {
				msg = "Successfully registered for " + course.Code
			} else if exists {
				msg = "Registration failed: Course is full"
			}
			response, _ = json.Marshal(map[string]string{"message": msg})
		}
		
		sys.mu.Unlock()
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
	
	fmt.Println("[LISTENING] Server is running on port 65432...")
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// Concurrency: Goroutine allows multiple students to connect at once
		go handleRequest(conn, sys) 
	}
}