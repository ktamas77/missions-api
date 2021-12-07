#!/bin/bash

go fmt
go fmt ./internal/cli
go fmt ./internal/config
go fmt ./internal/data
go fmt ./internal/data/pg
go fmt ./internal/service
go fmt ./internal/service/checker-svc/checker
go fmt ./internal/service/handlers
go fmt ./internal/service/helpers
go fmt ./internal/service/requests
