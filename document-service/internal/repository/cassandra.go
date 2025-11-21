package repository

import (
	"fmt"
	"time"

	"github.com/Eomaxl/DMS-Services/document-service/internal/model"
	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

type CassandraRepository struct {
	session *gocql.Session
	logger  *zap.Logger
}

func NewCassandraRepository(hosts []string, keyspace string, logger *zap.Logger) (*CassandraRepository, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 10 * time.Second
	cluster.ConnectTimeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create Cassandra session : %w", err)
	}

	return &CassandraRepository{
		session: session,
		logger:  logger,
	}, nil
}

func (r *CassandraRepository) Close() {
	if r.session != nil {
		r.session.Close()
	}
}
func (r *CassandraRepository) CreateDocument(doc *model.Document) error {
	query := `INSERT INTO documents (
		document_id, file_name, file_type, file_size, s3_key, s3_bucket,
		department, document_type, reference_number, created_at, created_by,
		modified_at, modified_by, current_version, status, tags, custom_metadata,
		ocr_status, ocr_confidence, retention_policy, archival_date, is_deleted
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	err := r.session.Query(query,
		doc.DocumentID, doc.FileName, doc.FileType, doc.FileSize, doc.S3Key, doc.S3Bucket,
		doc.Department, doc.DocumentType, doc.ReferenceNumber, doc.CreatedAt, doc.CreatedBy,
		doc.ModifiedAt, doc.ModifiedBy, doc.CurrentVersion, doc.Status, doc.Tags, doc.CustomMetadata,
		doc.OCRStatus, doc.OCRConfidence, doc.RetentionPolicy, doc.ArchivalDate, doc.IsDeleted,
	).Exec()

	if err != nil {
		r.logger.Error("Failed to create document", zap.Error(err), zap.String("document_id", doc.DocumentID.String()))
		return fmt.Errorf("Failed to create document: %w", err)
	}
	return nil
}
