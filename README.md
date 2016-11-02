# Temperature-Monitoring-Application
This application simulates the monitoring system for a power plant using Go

It uses RabbitMQ to access message queues, etc., the distributed nature of the resulting suite of applications are composed of as much core Go as possible with little or no support from other third-party libraries.

The purpose of this application is to learn the Go language.

This will be composed of following components.

1. It will simulate the behavior of sensors to recieve temperature data.
2. Message broker - decouples the messages source from its consumer.
3. The next component is the coordinator that will recieve messages from message broker and pass on to database manager which will be handling database operations.
4. Web application that will recieve data from coordinator, so user can navigate to it and recieve real time data.

This system uses Postgre SQL as a database.
