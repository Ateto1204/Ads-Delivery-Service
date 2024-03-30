curl -X POST -H "Content-Type: application/json" \
   http://localhost:8080/api/v1/ad \
   --data '{
      "title": "AD 1",
      "startAt": "2023-01-10T03:00:00.000Z",
      "endAt": "2024-12-31T16:00:00.000Z",
      "conditions": {
         "ageStart": 20,
         "ageEnd": 30,
         "gender" : "F",
         "country": ["TW", "JP"],
         "platform": ["ios"]
      }
   }'
curl -X POST -H "Content-Type: application/json" \
   http://localhost:8080/api/v1/ad \
   --data '{
      "title": "AD 31",
      "startAt": "2024-12-10T03:00:00.000Z",
      "endAt": "2024-12-31T16:00:00.000Z",
      "conditions": {
         "ageStart": 20,
         "ageEnd": 30,
         "gender" : "F",
         "country": ["TW", "JP"],
         "platform": ["android", "ios"]
      }
   }'
curl -X POST -H "Content-Type: application/json" \
   http://localhost:8080/api/v1/ad \
   --data '{
      "title": "AD 27",
      "startAt": "2023-12-10T03:00:00.000Z",
      "endAt": "2024-12-31T16:00:00.000Z",
      "conditions": {
         "ageStart": 20,
         "ageEnd": 30,
         "gender" : "F",
         "country": ["TW", "JP"],
         "platform": ["android"]
      }
   }'
curl -X POST -H "Content-Type: application/json" \
   http://localhost:8080/api/v1/ad \
   --data '{
      "title": "AD 55",
      "startAt": "2023-12-10T03:00:00.000Z",
      "endAt": "2025-12-31T16:00:00.000Z",
      "conditions": {
         "ageStart": 28,
         "ageEnd": 30,
         "gender" : "F",
         "country": ["TW", "JP"],
         "platform": ["android"]
      }
   }'
curl -X POST -H "Content-Type: application/json" \
   http://localhost:8080/api/v1/ad \
   --data '{
      "title": "AD 34",
      "startAt": "2023-12-10T03:00:00.000Z",
      "endAt": "2024-12-31T16:00:00.000Z",
      "conditions": {
         "ageStart": 20,
         "ageEnd": 21,
         "gender" : "M",
         "country": ["TW", "JP"],
         "platform": ["android", "ios"]
      }
   }'
curl -X POST -H "Content-Type: application/json" \
   http://localhost:8080/api/v1/ad \
   --data '{
      "title": "AD 17",
      "startAt": "2023-12-10T03:00:00.000Z",
      "endAt": "2023-12-31T16:00:00.000Z",
      "conditions": {
         "ageStart": 20,
         "ageEnd": 30,
         "gender" : "F", 
         "country": ["TW", "JP"],
         "platform": ["android", "ios"]
      }
   }'
curl -X GET -H "Content-Type: application/json" \
   "http://localhost:8080/api/v1/ad?offset=10&limit=3&age=24&gender=F&country=TW&platform=ios"
curl -X GET -H "Content-Type: application/json" \
   "http://localhost:8080/api/v1/ad?offset=10&limit=3&age=24&gender=M&country=JP&platform=android"