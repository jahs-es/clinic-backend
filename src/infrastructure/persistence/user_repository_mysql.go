package persistence

import (
	"database/sql"
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	entity2 "github.com/jahs/clinic-backend/src/shared/entity"
	"github.com/jahs/clinic-backend/src/shared/util"
	"time"
)

type MySQLUserRepo struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepo {
	return &MySQLUserRepo{
		db: db,
	}
}

func (r *MySQLUserRepo) Create(e *domain.User) (entity2.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into user (id, email, password, name, rol, created_at) 
		values(?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Email,
		e.Password,
		e.Name,
		e.Rol,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

func (r *MySQLUserRepo) Get(id entity2.ID) (*domain.User, error) {
	stmt, err := r.db.Prepare(`select id, email, name, rol, active from user where id = ?`)
	if err != nil {
		return nil, err
	}
	var u domain.User
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Email, &u.Name, &u.Rol, &u.Active)
	}

	if u.Email != "" {
		return &u, nil
	} else {
		return nil, exception.ErrNotFound
	}
}

func (r *MySQLUserRepo) GetByEmail(email string) (*domain.User, error) {
	stmt, err := r.db.Prepare(`select id, password, email from user where email = ?`)
	if err != nil {
		return nil, err
	}

	var u domain.User

	rows, err := stmt.Query(email)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Password, &u.Email)
	}

	return &u, nil
}

func (r *MySQLUserRepo) Update(e *domain.User) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update user set email = ?, password = ?, name = ?, rol = ?, updated_at = ? where id = ?", e.Email, e.Password, e.Name, e.Rol, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepo) Search(e *domain.User) ([]*domain.User, error) {
	sql := "select id, email, name, rol, active from user where "
	sql += util.GetExpresion("name", e.Name)
	sql += util.GetExpresion("email", e.Email)
	sql += util.GetExpresion("rol", e.Rol)

	sql = sql[0:len(sql) - 3]

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for rows.Next() {
		user := new(domain.User)

		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Rol, &user.Active)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (r *MySQLUserRepo) Delete(id entity2.ID) error {
	_, err := r.db.Exec("delete from user where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
