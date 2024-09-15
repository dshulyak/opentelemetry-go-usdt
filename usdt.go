package usdt

/*
#include <stdlib.h>
#include "./usdt/bindings.h"
#cgo LDFLAGS: ${SRCDIR}/usdt/libusdt.a
*/
import "C"

import (
	"context"
	"encoding/binary"
	"unsafe"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func New() *USDTTracer {
	return &USDTTracer{}
}

var _ sdktrace.SpanProcessor = (*USDTTracer)(nil)

type USDTTracer struct{}

func (u *USDTTracer) OnStart(parent context.Context, s sdktrace.ReadWriteSpan) {
	spanId := s.SpanContext().SpanID()
	parentId := s.Parent().SpanID()
	var id uint64 = 0
	var amount uint64 = 0
	name := C.CString(s.Name())
	defer C.free(unsafe.Pointer(name))
	C.stacks_enter(C.ulong(binary.BigEndian.Uint64(spanId[:])),
		C.ulong(binary.BigEndian.Uint64(parentId[:])), C.ulong(id), C.ulong(amount), name)
}

func (u *USDTTracer) OnEnd(s sdktrace.ReadOnlySpan) {
	spanId := s.SpanContext().SpanID()
	C.stacks_close(C.ulong(binary.BigEndian.Uint64(spanId[:])))
}

func (u *USDTTracer) ForceFlush(ctx context.Context) error {
	return nil
}

func (u *USDTTracer) Shutdown(ctx context.Context) error {
	return nil
}
