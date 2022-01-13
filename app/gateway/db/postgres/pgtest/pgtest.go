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
	instances       int64
	concurrent_conn *pgxpool.Pool
)

const (
	defaultDbName        = "pg-test-db-name"
	defaultContainerName = "pg-test-challenge-accounts"
)

func StartupNewPool() (teardownFn func(), err error) {
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

	resource, err := getDockerResource(pool, dbName)
	if err != nil {
		return nil, fmt.Errorf("on docker resource: %w", err)
	}

	port := resource.GetPort("5432/tcp")
	if err = pool.Retry(retryDbHelper(port, defaultDbName, log)); err != nil {
		return nil, fmt.Errorf("on wait living pool: could not connect to docker: %w", err)
	}

	dbDefaultURL := getPgConnURL(port, defaultDbName)
	dbDefaultPool, err := postgres.ConnectPoolWithoutMigrations(dbDefaultURL, log, postgres.LogLevelWarn)
	if err != nil {
		return nil, err
	}

	// create default database to unit test
	_, err = dbDefaultPool.Exec(context.Background(), fmt.Sprintf("create database %s", dbName))
	if err != nil {
		return nil, err
	}

	dbURL := getPgConnURL(port, dbName)
	dbPool, err := postgres.ConnectPoolWithMigrations(dbURL, log, postgres.LogLevelWarn)
	if err != nil {
		return nil, err
	}

	concurrent_conn = dbPool

	teardownFn = func() {
		pool.RemoveContainerByName(defaultContainerName)
		dropDB(dbName, dbPool)
		dbPool.Close()
		resource.Close()
	}

	return teardownFn, nil
}

func getDockerResource(pool *dockertest.Pool, dbName string) (*dockertest.Resource, error) {
	container, _ := pool.Client.InspectContainer(defaultContainerName)

	if container != nil {
		if container.State.Running {
			resource := &dockertest.Resource{Container: container}
			return resource, nil
		}

		if !container.State.Running {
			pool.RemoveContainerByName(defaultContainerName)
		}
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       defaultContainerName,
		Repository: "postgres",
		Tag:        "14-alpine",
		Env:        []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres", "POSTGRES_DB=" + defaultDbName},
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
	cconn := concurrent_conn

	_, err := concurrent_conn.Exec(
		context.Background(),
		fmt.Sprintf("create database %s", dbName),
	)
	require.NoError(t, err)

	orig := cconn.Config().ConnString()
	dbOrig := cconn.Config().ConnConfig.Database

	connString := strings.Replace(orig, dbOrig, dbName, 1)
	newPool, err := postgres.ConnectPoolWithMigrations(connString, log, postgres.LogLevelWarn)
	require.NoError(t, err)

	t.Cleanup(func() {
		newPool.Close()
		concurrent_conn.Exec(context.Background(), fmt.Sprintf("drop database %s", dbName))
	})

	return newPool
}

func retryDbHelper(port, dbName string, log *logrus.Entry) func() error {
	return func() error {
		dbURL := getPgConnURL(port, dbName)
		connPool, err := postgres.ConnectPoolWithoutMigrations(dbURL, log, postgres.LogLevelWarn)
		if err != nil {
			return err
		}
		defer connPool.Close()

		return connPool.Ping(context.Background())
	}
}

func getPgConnURL(port, dbName string) string {
	return fmt.Sprintf("postgres://postgres:postgres@localhost:%s/%s?sslmode=disable", port, dbName)
}

func dropDB(dbName string, pool *pgxpool.Pool) error {
	query := fmt.Sprintf(`DROP DATABASE IF EXISTS %s;`, dbName)

	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("on drop database %s: %w", dbName, err)
	}

	return nil
}
