// Copyright 2025- The sacloud/eventbus-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventbus

import (
	"context"
	"fmt"
	"runtime"

	client "github.com/sacloud/api-client-go"
	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
)

// DefaultAPIRootURL デフォルトのAPIルートURL
const DefaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/commonserviceitem/"

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"eventbus-api-go/%s (%s/%s; +https://github.com/sacloud/eventbus-api-go) %s",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
	client.DefaultUserAgent,
)

// SecuritySourceはOpenAPI定義で使用されている認証のための仕組み。api-client-goが処理するので、ogen用はダミーで誤魔化す
type DummySecuritySource struct {
	Token string
}

func (ss DummySecuritySource) CloudCtrlAuth(ctx context.Context, operationName v1.OperationName) (v1.CloudCtrlAuth, error) {
	return v1.CloudCtrlAuth{Token: ss.Token}, nil
}

func NewClient(params ...client.ClientParam) (*v1.Client, error) {
	return NewClientWithApiUrl(DefaultAPIRootURL, params...)
}

func NewClientWithApiUrl(apiUrl string, params ...client.ClientParam) (*v1.Client, error) {
	params = append(params, client.WithUserAgent(UserAgent))
	c, err := client.NewClient(apiUrl, params...)
	if err != nil {
		return nil, NewError("NewClientWithApiUrl", err)
	}

	v1Client, err := v1.NewClient(c.ServerURL(), DummySecuritySource{Token: "eventbus-client"}, v1.WithClient(c.NewHttpRequestDoer()))
	if err != nil {
		return nil, NewError("NewClientWithApiUrl", err)
	}

	return v1Client, nil
}
