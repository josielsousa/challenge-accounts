package pgtest

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres"
)

var (
	instances      int64
	concurrentConn *pgxpool.Pool
)

const (
	defaultDBName        = "pg-test-db-name"
	defaultContainerName = "pg-test-challenge-accounts"
)

func StartupNewPool() (func(), error) {
	logger := logrus.New()
	logger.SetLevel(postgres.LogLevelWarn)

	log := logger.WithField("environment", "integration test")

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("on new pool: could not connect to docker: %w", err)
	}

	if err = pool.Client.Ping(); err != nil {
		return nil, fmt.Errorf("on ping docker client: %w", err)
	}

	atomic.AddInt64(&instances, 1)
	ttName := fmt.Sprintf("db_%d_%d_test", atomic.LoadInt64(&instances), time.Now().UnixNano())
	dbName := strings.ToLower(ttName)

	resource, err := getDockerResource(pool)
	if err != nil {
		return nil, fmt.Errorf("on docker resource: %w", err)
	}

	port := resource.GetPort("5432/tcp")
	if err = pool.Retry(retryDBHelper(port, defaultDBName, log)); err != nil {
		return nil, fmt.Errorf("on wait living pool: could not connect to docker: %w", err)
	}

	dbDefaultURL := getPgConnURL(port, defaultDBName)

	dbDefaultPool, err := postgres.ConnectPoolWithoutMigrations(dbDefaultURL, log, postgres.LogLevelWarn)
	if err != nil {
		return nil, fmt.Errorf("on connect pool: %w", err)
	}

	// create default database to unit test
	_, err = dbDefaultPool.Exec(context.Background(), "create database "+dbName)
	if err != nil {
		return nil, fmt.Errorf("on create database %s: %w", dbName, err)
	}

	dbURL := getPgConnURL(port, dbName)

	dbPool, err := postgres.ConnectPoolWithMigrations(dbURL, log, postgres.LogLevelWarn)
	if err != nil {
		return nil, fmt.Errorf("on connect pool with migrations: %w", err)
	}

	concurrentConn = dbPool

	teardownFn := func() {
		if err := pool.RemoveContainerByName(defaultContainerName); err != nil {
			panic(fmt.Errorf("could not remove container: %w", err))
		}

		dropDB(dbName, dbPool)
		dbPool.Close()
		resource.Close()
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
		Tag:        "14-alpine",
		Env:        []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres", "POSTGRES_DB=" + defaultDBName},
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

	logger := logrus.New()
	logger.SetLevel(postgres.LogLevelWarn)
	log := logger.WithField("environment", "new db integration test")

	atomic.AddInt64(&instances, 1)
	dbName := fmt.Sprintf("db_%d_%d_test", atomic.LoadInt64(&instances), time.Now().UnixNano())
	cconn := concurrentConn

	_, err := concurrentConn.Exec(
		context.Background(),
		"create database "+dbName,
	)
	require.NoError(t, err)

	orig := cconn.Config().ConnString()
	dbOrig := cconn.Config().ConnConfig.Database

	connString := strings.Replace(orig, dbOrig, dbName, 1)
	newPool, err := postgres.ConnectPoolWithMigrations(connString, log, postgres.LogLevelWarn)
	require.NoError(t, err)

	t.Cleanup(func() {
		newPool.Close()

		_, err := concurrentConn.Exec(context.Background(), "drop database "+dbName)
		require.NoError(t, err)
	})

	return newPool
}

func retryDBHelper(port, dbName string, log *logrus.Entry) func() error {
	return func() error {
		dbURL := getPgConnURL(port, dbName)

		connPool, err := postgres.ConnectPoolWithoutMigrations(dbURL, log, postgres.LogLevelWarn)
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
	query := fmt.Sprintf(`DROP DATABASE IF EXISTS %s;`, dbName)

	if _, err := pool.Exec(context.Background(), query); err != nil {
		panic(fmt.Errorf("on drop database %s: %w", dbName, err))
	}
}
