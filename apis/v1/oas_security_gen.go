// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// CloudCtrlAuth provides CloudCtrlAuth security value.
	// Cloud-ctrl認証用のAPIキー.
	CloudCtrlAuth(ctx context.Context, operationName OperationName) (CloudCtrlAuth, error)
}

func (s *Client) securityCloudCtrlAuth(ctx context.Context, operationName OperationName, req *http.Request) error {
	t, err := s.sec.CloudCtrlAuth(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"CloudCtrlAuth\"")
	}
	req.Header.Set("Authorization", "Bearer "+t.Token)
	return nil
}
