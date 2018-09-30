package gRPC

import (
	"golang-tdd-rest-api/todo"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
)

type TodoGRPCServer struct {
	todoController todo.Controller
}

func NewHandler(todoController todo.Controller) *TodoGRPCServer {
	return &TodoGRPCServer{
		todoController: todoController,
	}
}

func (s *TodoGRPCServer) Get(ctx context.Context, in *TodosRequest) (*TodosReply, error) {
	todos, err := s.todoController.Get()
	if err != nil {
		return nil, err
	}

	todosReply := TodosReply{}

	for _, todo := range todos {
		var convertedCreatedAt, convertedUpdatedAt, convertedDeletedAt *timestamp.Timestamp
		if todo.CreatedAt != nil {
			convertedCreatedAt, _ = ptypes.TimestampProto(*todo.CreatedAt)
		}
		if todo.UpdatedAt != nil {
			convertedUpdatedAt, _ = ptypes.TimestampProto(*todo.UpdatedAt)
		}
		if todo.DeletedAt != nil {
			convertedDeletedAt, _ = ptypes.TimestampProto(*todo.DeletedAt)
		}

		todosReply.Todo = append(todosReply.Todo, &TodoReply{
			ID:        todo.ID,
			Content:   todo.Content,
			Done:      todo.Done,
			CreatedAt: convertedCreatedAt,
			UpdatedAt: convertedUpdatedAt,
			DeletedAt: convertedDeletedAt,
		})
	}

	return &todosReply, nil
}
