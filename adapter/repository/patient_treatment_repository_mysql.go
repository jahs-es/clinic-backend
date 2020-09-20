package repository

import (
	"database/sql"
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
	"time"
)

type MySQLPatientTreatmentRepo struct {
	db *sql.DB
}

func NewMySQLPatientTreatmentRepository(db *sql.DB) *MySQLPatientTreatmentRepo {
	return &MySQLPatientTreatmentRepo{
		db: db,
	}
}

func (r *MySQLPatientTreatmentRepo) Create(e *model.PatientTreatment) (entity2.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into patient_treatment (id, patient_id, treatment_id, detail, created_at, created_by) 
		values(?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.PatientId,
		e.TreatmentId,
		e.Detail,
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

func (r *MySQLPatientTreatmentRepo) Get(id entity2.ID) (*model.PatientTreatment, error) {
	stmt, err := r.db.Prepare(`select id, patient_id, treatment_id, detail, active from patient_treatment where id = ?`)
	if err != nil {
		return nil, err
	}
	var u model.PatientTreatment
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.PatientId, &u.TreatmentId, &u.Detail, &u.Active)
	}

	if u.Detail != "" {
		return &u, nil
	} else {
		return nil, exception.ErrNotFound
	}
}

func (r *MySQLPatientTreatmentRepo) Update(e *model.PatientTreatment) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update patient_treatment set detail = ?, updated_at = ?, updated_by = ? where id = ?", e.Detail, e.UpdatedAt.Format("2006-01-02"),e.UpdatedBy, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLPatientTreatmentRepo) Search(e *presenter.PatientTreatmentDTO) ([]*presenter.PatientTreatmentDTO, error) {
	sql := "select pt.id, patient_id, treatment_id, pt.detail, pt.active, p.name, t.name "
	sql += "from patient_treatment pt inner join patient p on pt.patient_id = p.id "
	sql += "inner join treatment t on pt.treatment_id = t.id where "
	sql += getExpresion("detail", e.Detail)
	sql += getExpresion("p.name", e.Patient)
	sql += getExpresion("p.name", e.Treatment)

	sql = sql[0:len(sql) - 3]

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	patient_treatments := make([]*presenter.PatientTreatmentDTO, 0)

	for rows.Next() {
		u := new(presenter.PatientTreatmentDTO)

		err = rows.Scan(&u.ID, &u.PatientId, &u.TreatmentId, &u.Detail, &u.Active, &u.Patient, &u.Treatment)
		if err != nil {
			return nil, err
		}

		patient_treatments = append(patient_treatments, u)
	}

	return patient_treatments, nil
}

func (r *MySQLPatientTreatmentRepo) Delete(id entity2.ID) error {
	_, err := r.db.Exec("delete from patient_treatment where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
