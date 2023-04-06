package persistence

import "github.com/go-jet/jet/v2/qrm"

type Repository interface {
	GetById(id int64) (*interface{}, error)
	Insert(m interface{}, db qrm.DB, attributes interface{}) (*int64, error)
	InsertBulk(models interface{}, db qrm.DB) (interface{}, error)
}

// TODO: think through this possibly move all repo stuff into a generic interface or something like that?
// func GetRepository(repo Repository) Repository {
// 	switch repo.(type) {
// 	case ExerciseRepository:
// 		return GetSetRepository()
// 	}
// }
