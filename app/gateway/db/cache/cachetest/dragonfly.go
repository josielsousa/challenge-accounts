package cachetest

import (
	"context"
	"fmt"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/redis/go-redis/v9"
)

const _defaulResourceExpiry = 60

type Client struct {
	Client   *redis.Client
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func NewClient() (*Client, error) {
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf(`could create pool docker: %w`, err)
	}

	if err = dockerPool.Client.Ping(); err != nil {
		return nil, fmt.Errorf(`could not connect to docker: %w`, err)
	}

	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "docker.dragonflydb.io/dragonflydb/dragonfly",
		Tag:        "latest",
	}, func(hostCfg *docker.HostConfig) {
		hostCfg.AutoRemove = true
		hostCfg.RestartPolicy = docker.NeverRestart()
	})
	if err != nil {
		return nil, fmt.Errorf(`could not start resource: %w`, err)
	}

	if err := resource.Expire(_defaulResourceExpiry); err != nil {
		return nil, fmt.Errorf(`could not set expiry resource: %w`, err)
	}

	cli := &Client{
		pool:     dockerPool,
		resource: resource,
	}

	err = dockerPool.Retry(func() error {
		client := redis.NewClient(&redis.Options{
			Addr: resource.GetHostPort("6379/tcp"),
		})

		if err := client.Ping(context.Background()).Err(); err != nil {
			return fmt.Errorf(`could not ping redis: %w`, err)
		}

		cli.Client = client

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf(`could not retry to redis: %w`, err)
	}

	return cli, nil
}

func (c *Client) Teardown() error {
	if err := c.Client.Close(); err != nil {
		return fmt.Errorf(`could not close client: %w`, err)
	}

	if err := c.pool.Purge(c.resource); err != nil {
		return fmt.Errorf(`could not purge resource: %w`, err)
	}

	return nil
}
