package repository

import (
	"database/sql"
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
	"time"
)

type MySQLPatientRepo struct {
	db *sql.DB
}

func NewMySQLPatientRepository(db *sql.DB) *MySQLPatientRepo {
	return &MySQLPatientRepo{
		db: db,
	}
}

func (r *MySQLPatientRepo) Create(e *model.Patient) (entity2.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into patient (id, name, address, email, phone, created_at, created_by) 
		values(?,?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Name,
		e.Address,
		e.Email,
		e.Phone,
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

func (r *MySQLPatientRepo) Get(id entity2.ID) (*model.Patient, error) {
	stmt, err := r.db.Prepare(`select id, name, address, email, phone, active from patient where id = ?`)
	if err != nil {
		return nil, err
	}
	var u model.Patient
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.Address, &u.Email, &u.Phone, &u.Active)
	}

	if u.Email != "" {
		return &u, nil
	} else {
		return nil, exception.ErrNotFound
	}
}

func (r *MySQLPatientRepo) GetByEmail(email string) (*model.Patient, error) {
	stmt, err := r.db.Prepare(`select id, email from patient where email = ?`)
	if err != nil {
		return nil, err
	}

	var u model.Patient

	rows, err := stmt.Query(email)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Email)
	}

	return &u, nil
}

func (r *MySQLPatientRepo) Update(e *model.Patient) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update patient set name = ?, address = ?, email = ?, phone = ?, updated_at = ?, updated_by = ? where id = ?", e.Name, e.Address, e.Email, e.Phone, e.UpdatedAt.Format("2006-01-02"),e.UpdatedBy, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLPatientRepo) Search(e *model.Patient) ([]*model.Patient, error) {
	sql := "select id, name, address, email, phone, active from patient where "
	sql += getExpresion("name", e.Name)
	sql += getExpresion("email", e.Email)
	sql += getExpresion("address", e.Address)

	sql = sql[0:len(sql) - 3]

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var patients []*model.Patient

	for rows.Next() {
		u := new(model.Patient)

		err = rows.Scan(&u.ID, &u.Name, &u.Address, &u.Email, &u.Phone, &u.Active)
		if err != nil {
			return nil, err
		}

		patients = append(patients, u)
	}

	return patients, nil
}

func (r *MySQLPatientRepo) Delete(id entity2.ID) error {
	_, err := r.db.Exec("delete from patient where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
