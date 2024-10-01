# mysql container usage

- `make up`: Starts the services in the background using Docker Compose. The MySQL container will be started and listen on the specified port.
- `make down`: Stops the services and removes the containers using Docker Compose. This will stop and remove all running containers.
- `make logs`: Displays the logs of the services in real-time using Docker Compose. This allows you to monitor the logs of the containers.
- `make mysql`: Connects to the MySQL server using the MySQL CLI. Assumes that the MySQL container is already running.
- `make clean`: Stops the services, removes the containers, and deletes the associated volumes and networks using Docker Compose. This will also remove the volumes where data is persisted.
