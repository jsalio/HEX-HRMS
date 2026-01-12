package models

type PermissionView string

type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

type CreateRole struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

type RoleItem struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Permissions []PermissionView `json:"permissions"`
}

func (r *Role) ToRoleItem() *RoleItem {
	var rolesList []PermissionView
	for _, permission := range r.Permissions {
		rolesList = append(rolesList, PermissionView(permission.Name))
	}
	return &RoleItem{
		ID:          r.ID,
		Name:        r.Name,
		Permissions: rolesList,
	}
}
