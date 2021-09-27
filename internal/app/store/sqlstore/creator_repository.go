package sqlstore

import (
	"database/sql"
	"patreon/internal/app/store"
	"patreon/internal/models"

	log "github.com/sirupsen/logrus"
)

type CreatorRepository struct {
	store          *Store
	UserRepository UserRepository
}

func NewCreatorRepository(st *Store) *CreatorRepository {
	return &CreatorRepository{
		store:          st,
		UserRepository: UserRepository{st},
	}
}

func (repo *CreatorRepository) Create(cr *models.Creator) error {
	if err := repo.store.db.QueryRow("INSERT INTO creator_profile (creator_id, category, description, avatar, cover) VALUES ($1, $2, $3, $4, $5)"+
		"RETURNING creator_id", cr.ID, cr.Category, cr.Description, cr.Avatar, cr.Cover).Scan(&cr.ID); err != nil {
		return err
	}
	return nil
}
func (repo *CreatorRepository) GetCreators() ([]models.ResponseCreator, error) {
	count := 0
	if err := repo.store.db.QueryRow("SELECT count(*) from creator_profile").Scan(&count); err != nil {
		return nil, store.InternalError
	}
	res := make([]models.ResponseCreator, count)
	rows, err := repo.store.db.Query("SELECT * from creator_profile")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	i := 0
	for empty := rows.Next(); !empty; i++ {
		var creator models.ResponseCreator
		if err = rows.Scan(&creator.ID, &creator.Category, &creator.Description,
			&creator.Avatar, &creator.Cover); err != nil {
			return nil, err
		}
		var user *models.User
		if user, err = repo.UserRepository.FindByID(int64(creator.ID)); err != nil {
			return nil, err
		}
		creator.Nickname = user.Nickname
		res[i] = creator

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return res, err

}

//func (repo *CreatorRepository) FindByLogin(login string) (*models.User, error) {
//	user := models.User{}
//
//	if err := repo.store.db.QueryRow("SELECT user_id, login, encrypted_password from users where login=$1", login).
//		Scan(&user.ID, &user.Login, &user.EncryptedPassword); err != nil {
//		return nil, store.NotFound
//	}
//
//	return &user, nil
//}

//func (repo *CreatorRepository) FindByID(id int64) (*models.User, error) {
//	user := models.User{}
//
//	if err := repo.store.db.QueryRow("SELECT user_id, nickname, avatar from users where user_id=$1", id).
//		Scan(&user.ID, &user.Nickname, &user.Avatar); err != nil {
//		return nil, store.NotFound
//	}
//
//	return &user, nil
//}