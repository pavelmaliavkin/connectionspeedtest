package ookla

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

func getClient() IApiClient {
	return &client{}
}

type IApiClient interface {
	DownloadTest() (float64, error)
	UploadTest() (float64, error)
}

type client struct {
}

func (a *client) getServer() (*speedtest.Server, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, fmt.Errorf(`user info cannot be fetched, err: %w`, err)
	}
	servers, err := speedtest.FetchServers(user)
	if err != nil {
		return nil, fmt.Errorf(`servers cannot be fetched, err: %w`, err)
	}
	if len(servers) == 0 {
		return nil, fmt.Errorf(`servers list is empty`)
	}
	return servers[0], nil
}

func (a *client) DownloadTest() (float64, error) {
	server, err := a.getServer()
	if err != nil {
		return 0, fmt.Errorf(`cannot fetch server to run test, err: %w`, err)
	}
	if err := server.DownloadTest(false); err != nil {
		return 0, fmt.Errorf(`download test failed, err: %w`, err)
	}
	return server.DLSpeed, nil
}

func (a *client) UploadTest() (float64, error) {
	server, err := a.getServer()
	if err != nil {
		return 0, fmt.Errorf(`cannot fetch server to run test, err: %w`, err)
	}
	if err := server.UploadTest(false); err != nil {
		return 0, fmt.Errorf(`upload test failed, err: %w`, err)
	}
	return server.DLSpeed, nil
}
