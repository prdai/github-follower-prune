ENV_PATH = "./.env"

init:
	if [ -f ENV_PATH ]; then
		echo ".env file already exists"
	else
		cp .env.example .env
	fi
	echo "update the config.json to your liking and your user name"

lint:
	golangci-lint run ./test/

format:
	golangci-lint fmt ./test/

run:
	go run main.go
