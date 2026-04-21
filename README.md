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

---

### 5. Final Touch: Add a Screenshot
GitHub allows you to show off your work. 
1. Run your program and take a screenshot of the **GUI window** and the **Server terminal** side-by-side.
2. Upload the image to your repository.
3. You can display it in your README by adding this line:
   `![Project Screenshot](screenshot.png)`

Now your project is live! Do you need help setting up a `.gitignore` file to keep your folder clean of compiled binaries?
