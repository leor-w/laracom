package main

import (
	"context"
	"os"

	pb "github.com/leor-w/laracom/product-service/proto/product"
	"github.com/leor-w/laracom/product-service/trace"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	tracePlugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func main() {
	t, io, err := trace.NewTracer("laracom.client.product", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		logrus.Fatalf("Init Tracer failed: err = %v", err)
	}
	defer io.Close()

	service := micro.NewService(
		micro.Name("laracom.client.product"),
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "brand name",
			},
		),
		micro.WrapClient(tracePlugin.NewClientWrapper(t)),
	)
	client := pb.NewBrandService("laracom.service.product", service.Client())

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "call")
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	defer span.Finish()

	opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	ctx = opentracing.ContextWithSpan(ctx, span)
	ctx = metadata.NewContext(ctx, md)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			logrus.Infof("Request brand : name %s", name)
			req := &pb.Brand{
				Name: name,
			}
			span.SetTag("req", req)
			r, err := client.Create(context.TODO(), req)
			if err != nil {
				logrus.Errorf("Create brand failed : err = %v", err)
				return err
			}

			span.SetTag("resp", r)
			logrus.Infof("Create brand success: brand = %v", *r.Brand)

			getAllReq := &pb.Request{}
			span.SetTag("GetAllReq", getAllReq)
			r, err = client.GetAll(context.TODO(), getAllReq)
			if err != nil {
				logrus.Errorf("Get all brand failed: %v", err)
				return err
			}
			span.SetTag("GetAllResp", r)
			logrus.Infof("Get all brand result: brands = %v", r.Brands)
			return nil
		}),
	)
}
