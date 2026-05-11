ENV_PATH = "./.env"

init:
	if [ -f ENV_PATH ]; then
		echo ".env file already exists"
	else
		cp .env.example .env
	fi
	echo "update the config.json to your liking and your user name"

lint:
	golangci-lint run 

format:
	goimports -w .
	golangci-lint fmt

run:
	go run ./cmd/blocker/main.go 
