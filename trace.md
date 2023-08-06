https://github.com/buddyxiao/go-micro-demo



JaegerHostPort = "127.0.0.1:6831"

jaeger, closer, err := config.NewJaegerTracer("api-gateway", consts.JaegerHostPort)
defer closer.Close()

middleware.JaegerGatewayMiddleware(jaeger),
```
package middleware
import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

func JaegerGatewayMiddleware(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var md = make(metadata.Metadata, 1)
		opName := ctx.Request.URL.Path + "-" + ctx.Request.Method
		parentSpan := tracer.StartSpan(opName)
		defer parentSpan.Finish()
		injectErr := tracer.Inject(parentSpan.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
		if injectErr != nil {
			logger.Fatalf("%s: Couldn't inject metadata", injectErr)
		}
		newCtx := metadata.NewContext(ctx.Request.Context(), md)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
```


micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
micro.WrapClient(opentracing.NewClientWrapper(tracer)),


```
package config

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

var Tracer opentracing.Tracer

func NewJaegerTracer(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  jaegerHostPort,
		},
	}
	var closer io.Closer
	var err error
	Tracer, closer, err = cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(Tracer)
	return Tracer, closer, err
}
```

