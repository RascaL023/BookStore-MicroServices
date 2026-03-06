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

func writeResponse(w http.ResponseWriter, response []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}


func (c *WriterController) GetByID(
	w http.ResponseWriter, 
	r *http.Request,
) {
	id := resolveID(r, w)
    writer, err := c.service.GetByID(r.Context(), id)

    if err != nil {
		fmt.Println("Server err")
		dto.ServerError(w, nil)
        return
    } else if writer == nil {
		dto.WriterNotFoundError(w, nil)
        return
    }

    response, _ := json.Marshal(dto.ToResponse(writer))
	writeResponse(w, response)
}


func (c *WriterController) GetByIDs(
	w http.ResponseWriter,
	r *http.Request,
) {
	ids := resolveIDs(r, w)
	writers, err := c.service.GetByIDs(r.Context(), ids)
	if err != nil {
		fmt.Println("Server err")
		dto.ServerError(w, nil)
        return
	} else if writers == nil {
		fmt.Println("Not found")
		dto.WriterNotFoundError(w, nil)
        return
	}

	responses, _ := json.Marshal(dto.ToResponses(writers))
	writeResponse(w, responses)
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
	if name != "" {
		writers, meta, err = c.service.GetByName(
			httpRequest.Context(),
			name, page, size,
		)
	} else {
		ids := httpRequest.URL.Query().Get("ids")
		if ids != "" {
			c.GetByIDs(responseWriter, httpRequest)
			return
		} else {
			writers, meta, err = c.service.GetAll(
				httpRequest.Context(),
				page, size,
			)
		}
	}

	if err != nil {
		dto.ServerError(responseWriter, nil)
		return
	}

	response := dto.NewPagedResponse(dto.ToResponses(writers), meta)
	responses, _ := json.Marshal(response)
	writeResponse(responseWriter, responses)
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


func (c *WriterController) CheckHealth(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
