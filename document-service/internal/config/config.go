package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port              int
	AWSRegion         string
	S3Bucket          string
	CassandraHosts    []string
	CassandraKeyspace string
	KafkaBrokers      []string
}

func Load() *Config {
	return &Config{
		Port:              getEnvAsInt("PORT", 8080),
		AWSRegion:         getEnv("AWS_REGION", "us-east-1"),
		S3Bucket:          getEnv("S3_BUCKET", "dms-documents"),
		CassandraHosts:    getEnvAsSlice("CASSANDRA_HOSTS", []string{"cassandra-0.cassandra.data.svc.cluster.local"}),
		CassandraKeyspace: getEnv("CASSANDRA_KEYSPACE", "metadata_keyspace"),
		KafkaBrokers:      getEnvAsSlice("KAFKA_BROKERS", []string{"kafka-0.kafka.data.svc.cluster.local:9092"}),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	return strings.Split(valueStr, ",")
}
