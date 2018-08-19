package admin

import "gitlab.com/computingfun/computingfun.org/admin/permission"

// Store TODO: Comment
type Store interface {
	Close() error

	Add(user, pass string) error
	Remove(user string) error

	CheckPass(user, pass string) error
	ChangePass(user, pass string) error

	GetName(user string) (string, error)
	SetName(user, name string) error

	GetPermissionMap(user string) (permission.Map, error)
	SetPermissionMap(user string, permissions permission.Map) error

	GetPermission(user string, permission permission.Type) (permission.Level, error)
	SetPermission(user string, permission permission.Type, level permission.Level) error
	CheckPermission(user string, permission permission.Type, level permission.Level) error
}
