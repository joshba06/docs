package server

import (
	"context"
	"log"
	"sync"

	"github.com/gofrs/uuid"
	documentsv1 "github.com/joshba06/docs/proto/documents/v1"
	usersv1 "github.com/joshba06/docs/proto/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Backend implements the protobuf interface
type Backend struct {
	mu        *sync.RWMutex
	users     []*usersv1.User
	documents []*documentsv1.Document
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

// AddUser adds a user to the in-memory store.
func (b *Backend) AddUser(ctx context.Context, req *usersv1.AddUserRequest) (*usersv1.AddUserResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	user := &usersv1.User{
		Id:    uuid.Must(uuid.NewV4()).String(),
		Email: req.GetEmail(),
	}
	log.Println("New user: ", user)
	b.users = append(b.users, user)

	return &usersv1.AddUserResponse{
		User: user,
	}, nil
}

// AddDocument adds a document to the in-memory store.
func (b *Backend) AddDocument(ctx context.Context, req *documentsv1.AddDocumentRequest) (*documentsv1.Document, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	document := &documentsv1.Document{
		Id:      uuid.Must(uuid.NewV4()).String(),
		Content: req.GetContent(),
	}
	log.Println("New document: ", document)
	b.documents = append(b.documents, document)

	return document, nil
}

// GetDocument gets a document in the store
func (b *Backend) GetDocument(ctx context.Context, req *documentsv1.GetDocumentRequest) (*documentsv1.Document, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, document := range b.documents {
		if document.GetId() == req.GetDocumentId() {
			return document, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "document %q could not be found", req.GetDocumentId())
}

// ListDocuments streams a list of documents in the store
func (b *Backend) ListDocuments(_ *documentsv1.ListDocumentsRequest, srv documentsv1.DocumentsService_ListDocumentsServer) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, document := range b.documents {
		err := srv.Send(document)
		if err != nil {
			return err
		}
	}
	return nil
}
