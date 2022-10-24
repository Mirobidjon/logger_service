# logger_service

Logger service is Stores 4 RMQ consumers listening to the 
queues exchange "logger" routes to. Messages differ in the 
level of logging:
    1. consumer1 - consumes messages with level "debug" and
    2. consumer2 - consumes messages with level "info" and
    3. consumer3 - consumes messages with level "error" and
    4. consumer4 - consumes messages with all level of logging.

## Installation
Use the following command to run the application:

    docker run us.gcr.io/learn-cloud-0809/logger_service:latest

# ENVIRONMENT VARIABLES
    1. RABBIT_MQ_URL - RabbitMQ URL (default: amqp://guest:guest@localhost:5672)
    2. EXCHANGE_NAME - RabbitMQ exchange name (default: v1.phone)

