// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeDataGetFormatResponse(response string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	e.Str(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeDefaultTestResponse(response int32, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	e.Int32(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeErrorGetResponse(response ErrorStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	st := http.StatusText(response.StatusCode)
	if response.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeFoobarGetResponse(response FoobarGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *NotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		return nil

	default:
		return errors.Errorf("/foobar"+`: unexpected response type: %T`, response)
	}
}
func encodeFoobarPostResponse(response FoobarPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *NotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/foobar"+`: unexpected response type: %T`, response)
	}
}
func encodeFoobarPutResponse(response FoobarPutDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(response.StatusCode)
	st := http.StatusText(response.StatusCode)
	if response.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}
	return nil

}
func encodeGetHeaderResponse(response Hash, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeNoAdditionalPropertiesTestResponse(response NoAdditionalPropertiesTest, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeNullableDefaultResponseResponse(response NilIntStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	st := http.StatusText(response.StatusCode)
	if response.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeOneofBugResponse(response OneofBugOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	return nil

}
func encodePatternRecursiveMapGetResponse(response PatternRecursiveMap, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodePetCreateResponse(response Pet, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodePetFriendsNamesByIDResponse(response []string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	e.ArrStart()
	for _, elem := range response {
		e.Str(elem)
	}
	e.ArrEnd()
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodePetGetResponse(response PetGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *PetGetDefStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pet"+`: unexpected response type: %T`, response)
	}
}
func encodePetGetAvatarByIDResponse(response PetGetAvatarByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetGetAvatarByIDOK:
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		if _, err := io.Copy(w, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *NotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pet/avatar"+`: unexpected response type: %T`, response)
	}
}
func encodePetGetAvatarByNameResponse(response PetGetAvatarByNameRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetGetAvatarByNameOK:
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		if _, err := io.Copy(w, response); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *NotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pet/{name}/avatar"+`: unexpected response type: %T`, response)
	}
}
func encodePetGetByNameResponse(response Pet, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodePetNameByIDResponse(response string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	e.Str(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodePetUpdateNameAliasPostResponse(response PetUpdateNameAliasPostDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(response.StatusCode)
	st := http.StatusText(response.StatusCode)
	if response.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}
	return nil

}
func encodePetUpdateNamePostResponse(response PetUpdateNamePostDefStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(response.StatusCode)
	st := http.StatusText(response.StatusCode)
	if response.StatusCode >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}
	return nil

}
func encodePetUploadAvatarByIDResponse(response PetUploadAvatarByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PetUploadAvatarByIDOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		return nil

	case *NotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pet/avatar"+`: unexpected response type: %T`, response)
	}
}
func encodeRecursiveArrayGetResponse(response RecursiveArray, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeRecursiveMapGetResponse(response RecursiveMap, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeSecurityTestResponse(response string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	e.Str(response)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeStringIntMapGetResponse(response StringIntMap, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
func encodeTestFloatValidationResponse(response TestFloatValidationOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	return nil

}
func encodeTestFormURLEncodedResponse(response TestFormURLEncodedOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	return nil

}
func encodeTestObjectQueryParameterResponse(response TestObjectQueryParameterOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}
