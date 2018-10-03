# Example Schema
```
{
  "nodes": [
    {
      "count": 1,
      "properties": {
        "name": "name"
      },
      "label": "user"
    },
    {
      "count": 20,
      "properties": {
        "username": "UserName",
        "ccNumber": "creditCardNum"
      },
      "label": "account"
    },
    {
      "count": 200,
      "properties": {
        "username": "UserName"
      },
      "label": "game"
    },
    {
      "count": 200,
      "properties": {
        "username": "UserName"
      },
      "label": "save"
    }
  ],
  "relationships": [
    {
      "source": "user",
      "target": "account",
      "label": "has",
      "max": 3,
      "type": "1..n"
    },
    {
      "source": "account",
      "target": "game",
      "label": "purchased",
      "max": 5,
      "type": "n..n"
    },
    {
      "source": "game",
      "target": "save",
      "label": "uses",
      "max": 10,
      "type": "n..n",
      "probability": 100
    }
  ]
}
```
# Test Request
```
curl -X POST \
  https://utmn8prfg1.execute-api.us-west-2.amazonaws.com/prod \
  -H 'Cache-Control: no-cache' \
  -H 'Postman-Token: 9d56bf00-c46f-462d-953d-ac45d4155d47' \
  -d '{
  "nodes": [
    {
      "count": 1,
      "properties": {
        "name": "name"
      },
      "label": "user"
    },
    {
      "count": 20,
      "properties": {
        "username": "UserName",
        "ccNumber": "creditCardNum"
      },
      "label": "account"
    },
    {
      "count": 200,
      "properties": {
        "username": "UserName"
      },
      "label": "game"
    },
    {
      "count": 200,
      "properties": {
        "username": "UserName"
      },
      "label": "save"
    }
  ],
  "relationships": [
    {
      "source": "user",
      "target": "account",
      "label": "has",
      "max": 3,
      "type": "1..n"
    },
    {
      "source": "account",
      "target": "game",
      "label": "purchased",
      "max": 5,
      "type": "n..n"
    },
    {
      "source": "game",
      "target": "save",
      "label": "uses",
      "max": 10,
      "type": "n..n",
      "probability": 100
    }
  ]
}'
```

# Error example
```
{
    "message": "relationship definition must include type: 1..n, 1..1, n..n",
    "code": 400
}
```