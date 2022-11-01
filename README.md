# Go-React Todo Application
A modern, full-stack todo application built with Go and React, featuring real-time updates, dark mode, and CRUD operations.

## 🚀 Tech Stack

### Backend
- **Go** - Server-side language
- **Fiber** - Web framework
- **MongoDB** - Database
- **godotenv** - Environment management

### Frontend
- **React** - UI library
- **TypeScript** - Type safety
- **TanStack Query** - Data fetching & caching
- **ChakraUI** - Component library

## ✨ Features

- ✅ Complete CRUD operations for todos
- 🌓 Light/Dark theme support
- 📱 Responsive design
- 🔄 Real-time updates with TanStack Query
- 🔒 Environment variable configuration
- 🌐 CORS enabled for local development
- 📊 MongoDB Atlas integration

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/go-react-todo
   cd go-react-todo
   ```

2. **Set up environment variables**
   Create a `.env` file in the root directory:
   ```env
   MONGODB_URI=your_mongodb_connection_string
   PORT=5050
   ENV=development
   ```

3. **Install backend dependencies**
   ```bash
   go mod download
   ```

4. **Install frontend dependencies**
   ```bash
   cd client
   npm install
   ```

## 🚀 Running the Application

### Development Mode

1. **Start the backend server**
   ```bash
   go run main.go
   ```
   The server will start at `http://localhost:5050`

2. **Start the frontend development server**
   ```bash
   cd client
   npm run dev
   ```
   The frontend will be available at `http://localhost:5173`

### Production Mode

1. **Build the frontend**
   ```bash
   cd client
   npm run build
   ```

2. **Start the server with production environment**
   ```bash
   ENV=production go run main.go
   ```

## 📁 Project Structure

```
├── main.go              # Main server file
├── .env                 # Environment variables
├── client/             # Frontend React application
│   ├── src/
│   ├── package.json
│   └── vite.config.ts
└── README.md
```

## 🔗 API Endpoints

| Method | Endpoint         | Description      |
|--------|-----------------|------------------|
| GET    | /api/todos      | Get all todos    |
| POST   | /api/todos      | Create new todo  |
| PATCH  | /api/todos/:id  | Update todo      |
| DELETE | /api/todos/:id  | Delete todo      |

## 💡 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [TanStack Query](https://tanstack.com/query)
- [ChakraUI](https://chakra-ui.com/)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

---
Built with ❤️ using Go and React
