// Code generated by ogen, DO NOT EDIT.

package techempower

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

func decodeCachingResponse(resp *http.Response) (*WorldObjects, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response WorldObjects
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

func decodeDBResponse(resp *http.Response) (*WorldObject, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response WorldObject
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

func decodeJSONResponse(resp *http.Response) (*HelloWorld, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response HelloWorld
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

func decodeQueriesResponse(resp *http.Response) (*WorldObjects, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response WorldObjects
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

func decodeUpdatesResponse(resp *http.Response) (*WorldObjects, error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response WorldObjects
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
