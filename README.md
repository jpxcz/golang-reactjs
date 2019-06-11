# Golang and ReactJS webapp
Webserver in Golang using fswatch for the hotreload  
Client Side on ReactJS using nodemon and react-create-app  
NGINX for serving routes and files  

## Development
Must have installed docker and docker compose. Follow steps on https://docs.docker.com/  
This should only be used for development only as it will start the application with hot reloads on port `4000`  
``` sh
$ cd appfolder
$ docker-compose -f docker-compose.dev.yml up --build
```
The docker-compose might not recognize the applications the first time you run the command. It's suggested to run it again if any applications fails.

To make requests to the Webserver, use the route `localhost:4000/api/vX/`  
This was made so in the future you can use diferent api versions to make requests. Making the CI/CD more efficient.

## Start on Production Style
``` sh
$ cd appfolder
$ docker-compose up
```

## Next Steps:
- golang dependencies
- AWS Beanstalk integration
- env variables

Suggestions to jop0693@gmail.com or via github 