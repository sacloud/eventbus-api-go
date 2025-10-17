package eventbus

import (
	"context"
	"fmt"
	"log"
	"net/http"

	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
)

var _ http.RoundTripper = (*filterInjector)(nil)

type filterInjector struct{}

func (t *filterInjector) RoundTrip(req *http.Request) (*http.Response, error) {
	// NOTE: OpenAPIで表現できないクエリの書き込みを行う
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"eventbusschedule"}}`
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"eventbustrigger"}}`
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"eventbusprocessconfiguration"}}`.
	if req.URL.Path == "/commonserviceitem" && req.Method == http.MethodGet {
		pc := getFilterProviderClass(req.Context())
		req.URL.RawQuery = fmt.Sprintf(`{"Filter":{"Provider.Class":"%s"}}`, pc)
	}
	return http.DefaultTransport.RoundTrip(req)
}

type ctxKeyFilterProviderClass struct{}

func setFilterProviderClass(ctx context.Context, v v1.ProviderClass) context.Context {
	return context.WithValue(ctx, ctxKeyFilterProviderClass{}, v)
}

func getFilterProviderClass(ctx context.Context) v1.ProviderClass {
	return ctx.Value(ctxKeyFilterProviderClass{}).(v1.ProviderClass)
}
