package pgtest

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
)

var (
	instances       int64
	concurrent_conn *pgxpool.Pool
)

func StartupNewPool() (teardownFn func(), err error) {
	logger := logrus.New()
	logger.SetLevel(postgres.LogLevelWarn)

	log := logger.WithField("environment", "integration test")

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("on new pool: could not connect to docker: %w", err)
	}

	atomic.AddInt64(&instances, 1)
	ttName := fmt.Sprintf("db_%d_%d_test", atomic.LoadInt64(&instances), time.Now().UnixNano())
	dbName := strings.ToLower(ttName)

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14-alpine",
		Env:        []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres", "POSTGRES_DB=" + dbName},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		return nil, fmt.Errorf("could not start resource: %w", err)
	}

	port := resource.GetPort("5432/tcp")
	if err = pool.Retry(pingDatabase(port, log)); err != nil {
		return nil, fmt.Errorf("on wait living pool: could not connect to docker: %w", err)
	}

	// default max age is 120 seconds
	resource.Expire(120)

	dbURL := getPostgresConnString(port, "postgres")
	defaultPGPool, err := postgres.ConnectPoolWithoutMigrations(dbURL, log, postgres.LogLevelWarn)
	if err != nil {
		return nil, err
	}

	err = createDB(dbName, defaultPGPool)
	if err != nil {
		return nil, fmt.Errorf("on create database: %w", err)
	}

	// create default database
	dbPool, err := postgres.ConnectPoolWithMigrations(dbURL, log, postgres.LogLevelWarn)
	if err != nil {
		return nil, err
	}

	concurrent_conn = dbPool

	teardownFn = func() {
		dbPool.Close()

		dropDB(dbName, defaultPGPool)
		defaultPGPool.Close()

		resource.Close()
	}

	return teardownFn, nil
}

func NewDB(t *testing.T, dbName string) *pgxpool.Pool {
	t.Helper()

	dbName = strings.ToLower(dbName)
	_, _ = concurrent_conn.Exec(context.Background(), fmt.Sprintf("drop database %s", dbName))

	connString := strings.Replace(concurrent_conn.Config().ConnString(), concurrent_conn.Config().ConnConfig.Database, dbName, 1)
	newPool, err := pgxpool.Connect(context.Background(), connString)
	require.NoError(t, err)

	t.Cleanup(func() {
		newPool.Close()
		concurrent_conn.Exec(context.Background(), fmt.Sprintf("drop database %s", dbName))
	})

	return newPool
}

func pingDatabase(port string, log *logrus.Entry) func() error {
	return func() error {
		dbURL := getPostgresConnString(port, "postgres")
		connPool, err := postgres.ConnectPoolWithoutMigrations(dbURL, log, postgres.LogLevelWarn)
		if err != nil {
			return err
		}
		defer connPool.Close()

		return connPool.Ping(context.Background())
	}
}

func getPostgresConnString(port, dbName string) string {
	return fmt.Sprintf("postgres://postgres:postgres@localhost:%s/%s?sslmode=disable", port, dbName)
}

func dropDB(dbName string, pool *pgxpool.Pool) error {
	if _, err := pool.Exec(context.Background(), fmt.Sprintf(`DROP DATABASE IF EXISTS %s;`, dbName)); err != nil {
		return fmt.Errorf("on drop database %s: %w", dbName, err)
	}

	return nil
}

func createDB(dbName string, pool *pgxpool.Pool) error {
	_ = dropDB(dbName, pool)

	if _, err := pool.Exec(context.Background(), fmt.Sprintf(`CREATE DATABASE %s;`, dbName)); err != nil {
		return fmt.Errorf("on create database %s: %w", dbName, err)
	}

	return nil
}
