package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var kafkaProducer *kafka.Writer

const (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// For Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// For Consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset, // khi 1 mà 1 user vào thì lấy giá trị đầu tiên
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoan
func NewStock(msg, typeMsg string) *StockInfo {
	return &StockInfo{
		Message: msg,
		Type:    typeMsg,
	}
}

func actionStock(c *gin.Context) {
	stock := NewStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = stock

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

func RegisterConsumerATC(id int) {
	// group consumer ??
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer (%d) is Hong Phien ATC::\n", id)
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) error: %v", id, err)
		}

		fmt.Printf("Consumer (%d), hong topic: %v, partition: %v, offset: %v, time:%d %s \n",
			id, msg.Topic, msg.Partition, msg.Offset, msg.Time.Unix(), string(msg.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("/action/stock", actionStock)

	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4)

	r.Run(":8000")
}
