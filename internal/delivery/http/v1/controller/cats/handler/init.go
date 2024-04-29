package catHandler

import catsService "cats-social/internal/service/cats"

type catHandler struct {
	catService catsService.CatService
}

func NewHandler(catSvc catsService.CatService) *catHandler {
	return &catHandler{
		catService: catSvc,
	}
}
