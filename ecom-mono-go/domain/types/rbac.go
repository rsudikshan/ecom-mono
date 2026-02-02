package types

type Permission string
// Easier to check permission like this then []Permissions
type PermissionSet map[Permission]struct{}
type Role string 

const (
	ROLE_SUPER_ADMIN Role = "SUPER_ADMIN"
	ROLE_ADMIN Role = "ROLE_ADMIN"
	ROLE_USER Role = "ROLE_USER"
	ROLE_VENDOR Role = "ROLE_VENDOR"
)

const (
	PERMISSION_ALL Permission = "ALL"

	PERMISSION_UPDATE_USER Permission = "UPDATE_USER"
	PERMISSION_DELETE_USER Permission = "DELETE_USER"

	PERMISSION_ADD_PRODUCT Permission = "ADD_PRODUCT"
	PERMISSION_DELETE_PRODUCT Permission = "DELETE_PRODUCT"
	PERMISSION_UPDATE_PRODUCT Permission = "UPDATE_PRODUCT"
)

var predefinedRolesAndPermissions = map[Role]PermissionSet{
	ROLE_SUPER_ADMIN : {
		PERMISSION_ALL:{},
	},
	ROLE_ADMIN : {
		PERMISSION_DELETE_USER:{},
		PERMISSION_UPDATE_USER:{},
	},
	ROLE_VENDOR : {
		PERMISSION_ADD_PRODUCT:{},
		PERMISSION_DELETE_PRODUCT:{},
		PERMISSION_UPDATE_PRODUCT:{},
	},
}

func HasPermission(role Role, p Permission) bool {

	// Role check and fetch permissions if role exists
	perms,ok := predefinedRolesAndPermissions[role]
	if !ok {
		return false
	}

	// Permission check in fetched role for super admin
	if _,ok=perms[PERMISSION_ALL]; ok {
		return true
	}

	// Permission check in a role
	_,ok = perms[p]
	return ok
}
