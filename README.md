# MediView
A small back end implementation of a in-memory message queue that stores medical records. 
This application was written in mind for a small to mid-sized hospital with 20-750 patients. 
This is an application that was visualized first as a standalone HTTP server, but due to the 
usage of a message queue could be theoretically split into a front end and backend listening for
messages. 

## Packages

- **cmd/main** - Logic to start up basic HTTP listener, helps to access HTTP handlers which call services
- **data** - Data layer for the application, contains logic for DB models and associated operations to access data
- **di** - Short for Dependency Injection, contains logic for setup of the service via main
- **http** - HTTP services, including handlers that hook into service layer
- **queue** - Queue services, holding logic for both senders and receivers. Uses RabbitMQ
- **service** - Basic business logic that handles calls to data layer

## Usage
There is a provided Makefile. This application uses **dep** to ensure the packages its dependent on
are provided. To ensure the packages go to where the application was cloned to and run

``
make dep
``
  
To run the application RabbitMQ should be initialized. To run locally, you can do the following using HomeBrew (for additional instructions for different tools, refer to [the installation guide](https://www.rabbitmq.com/download.html))

```$bash
brew update
brew install rabbitmq
```

Export to the PATH
```bash
export PATH=$PATH:/usr/local/opt/rabbitmq/sbin
```

Then run the server with
```bash
rabbitmq-server
```

Once all the tools are in place you can initialize the server in a shell by doing
```bash
make main [short] //This option allows for 'debugging' and shortens the timers on deletion/reset
```

## Design Choices
**dep** is used by **Go**
 - **Why Go?**
    - Open source
    - Scalable
    - Cross-Platform
    - Good Concurrency which is a must for the application
    - Built-in Garbage Collection and Memory management
    - ***(Personal)*** Wanted to continue practicing and using the language. 

**RabbitMQ** is used by the application for a message queue
 - **Why Message Queue?**
    - A queue would in theory allow the decoupling of the backend and frontend logic
        - The HTTP package for example could remove the service instance and then be completely reliant on the sender
    - Allows for a separation of concerns
        - Attempted to design and implement packages after implementing the queue with this in mind
 - **Why RabbitMQ?**
    - Open source
    - Good support libraries in Golang
    - Reliability and flexible routing make it a good flexible choice
    - The ability to scale and federate in case the application needs to be used by a larger hospital makes scaling a possibility
    - Considerations were make for **Kafka** and **Redis** but given the size and complexity of the application these felt like overkill
  
The application initially comes to run with an in-memory cache, although it has been designed to where a different implementation (SQL, NoSQL) is possible
 - **Why Memcache?**
    - Scale and scope of the issue didn't seem to necessitate a large table or database
        - The data is simplistic enough to be held in most modern servers under 100MB simply in memory
    - Data was simplistic
        - The use of simple Go structs held most of the relevant data
    - With the implication of data being 'invalid' after 24 hours, a large scale Database store seemed ineffective

The application also makes used of Google's UUID library
 - **Why UUID?**
    - Uniqueness of the key allows for no collisions to occur in the system
        - Rand was considered, but the library is only 'pseudo-random' and wouldn't make for a good key
    - After 24 hours the keys are easily remade and doesn't impact the memcache that harshly 

 - **Considerations**
    - Implementing a rule in Makefile to allow for the initialization of RabbitMQ and the service seems like an easy enough change
    - Error handling needs to be improved: there are some simple enough messages to allow for flow of the data but need to be able to handle exceptions
    - In theory now, http could be decoupled and moved to its own project. This would make the application flow a bit more complicated, would need to implement the logic and then update the README
    
