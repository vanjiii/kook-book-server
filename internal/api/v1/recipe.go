package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"vanjiii/kook-book-server/internal/httpx"
	"vanjiii/kook-book-server/internal/recipe"

	"github.com/go-chi/chi/v5"
)

type RecipeHandler struct {
	service *recipe.Service
}

func NewRecipe(s *recipe.Service) *RecipeHandler {
	return &RecipeHandler{
		service: s,
	}
}

type RecipeResponse struct {
	ID          uint   `json:"id"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

func (h *RecipeHandler) Get(w http.ResponseWriter, r *http.Request) {
	idstr := chi.RouteContext(r.Context()).URLParam("recipeID")
	if idstr == "" {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(httpx.APIResponse{
			Code:  http.StatusBadRequest,
			Error: "missing id",
		})

		return
	}

	recipeID, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(httpx.APIResponse{
			Code:  http.StatusBadRequest,
			Error: "bad id",
		})

		return
	}

	recipe, err := h.service.GetByID(uint(recipeID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(httpx.APIResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})

		return
	}

	// TODO: handle error
	if err := json.NewEncoder(w).Encode(httpx.APIResponse{
		Code: http.StatusOK,
		Data: RecipeResponse{
			ID:          recipe.ID,
			Photo:       recipe.Photo,
			Description: recipe.Description,
		},
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(httpx.APIResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})

		return
	}
}

type CreateRequest struct {
	ID          uint   `json:"id"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO validator

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(httpx.APIResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})

		return
	}

	var ent *CreateRequest

	if err := json.Unmarshal(body, &ent); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(httpx.APIResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})

		return
	}

	if err := h.service.Insert(recipe.Recipe{
		ID:          ent.ID,
		Description: ent.Description,
		Photo:       ent.Photo,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(httpx.APIResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
}
