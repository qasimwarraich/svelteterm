CC = go
MAIN = ./main.go

all:
	${CC} run ${MAIN}

build:
	${CC} build -o ./bin ${MAIN}

test:
	gotest -v ./src/...
