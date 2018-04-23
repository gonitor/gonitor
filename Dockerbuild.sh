CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t gonitor/gonitor .
docker tag  gonitor/gonitor gonitor/gonitor:0.1.0