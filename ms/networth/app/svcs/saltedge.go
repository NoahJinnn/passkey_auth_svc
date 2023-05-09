package svcs

import "github.com/hellohq/hqservice/ms/networth/config"

type ISeSvc interface {
}

type seSvc struct {
	cfg *config.Config
}
