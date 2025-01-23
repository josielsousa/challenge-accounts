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
	defaultDBName        = "pg-test-db-name"
	defaultContainerName = "pg-test-challenge-accounts"
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
	dbName := fmt.Sprintf("db_%d_%d", time.Now().UnixNano(), atomic.LoadInt32(&instances))

	dockerResource, err := getDockerResource(dockerPool)
	if err != nil {
		return nil, fmt.Errorf("on get docker resource: %w", err)
	}

	dbPort := dockerResource.GetPort("5432/tcp")

	if err = dockerPool.Retry(retryDBHelper(dbPort, defaultDBName)); err != nil {
		return nil, fmt.Errorf("on wait living pool: could not connect to docker: %w", err)
	}

	dbDefaultURL := getPgConnURL(dbPort, defaultDBName)
	ctx := context.Background()

	dbDefaultPool, err := postgres.ConnectPoolWithoutMigrations(ctx, dbDefaultURL)
	if err != nil {
		return nil, fmt.Errorf("on connect pool: %w", err)
	}

	// create default database to unit test
	_, err = dbDefaultPool.Exec(context.Background(), "create database "+dbName)
	if err != nil {
		return nil, fmt.Errorf("on create database %s: %w", dbName, err)
	}

	dbURL := getPgConnURL(dbPort, dbName)

	dbPool, err := postgres.ConnectPoolWithMigrations(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("on connect pool with migrations: %w", err)
	}

	concurrentConn = dbPool

	teardownFn := func() {
		dbPool.Close()

		dropDB(dbName, dbDefaultPool)

		dbDefaultPool.Close()
	}

	return teardownFn, nil
}

func getDockerResource(pool *dockertest.Pool) (*dockertest.Resource, error) {
	container, _ := pool.Client.InspectContainer(defaultContainerName)

	if container != nil {
		if container.State.Running {
			resource := &dockertest.Resource{Container: container}

			return resource, nil
		}

		if !container.State.Running {
			if err := pool.RemoveContainerByName(defaultContainerName); err != nil {
				return nil, fmt.Errorf("could not remove container: %w", err)
			}
		}
	}

	//nolint:exhaustruct
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       defaultContainerName,
		Repository: "postgres",
		Tag:        "16-alpine",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=" + defaultDBName,
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
	cconn := concurrentConn

	_, err := concurrentConn.Exec(ctx, "create database "+dbName)
	require.NoError(t, err)

	orig := cconn.Config().ConnString()
	dbOrig := cconn.Config().ConnConfig.Database

	connString := strings.Replace(orig, dbOrig, dbName, 1)
	newPool, err := postgres.ConnectPoolWithMigrations(ctx, connString)
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
		dbURL := getPgConnURL(port, dbName)

		connPool, err := postgres.ConnectPoolWithoutMigrations(context.Background(), dbURL)
		if err != nil {
			return fmt.Errorf("on connect pool: %w", err)
		}

		defer connPool.Close()

		return connPool.Ping(context.Background())
	}
}

func getPgConnURL(port, dbName string) string {
	return fmt.Sprintf("postgres://postgres:postgres@localhost:%s/%s?sslmode=disable", port, dbName)
}

func dropDB(dbName string, pool *pgxpool.Pool) {
	if dbName == defaultDBName {
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
