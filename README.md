# Rabbit mq test for kubernetes 
This repository is for testing rabbit mq with golang and also test the consumer with rproducer


# First things
```sh
docker network create rabbits

```

# Launch rabbitmq
```sh
docker build -t my-rabbit .

docker run -d \
  --net rabbits \
  --name rabbitmq \
  --hostname rabbit-1 \
  -p 15672:15672 \
  -p 5672:5672 \
  my-rabbit
```
Identifier of rabbit mq <identifier>@<hostname>
# Run in statefulset in kubernetes


# 4369 : PORT FOR CLUSTER
# 5671-5672 : PORT COMMUNICATION
