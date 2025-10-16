package eventbus

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
)

var _ http.RoundTripper = (*filterInjector)(nil)

type filterInjector struct{}

func (t *filterInjector) RoundTrip(req *http.Request) (*http.Response, error) {
	// NOTE: OpenAPIで表現できないクエリの書き込みを行う
	// 同じエンドポイントに3種類のProvider.Classでフィルタしたいため、生成コードの書き換えでなくclient middlewareで対応
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"eventbusschedule"}}`
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"eventbustrigger"}}`
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"eventbusprocessconfiguration"}}`.
	if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/commonserviceitem") {
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
