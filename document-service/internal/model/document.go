package model

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	DocumentID      uuid.UUID         `json:"document_id" cql:"document_id"`
	FileName        string            `json:"file_name" cql:"file_name"`
	FileType        string            `json:"file_type" cql:"file_type"`
	FileSize        int64             `json:"file_size" cql:"file_size"`
	S3Key           string            `json:"s3_key" cql:"s3_key"`
	S3Bucket        string            `json:"s3_bucket" cql:"s3_bucket"`
	Department      string            `json:"department" cql:"department"`
	DocumentType    string            `json:"document_type" cql:"document_type"`
	ReferenceNumber string            `json:"reference_number" cql:"reference_number"`
	CreatedAt       time.Time         `json:"created_at" cql:"created_at"`
	CreatedBy       string            `json:"created_by" cql:"created_by"`
	ModifiedAt      time.Time         `json:"modified_at" cql:"modified_at"`
	ModifiedBy      string            `json:"modified_by" cql:"modified_by"`
	CurrentVersion  int               `json:"current_version" cql:"current_version"`
	Status          string            `json:"status" cql:"status"`
	Tags            []string          `json:"tags" cql:"tags"`
	CustomMetadata  map[string]string `json:"custom_metadata" cql:"custom_metadata"`
	OCRStatus       string            `json:"ocr_status" cql:"ocr_status"`
	OCRConfidence   float64           `json:"ocr_confidence" cql:"ocr_confidence"`
	RetentionPolicy string            `json:"retention_policy" cql:"retention_policy"`
	ArchivalDate    time.Time         `json:"archival_date" cql:"archival_date"`
	IsDeleted       bool              `json:"is_deleted" cql:"is_deleted"`
	CheckedOutBy    string            `json:"checked_out_by,omitempty"`
	CheckedOutAt    time.Time         `json:"checked_out_at,omitempty"`
}

type DocumentVersion struct {
	DocumentID        uuid.UUID `json:"document_id" cql:"document_id"`
	VersionNumber     int       `json:"version_number" cql:"version_number"`
	S3Key             string    `json:"s3_key" cql:"s3_key"`
	CreatedAt         time.Time `json:"created_at" cql:"created_at"`
	CreatedBy         string    `json:"created_by" cql:"created_by"`
	ChangeDescription string    `json:"change_description" cql:"change_description"`
	FileSize          int64     `json:"file_size" cql:"file_size"`
	Checksum          string    `json:"checksum" cql:"checksum"`
}

type DocumentEvent struct {
	EventID    uuid.UUID         `json:"event_id"`
	EventType  string            `json:"event_type"`
	DocumentID uuid.UUID         `json:"document_id"`
	UserID     string            `json:"user_id"`
	Timestamp  time.Time         `json:"timestamp"`
	Metadata   map[string]string `json:"metadata"`
}

type UploadRequest struct {
	FileName        string            `json:"file_name" binding:"required"`
	FileType        string            `json:"file_type" binding:"required"`
	Department      string            `json:"department" binding:"required"`
	DocumentType    string            `json:"document_type" binding:"required"`
	ReferenceNumber string            `json:"reference_number"`
	Tags            []string          `json:"tags"`
	CustomMetadata  map[string]string `json:"custom_metadata"`
	CreatedBy       string            `json:"created_by" binding:"required"`
}

type UpdateRequest struct {
	FileName          string            `json:"file_name"`
	Department        string            `json:"department"`
	DocumentType      string            `json:"document_type"`
	ReferenceNumber   string            `json:"reference_number"`
	Tags              []string          `json:"tags"`
	CustomMetadata    map[string]string `json:"custom_metadata"`
	ModifiedBy        string            `json:"modified_by" binding:"required"`
	ChangeDescription string            `json:"change_description"`
}

type CheckoutRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type CheckinRequest struct {
	UserID            string `json:"user_id" binding:"required"`
	ChangeDescription string `json:"change_description"`
}
