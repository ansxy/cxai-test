package custome_kafka

import (
	"log"

	"github.com/ansxy/golang-boilerplate-gin/config"
	"github.com/segmentio/kafka-go"
)

func KafkaWriter(configKafka *config.KafkaConfigWriter) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: configKafka.Brokers,
		Topic:   configKafka.Topic,
	})
}

func KafkaReader(configKafka *config.KafkaConfigReader) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     configKafka.Brokers,
		Topic:       configKafka.Topic,
		Partition:   configKafka.Partition,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		Logger:      kafka.LoggerFunc(log.Printf),
		ErrorLogger: kafka.LoggerFunc(log.Printf),
	})
}

type KafaClient struct {
	Writer *kafka.Writer
}

func InitKafka(configKafka *config.KafkaConfig) (*KafaClient, error) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: configKafka.KafkaConfigWriter.Brokers,
		Topic:   configKafka.KafkaConfigWriter.Topic,
	})

	// r := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers:     configKafka.KafkaConfigReader.Brokers,
	// 	Topic:       configKafka.KafkaConfigReader.Topic,
	// 	Partition:   configKafka.KafkaConfigReader.Partition,
	// 	MinBytes:    10e3,
	// 	MaxBytes:    10e6,
	// 	Logger:      kafka.LoggerFunc(log.Printf),
	// 	ErrorLogger: kafka.LoggerFunc(log.Printf),
	// })

	// msg := kafka.Message{
	// 	Key:   []byte("1"),
	// 	Value: []byte("JIBTIK"),
	// }

	// if err := w.WriteMessages(context.Background(), msg); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := w.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// defer r.Close()

	// fmt.Println("start reading")

	// for {
	// 	msg, err := r.ReadMessage(context.Background())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(string(msg.Value))
	// }
	return &KafaClient{
		Writer: w,
	}, nil
}
