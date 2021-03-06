package data

import (
	"github.com/MediView/data/model"

	"github.com/google/uuid"
)

//DAO is the Data Access Object(DAO) interface, and is meant to be an
//implementation of the application's data needs. The idea behind putting
//it behind an interface is because it allows for multiple implementations
//to be used to serve the client's needs.
type DAO interface {
	//Gets a Patient from the data store
	//If the key doesn't exist in the data store an error is returned
	//id - the UUID of the Patient
	GetPatient(id uuid.UUID) (*model.Patient, error)

	//Adds a Patient to the data store
	//p - the Patient to add to the data store
	AddPatient(name string, age int) error

	//Gets the list of Patients from the data store
	GetPatients() model.PatientRecords

	//Removes a Patient from the data store
	//id - the UUID of the Patient
	DeletePatient(id uuid.UUID)

	//Adds a Record for the associated Patient
	//If the Patient has an existing record, this record will be
	//stored in aggregation data for the patient's history and then overwritten
	AddRecord(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error)

	//Returned the associated history for the Patients
	GetPatientHistories() model.PatientVitalHistories

	//Resets all Patient history
	//This should be called once every hour (60 minutes)
	ResetPatientHistory()

	//Removes all PatientVitalsHistory from the data store
	//This should be called once a day (24 hours)
	DeleteAllHistory()
}
