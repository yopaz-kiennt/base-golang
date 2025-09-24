package drbiz

import (
	"context"
	"errors"
	"backend/module/dailyreport/model"
)

// Interface nên được khai báo ở tại nơi sử dụng nó, tránh lỗi cycle import (các pakage import lẫn nhau).
type EventTodoStorage interface {
	CreateEvent(ctx context.Context, data *model.SaleReport) error
}

// Khai báo thuộc tính và phương thức
type saleReportBiz struct {
	create CreateTodoItemStorage
	get CreateTodoItemStorage
	update CreateTodoItemStorage
	delete CreateTodoItemStorage
}

func NewCreateToDoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *saleReportBiz) CreateNewItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if data.Title == "" {
		return errors.New("title can not be blank")
	}

	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
}

	