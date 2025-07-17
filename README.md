# Message Api with go

> Hello, in this project, I made a sample message sending application with Go. The messages go to a single source and are stored as a source variable. No database was used. It was developed with Fiber v2.

## 🔎 Using

### Clone this repository
```bash
git clone "https://github.com/Hasankarabulut4/messaging-api.git"
```
---

### Initialize Golang project
```bash
go mod tidy
```

---

### Start main.go file
```bash
go run .
```

---

### ⚒️ Endpoints

> GET | http://localhost:8000/message

<p>Get all messages</p>

---

> POST | http://localhost:8000/message/send


#### Example json raw data
```
{
    "message":"Your message content",
    "username":"John Doe"
}
```

<p>Send a message</p>

---

> DELETE | http://localhost:8000/message/delete/{id}

<p>Delete a message</p>
