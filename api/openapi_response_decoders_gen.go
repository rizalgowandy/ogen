// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
)

func DecodeFoobarGetResponse(resp *http.Response) (FoobarGetResponse, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			return &response, nil
		default:
			return nil, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &NotFound{}, nil
	default:
		return nil, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func DecodeFoobarPostResponse(resp *http.Response) (FoobarPostResponse, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			return &response, nil
		default:
			return nil, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		return &NotFound{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorResponseDefault
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return nil, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func DecodePetGetResponse(resp *http.Response) (PetGetResponse, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			return &response, nil
		default:
			return nil, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response PetGetDefaultResponse
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return nil, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func DecodePetPostResponse(resp *http.Response) (*Pet, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			return &response, nil
		default:
			return nil, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return nil, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func DecodePetNameGetResponse(resp *http.Response) (*Pet, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return nil, err
			}

			return &response, nil
		default:
			return nil, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return nil, fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}
