cd server && screen -dmS server go run server.go && cd ..
cd client && go build -o client
./client create --authorid "101010010" --content "Some great blog content" --title "Very great blog post"
./client read --blogid 62fc6fd9e6a2433eaa56fda7
./client update --id 62fc6fd9e6a2433eaa56fda7 --title "Some great title" --content "Some great content" --authorid "100203900392840932875"
./client delete --blogid 62fc6fd9e6a2433eaa56fda7
