# Document Service
The Document service is a Go-based microservice that handles document upload, retrieval, version management and check-out/check-in functionality for this system.

## Features
- Document upload with streaming support
- Document retrieval with pre-signed S3 URLS
- Document update and soft delete
- Version management
- Check-in/Check-out functionality with locking
- Integration with Cassandra for metadata storage
- Integration with S3 for document storage
- Integration with Kafka for event publishing
- Access control checks
- Structured logging with Zap
- Prometheus metrics instrumentation

## Architecture
- **HTTP Framework**: Gin
- **Database**: Cassandra (gocql driver)
- **Object Storage**: AWS S3 (aws-sdk-go-v2)
- **Message Broker**: Kafka (sarama)
- **Logging**: Zap
- **Metrics**: Promethesus

## API Endpoints

### Health Checks
- `GET /health` - health check endpoint
- `GET /ready` - Readiness check endpoint
- `GET /metrics` - Promethesus metrics endpoint

### Document Operations
- `POST /api/v1/documents` - Upload a new document
- `GET /api/v1/documents/:id` - Retrieve a document
- `PUT /api/v1/documents/:id` - Update a document
- `DELETE /api/v1/documents/:id` - Soft delete a document

### Configuration

The service is configured via environment variables:
- `PORT` - HTTP server port (default: 8080)
- `AWS_REGION` - AWS region (default: us-east-1)
- `S3_BUCKET` - S3 Bucket name (default: dms-documents)
- `CASSANDRA_HOSTS` - Comma-separated list of Cassandra hosts
- `CASSANDRA_KEYSPACE` - Cassandra keyspace (default: metadata_keyspace)
- `KAFKA_BROKERS` - Comma-separated list of Kafka brokers

## Running locally
```bash
    go mod download
    
    go run cmd/main.go
```

## Building
```bash
    go build -o document-service cmd/main.go
    
    docker build -t document-service:latest .
```

## Testing

```bash
# Run unit tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Deployment

The service is deployed on Kubernetes using the manifests in the `kubernetes/` directory.
See the Dockerfile and Kubernetes manifests for deployment configuration.