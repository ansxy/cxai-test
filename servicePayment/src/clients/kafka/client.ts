import dotenv from "dotenv";
import { Kafka } from "kafkajs";
import { ProcessTransaction } from "../../controller/transaction";

dotenv.config();

const kafka = new Kafka({
  brokers: [process.env.KAFKA_BROKER || "localhost:9092"],
});

const consumer = kafka.consumer({
  groupId: "transaction-group",
  allowAutoTopicCreation: true,
});
const initConsumer = async () => {
  try {
    await consumer.connect();
    await consumer.subscribe({
      topic: "transaction-log",
      fromBeginning: true,
    });

    await consumer.run({
      eachMessage: async ({ topic, partition, message }) => {
        try {
          const key = message.key?.toString();
          const value = JSON.parse(message.value?.toString() || "");
          console.log("Message", key, value);
          ProcessTransaction(value);
        } catch (err) {
          console.error("Error processing message:", err);
        }
      },
    });
  } catch (err) {
    console.error("Error initializing consumer:", err);
  }
};

export default initConsumer;
