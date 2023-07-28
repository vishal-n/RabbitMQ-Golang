# RabbitMQ-Golang
A simple use-case of RabbitMQ implemented in Golang


# How to setup a RabbitMQ instance running in the background
1. If you are a mac user, you can run the command: brew install rabbitmq
2. In case you prefer docker, you can just run the command: docker-compose up.
3. The docker-compose command uses the docker-compose.yml located inside the repo, in order to have setup the rabbitmq in the container.


# Local machine setup
To setup the repo, on your local machine, follow these steps:
1. git clone git@github.com:vishal-n/RabbitMQ-Golang.git
2. Make sure that the golang is installed on your machine.
3. You need have a rabbitmq instance running in your background, such that the program can use that rabbitmq instance.
4. Once the repo and rabbitmq instance is setup, run the following commands in two seperate terminals:

   -> go run message_publisher.go
   
   -> go run message_consumer.go
6. Once the above commands are successfully excetuted, the message_publisher.go prompts the user to enter a message each time it's exceuted.
7. Once the user enters the message, the message_consumer.go which keeps listening continuously receives the message and the message is displayed on the terminal.
