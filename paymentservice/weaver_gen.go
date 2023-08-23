// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package paymentservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/onlineboutique/types/money"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/onlineboutique/paymentservice/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, chargeMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/paymentservice/T", Method: "Charge", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, chargeMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/paymentservice/T", Method: "Charge", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type t_local_stub struct {
	impl          T
	tracer        trace.Tracer
	chargeMetrics *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) Charge(ctx context.Context, a0 money.T, a1 CreditCardInfo) (r0 string, err error) {
	// Update metrics.
	begin := s.chargeMetrics.Begin()
	defer func() { s.chargeMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "paymentservice.T.Charge", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Charge(ctx, a0, a1)
}

// Client stub implementations.

type t_client_stub struct {
	stub          codegen.Stub
	chargeMetrics *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) Charge(ctx context.Context, a0 money.T, a1 CreditCardInfo) (r0 string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.chargeMetrics.Begin()
	defer func() { s.chargeMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "paymentservice.T.Charge", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	(a0).WeaverMarshal(enc)
	(a1).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.20.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// Check that t_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*t_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Charge":
		return s.charge
	default:
		return nil
	}
}

func (s t_server_stub) charge(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 money.T
	(&a0).WeaverUnmarshal(dec)
	var a1 CreditCardInfo
	(&a1).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Charge(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that t_reflect_stub implements the T interface.
var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) Charge(ctx context.Context, a0 money.T, a1 CreditCardInfo) (r0 string, err error) {
	err = s.caller("Charge", ctx, []any{a0, a1}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*CreditCardInfo)(nil)

type __is_CreditCardInfo[T ~struct {
	weaver.AutoMarshal
	Number          string
	CVV             int32
	ExpirationYear  int
	ExpirationMonth time.Month
}] struct{}

var _ __is_CreditCardInfo[CreditCardInfo]

func (x *CreditCardInfo) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("CreditCardInfo.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Number)
	enc.Int32(x.CVV)
	enc.Int(x.ExpirationYear)
	enc.Int((int)(x.ExpirationMonth))
}

func (x *CreditCardInfo) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("CreditCardInfo.WeaverUnmarshal: nil receiver"))
	}
	x.Number = dec.String()
	x.CVV = dec.Int32()
	x.ExpirationYear = dec.Int()
	*(*int)(&x.ExpirationMonth) = dec.Int()
}
