package pgtest

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
)

var (
	instances      int32
	concurrentConn *pgxpool.Pool
)

const (
	baseDBName   = "pg-test-db-name"
	baseContName = "pg-test-challenge-accounts"
)

type DockerContainerConfig struct {
	DBName string
}

func StartupNewPool() (func(), error) {
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf(`could not connect to docker: %w`, err)
	}

	if err = dockerPool.Client.Ping(); err != nil {
		return nil, fmt.Errorf(`could not connect to docker: %w`, err)
	}

	atomic.AddInt32(&instances, 1)
	baseInstaceName := fmt.Sprintf("%d_%d", time.Now().UnixNano(), atomic.LoadInt32(&instances))

	dbName := "db_" + baseInstaceName

	dockerResource, err := getDockerResource(baseInstaceName, dockerPool)
	if err != nil {
		return nil, fmt.Errorf("on get docker resource: %w", err)
	}

	dbPort := dockerResource.GetPort("5432/tcp")

	if err = dockerPool.Retry(retryDBHelper(dbPort, baseDBName)); err != nil {
		return nil, fmt.Errorf("on wait living pool: could not connect to docker: %w", err)
	}

	defaultConnConfig := getConnConfig(dbPort, baseDBName)
	ctx := context.Background()

	dbDefaultPool, err := postgres.NewPool(ctx, postgres.WithConnString(
		defaultConnConfig,
	))
	if err != nil {
		return nil, fmt.Errorf("on connect pool: %w", err)
	}

	// create default database to unit test
	_, err = dbDefaultPool.Exec(context.Background(), "create database "+dbName)
	if err != nil {
		return nil, fmt.Errorf("on create database %s: %w", dbName, err)
	}

	dbConnConfig := getConnConfig(dbPort, dbName)

	dbPool, err := postgres.NewPool(ctx, postgres.WithConnString(dbConnConfig))
	if err != nil {
		return nil, fmt.Errorf("on connect pool with migrations: %w", err)
	}

	concurrentConn = dbPool

	teardownFn := func() {
		dbPool.Close()

		dropDB(dbName, dbDefaultPool)

		dbDefaultPool.Close()

		_ = dockerPool.Purge(dockerResource)
	}

	return teardownFn, nil
}

func getDockerResource(baseInstaceName string, pool *dockertest.Pool) (*dockertest.Resource, error) {
	//nolint:exhaustruct
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       baseContName + "_" + baseInstaceName,
		Repository: "postgres",
		Tag:        "16-alpine",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=" + baseDBName,
		},
	}, func(c *docker.HostConfig) {
		c.AutoRemove = true
		c.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return nil, fmt.Errorf("could not start resource: %w", err)
	}

	return resource, nil
}

func NewDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	ctx := context.Background()

	atomic.AddInt32(&instances, 1)
	dbName := fmt.Sprintf("db_%d_%d_test", atomic.LoadInt32(&instances), time.Now().UnixNano())

	_, err := concurrentConn.Exec(ctx, "create database "+dbName)
	require.NoError(t, err)

	orig := concurrentConn.Config().ConnString()
	dbOrig := concurrentConn.Config().ConnConfig.Database

	connString := strings.Replace(orig, dbOrig, dbName, 1)

	newPool, err := postgres.NewPool(ctx, postgres.WithConnString(connString))

	require.NoError(t, err)

	t.Cleanup(func() {
		newPool.Close()

		_, err := concurrentConn.Exec(context.Background(), "drop database "+dbName)
		require.NoError(t, err)
	})

	return newPool
}

func retryDBHelper(port, dbName string) func() error {
	return func() error {
		dbConnConfig := getConnConfig(port, dbName)

		connPool, err := postgres.NewPool(
			context.Background(),
			postgres.WithConnString(dbConnConfig),
		)
		if err != nil {
			return fmt.Errorf("on connect pool: %w", err)
		}

		defer connPool.Close()

		return connPool.Ping(context.Background())
	}
}

func getConnConfig(port, dbName string) string {
	return fmt.Sprintf(
		"user=postgres password=postgres host=localhost port=%s dbname=%s sslmode=disable",
		port,
		dbName,
	)
}

func dropDB(dbName string, pool *pgxpool.Pool) {
	if dbName == baseDBName {
		return
	}

	query := fmt.Sprintf(`DROP DATABASE IF EXISTS %s;`, dbName)

	if _, err := pool.Exec(context.Background(), query); err != nil {
		slog.Error(
			"on drop database",
			slog.String("dbName", dbName),
			slog.Any("error", err),
		)
	}
}
