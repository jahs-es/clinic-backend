ALTER TABLE clinic.patient_treatment ADD CONSTRAINT patient_treatment_treatment_FK FOREIGN KEY (treatment_id) REFERENCES clinic.treatment(id);
