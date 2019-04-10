package data

import (
	"MediView/data/model"

	"github.com/google/uuid"
)

//The Data Access Object(DAO) interface is meant to be an implementation of
//the application's data needs. The idea behind putting it behind an interface
//is because it allows for multiple implementations to be used to serve the
//client's needs.
type DAO interface {
	//Gets a Patient from the data store
	//If the key doesn't exist in the data store an error is returned
	//id - the UUID of the Patient
	GetPatient(id uuid.UUID) (*model.Patient, error)

	//Adds a Patient to the data store
	//p - the Patient to add to the data store
	AddPatient(name string, age int) (uuid.UUID, error)

	//Gets the list of Patients from the data store
	GetPatients() model.PatientRecords

	//Removes a Patient from the data store
	//id - the UUID of the Patient
	DeletePatient(id uuid.UUID)

	AddRecord(pid uuid.UUID, vitals model.Vitals) (*model.Patient, error)
}
