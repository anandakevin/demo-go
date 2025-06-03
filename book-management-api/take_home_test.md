## 📘 **Take Home Test – Book Management REST API (Golang)**

### 🧾 Description

## Create a simple **REST API** using **Golang** to manage book data without a database (in-memory only). The application must support full **CRUD** operations with **unique ISBN validation**, and **pagination**.

### 📚 Book Data

Each book must contain:

- `title` (string): Book title
- `author` (string): Author name
- `isbn` (string): ISBN number, must be **unique**
- `release_date` (string, format: `YYYY-MM-DD`): Release date

---

### 🔧 Required Features

1. **Create New Book**

   - Endpoint: `POST /books`
   - Validates uniqueness of ISBN
   - Simultaneously, simulate writing a log (e.g., "Book created: {isbn}")

2. **Get All Books (with Pagination)**

   - Endpoint: `GET /books?page=1&limit=10`

3. **Get Book by ISBN**

   - Endpoint: `GET /books/{isbn}`

4. **Update Book by ISBN**

   - Endpoint: `PUT /books/{isbn}`

5. **Delete Book by ISBN**

   - Endpoint: `DELETE /books/{isbn}`
   - Simultaneously, simulate writing a log (e.g., "Book deleted: {isbn}")

---

### ⚙️ Technical Requirements

- Language: **Go (Golang)**
- Must use Go’s standard library net/http only

  - No external routing frameworks (e.g., no Gin, Chi, Echo, etc.)

- Store all data in memory
- Clean, idiomatic Go code with proper error handling
- Modular structure and clear separation of concerns
- Include comments or README explanation
- Writing a log to simulated log file asycrhronously

---

### ✅ Evaluation Criteria

- Correct and complete implementation
- Appropriate and effective use of goroutines
- Clean, maintainable, idiomatic Go code
- Error handling and validations
- Bonus: use of channels or worker pattern for background processing

---

### 📦 Deliverables

- Source code (zip)
- `README.md` including:

  - How to run the app
  - API endpoint documentation
