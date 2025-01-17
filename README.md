# lms_gameLife -  primitive game realization

## Description

lms_gameLife is a primitive implementation of the game “Life” by John Conway.   
The project includes an HTTP server to which the game logic is connected.

## Technologies used

- Programming language: Go (1.23.4)
- Web server: net/http (standard Go library)
- Logging: zap (Uber)
- Testing: testing (standard Go library)

## Installation and startup

### Требования

- Go 1.19 or higher

### Installation

1.Clone the repository:
```
git clone https://github.com/flexer2006/lms_gameLife.git
```

2.Navigate to the project directory:  
```
cd lms_gameLife
```

3.Install dependencies:  
```
go mod tidy
```

### Launch

Start the server:
```
go run cmd/life/main.go
```
The server will be available at http://localhost:8081


### Utilization

POST /setstate: Set the new game state.  
	Example request:
```
	{
  "fill": 30
	}
```
	Response Example: HTTP 200 OK

PUT /reset: Restore the last saved game state.  
	Response Example:
```
	{
  "fill": 30
	}
```

GET /next-state: Get the next state of the game.  
	Example response:
```
	[
  [false, true, false],
  [true, true, true, true],
  [false, true, false]
	]
```

### Testing

```
go test ./...
```


