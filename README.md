# University Registration System (Go)

A concurrent, client-server university registration system featuring a **Fyne GUI** and a thread-safe **TCP Backend**.

## 🚀 Features
- **Concurrency:** Handles multiple student connections using Goroutines.
- **Thread Safety:** Uses Mutex locking to prevent over-enrollment.
- **Networking:** Communicates via TCP/IP using JSON serialization.
- **GUI:** User-friendly interface for course selection and status updates.

## 🛠 Installation & Usage
1. **Clone the repo:**
   ```bash
   git clone [https://github.com/YOUR_USERNAME/University-Registration-Go.git](https://github.com/YOUR_USERNAME/University-Registration-Go.git)
   ```
2. **Run the Server:**
   ```bash
   go run server.go
   ```
3. **Run the Client:**
   ```bash
   go run client.go
   ```

## 🏗 System Architecture
The system uses an Object-Oriented approach in Go, where the `Course` struct encapsulates registration logic, ensuring that capacity limits are strictly enforced across all concurrent threads.
```
