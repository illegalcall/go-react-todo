package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/illegalcall/go-react-todo/server/database"
	"github.com/illegalcall/go-react-todo/server/models"
)

// Mock database setup
func setupMockDB(mt *mtest.T) {
	database.Collection = mt.Coll
}

// Create a test Fiber app
func setupTestApp() *fiber.App {
	app := fiber.New()
	app.Get("/api/todos", GetTodos)
	app.Post("/api/todos", CreateTodo)
	app.Patch("/api/todos/:id", UpdateTodo)
	app.Delete("/api/todos/:id", DeleteTodo)
	return app
}

// ✅ Test GET /api/todos (empty database)
func TestGetTodos_EmptyDB(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer func() {
		if mt.Client != nil {
			_ = mt.Client.Disconnect(context.Background())
		}
	}()

	mt.Run("Get todos - empty database", func(mt *mtest.T) {
		setupMockDB(mt)

		// Mock an empty response
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "golang_db.todos", mtest.FirstBatch))

		app := setupTestApp()
		req := httptest.NewRequest(http.MethodGet, "/api/todos", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var todos []models.Todo
		_ = json.NewDecoder(resp.Body).Decode(&todos)
		assert.Len(t, todos, 0) // Expect empty list
	})
}

// ✅ Test GET /api/todos (Database error)
func TestGetTodos_DBError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer func() {
		if mt.Client != nil {
			_ = mt.Client.Disconnect(context.Background())
		}
	}()

	mt.Run("Database error on fetching todos", func(mt *mtest.T) {
		setupMockDB(mt)

		// Mock a database error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    1234,
			Message: "Database error",
		}))

		app := setupTestApp()
		req := httptest.NewRequest(http.MethodGet, "/api/todos", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}

// ✅ Test POST /api/todos (empty request body)
func TestCreateTodo_EmptyBody(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodPost, "/api/todos", nil) // No body
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// ✅ Test POST /api/todos (invalid JSON)
func TestCreateTodo_InvalidJSON(t *testing.T) {
	app := setupTestApp()
	invalidJSON := []byte(`{invalid json}`) // Corrupt JSON

	req := httptest.NewRequest(http.MethodPost, "/api/todos", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// ✅ Test PATCH /api/todos/:id (invalid ID)
func TestUpdateTodo_InvalidID(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodPatch, "/api/todos/invalid_id", nil) // Invalid ID
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// ✅ Test PATCH /api/todos/:id (database failure)
func TestUpdateTodo_DBFailure(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer func() {
		if mt.Client != nil {
			_ = mt.Client.Disconnect(context.Background())
		}
	}()

	mt.Run("Database update failure", func(mt *mtest.T) {
		setupMockDB(mt)

		todoID := primitive.NewObjectID().Hex()

		// Mock MongoDB UpdateOne failure
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "Update failed",
		}))

		app := setupTestApp()
		req := httptest.NewRequest(http.MethodPatch, "/api/todos/"+todoID, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}

// ✅ Test DELETE /api/todos/:id (invalid ID)
func TestDeleteTodo_InvalidID(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodDelete, "/api/todos/invalid_id", nil) // Invalid ID
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// ✅ Test DELETE /api/todos/:id (database failure)
func TestDeleteTodo_DBFailure(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer func() {
		if mt.Client != nil {
			_ = mt.Client.Disconnect(context.Background())
		}
	}()

	mt.Run("Database delete failure", func(mt *mtest.T) {
		setupMockDB(mt)

		todoID := primitive.NewObjectID().Hex()

		// Mock MongoDB DeleteOne failure
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "Delete failed",
		}))

		app := setupTestApp()
		req := httptest.NewRequest(http.MethodDelete, "/api/todos/"+todoID, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
