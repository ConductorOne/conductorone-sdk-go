package tools

import (
	"context"
	"fmt"
	"reflect"
)

type ResponseType interface {
	GetStatusCode() int
}

type Result[T any] struct {
	Message T
	Error   error
}

func Unroll[reqT any, respT ResponseType, itemT any](
	ctx context.Context,
	preq reqT,
	listFunc func(ctx context.Context, req reqT) (respT, error),
	extractFunc func(respT) (*string, []itemT, error),
) ([]itemT, *reqT, error) {
	var rv = make([]itemT, 0, 100)

	resp, err := listFunc(ctx, preq)
	if err != nil {
		return nil, nil, err
	}
	if resp.GetStatusCode() != 200 {
		// TOOD(pquerna): better structured error
		return nil, nil, fmt.Errorf("unexpected status code: %d", resp.GetStatusCode())
	}
	nextPageToken, items, err := extractFunc(resp)
	if err != nil {
		return nil, nil, err
	}
	rv = append(rv, items...)
	if nextPageToken == nil || *nextPageToken == "" {
		return rv, nil, nil
	}

	// https://play.golang.com/p/1q-RAyuvL2K
	newValueReflect := reflect.ValueOf(nextPageToken)
	reflect.ValueOf(&preq).Elem().FieldByName("PageToken").Set(newValueReflect)
	return rv, &preq, nil
}

func Paginate[reqT any, respT ResponseType, itemT any](
	ctx context.Context,
	preq reqT,
	listFunc func(ctx context.Context, req reqT) (respT, error),
	extractFunc func(respT) (*string, []itemT, error),
	quit chan int,
) <-chan Result[[]itemT] {
	ch := make(chan Result[[]itemT], 1)
	var req = preq
	go func() {
		defer close(ch)
		for {
			select {
			case <-quit:
				return
			default:
				rv, preq, err := Unroll(ctx, req, listFunc, extractFunc)
				if err != nil {
					ch <- Result[[]itemT]{Message: rv, Error: err}
					return
				}
				ch <- Result[[]itemT]{Message: rv, Error: err}
				if preq == nil {
					return
				}
				req = *preq
			}

		}
	}()
	return ch
}
