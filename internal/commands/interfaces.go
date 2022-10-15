package commands

import (
	"context"

	"github.com/Permify/permify/internal/repositories"
	base "github.com/Permify/permify/pkg/pb/base/v1"
)

type ICommand interface {
	GetRelationTupleRepository() repositories.IRelationTupleRepository
}

// ICheckCommand -
type ICheckCommand interface {
	Execute(ctx context.Context, q *CheckQuery, child *base.Child) (response CheckResponse, err error)
}

// IExpandCommand -
type IExpandCommand interface {
	Execute(ctx context.Context, q *ExpandQuery, child *base.Child) (response ExpandResponse, err error)
}

// ISchemaLookupCommand -
type ISchemaLookupCommand interface {
	Execute(ctx context.Context, q *SchemaLookupQuery, actions map[string]*base.ActionDefinition) (response SchemaLookupResponse, err error)
}

// ILookupQueryCommand -
type ILookupQueryCommand interface {
	Execute(ctx context.Context, q *LookupQueryQuery, child *base.Child) (response LookupQueryResponse, err error)
}