#!/bin/bash

# List of topics to create
TOPICS=(transaction-log)

# Kafka broker
KAFKA_BROKER=kafka:9092

for topic in "${TOPICS[@]}"; do
  kafka-topics --create --topic "$topic" --bootstrap-server "$KAFKA_BROKER" --partitions 1 --replication-factor 1
done
