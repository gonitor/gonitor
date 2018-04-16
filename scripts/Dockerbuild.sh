echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build -t gonitor/gonitor:0.0.0 .
docker push gonitor/gonitor:0.0.0