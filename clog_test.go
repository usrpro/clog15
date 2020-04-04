// Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
// Use of this source code is governed by a License that can be found in the LICENSE file.
// SPDX-License-Identifier: BSD-3-Clause

package clog15

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"testing"

	log "github.com/inconshreveable/log15"
)

func TestSetLogger(t *testing.T) {
	type args struct {
		ctx  context.Context
		l    log.Logger
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"No args",
			args{
				context.Background(),
				log.New(),
				nil,
			},
		},
		{
			"With args",
			args{
				context.Background(),
				log.New(),
				[]interface{}{"foo", "bar"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := SetLogger(tt.args.ctx, tt.args.l, tt.args.args...)

			want := tt.args.l
			if len(tt.args.args) > 0 {
				want = tt.args.l.New(tt.args.args...)
			}

			if got := GetLogger(ctx); !reflect.DeepEqual(got, want) {
				t.Errorf("GetLogger() = %v, want %v", got, want)
			}
		})
	}
}

func TestNewLogger(t *testing.T) {
	ctx := NewLogger(context.Background(), "foo", "bar")
	if _, ok := ctx.Value(CtxLogger).(log.Logger); !ok {
		t.Error("NewLogger: logger not in context")
	}
}

func TestGetLogger(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want log.Logger
	}{
		{
			"No logger",
			context.Background(),
			log.New(),
		},
		{
			"Wrong type for entry",
			context.WithValue(context.Background(), CtxLogger, 22),
			log.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogger(tt.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddArgs(t *testing.T) {
	ctx := AddArgs(context.Background(), "foo", "bar")
	if _, ok := ctx.Value(CtxLogger).(log.Logger); !ok {
		t.Error("AddArgs: logger not in context")
	}
}

// testHandler implements log.Handler
type testHandler struct {
	bytes.Buffer
}

func (h *testHandler) Log(r *log.Record) error {
	if _, err := fmt.Fprintf(h, "lvl=%s msg=%s", r.Lvl, r.Msg); err != nil {
		return err
	}
	for i := 0; i < len(r.Ctx); i += 2 {
		if _, err := fmt.Fprintf(h, " %v=%v", r.Ctx[i], r.Ctx[i+1]); err != nil {
			return err
		}
	}
	return nil
}

var logOutput testHandler

func init() {
	log.Root().SetHandler(&logOutput)
}

func TestOutput(t *testing.T) {
	ctx := context.WithValue(context.Background(), CtxLogger, log.New())

	tests := []struct {
		level string
		fn    func(context.Context, string, ...interface{})
	}{
		{
			"dbug",
			Debug,
		},
		{
			"info",
			Info,
		},
		{
			"warn",
			Warn,
		},
		{
			"eror",
			Error,
		},
		{
			"crit",
			Crit,
		},
	}

	for _, tt := range tests {
		logOutput.Reset()
		want := fmt.Sprintf("lvl=%s msg=hello foo=bar", tt.level)

		tt.fn(ctx, "hello", "foo", "bar")

		if got := logOutput.String(); got != want {
			t.Errorf("Level %s got:\n%s\nwant:\n%s", tt.level, got, want)
		}
	}
}
