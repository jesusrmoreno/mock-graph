A small application to generate mock graph structures. 

All relationships are directed and only connected nodes are returned. 

Relationship types are currently `1..n`, `n..n`, `1..1` and are treated as _at least one_.
This very much an early WIP so if any issues are found please open an issue.
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

# Example response
```
{
    "nodes": [
        {
            "properties": {
                "ccNumber": "0000-5623-5822-4471",
                "username": "repellat_mollitia"
            },
            "label": "account",
            "id": "467aa8a3-3602-44d5-9ce9-a5b7f9fcec74"
        },
        {
            "properties": {
                "username": "gGriffin"
            },
            "label": "game",
            "id": "fa07299e-6000-4091-8d54-34969871e310"
        },
        {
            "properties": {
                "username": "AntonioBanks"
            },
            "label": "game",
            "id": "c91bbbf0-f41b-4315-8871-c091dbce2931"
        },
        {
            "properties": {
                "username": "veritatis_qui_aut"
            },
            "label": "save",
            "id": "4fa74103-ef77-4e96-9bb3-841f7b43c871"
        },
        {
            "properties": {
                "ccNumber": "0000-8646-2668-3594",
                "username": "bArmstrong"
            },
            "label": "account",
            "id": "9cf402ae-a80b-401a-b87c-1be3d39b0cab"
        },
        {
            "properties": {
                "username": "AnnaOwens"
            },
            "label": "save",
            "id": "98e5d0e2-735b-4601-ac94-9cb81629f827"
        },
        {
            "properties": {
                "username": "quos_et"
            },
            "label": "save",
            "id": "8d6958ce-b82f-4093-b5fb-4f8615f5e162"
        },
        {
            "properties": {
                "username": "qGraham"
            },
            "label": "save",
            "id": "fd1060f2-b1e4-4826-bd09-07b71de671c4"
        },
        {
            "properties": {
                "name": "Ronald Ross"
            },
            "label": "user",
            "id": "220488df-ce6c-4636-9492-25c3fc5d4b70"
        },
        {
            "properties": {
                "ccNumber": "0000-9684-7209-1131",
                "username": "BrendaMitchell"
            },
            "label": "account",
            "id": "d34e4001-aaa6-42e9-8f1d-1b571a24e052"
        },
        {
            "properties": {
                "username": "qui_ut_vel"
            },
            "label": "game",
            "id": "8e1b7603-cf67-4324-934e-e005a0d8536b"
        },
        {
            "properties": {
                "username": "CarlosRussell"
            },
            "label": "game",
            "id": "b59de0da-d713-447a-a7b7-db7e45a018fa"
        },
        {
            "properties": {
                "username": "et_eligendi"
            },
            "label": "save",
            "id": "ccf367d9-9ea1-42a2-8d8d-e34469e7683e"
        },
        {
            "properties": {
                "username": "et_eos_pariatur"
            },
            "label": "save",
            "id": "2e941259-9a74-439d-a4fb-3caaa865cb62"
        },
        {
            "properties": {
                "username": "oFuller"
            },
            "label": "game",
            "id": "61af1181-e354-4ac4-8a24-a011bb901417"
        },
        {
            "properties": {
                "username": "2Henderson"
            },
            "label": "save",
            "id": "01fd688a-5051-4db5-9853-57cbf094394d"
        },
        {
            "properties": {
                "username": "jLittle"
            },
            "label": "save",
            "id": "7be5cfd7-5fd7-47b3-ad94-53e54bffaaf9"
        },
        {
            "properties": {
                "username": "iWood"
            },
            "label": "save",
            "id": "34ceaacb-2577-4994-9e70-944de95e4c84"
        },
        {
            "properties": {
                "username": "0Torres"
            },
            "label": "save",
            "id": "3624f379-736c-45c3-8187-be689431cb63"
        }
    ],
    "edges": [
        {
            "source": "220488df-ce6c-4636-9492-25c3fc5d4b70",
            "target": "d34e4001-aaa6-42e9-8f1d-1b571a24e052",
            "label": "has",
            "sourceLabel": "user",
            "targetLabel": "account"
        },
        {
            "source": "220488df-ce6c-4636-9492-25c3fc5d4b70",
            "target": "467aa8a3-3602-44d5-9ce9-a5b7f9fcec74",
            "label": "has",
            "sourceLabel": "user",
            "targetLabel": "account"
        },
        {
            "source": "220488df-ce6c-4636-9492-25c3fc5d4b70",
            "target": "9cf402ae-a80b-401a-b87c-1be3d39b0cab",
            "label": "has",
            "sourceLabel": "user",
            "targetLabel": "account"
        },
        {
            "source": "d34e4001-aaa6-42e9-8f1d-1b571a24e052",
            "target": "fa07299e-6000-4091-8d54-34969871e310",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "d34e4001-aaa6-42e9-8f1d-1b571a24e052",
            "target": "61af1181-e354-4ac4-8a24-a011bb901417",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "d34e4001-aaa6-42e9-8f1d-1b571a24e052",
            "target": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "d34e4001-aaa6-42e9-8f1d-1b571a24e052",
            "target": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "d34e4001-aaa6-42e9-8f1d-1b571a24e052",
            "target": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "467aa8a3-3602-44d5-9ce9-a5b7f9fcec74",
            "target": "61af1181-e354-4ac4-8a24-a011bb901417",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "467aa8a3-3602-44d5-9ce9-a5b7f9fcec74",
            "target": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "467aa8a3-3602-44d5-9ce9-a5b7f9fcec74",
            "target": "fa07299e-6000-4091-8d54-34969871e310",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "467aa8a3-3602-44d5-9ce9-a5b7f9fcec74",
            "target": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "467aa8a3-3602-44d5-9ce9-a5b7f9fcec74",
            "target": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "9cf402ae-a80b-401a-b87c-1be3d39b0cab",
            "target": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "9cf402ae-a80b-401a-b87c-1be3d39b0cab",
            "target": "fa07299e-6000-4091-8d54-34969871e310",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "9cf402ae-a80b-401a-b87c-1be3d39b0cab",
            "target": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "9cf402ae-a80b-401a-b87c-1be3d39b0cab",
            "target": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "9cf402ae-a80b-401a-b87c-1be3d39b0cab",
            "target": "61af1181-e354-4ac4-8a24-a011bb901417",
            "label": "purchased",
            "sourceLabel": "account",
            "targetLabel": "game"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "01fd688a-5051-4db5-9853-57cbf094394d",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "4fa74103-ef77-4e96-9bb3-841f7b43c871",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "ccf367d9-9ea1-42a2-8d8d-e34469e7683e",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "2e941259-9a74-439d-a4fb-3caaa865cb62",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "8d6958ce-b82f-4093-b5fb-4f8615f5e162",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "7be5cfd7-5fd7-47b3-ad94-53e54bffaaf9",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "98e5d0e2-735b-4601-ac94-9cb81629f827",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "34ceaacb-2577-4994-9e70-944de95e4c84",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "3624f379-736c-45c3-8187-be689431cb63",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "8e1b7603-cf67-4324-934e-e005a0d8536b",
            "target": "fd1060f2-b1e4-4826-bd09-07b71de671c4",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "ccf367d9-9ea1-42a2-8d8d-e34469e7683e",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "7be5cfd7-5fd7-47b3-ad94-53e54bffaaf9",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "34ceaacb-2577-4994-9e70-944de95e4c84",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "fd1060f2-b1e4-4826-bd09-07b71de671c4",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "2e941259-9a74-439d-a4fb-3caaa865cb62",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "01fd688a-5051-4db5-9853-57cbf094394d",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "98e5d0e2-735b-4601-ac94-9cb81629f827",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "4fa74103-ef77-4e96-9bb3-841f7b43c871",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "8d6958ce-b82f-4093-b5fb-4f8615f5e162",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "fa07299e-6000-4091-8d54-34969871e310",
            "target": "3624f379-736c-45c3-8187-be689431cb63",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "4fa74103-ef77-4e96-9bb3-841f7b43c871",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "3624f379-736c-45c3-8187-be689431cb63",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "8d6958ce-b82f-4093-b5fb-4f8615f5e162",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "2e941259-9a74-439d-a4fb-3caaa865cb62",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "98e5d0e2-735b-4601-ac94-9cb81629f827",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "01fd688a-5051-4db5-9853-57cbf094394d",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "ccf367d9-9ea1-42a2-8d8d-e34469e7683e",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "fd1060f2-b1e4-4826-bd09-07b71de671c4",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "7be5cfd7-5fd7-47b3-ad94-53e54bffaaf9",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "b59de0da-d713-447a-a7b7-db7e45a018fa",
            "target": "34ceaacb-2577-4994-9e70-944de95e4c84",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "fd1060f2-b1e4-4826-bd09-07b71de671c4",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "98e5d0e2-735b-4601-ac94-9cb81629f827",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "2e941259-9a74-439d-a4fb-3caaa865cb62",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "ccf367d9-9ea1-42a2-8d8d-e34469e7683e",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "3624f379-736c-45c3-8187-be689431cb63",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "01fd688a-5051-4db5-9853-57cbf094394d",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "34ceaacb-2577-4994-9e70-944de95e4c84",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "7be5cfd7-5fd7-47b3-ad94-53e54bffaaf9",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "4fa74103-ef77-4e96-9bb3-841f7b43c871",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "c91bbbf0-f41b-4315-8871-c091dbce2931",
            "target": "8d6958ce-b82f-4093-b5fb-4f8615f5e162",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "01fd688a-5051-4db5-9853-57cbf094394d",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "8d6958ce-b82f-4093-b5fb-4f8615f5e162",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "7be5cfd7-5fd7-47b3-ad94-53e54bffaaf9",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "4fa74103-ef77-4e96-9bb3-841f7b43c871",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "fd1060f2-b1e4-4826-bd09-07b71de671c4",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "2e941259-9a74-439d-a4fb-3caaa865cb62",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "98e5d0e2-735b-4601-ac94-9cb81629f827",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "3624f379-736c-45c3-8187-be689431cb63",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "ccf367d9-9ea1-42a2-8d8d-e34469e7683e",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        },
        {
            "source": "61af1181-e354-4ac4-8a24-a011bb901417",
            "target": "34ceaacb-2577-4994-9e70-944de95e4c84",
            "label": "uses",
            "sourceLabel": "game",
            "targetLabel": "save"
        }
    ]
}
```

# Error example
```
{
    "message": "relationship definition must include type: 1..n, 1..1, n..n",
    "code": 400
}
```
