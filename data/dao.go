package data

//The Data Access Object(DAO) interface is meant to be an implementation of
//the application's data needs. The idea behind putting it behind an interface
//is because it allows for multiple implementations to be used to serve the
//client's needs.
//
//Attempts to be as generic as possible, and allow for multiple types
//to implement the interface and not be worried about type
type DAO interface {
	//Adds a record to the data store
	//id - the unique key that is associated with the interface
	//data - the data object to add to the data store
	//TODO: look into using the UUID package instead
	Add(id int64, data interface{}) error

	//Removes a record from the data store
	//id - the unique key that is associated with the interface
	//TODO: Look into using the UUID package instead
	Delete(id int64) error
}
