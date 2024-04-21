package client_test

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"
	"time"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/client"
	"github.com/alexfalkowski/bezeichner/cmd"
	"github.com/alexfalkowski/bezeichner/config"
	"github.com/alexfalkowski/bezeichner/transport"
	"github.com/alexfalkowski/go-service/compressor"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	. "github.com/smartystreets/goconvey/convey" //nolint:revive
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"google.golang.org/grpc"
)

var options = []fx.Option{
	fx.NopLogger, cmd.Module, debug.Module,
	config.Module, transport.Module,
	compressor.Module, marshaller.Module,
	telemetry.Module, metrics.Module,
	client.Module, fx.Invoke(register),
}

func TestValidSetup(t *testing.T) {
	Convey("Given I have a app", t, func() {
		So(os.Setenv("CONFIG_FILE", "../test/.config/client.yml"), ShouldBeNil)

		l, err := net.Listen("tcp", "localhost:12000")
		So(err, ShouldBeNil)

		server := grpc.NewServer()

		go server.Serve(l) //nolint:errcheck

		time.Sleep(time.Second)

		app := fxtest.New(t, options...)

		Convey("When I start the app", func() {
			app.RequireStart()

			Convey("Then I should have a started app", func() {
				app.RequireStop()
			})
		})

		server.Stop()

		So(os.Unsetenv("CONFIG_FILE"), ShouldBeNil)
	})
}

func register(_ *client.Client) {}

func TestValidGenerateIdentifiers(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		v := &validClient{}
		c := client.NewClient(v)

		Convey("When I generate identifiers", func() {
			t, err := c.GenerateIdentifiers(context.Background(), "test", 1)
			So(err, ShouldBeNil)

			Convey("Then I should have a identifiers", func() {
				So(t, ShouldResemble, []string{"test"})
			})
		})
	})
}

func TestInvalidGenerateIdentifiers(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		v := &invalidClient{}
		c := client.NewClient(v)

		Convey("When I generate identifiers", func() {
			_, err := c.GenerateIdentifiers(context.Background(), "test", 1)

			Convey("Then I should have an error", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestValidMapIdentifiers(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		v := &validClient{}
		c := client.NewClient(v)

		Convey("When I map identifiers", func() {
			t, err := c.MapIdentifiers(context.Background(), []string{"test"})
			So(err, ShouldBeNil)

			Convey("Then I should have a valid identifiers", func() {
				So(t, ShouldResemble, []string{"test"})
			})
		})
	})
}

func TestInvalidMapIdentifiers(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		v := &invalidClient{}
		c := client.NewClient(v)

		Convey("When I map identifiers", func() {
			_, err := c.MapIdentifiers(context.Background(), []string{"test"})

			Convey("Then I should have an error", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

type validClient struct{}

func (*validClient) GenerateIdentifiers(_ context.Context, _ *v1.GenerateIdentifiersRequest, _ ...grpc.CallOption) (*v1.GenerateIdentifiersResponse, error) {
	return &v1.GenerateIdentifiersResponse{Ids: []string{"test"}}, nil
}

func (*validClient) MapIdentifiers(_ context.Context, _ *v1.MapIdentifiersRequest, _ ...grpc.CallOption) (*v1.MapIdentifiersResponse, error) {
	return &v1.MapIdentifiersResponse{Ids: []string{"test"}}, nil
}

type invalidClient struct{}

func (*invalidClient) GenerateIdentifiers(_ context.Context, _ *v1.GenerateIdentifiersRequest, _ ...grpc.CallOption) (*v1.GenerateIdentifiersResponse, error) {
	return &v1.GenerateIdentifiersResponse{}, errors.New("invalid")
}

func (*invalidClient) MapIdentifiers(_ context.Context, _ *v1.MapIdentifiersRequest, _ ...grpc.CallOption) (*v1.MapIdentifiersResponse, error) {
	return &v1.MapIdentifiersResponse{}, errors.New("invalid")
}
