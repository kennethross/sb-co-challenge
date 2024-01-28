# SB-CO Challenge

A simple web server to solve the SB-CO challenge.

## Endpoints
Start a new game
```
curl  http://localhost:3000/new\?width\=10\&height\=10
```

Validate Game
```
curl --location 'http://localhost:3000/validator' \
--header 'Content-Type: application/json' \
--data '{
  "state": {
    "gameId": "924ddb07-b9f1-4e36-a2cd-35584755670d",
    "width": 10,
    "height": 10,
    "score": 2,
    "fruit": {
      "x": 2,
      "y": 2
    },
    "snake": {
      "x": 0,
      "y": 0,
      "velX": 1,
      "velY": 0
    }
  },
  "ticks": [
    {
      "velX": 1,
      "velY": 0
    },
    {
      "velX": 0,
      "velY": -1 
    },
    {
      "velX": 0,
      "velY": -1 
    }
  ]
}'
```