
CREATE TABLE Patients (
	ID serial PRIMARY KEY,
  	FirstName varchar(255),
    MiddleName varchar(255),
    LastName varchar(255),
  	Email varchar(255),
  	Phone varchar(255),
  	DateOfBirst varchar(255)
);

CREATE TABLE ConsultationRequest (
	ID serial PRIMARY KEY,
  	PatientID integer,
  	Text text,
  	CreatedAt timestamp,
    CONSTRAINT fk_patients FOREIGN KEY(PatientID) REFERENCES Patients(ID)
);

CREATE TABLE Recommendations (
	ID serial PRIMARY KEY,
  	ConsultationRequestID integer,
  	Text text,
    CONSTRAINT fk_consultation_request_id FOREIGN KEY(ConsultationRequestID) REFERENCES ConsultationRequest(ID)
);
