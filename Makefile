build:
	go build -o  main.go

run: 
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

