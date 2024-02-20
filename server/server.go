package server

import (
	"context"
	"log"
	"sync"

	"github.com/gofrs/uuid"
	invoicesv1 "github.com/joshba06/docs/proto/invoices/v1"
	usersv1 "github.com/joshba06/docs/proto/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var invoices = []*invoicesv1.Invoice {
	{
		Id: "INV001",
		Content: "Product A - $50",
	},
	{
		Id: "INV002",
		Content: "Service B - $100",
	},
	{
		Id: "INV003",
		Content: "Product C - $75",
	},
};


// Backend implements the protobuf interface
type Backend struct {
	mu       *sync.RWMutex
	users    []*usersv1.User
	invoices []*invoicesv1.Invoice
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

// GetInvoice gets all invoices in the store
func (b *Backend) GetInvoice(ctx context.Context, req *invoicesv1.GetInvoiceRequest) (*invoicesv1.Invoice, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, invoice := range b.invoices {
		if invoice.GetId() == req.GetInvoiceId() {
			return invoice, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "invoice %q could not be found", req.GetInvoiceId())
}

// Addinvoice adds an invoice to the in-memory store.
func (b *Backend) AddInvoice(ctx context.Context, req *invoicesv1.AddInvoiceRequest) (*invoicesv1.AddInvoiceResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	log.Println("New invoice: ", req)

	invoice := &invoicesv1.Invoice{
		Id:      uuid.Must(uuid.NewV4()).String(),
		Content: req.GetContent(),
	}
	log.Println("New invoice: ", invoice)
	b.invoices = append(b.invoices, invoice)

	return &invoicesv1.AddInvoiceResponse{
		Invoice: invoice,
	}, nil
}
