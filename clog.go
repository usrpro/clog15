// Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
// Use of this source code is governed by a License that can be found in the LICENSE file.
// SPDX-License-Identifier: BSD-3-Clause

/*
Package clog15 provides utilities to embed and extract a log15.Logger to a context.
This might be helpfull to preserve logger context while being restricted by funtction signatures.
For instance in http.HanderFunc, middleware or gRPC interceptors.
It allows you to define logging context and attach the configured logger to a context
passed down the executions chain.
*/
package clog15

import (
	"context"

	log "github.com/inconshreveable/log15"
)

type logKey int

// CtxLogger is the context Value key which must be used.
const CtxLogger logKey = 1

// SetLogger embeds the logger into the returned context.
// args are added to the logger's context.
func SetLogger(ctx context.Context, l log.Logger, args ...interface{}) context.Context {
	if len(args) > 0 {
		l = l.New(args...)
	}

	return context.WithValue(ctx, CtxLogger, l)
}

// NewLogger creates a new logger from the Root logger.
// It embeds the logger into the returned context.
// args are added to the logger's context.
func NewLogger(ctx context.Context, args ...interface{}) context.Context {
	return context.WithValue(ctx, CtxLogger, log.New(args...))
}

// GetLogger retrieves a logger from the context.
// If there is no logger set to context,
// a new root logger is returned.
func GetLogger(ctx context.Context) log.Logger {
	l, ok := ctx.Value(CtxLogger).(log.Logger)
	if !ok {
		return log.New()
	}

	return l
}

// Debug is a wrapper for GetLogger().Debug()
func Debug(ctx context.Context, msg string, args ...interface{}) {
	GetLogger(ctx).Debug(msg, args...)
}

// Info is a wrapper for GetLogger().Info()
func Info(ctx context.Context, msg string, args ...interface{}) {
	GetLogger(ctx).Info(msg, args...)
}

// Warn is a wrapper for GetLogger().Warn()
func Warn(ctx context.Context, msg string, args ...interface{}) {
	GetLogger(ctx).Warn(msg, args...)
}

// Error is a wrapper for GetLogger().Error()
func Error(ctx context.Context, msg string, args ...interface{}) {
	GetLogger(ctx).Error(msg, args...)
}

// Crit is a wrapper for GetLogger().Crit()
func Crit(ctx context.Context, msg string, args ...interface{}) {
	GetLogger(ctx).Crit(msg, args...)
}
