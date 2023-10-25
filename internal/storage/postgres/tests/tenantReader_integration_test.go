//go:build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Permify/permify/internal/storage"
	"github.com/Permify/permify/internal/storage/postgres"
	"github.com/Permify/permify/pkg/database"
	PQDatabase "github.com/Permify/permify/pkg/database/postgres"
)

func TestTenantReader_Integration(t *testing.T) {
	ctx := context.Background()

	err := storage.Migrate(cfg)
	require.NoError(t, err)

	var db database.Database
	db, err = PQDatabase.New(cfg.URI,
		PQDatabase.MaxOpenConnections(cfg.MaxOpenConnections),
		PQDatabase.MaxIdleConnections(cfg.MaxIdleConnections),
		PQDatabase.MaxConnectionIdleTime(cfg.MaxConnectionIdleTime),
		PQDatabase.MaxConnectionLifeTime(cfg.MaxConnectionLifetime),
	)
	require.NoError(t, err)

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Create a Tenant instances
	tenantWriter := postgres.NewTenantWriter(db.(*PQDatabase.Postgres))
	tenantReader := postgres.NewTenantReader(db.(*PQDatabase.Postgres))

	// Test the CreateTenant method
	createdTenant, err := tenantWriter.CreateTenant(ctx, "2", "Test Tenant")
	require.NoError(t, err)
	assert.Equal(t, "2", createdTenant.Id)
	assert.Equal(t, "Test Tenant", createdTenant.Name)

	pagination := database.NewPagination()

	// Test the DeleteTenant method
	listTenant, _, err := tenantReader.ListTenants(ctx, pagination)

	require.NoError(t, err)
	assert.Equal(t, "t1", listTenant[1].Id)
	assert.Equal(t, "example tenant", listTenant[1].Name)
	assert.Equal(t, "2", listTenant[0].Id)
	assert.Equal(t, "Test Tenant", listTenant[0].Name)
}