package service

import (
    "net"
    "net/http"

    "gitlab.com/distributed_lab/logan/v3/errors"
    "gitlab.com/distributed_lab/kit/copus/types"
    "githab.com/redcuckoo/bsc-checker-events/internal/config"
    "gitlab.com/distributed_lab/logan/v3"
)

type service struct {
    log      *logan.Entry
    copus    types.Copus
    listener net.Listener
}

func (s *service) run() error {
    // TODO implement custom logic here
    r := s.router()

    if err := s.copus.RegisterChi(r); err != nil {
        return errors.Wrap(err, "cop failed")
    }


    return http.Serve(s.listener, r)
    return nil
}

func newService(cfg config.Config) *service {
    return &service{
        log:        cfg.Log(),
        copus:      cfg.Copus(),
        listener:   cfg.Listener(),
    }
}

func Run(cfg config.Config) {
    if err := newService(cfg).run(); err != nil {
        panic(err)
    }
}
