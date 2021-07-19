#Test project with small api in Golang and front based on Angular

## Setup

This project contains two application, you can run them separately by Dockerfile in each directory, or run docker-compose.yml

### Front side

You must have NodeJS installed on your machine with yarn package manager

For MacOS "brew install nodejs"
For Linux(Debian based) "sudo apt install nodejs"

Then "node install -g yarn"

1. yarn add global @angular/cli
2. cd ./front
3. yarn install
4. ng serve
5. Visit localhost:4200 on your browser

### Backend side

For MacOS "brew install golang"

1. cd ./back
2. go mod download
3. go build
4. ./back

Hit api with POST: localhost:3000/pages 
```
{
    "url": "https://example.com"
}
```

Or GET: localhost:3000/pages to download list of pages


#### Container version

If you don't want to manage local deployment manually you can use ``docker compose up`` and frontend application with be running on localhost:4000