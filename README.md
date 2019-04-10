# MediView
A small back end implementation of a in-memory message queue that stores medical records

## Packages

- **cmd/main** - Logic to start up basic HTTP listener, helps to access HTTP handlers which call services
- **data** - Data layer for the application, contains logic for DB models and associated operations to access data
- **di** - Short for Dependency Injection, contains logic for setup of the service via main
- **service** - Basic business logic that handles calls to data layer
- **http** - HTTP services, including handlers that hook into service layer
