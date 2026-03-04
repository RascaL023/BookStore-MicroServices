package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"writer/internal/dto"
	"writer/internal/model"
	"writer/internal/service"
)

type WriterController struct { service *service.WriterService }

func New(s *service.WriterService) *WriterController {
	return &WriterController{service: s}
}


func (c *WriterController) GetByID(
	w http.ResponseWriter, 
	r *http.Request,
) {
	id := resolveID(r, w)
    writer, err := c.service.GetByID(r.Context(), id)

    if err != nil {
		fmt.Println("Server err")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } else if writer == nil {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }

    response, _ := json.Marshal(dto.ToResponse(writer))
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func (c *WriterController) GetAll(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
) {
	page, size := resolvePage(httpRequest)
	name := httpRequest.URL.Query().Get("name")

	var err error
	var writers []*model.Writer
	var meta *dto.Meta
	if name == "" {
		writers, meta, err = c.service.GetAll(
			httpRequest.Context(),
			page, size,
		)
	} else {
		writers, meta, err = c.service.GetByName(
			httpRequest.Context(),
			name, page, size,
		)
	}
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.NewPagedResponse(dto.ToResponses(writers), meta)
	responses, _ := json.Marshal(response)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(responses)
}


func (c *WriterController) Upsert(
	res http.ResponseWriter,
	req *http.Request,
) {
	var body dto.WriterRequest
	id := resolveID(req, res)

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(res, "invalid JSON Body", http.StatusBadRequest)
		return
	}

	writer, err := c.service.Upsert(req.Context(), dto.ToEntity(&body), id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(dto.ToResponse(writer))
}


func (c *WriterController) PatchUpdate(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body dto.WriterPatchRequest
	id := resolveID(r, w)

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON Body", http.StatusBadRequest)
		return
	}


	writer, err := c.service.PatchUpdate(r.Context(), &body, id)
	if err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ToResponse(writer))
}
