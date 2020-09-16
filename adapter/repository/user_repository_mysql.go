package repository

import (
	"database/sql"
	"fmt"
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
	"time"
)

//MySQLUserRepo mysql repo
type MySQLUserRepo struct {
	db *sql.DB
}

//NewMySQLUserRepository create new repository
func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepo {
	return &MySQLUserRepo{
		db: db,
	}
}

//Create an user
func (r *MySQLUserRepo) Create(e *model.User) (entity2.ID, error) {
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

//Get an user
func (r *MySQLUserRepo) Get(id entity2.ID) (*model.User, error) {
	stmt, err := r.db.Prepare(`select id, email, name, rol, active from user where id = ?`)
	if err != nil {
		return nil, err
	}
	var u model.User
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

//Get an user by email
func (r *MySQLUserRepo) GetByEmail(email string) (*model.User, error) {
	stmt, err := r.db.Prepare(`select id, password, email from user where email = ?`)
	if err != nil {
		return nil, err
	}

	var u model.User

	rows, err := stmt.Query(email)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Password, &u.Email)
	}

	return &u, nil
}

func (r *MySQLUserRepo) Update(e *model.User) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update user set email = ?, password = ?, name = ?, rol = ?, updated_at = ? where id = ?", e.Email, e.Password, e.Name, e.Rol, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

func getExpresion(column string, value string) string {
	if column != "" {
		return column + " like '%" + value + "%' OR "
	} else {
		return ""
	}
}

//Search users
func (r *MySQLUserRepo) Search(e *model.User) ([]*model.User, error) {
	sql := "select id, email, name, rol, active from user where "
	sql += getExpresion("name", e.Name)
	sql += getExpresion("email", e.Email)
	sql += getExpresion("rol", e.Rol)

	sql = sql[0:len(sql) - 3]

	fmt.Println(sql)

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for rows.Next() {
		user := new(model.User)

		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Rol, &user.Active)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

//Delete an user
func (r *MySQLUserRepo) Delete(id entity2.ID) error {
	_, err := r.db.Exec("delete from user where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
