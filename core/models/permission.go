package models

type Permission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	RoleId      string `json:"role_id"`
	Role        Role   `json:"role"`
}

const (
	PermissionViewMenuDashboard     = "view_menu_dashboard"
	PermissionViewMenuEmployees     = "view_menu_employees"
	PermissionEditEmployees         = "edit_employees"
	PermissionViewEmployees         = "view_employees"
	PermissionViewMenuDepartments   = "view_menu_departments"
	PermissionViewMenuPosition      = "view_menu_position"
	PermissionViewMenuAttendance    = "view_menu_attendance"
	PermissionViewMenuPayroll       = "view_menu_payroll"
	PermissionViewMenuLeaveRequests = "view_menu_leave_requests"
	PermissionViewMenuSettings      = "view_menu_settings"
	PermissionAllAccess             = "all_access"
	PermissionViewRoles             = "view_roles"
	PermissionEditRoles             = "edit_roles"
	PermissionEditUsers             = "edit_users"
	PermissionViewUsers             = "view_users"
)
