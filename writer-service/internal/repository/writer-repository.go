package repository

import (
	"context"
	"writer/internal/model"
)

type WriterRepository interface {
	 Save(context.Context, model.Writer) (*model.Writer, error)
	 FindByID(context.Context, int64) (*model.Writer, error)
	 FindByIDs(context.Context, []int64) ([]*model.Writer, error)
	 FindAll(context.Context, int64, int64) ([]*model.Writer, error)
	 FindByName(
		context.Context, 
		string, int64, int64,
	) ([]*model.Writer, error)
	 Count(context.Context) (int64, error)
}
