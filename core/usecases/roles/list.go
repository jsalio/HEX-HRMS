package roles

import (
	"hrms.local/core/contracts"
	"hrms.local/core/models"
)

type ListRolesUsecase struct {
	repo    contracts.RoleContract
	request *contracts.GenericRequest[models.SearchQuery]
}

func NewListRolesUsecase(repo contracts.RoleContract, request *contracts.GenericRequest[models.SearchQuery]) *ListRolesUsecase {
	return &ListRolesUsecase{repo: repo, request: request}
}

func (u *ListRolesUsecase) Validate() *models.SystemError {
	// No specific validation needed for list
	return nil
}

func (u *ListRolesUsecase) Execute() (*models.PaginatedResponse[models.Role], *models.SystemError) {
	query := u.request.Build()
	return u.repo.GetByFilter(query)
}
