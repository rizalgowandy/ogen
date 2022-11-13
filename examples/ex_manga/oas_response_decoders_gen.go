// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/validate"
)

func decodeGetBookResponse(resp *http.Response) (res GetBookRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Book
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetBookForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetPageCoverImageResponse(resp *http.Response) (res GetPageCoverImageRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ht.MatchContentType("image/*", ct):
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := GetPageCoverImageOK{Data: bytes.NewReader(b)}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetPageCoverImageForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetPageImageResponse(resp *http.Response) (res GetPageImageRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ht.MatchContentType("image/*", ct):
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := GetPageImageOK{Data: bytes.NewReader(b)}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetPageImageForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetPageThumbnailImageResponse(resp *http.Response) (res GetPageThumbnailImageRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ht.MatchContentType("image/*", ct):
			reader := resp.Body
			b, err := io.ReadAll(reader)
			if err != nil {
				return res, err
			}

			response := GetPageThumbnailImageOK{Data: bytes.NewReader(b)}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &GetPageThumbnailImageForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeSearchResponse(resp *http.Response) (res SearchRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response SearchOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &SearchForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeSearchByTagIDResponse(resp *http.Response) (res SearchByTagIDRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response SearchByTagIDOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		// Code 403.
		return &SearchByTagIDForbidden{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
