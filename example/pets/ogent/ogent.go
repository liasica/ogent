// Code generated by entc, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"github.com/ariga/ogent/example/pets"
	"github.com/ariga/ogent/example/pets/category"
	"github.com/ariga/ogent/example/pets/pet"
	"github.com/ariga/ogent/example/pets/user"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *pets.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *pets.Client) *OgentHandler { return &OgentHandler{c} }

// CreateCategory handles POST /categories requests.
func (h *OgentHandler) CreateCategory(ctx context.Context, req CreateCategoryReq) (CreateCategoryRes, error) {
	b := h.client.Category.Create()
	// Add all fields.
	b.SetName(req.Name)
	// Add all edges.
	b.AddPetIDs(req.Pets...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Category.Query().Where(category.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		switch {
		case pets.IsNotFound(err):
			return &R400{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewCategoryCreate(e), nil
}

// ReadCategory handles GET /categories/{id} requests.
func (h *OgentHandler) ReadCategory(ctx context.Context, params ReadCategoryParams) (ReadCategoryRes, error) {
	q := h.client.Category.Query().Where(category.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case pets.IsNotFound(err):
			return &R400{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewCategoryRead(e), nil
}

// UpdateCategory handles PATCH /categories/{id} requests.
func (h *OgentHandler) UpdateCategory(ctx context.Context, req UpdateCategoryReq, params UpdateCategoryParams) (UpdateCategoryRes, error) {
	panic("unimplemented")
}

// DeleteCategory handles DELETE /categories/{id} requests.
func (h *OgentHandler) DeleteCategory(ctx context.Context, params DeleteCategoryParams) (DeleteCategoryRes, error) {
	panic("unimplemented")
}

// ListCategory handles GET /categories requests.
func (h *OgentHandler) ListCategory(ctx context.Context, params ListCategoryParams) (ListCategoryRes, error) {
	panic("unimplemented")
}

// CreateCategoryPets handles POST /categories/{id}/pets requests.
func (h *OgentHandler) CreateCategoryPets(ctx context.Context, req CreateCategoryPetsReq, params CreateCategoryPetsParams) (CreateCategoryPetsRes, error) {
	panic("unimplemented")
}

// ListCategoryPets handles GET /categories/{id}/pets requests.
func (h *OgentHandler) ListCategoryPets(ctx context.Context, params ListCategoryPetsParams) (ListCategoryPetsRes, error) {
	panic("unimplemented")
}

// CreatePet handles POST /pets requests.
func (h *OgentHandler) CreatePet(ctx context.Context, req CreatePetReq) (CreatePetRes, error) {
	b := h.client.Pet.Create()
	// Add all fields.
	b.SetName(req.Name)
	if v, ok := req.Weight.Get(); ok {
		b.SetWeight(v)
	}
	if v, ok := req.Birthday.Get(); ok {
		b.SetBirthday(v)
	}
	// Add all edges.
	b.AddCategoryIDs(req.Categories...)
	b.SetOwnerID(req.Owner)
	b.AddFriendIDs(req.Friends...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Pet.Query().Where(pet.ID(e.ID))
	// Eager load edges that are required on create operation.
	q.WithCategories().WithOwner()
	e, err = q.Only(ctx)
	if err != nil {
		switch {
		case pets.IsNotFound(err):
			return &R400{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewPetCreate(e), nil
}

// ReadPet handles GET /pets/{id} requests.
func (h *OgentHandler) ReadPet(ctx context.Context, params ReadPetParams) (ReadPetRes, error) {
	q := h.client.Pet.Query().Where(pet.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case pets.IsNotFound(err):
			return &R400{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewPetRead(e), nil
}

// UpdatePet handles PATCH /pets/{id} requests.
func (h *OgentHandler) UpdatePet(ctx context.Context, req UpdatePetReq, params UpdatePetParams) (UpdatePetRes, error) {
	panic("unimplemented")
}

// DeletePet handles DELETE /pets/{id} requests.
func (h *OgentHandler) DeletePet(ctx context.Context, params DeletePetParams) (DeletePetRes, error) {
	panic("unimplemented")
}

// ListPet handles GET /pets requests.
func (h *OgentHandler) ListPet(ctx context.Context, params ListPetParams) (ListPetRes, error) {
	panic("unimplemented")
}

// CreatePetCategories handles POST /pets/{id}/categories requests.
func (h *OgentHandler) CreatePetCategories(ctx context.Context, req CreatePetCategoriesReq, params CreatePetCategoriesParams) (CreatePetCategoriesRes, error) {
	panic("unimplemented")
}

// ListPetCategories handles GET /pets/{id}/categories requests.
func (h *OgentHandler) ListPetCategories(ctx context.Context, params ListPetCategoriesParams) (ListPetCategoriesRes, error) {
	panic("unimplemented")
}

// ReadPetOwner handles GET /pets/{id}/owner requests.
func (h *OgentHandler) ReadPetOwner(ctx context.Context, params ReadPetOwnerParams) (ReadPetOwnerRes, error) {
	panic("unimplemented")
}

// DeletePetOwner handles DELETE /pets/{id}/owner requests.
func (h *OgentHandler) DeletePetOwner(ctx context.Context, params DeletePetOwnerParams) (DeletePetOwnerRes, error) {
	panic("unimplemented")
}

// CreatePetOwner handles POST /pets/{id}/owner requests.
func (h *OgentHandler) CreatePetOwner(ctx context.Context, req CreatePetOwnerReq, params CreatePetOwnerParams) (CreatePetOwnerRes, error) {
	panic("unimplemented")
}

// CreatePetFriends handles POST /pets/{id}/friends requests.
func (h *OgentHandler) CreatePetFriends(ctx context.Context, req CreatePetFriendsReq, params CreatePetFriendsParams) (CreatePetFriendsRes, error) {
	panic("unimplemented")
}

// ListPetFriends handles GET /pets/{id}/friends requests.
func (h *OgentHandler) ListPetFriends(ctx context.Context, params ListPetFriendsParams) (ListPetFriendsRes, error) {
	panic("unimplemented")
}

// CreateUser handles POST /users requests.
func (h *OgentHandler) CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error) {
	b := h.client.User.Create()
	// Add all fields.
	b.SetName(req.Name)
	b.SetAge(req.Age)
	// Add all edges.
	b.AddPetIDs(req.Pets...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		switch {
		case pets.IsNotFound(err):
			return &R400{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewUserCreate(e), nil
}

// ReadUser handles GET /users/{id} requests.
func (h *OgentHandler) ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case pets.IsNotFound(err):
			return &R400{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: NewOptString(err.Error()),
			}, nil
		case pets.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: NewOptString(err.Error()),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewUserRead(e), nil
}

// UpdateUser handles PATCH /users/{id} requests.
func (h *OgentHandler) UpdateUser(ctx context.Context, req UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error) {
	panic("unimplemented")
}

// DeleteUser handles DELETE /users/{id} requests.
func (h *OgentHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error) {
	panic("unimplemented")
}

// ListUser handles GET /users requests.
func (h *OgentHandler) ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error) {
	panic("unimplemented")
}

// CreateUserPets handles POST /users/{id}/pets requests.
func (h *OgentHandler) CreateUserPets(ctx context.Context, req CreateUserPetsReq, params CreateUserPetsParams) (CreateUserPetsRes, error) {
	panic("unimplemented")
}

// ListUserPets handles GET /users/{id}/pets requests.
func (h *OgentHandler) ListUserPets(ctx context.Context, params ListUserPetsParams) (ListUserPetsRes, error) {
	panic("unimplemented")
}

var _ Handler = (*OgentHandler)(nil)