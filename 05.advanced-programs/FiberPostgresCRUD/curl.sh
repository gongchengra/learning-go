curl -X POST -H "Content-Type: application/json" -d '{"Title":"Love","Owner":"Alan","Rating":30}' 127.0.0.1:3000/api/v1/item
curl -X POST -H "Content-Type: application/json" -d '{"Title":"Good","Owner":"Alan","Rating":40}' 127.0.0.1:3000/api/v1/item
curl -X POST -d 'Title=test&Owner=Alan&Rating=50' 127.0.0.1:3000/api/v1/item
curl 127.0.0.1:3000/api/v1/item
curl -X DELETE 127.0.0.1:3000/api/v1/item/2
