ALTER TABLE clinic.patient_treatment ADD CONSTRAINT patient_treatment_patient_FK FOREIGN KEY (patient_id) REFERENCES clinic.patient(id);
