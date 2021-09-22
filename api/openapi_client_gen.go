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

type Client struct {
	serverURL string
	http      *http.Client
}

func NewClient(serverURL string) *Client {
	return &Client{
		serverURL: serverURL,
		http: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c *Client) FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetResponse, error) {
	path := c.serverURL
	path += "/foobar"

	r, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := r.URL.Query()
	{
		s := conv.Int64ToString(params.Query.InlinedParam)
		q.Set("inlinedParam", s)
	}
	{
		s := conv.Int32ToString(params.Query.Skip)
		q.Set("skip", s)
	}
	r.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	decoded, err := DecodeFoobarGetResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return decoded, nil
}

func (c *Client) FoobarPost(ctx context.Context, req *Pet) (FoobarPostResponse, error) {
	body, contentType, err := EncodeFoobarPostRequest(req)
	if err != nil {
		return nil, err
	}

	path := c.serverURL
	path += "/foobar"

	r, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	decoded, err := DecodeFoobarPostResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return decoded, nil
}

func (c *Client) PetGet(ctx context.Context, params PetGetParams) (PetGetResponse, error) {
	path := c.serverURL
	path += "/pet"

	r, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	{
		value := conv.StringArrayToString(params.Header.XScope)
		for _, v := range value {
			r.Header.Add("x-scope", v)
		}
	}

	q := r.URL.Query()
	{
		s := conv.Int64ToString(params.Query.PetID)
		q.Set("petID", s)
	}
	r.URL.RawQuery = q.Encode()

	{
		value := conv.StringToString(params.Cookie.Token)
		r.AddCookie(&http.Cookie{
			Name:   "token",
			Value:  value,
			MaxAge: 1337,
		})
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	decoded, err := DecodePetGetResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return decoded, nil
}

func (c *Client) PetPost(ctx context.Context, req PetPostRequest) (*Pet, error) {
	body, contentType, err := EncodePetPostRequest(req)
	if err != nil {
		return nil, err
	}

	path := c.serverURL
	path += "/pet"

	r, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	decoded, err := DecodePetPostResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return decoded, nil
}

func (c *Client) PetNameGet(ctx context.Context, params PetNameGetParams) (*Pet, error) {
	path := c.serverURL
	path += "/pet"
	{
		value := conv.StringToString(params.Path.Name)
		path += "/" + value
	}

	r, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	decoded, err := DecodePetNameGetResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return decoded, nil
}
