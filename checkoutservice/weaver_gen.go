// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package checkoutservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/onlineboutique/paymentservice"
	"github.com/ServiceWeaver/onlineboutique/shippingservice"
	"github.com/ServiceWeaver/onlineboutique/types"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/onlineboutique/checkoutservice/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, placeOrderMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/checkoutservice/T", Method: "PlaceOrder", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, placeOrderMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/checkoutservice/T", Method: "PlaceOrder", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "⟦bd41e485:wEaVeReDgE:github.com/ServiceWeaver/onlineboutique/checkoutservice/T→github.com/ServiceWeaver/onlineboutique/productcatalogservice/T⟧\n⟦d54c5b5d:wEaVeReDgE:github.com/ServiceWeaver/onlineboutique/checkoutservice/T→github.com/ServiceWeaver/onlineboutique/cartservice/T⟧\n⟦fc39f6a9:wEaVeReDgE:github.com/ServiceWeaver/onlineboutique/checkoutservice/T→github.com/ServiceWeaver/onlineboutique/currencyservice/T⟧\n⟦06129fa8:wEaVeReDgE:github.com/ServiceWeaver/onlineboutique/checkoutservice/T→github.com/ServiceWeaver/onlineboutique/shippingservice/T⟧\n⟦176aaf8f:wEaVeReDgE:github.com/ServiceWeaver/onlineboutique/checkoutservice/T→github.com/ServiceWeaver/onlineboutique/emailservice/T⟧\n⟦b148055f:wEaVeReDgE:github.com/ServiceWeaver/onlineboutique/checkoutservice/T→github.com/ServiceWeaver/onlineboutique/paymentservice/T⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type t_local_stub struct {
	impl              T
	tracer            trace.Tracer
	placeOrderMetrics *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) PlaceOrder(ctx context.Context, a0 PlaceOrderRequest) (r0 types.Order, err error) {
	// Update metrics.
	begin := s.placeOrderMetrics.Begin()
	defer func() { s.placeOrderMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "checkoutservice.T.PlaceOrder", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.PlaceOrder(ctx, a0)
}

// Client stub implementations.

type t_client_stub struct {
	stub              codegen.Stub
	placeOrderMetrics *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) PlaceOrder(ctx context.Context, a0 PlaceOrderRequest) (r0 types.Order, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.placeOrderMetrics.Begin()
	defer func() { s.placeOrderMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "checkoutservice.T.PlaceOrder", trace.WithSpanKind(trace.SpanKindClient))
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
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.22.0 (codegen
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
	case "PlaceOrder":
		return s.placeOrder
	default:
		return nil
	}
}

func (s t_server_stub) placeOrder(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 PlaceOrderRequest
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.PlaceOrder(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that t_reflect_stub implements the T interface.
var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) PlaceOrder(ctx context.Context, a0 PlaceOrderRequest) (r0 types.Order, err error) {
	err = s.caller("PlaceOrder", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*PlaceOrderRequest)(nil)

type __is_PlaceOrderRequest[T ~struct {
	weaver.AutoMarshal
	UserID       string
	UserCurrency string
	Address      shippingservice.Address
	Email        string
	CreditCard   paymentservice.CreditCardInfo
}] struct{}

var _ __is_PlaceOrderRequest[PlaceOrderRequest]

func (x *PlaceOrderRequest) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("PlaceOrderRequest.WeaverMarshal: nil receiver"))
	}
	enc.String(x.UserID)
	enc.String(x.UserCurrency)
	(x.Address).WeaverMarshal(enc)
	enc.String(x.Email)
	(x.CreditCard).WeaverMarshal(enc)
}

func (x *PlaceOrderRequest) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("PlaceOrderRequest.WeaverUnmarshal: nil receiver"))
	}
	x.UserID = dec.String()
	x.UserCurrency = dec.String()
	(&x.Address).WeaverUnmarshal(dec)
	x.Email = dec.String()
	(&x.CreditCard).WeaverUnmarshal(dec)
}
