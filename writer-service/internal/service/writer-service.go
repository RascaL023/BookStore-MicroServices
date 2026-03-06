package service

import (
	"context"
	"fmt"
	"writer/internal/cache"
	"writer/internal/dto"
	"writer/internal/model"
	"writer/internal/repository"
)

type WriterService struct { repo  repository.WriterRepository }

func New(repo repository.WriterRepository) *WriterService {
	return &WriterService{repo: repo}
}



func (s *WriterService) PatchUpdate(
	ctx context.Context,
	w *dto.WriterPatchRequest,
	id int64,
) (*model.Writer, error) {
	writer, err := s.GetByID(ctx, id)
	if err != nil { return nil, err }

	if w.Name != nil { writer.Name = *w.Name }
	if w.Email != nil { writer.Email = *w.Email }
	if w.City != nil { writer.City = *w.City }
	if w.IsActive != nil { writer.IsActive = *w.IsActive }

	writer, err = s.repo.Save(ctx, *writer)
	if err != nil { return nil, err }

	cache.SetJSON(
		ctx, fmt.Sprintf("writer:%d", writer.Id), 
		writer, 3,
	)
	return writer, nil
}


func (s *WriterService) Upsert(
	ctx context.Context,
	w *model.Writer,
	id int64,
) (*model.Writer, error) {
	if id != 0 { w.Id = id }

	writer, err := s.repo.Save(ctx, *w)
	if err != nil { return nil, err }

	cache.SetJSON(
		ctx, fmt.Sprintf("writer:%d", id), 
		writer, 3,
	)
	return writer, nil
}

func (s *WriterService) GetByID(
	ctx context.Context,
	id int64,
) (*model.Writer, error) {
	var writer *model.Writer

	cacheKey := fmt.Sprintf("writer:%d", id)
	writer, err := cache.GetJSON[*model.Writer](ctx, cacheKey)
	if err == nil { return writer, nil }
	
	fmt.Println("Cache miss => " + cacheKey)
	writer, err = s.repo.FindByID(ctx, id)
	if err != nil { return nil, err }

	cache.SetJSON(ctx, cacheKey, writer, 3)
	return writer, nil
}


func (s *WriterService) GetByIDs(
	ctx context.Context,
	ids []int64,
) ([]*model.Writer, error) {
	return s.repo.FindByIDs(ctx, ids)
}


func (s *WriterService) GetAll(
	ctx context.Context,
	page, size int64,
) ([]*model.Writer, *dto.Meta, error) {
	var total int64
	offset := (page - 1) * size

	writers, err := s.repo.FindAll(ctx, offset, size)
	if err != nil { return nil, nil, err }

	total, err = cache.GetInt(ctx, "writer:count")
	if err != nil { 
		total, err = s.repo.Count(ctx)
		fmt.Println("Cache miss => writer:count")
		if err != nil { return nil, nil, err }
	}

	meta := &dto.Meta{
		Page: 		int(page),
		Size: 		int(size),
		Total:		int(total),
		Current:	len(writers),
	}

	cache.SetWriterJSONs(ctx, "writer:%d", writers, 3)
	return writers, meta, nil
}


func (s *WriterService) GetByName(
	ctx context.Context,
	name string,
	page, size int64,
) ([]*model.Writer, *dto.Meta, error) {
	offset := (page - 1) * size

	writers, err := s.repo.FindByName(
		ctx, name, offset, size,
	)
	if err != nil { return nil, nil, err }

	total, err := s.repo.Count(ctx)
	if err != nil { return nil, nil, err }

	meta := &dto.Meta{
		Page: 		int(page),
		Size: 		int(size),
		Total:		int(total),
		Current:	len(writers),
	}

	return writers, meta, nil
}
