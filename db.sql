

CREATE TABLE counties(
    county_id INT GENERATED ALWAYS AS IDENTITY,
    county_name  VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (county_id)
);
CREATE TABLE patients(
    patient_id BIGINT GENERATED ALWAYS AS IDENTITY,
    county_id INT NOT NULL,
    first_name  VARCHAR(64) NOT NULL,
    surname  VARCHAR(64),
    other_name  VARCHAR(64),
    -- patient_number is to be generated accroding to organization numbering --
    patient_number  VARCHAR(64) UNIQUE NOT NULL,
    mobile_number  VARCHAR(12)  UNIQUE NOT NULL,
    email_address  VARCHAR(128) UNIQUE, -- some patients might not have email_address --
    is_disabled BOOLEAN DEFAULT false,
    gender  VARCHAR(1)  check (gender = 'F' or gender = 'M') NOT NULL ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (patient_id),
    CONSTRAINT fk_county FOREIGN KEY(county_id) REFERENCES counties(county_id)
);
-- ct_persons is  Contact persons table for storing alternative contact person for patient --
CREATE TABLE contact_persons(
    contact_person_id INT GENERATED ALWAYS AS IDENTITY,
    patient_id BIGINT UNIQUE NOT NULL,
    person_name  VARCHAR(64) NOT NULL,
    tel_number  VARCHAR(12)  UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (contact_person_id),
    CONSTRAINT fk_patient FOREIGN KEY(patient_id) REFERENCES patients(patient_id)
);
CREATE TABLE appointments(
    appointment_id BIGINT GENERATED ALWAYS AS IDENTITY,
    patient_id BIGINT NOT NULL,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (appointment_id),
    CONSTRAINT fk_patient FOREIGN KEY(patient_id) REFERENCES patients(patient_id)
);