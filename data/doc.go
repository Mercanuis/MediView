//Package data contains all files related to the implementation of
//the application data. This is the logic and the data fields that
//the application should work with and manipulate as needed to
//perform the services necessary for the application.
//
//Subpackage model contains the Data Access Objects (DAO) of the
//system, with the relevant data types and their respective fields.
//
//dao.go is an interface that defines the methods that are allowed
//by the data package to access the data layers, and should be used
//by an initialized DAO in the application. By making this an interface
//it allows for multiple implementations and to allow for something more
//than a memory cache implementation (the default)
//
//The data store is based on the use of UUIDs. This is due to the
//uniqueness of UUIDs and the fact that they can be "refreshed"
//given the applications constraints.
package data
