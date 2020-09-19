package repository

import (
	"database/sql"
	"fmt"
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
	"time"
)

type MySQLTreatmentRepo struct {
	db *sql.DB
}

func NewMySQLTreatmentRepository(db *sql.DB) *MySQLTreatmentRepo {
	return &MySQLTreatmentRepo{
		db: db,
	}
}

func (r *MySQLTreatmentRepo) Create(e *model.Treatment) (entity2.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into treatment (id, name, created_at, created_by) 
		values(?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Name,
		time.Now().Format("2006-01-02"),
		e.CreatedBy,
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

func (r *MySQLTreatmentRepo) Get(id entity2.ID) (*model.Treatment, error) {
	stmt, err := r.db.Prepare(`select id, name, active from treatment where id = ?`)
	if err != nil {
		return nil, err
	}
	var u model.Treatment
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.Active)
	}

	if u.Name != "" {
		return &u, nil
	} else {
		return nil, exception.ErrNotFound
	}
}

func (r *MySQLTreatmentRepo) GetByName(name string) (*model.Treatment, error) {
	stmt, err := r.db.Prepare(`select id, name from treatment where name = ?`)
	if err != nil {
		return nil, err
	}

	var u model.Treatment

	rows, err := stmt.Query(name)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name)
	}

	return &u, nil
}

func (r *MySQLTreatmentRepo) Update(e *model.Treatment) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update treatment set name = ?, updated_at = ?, updated_by = ? where id = ?", e.Name, e.UpdatedAt.Format("2006-01-02"),e.UpdatedBy, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLTreatmentRepo) Search(e *model.Treatment) ([]*model.Treatment, error) {
	sql := "select id, name, active from treatment where "
	sql += getExpresion("name", e.Name)

	sql = sql[0:len(sql) - 3]

	fmt.Println(sql)

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var patients []*model.Treatment
	for rows.Next() {
		u := new(model.Treatment)

		err = rows.Scan(&u.ID, &u.Name, &u.Active)
		if err != nil {
			return nil, err
		}

		patients = append(patients, u)
	}
	return patients, nil
}

func (r *MySQLTreatmentRepo) Delete(id entity2.ID) error {
	_, err := r.db.Exec("delete from treatment where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
