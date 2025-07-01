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
	"errors"

	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
)

type ScheduleAPI interface {
	List(ctx context.Context) ([]v1.Schedule, error)
	Read(ctx context.Context, id string) (*v1.Schedule, error)
	Create(ctx context.Context, request v1.ScheduleRequestSettings) (*v1.Schedule, error)
	Update(ctx context.Context, id string, request v1.ScheduleRequestSettings) (*v1.Schedule, error)
	Delete(ctx context.Context, id string) error
}

var _ ScheduleAPI = (*scheduleOp)(nil)

type scheduleOp struct {
	client *v1.Client
}

func NewScheduleOp(client *v1.Client) ScheduleAPI {
	return &scheduleOp{client: client}
}

func (op *scheduleOp) List(ctx context.Context) ([]v1.Schedule, error) {
	res, err := op.client.GetSchedules(ctx)
	if err != nil {
		return nil, NewError("List", err)
	}

	switch p := res.(type) {
	case *v1.GetSchedulesOK:
		return p.Schedules, nil
	case *v1.GetSchedulesUnauthorized:
		return nil, NewError("List", errors.New(p.Message.Value))
	case *v1.GetSchedulesBadRequest:
		return nil, NewError("List", errors.New(p.Message.Value))
	case *v1.GetSchedulesInternalServerError:
		return nil, NewError("List", errors.New(p.Message.Value))
	default:
		return nil, NewError("List", errors.New("unknown error"))
	}
}

func (op *scheduleOp) Read(ctx context.Context, id string) (*v1.Schedule, error) {
	res, err := op.client.GetScheduleById(ctx, v1.GetScheduleByIdParams{ID: id})
	if err != nil {
		return nil, NewError("Read", err)
	}

	switch p := res.(type) {
	case *v1.GetScheduleByIdOK:
		return &p.Schedule, nil
	case *v1.GetScheduleByIdUnauthorized:
		return nil, NewError("Read", errors.New(p.Message.Value))
	case *v1.GetScheduleByIdBadRequest:
		return nil, NewError("Read", errors.New(p.Message.Value))
	case *v1.GetScheduleByIdNotFound:
		return nil, NewError("Read", errors.New(p.Message.Value))
	case *v1.GetScheduleByIdInternalServerError:
		return nil, NewError("Read", errors.New(p.Message.Value))
	default:
		return nil, NewError("Read", errors.New("unknown error"))
	}
}

func (op *scheduleOp) Create(ctx context.Context, request v1.ScheduleRequestSettings) (*v1.Schedule, error) {
	res, err := op.client.CreateSchedule(ctx, &v1.CreateScheduleRequest{
		Schedule: request,
	})
	if err != nil {
		return nil, NewError("Create", err)
	}

	switch p := res.(type) {
	case *v1.CreateScheduleOK:
		return &p.Schedule, nil
	case *v1.CreateScheduleBadRequest:
		return nil, NewError("Create", errors.New(p.Message.Value))
	case *v1.CreateScheduleUnauthorized:
		return nil, NewError("Create", errors.New(p.Message.Value))
	case *v1.CreateScheduleInternalServerError:
		return nil, NewError("Create", errors.New(p.Message.Value))
	default:
		return nil, NewError("Create", errors.New("unknown error"))
	}
}

func (op *scheduleOp) Update(ctx context.Context, id string, request v1.ScheduleRequestSettings) (*v1.Schedule, error) {
	res, err := op.client.ConfigureSchedule(ctx, &v1.ConfigureScheduleRequest{
		Schedule: request,
	}, v1.ConfigureScheduleParams{ID: id})
	if err != nil {
		return nil, NewError("Update", err)
	}

	switch p := res.(type) {
	case *v1.ConfigureScheduleOK:
		return &p.Schedule, nil
	case *v1.ConfigureScheduleBadRequest:
		return nil, NewError("Update", errors.New(p.Message.Value))
	case *v1.ConfigureScheduleNotFound:
		return nil, NewError("Update", errors.New(p.Message.Value))
	case *v1.ConfigureScheduleInternalServerError:
		return nil, NewError("Update", errors.New(p.Message.Value))
	default:
		return nil, NewError("Update", errors.New("unknown error"))
	}
}

func (op *scheduleOp) Delete(ctx context.Context, id string) error {
	res, err := op.client.DeleteSchedule(ctx, v1.DeleteScheduleParams{ID: id})
	if err != nil {
		return NewError("Delete", err)
	}

	switch p := res.(type) {
	case *v1.DeleteScheduleOK:
		return nil
	case *v1.DeleteScheduleUnauthorized:
		return NewError("Delete", errors.New(p.Message.Value))
	case *v1.DeleteScheduleBadRequest:
		return NewError("Delete", errors.New(p.Message.Value))
	case *v1.DeleteScheduleNotFound:
		return NewError("Delete", errors.New(p.Message.Value))
	case *v1.DeleteScheduleInternalServerError:
		return NewError("Delete", errors.New(p.Message.Value))
	default:
		return NewError("Delete", errors.New("unknown error"))
	}
}
