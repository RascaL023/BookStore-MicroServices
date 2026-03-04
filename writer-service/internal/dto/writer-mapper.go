package dto

import "writer/internal/model"

func ToEntity(request *WriterRequest) *model.Writer {
	return &model.Writer{
		Name: request.Name,
		City: request.City,
		Email: request.Email,
		IsActive: request.IsActive,
	}
}

func ToRequest(writer *model.Writer) *WriterRequest {
	return &WriterRequest{
		Name: writer.Name,
		City: writer.City,
		Email: writer.Email,
		IsActive: writer.IsActive,
	}
}

func ToResponse(writer *model.Writer) Response[WriterResponse] {
	return Response[WriterResponse]{
		Data: WriterResponse{
			Id: writer.Id,
			Name: writer.Name,
			Email: writer.Email,
			City: writer.City,
		},
	}
}

func ToResponses(list []*model.Writer) Response[[]WriterResponse] {
	responses := make([]WriterResponse, 0, len(list))
	for _, w := range list {
		responses = append(responses, WriterResponse{
			Id: w.Id,
			Name: w.Name,
			Email: w.Email,
			City: w.City,
		})
	}

	return Response[[]WriterResponse]{Data: responses}
}

func NewPagedResponse[T any](data T, meta *Meta) Response[T] {
    return Response[T]{Data: data, Meta: meta}
}
