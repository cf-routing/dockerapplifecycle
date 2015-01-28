# docker_app_lifecycle

The docker app lifecycle implements a Docker deployment strategy for Cloud
Foundry on Diego.

The **Builder** extracts the start command and execution metadata from the docker image.

The **Launcher** executes the start command with the correct cloudfoundry and docker enviroment.

The **Healthcheck** performs a tcp port check, defaulting to port 8080.

Read about the app lifecycle spec here: https://github.com/cloudfoundry-incubator/diego-design-notes#app-lifecycles
