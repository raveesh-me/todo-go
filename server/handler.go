package server

import (
	"context"
	"sync"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	todov1 "github.com/raveesh-me/todo/gen/go/protos/todo/v1"
)

type TodoServer struct {
	mu    sync.Mutex
	todos []*todov1.Todo
}

func NewTodoServer() *TodoServer {
	return &TodoServer{}
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *connect.Request[todov1.CreateTodoRequest]) (*connect.Response[todov1.CreateTodoResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &todov1.Todo{
		Id:        uuid.NewString(),
		Title:     req.Msg.Title,
		Completed: false,
	}
	s.todos = append(s.todos, todo)

	return connect.NewResponse(&todov1.CreateTodoResponse{Todo: todo}), nil
}

func (s *TodoServer) ListTodos(ctx context.Context, req *connect.Request[todov1.ListTodosRequest]) (*connect.Response[todov1.ListTodosResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return connect.NewResponse(&todov1.ListTodosResponse{Todos: s.todos}), nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, req *connect.Request[todov1.DeleteTodoRequest]) (*connect.Response[todov1.DeleteTodoResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.todos {
		if todo.Id == req.Msg.Id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			break
		}
	}

	return connect.NewResponse(&todov1.DeleteTodoResponse{}), nil
}
