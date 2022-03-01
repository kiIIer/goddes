# goddes
A simple app designed to benchmark sites. It is written in go and uses fasthttp package to maximase the perfomance. Please keep in mind that request sending speed was the priority and I am just a student in KPI.

# Installation
I strongly suggest using Docker image like so:  
`docker pull IMAGE NAME`  
or (if you have go installed) use  
`go install github.com/codesenberg/bombardier@latest`

# Running the app
Usage:
  goddes \[command\]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  hell        Command to make the hell on site
  help        Help about any command

Flags:
  -h, --help   help for goddes

Use "goddes \[command\] --help" for more information about a command.

If you are new to all this stuff  
`goddes hell --url (url of target) --gophers (number of go routines to use) --method (method to use(GET, POST, PUT...))`  
# Exmaple  
This will start hell on your localhost using 5000 goroutines with http get requests  
`goddes hell --url http://localhost:80/ --gophers 5000 --method GET`  
For docker users(this does the same thing)   
`docker run IMAGENAME hell --url http://localhost:80/ --gophers 5000 --method GET`
