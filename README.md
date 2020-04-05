![Gopher](https://golang.org/lib/godoc/images/footer-gopher.jpg)

# bootstrap-go-httpserver
A minimumal bootstrap HTTP server written in [Go](https://golang.org)

## Features
- Follows minimal standard project layout specified in [project-layout](https://github.com/golang-standards/project-layout)
- Barebones HTTP server using the following libraries:
  - Logging: [uber-go/zap](https://github.com/uber-go/zap)
  - Routing and middleware: [go-chi/chi](https://github.com/go-chi/chi)
  - Environment configuration: [spf13/viper](https://github.com/spf13/viper)
- Easy addition of new routes + graceful shutdown
- Running the server using [Docker](https://www.docker.com)

## Requirements
- Docker
- Go 1.14

## Setting up
1. Clone this repository

   ```
    git clone git@github.com:irisgve/bootstrap-go-httpserver.git
   ```
2. Replace all instances of `bootstrap-go-httpserver` with the app name
3. To run the server 

   ```
    make run
   ```  

# License

The scripts and documentation in this project are released under the [MIT License](LICENSE)