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
	"encoding/json"
	"errors"

	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
)

type ProcessConfigurationAPI interface {
	List(ctx context.Context) (v1.ProcessConfigurations, error)
	Read(ctx context.Context, id string) (v1.ProcessConfiguration, error)
	Create(ctx context.Context, request v1.ProcessConfigurationRequestSettings) (v1.ProcessConfiguration, error)
	Update(ctx context.Context, id string, request v1.ProcessConfigurationRequestSettings) (v1.ProcessConfiguration, error)
	UpdateSecret(ctx context.Context, id string, secret v1.ProcessConfigurationSecret) error
	Delete(ctx context.Context, id string) (v1.ProcessConfiguration, error)
}

var _ ProcessConfigurationAPI = (*processConfigurationOp)(nil)

type processConfigurationOp struct {
	client *v1.Client
}

func NewProcessConfigurationOp(client *v1.Client) ProcessConfigurationAPI {
	return &processConfigurationOp{client: client}
}

func (op *processConfigurationOp) List(ctx context.Context) (v1.ProcessConfigurations, error) {
	res, err := op.client.GetProcessConfigurations(ctx)
	if err != nil {
		return nil, err
	}

	switch p := res.(type) {
	case *v1.GetProcessConfigurationsOK:
		return p.ProcessConfigurations, nil
	case *v1.GetProcessConfigurationsUnauthorized:
		return nil, errors.New(p.Message.Value)
	case *v1.GetProcessConfigurationsBadRequest:
		return nil, errors.New(p.Message.Value)
	case *v1.GetProcessConfigurationsInternalServerError:
		return nil, errors.New(p.Message.Value)
	default:
		return nil, errors.New("unknown error")
	}
}

func (op *processConfigurationOp) Read(ctx context.Context, id string) (v1.ProcessConfiguration, error) {
	var empty v1.ProcessConfiguration
	res, err := op.client.GetProcessConfigurationById(ctx, v1.GetProcessConfigurationByIdParams{ID: id})
	if err != nil {
		return empty, err
	}

	switch p := res.(type) {
	case *v1.GetProcessConfigurationByIdOK:
		return p.ProcessConfiguration, nil
	case *v1.GetProcessConfigurationByIdUnauthorized:
		return empty, errors.New(p.Message.Value)
	case *v1.GetProcessConfigurationByIdBadRequest:
		return empty, errors.New(p.Message.Value)
	case *v1.GetProcessConfigurationByIdNotFound:
		return empty, errors.New(p.Message.Value)
	case *v1.GetProcessConfigurationByIdInternalServerError:
		return empty, errors.New(p.Message.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *processConfigurationOp) Create(ctx context.Context, request v1.ProcessConfigurationRequestSettings) (v1.ProcessConfiguration, error) {
	var empty v1.ProcessConfiguration
	res, err := op.client.CreateProcessConfiguration(ctx, &v1.CreateProcessConfigurationRequest{
		ProcessConfiguration: request,
	})
	if err != nil {
		return empty, err
	}

	switch p := res.(type) {
	case *v1.CreateProcessConfigurationOK:
		return p.ProcessConfiguration, nil
	case *v1.CreateProcessConfigurationBadRequest:
		return empty, errors.New(p.Message.Value)
	case *v1.CreateProcessConfigurationUnauthorized:
		return empty, errors.New(p.Message.Value)
	case *v1.CreateProcessConfigurationInternalServerError:
		return empty, errors.New(p.Message.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *processConfigurationOp) Update(ctx context.Context, id string, request v1.ProcessConfigurationRequestSettings) (v1.ProcessConfiguration, error) {
	var empty v1.ProcessConfiguration
	res, err := op.client.ConfigureProcessConfiguration(ctx, &v1.ConfigureProcessConfigurationRequest{
		ProcessConfiguration: request,
	}, v1.ConfigureProcessConfigurationParams{ID: id})
	if err != nil {
		return empty, err
	}

	switch p := res.(type) {
	case *v1.ConfigureProcessConfigurationOK:
		return p.ProcessConfiguration, nil
	case *v1.ConfigureProcessConfigurationBadRequest:
		return empty, errors.New(p.Message.Value)
	case *v1.ConfigureProcessConfigurationNotFound:
		return empty, errors.New(p.Message.Value)
	case *v1.ConfigureProcessConfigurationInternalServerError:
		return empty, errors.New(p.Message.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *processConfigurationOp) UpdateSecret(ctx context.Context, id string, secret v1.ProcessConfigurationSecret) error {
	res, err := op.client.ConfigureProcessConfigurationSecret(ctx, &v1.ConfigureProcessConfigurationSecretRequest{
		Secret: secret,
	}, v1.ConfigureProcessConfigurationSecretParams{ID: id})
	if err != nil {
		return err
	}

	switch p := res.(type) {
	case *v1.ConfigureProcessConfigurationSecretOK:
		return nil
	case *v1.ConfigureProcessConfigurationSecretBadRequest:
		return errors.New(p.Message.Value)
	case *v1.ConfigureProcessConfigurationSecretNotFound:
		return errors.New(p.Message.Value)
	case *v1.ConfigureProcessConfigurationSecretInternalServerError:
		return errors.New(p.Message.Value)
	default:
		return errors.New("unknown error")
	}
}

func (op *processConfigurationOp) Delete(ctx context.Context, id string) (v1.ProcessConfiguration, error) {
	var empty v1.ProcessConfiguration
	res, err := op.client.DeleteProcessConfiguration(ctx, v1.DeleteProcessConfigurationParams{ID: id})
	if err != nil {
		return empty, err
	}

	switch p := res.(type) {
	case *v1.DeleteProcessConfigurationOK:
		return p.ProcessConfiguration, nil
	case *v1.DeleteProcessConfigurationUnauthorized:
		return empty, errors.New(p.Message.Value)
	case *v1.DeleteProcessConfigurationBadRequest:
		return empty, errors.New(p.Message.Value)
	case *v1.DeleteProcessConfigurationNotFound:
		return empty, errors.New(p.Message.Value)
	case *v1.DeleteProcessConfigurationInternalServerError:
		return empty, errors.New(p.Message.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func CreateSimpleNotificationSettings(groupId string, message string) v1.DestinationSettings {
	param, _ := json.Marshal(map[string]string{"group_id": groupId, "message": message})
	return v1.DestinationSettings{
		Destination: "simplenotification",
		Parameters:  string(param),
	}
}

func CreateSimpleMqSettings(queueName string, content string) v1.DestinationSettings {
	param, _ := json.Marshal(map[string]string{"queue_name": queueName, "content": content})
	return v1.DestinationSettings{
		Destination: "simplemq",
		Parameters:  string(param),
	}
}
