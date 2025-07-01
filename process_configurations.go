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
	List(ctx context.Context) ([]v1.ProcessConfiguration, error)
	Read(ctx context.Context, id string) (*v1.ProcessConfiguration, error)
	Create(ctx context.Context, request v1.ProcessConfigurationRequestSettings) (*v1.ProcessConfiguration, error)
	Update(ctx context.Context, id string, request v1.ProcessConfigurationRequestSettings) (*v1.ProcessConfiguration, error)
	UpdateSecret(ctx context.Context, id string, secret v1.ProcessConfigurationSecret) error
	Delete(ctx context.Context, id string) error
}

var _ ProcessConfigurationAPI = (*processConfigurationOp)(nil)

type processConfigurationOp struct {
	client *v1.Client
}

func NewProcessConfigurationOp(client *v1.Client) ProcessConfigurationAPI {
	return &processConfigurationOp{client: client}
}

func (op *processConfigurationOp) List(ctx context.Context) ([]v1.ProcessConfiguration, error) {
	res, err := op.client.GetProcessConfigurations(ctx)
	if err != nil {
		return nil, NewError("List", err)
	}

	switch p := res.(type) {
	case *v1.GetProcessConfigurationsOK:
		return p.ProcessConfigurations, nil
	case *v1.GetProcessConfigurationsUnauthorized:
		return nil, NewError("List", errors.New(p.Message.Value))
	case *v1.GetProcessConfigurationsBadRequest:
		return nil, NewError("List", errors.New(p.Message.Value))
	case *v1.GetProcessConfigurationsInternalServerError:
		return nil, NewError("List", errors.New(p.Message.Value))
	default:
		return nil, NewError("List", errors.New("unknown error"))
	}
}

func (op *processConfigurationOp) Read(ctx context.Context, id string) (*v1.ProcessConfiguration, error) {
	res, err := op.client.GetProcessConfigurationById(ctx, v1.GetProcessConfigurationByIdParams{ID: id})
	if err != nil {
		return nil, NewError("Read", err)
	}

	switch p := res.(type) {
	case *v1.GetProcessConfigurationByIdOK:
		return &p.ProcessConfiguration, nil
	case *v1.GetProcessConfigurationByIdUnauthorized:
		return nil, NewError("Read", errors.New(p.Message.Value))
	case *v1.GetProcessConfigurationByIdBadRequest:
		return nil, NewError("Read", errors.New(p.Message.Value))
	case *v1.GetProcessConfigurationByIdNotFound:
		return nil, NewError("Read", errors.New(p.Message.Value))
	case *v1.GetProcessConfigurationByIdInternalServerError:
		return nil, NewError("Read", errors.New(p.Message.Value))
	default:
		return nil, NewError("Read", errors.New("unknown error"))
	}
}

func (op *processConfigurationOp) Create(ctx context.Context, request v1.ProcessConfigurationRequestSettings) (*v1.ProcessConfiguration, error) {
	res, err := op.client.CreateProcessConfiguration(ctx, &v1.CreateProcessConfigurationRequest{
		ProcessConfiguration: request,
	})
	if err != nil {
		return nil, NewError("Create", err)
	}

	switch p := res.(type) {
	case *v1.CreateProcessConfigurationOK:
		return &p.ProcessConfiguration, nil
	case *v1.CreateProcessConfigurationBadRequest:
		return nil, NewError("Create", errors.New(p.Message.Value))
	case *v1.CreateProcessConfigurationUnauthorized:
		return nil, NewError("Create", errors.New(p.Message.Value))
	case *v1.CreateProcessConfigurationInternalServerError:
		return nil, NewError("Create", errors.New(p.Message.Value))
	default:
		return nil, NewError("Create", errors.New("unknown error"))
	}
}

func (op *processConfigurationOp) Update(ctx context.Context, id string, request v1.ProcessConfigurationRequestSettings) (*v1.ProcessConfiguration, error) {
	res, err := op.client.ConfigureProcessConfiguration(ctx, &v1.ConfigureProcessConfigurationRequest{
		ProcessConfiguration: request,
	}, v1.ConfigureProcessConfigurationParams{ID: id})
	if err != nil {
		return nil, NewError("Update", err)
	}

	switch p := res.(type) {
	case *v1.ConfigureProcessConfigurationOK:
		return &p.ProcessConfiguration, nil
	case *v1.ConfigureProcessConfigurationBadRequest:
		return nil, NewError("Update", errors.New(p.Message.Value))
	case *v1.ConfigureProcessConfigurationNotFound:
		return nil, NewError("Update", errors.New(p.Message.Value))
	case *v1.ConfigureProcessConfigurationInternalServerError:
		return nil, NewError("Update", errors.New(p.Message.Value))
	default:
		return nil, NewError("Update", errors.New("unknown error"))
	}
}

func (op *processConfigurationOp) UpdateSecret(ctx context.Context, id string, secret v1.ProcessConfigurationSecret) error {
	res, err := op.client.ConfigureProcessConfigurationSecret(ctx, &v1.ConfigureProcessConfigurationSecretRequest{
		Secret: secret,
	}, v1.ConfigureProcessConfigurationSecretParams{ID: id})
	if err != nil {
		return NewError("UpdateSecret", err)
	}

	switch p := res.(type) {
	case *v1.ConfigureProcessConfigurationSecretOK:
		return nil
	case *v1.ConfigureProcessConfigurationSecretBadRequest:
		return NewError("UpdateSecret", errors.New(p.Message.Value))
	case *v1.ConfigureProcessConfigurationSecretNotFound:
		return NewError("UpdateSecret", errors.New(p.Message.Value))
	case *v1.ConfigureProcessConfigurationSecretInternalServerError:
		return NewError("UpdateSecret", errors.New(p.Message.Value))
	default:
		return NewError("UpdateSecret", errors.New("unknown error"))
	}
}

func (op *processConfigurationOp) Delete(ctx context.Context, id string) error {
	res, err := op.client.DeleteProcessConfiguration(ctx, v1.DeleteProcessConfigurationParams{ID: id})
	if err != nil {
		return NewError("Delete", err)
	}

	switch p := res.(type) {
	case *v1.DeleteProcessConfigurationOK:
		return nil
	case *v1.DeleteProcessConfigurationUnauthorized:
		return NewError("Delete", errors.New(p.Message.Value))
	case *v1.DeleteProcessConfigurationBadRequest:
		return NewError("Delete", errors.New(p.Message.Value))
	case *v1.DeleteProcessConfigurationNotFound:
		return NewError("Delete", errors.New(p.Message.Value))
	case *v1.DeleteProcessConfigurationInternalServerError:
		return NewError("Delete", errors.New(p.Message.Value))
	default:
		return NewError("Delete", errors.New("unknown error"))
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
