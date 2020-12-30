package persistence

import (
	"database/sql"
	"fmt"
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	entity2 "github.com/jahs/clinic-backend/src/shared/entity"
	"github.com/jahs/clinic-backend/src/shared/util"
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

func (r *MySQLPatientRepo) Create(e *domain.Patient) (entity2.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into patient (id, avatar_path, name, address, email, phone, created_at, created_by) 
		values(?,?,?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.AvatarPath,
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

func (r *MySQLPatientRepo) Get(id entity2.ID) (*domain.Patient, error) {
	stmt, err := r.db.Prepare(`select id, avatar_path, name, address, email, phone, active from patient where id = ?`)
	if err != nil {
		return nil, err
	}
	var u domain.Patient
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.AvatarPath, &u.Name, &u.Address, &u.Email, &u.Phone, &u.Active)
	}

	if u.Email != "" {
		return &u, nil
	} else {
		return nil, exception.ErrNotFound
	}
}

func (r *MySQLPatientRepo) GetByEmail(email string) (*domain.Patient, error) {
	stmt, err := r.db.Prepare(`select id, email from patient where email = ?`)
	if err != nil {
		return nil, err
	}

	var u domain.Patient

	rows, err := stmt.Query(email)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Email)
	}

	return &u, nil
}

func (r *MySQLPatientRepo) Update(e *domain.Patient) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update patient set name = ?, address = ?, email = ?, phone = ?, updated_at = ?, updated_by = ? where id = ?", e.Name, e.Address, e.Email, e.Phone, e.UpdatedAt.Format("2006-01-02"),e.UpdatedBy, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLPatientRepo) Search(e *domain.Patient) ([]*domain.Patient, error) {
	sql := "select id, avatar_path, name, address, email, phone, active from patient where "
	sql += util.GetExpresion("name", e.Name)
	sql += util.GetExpresion("email", e.Email)
	sql += util.GetExpresion("address", e.Address)

	sql = sql[0:len(sql) - 3]

	fmt.Println(sql)

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var patients []*domain.Patient

	for rows.Next() {
		u := new(domain.Patient)

		err = rows.Scan(&u.ID, &u.AvatarPath, &u.Name, &u.Address, &u.Email, &u.Phone, &u.Active)
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
