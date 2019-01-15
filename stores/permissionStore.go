package stores

import (
	"database/sql"
	"github.com/avi1989/Identity/models"
)

type PermissionStore interface {
	GetPermissions(roleCode string) ([]models.Permission, error)
}

func (store *Store) GetPermissions(roleCode string) ([]models.Permission, error) {
	var permissions []models.Permission
	db := store.getDb()
	rows, err := db.Query(
		`SELECT p.* FROM identity.role_permission rp 
					INNER JOIN identity.permission p on rp.permission_code = p.code
				WHERE rp.role_code = $1`, roleCode)
	if err != nil {
		return permissions, err
	}

	for rows.Next() {
		var permission models.Permission
		err := rows.Scan(&permission.Code, &permission.Description)
		if err != nil {
			return permissions, err
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}

func NewPermissionStore(getDb func() *sql.DB) PermissionStore {
	return &Store{getDb: getDb}
}
