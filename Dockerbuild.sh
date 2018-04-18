echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build -t gonitor/gonitor .
docker tag  gonitor/gonitor gonitor/gonitor:0.1.0
docker push gonitor/gonitor