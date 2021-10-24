// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

// WriteJSON implements json.Marshaler.
func (s Book) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.ID.Set {
		more.More()
		j.WriteObjectField("id")
		s.ID.WriteJSON(j)
	}
	if s.Images.Set {
		more.More()
		j.WriteObjectField("images")
		s.Images.WriteJSON(j)
	}
	if s.MediaID.Set {
		more.More()
		j.WriteObjectField("media_id")
		s.MediaID.WriteJSON(j)
	}
	if s.NumFavorites.Set {
		more.More()
		j.WriteObjectField("num_favorites")
		s.NumFavorites.WriteJSON(j)
	}
	if s.NumPages.Set {
		more.More()
		j.WriteObjectField("num_pages")
		s.NumPages.WriteJSON(j)
	}
	if s.Scanlator.Set {
		more.More()
		j.WriteObjectField("scanlator")
		s.Scanlator.WriteJSON(j)
	}
	if s.Tags != nil {
		more.More()
		j.WriteObjectField("tags")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.Tags {
			more.More()
			elem.WriteJSON(j)
		}
		j.WriteArrayEnd()
		more.Up()
	}
	if s.Title.Set {
		more.More()
		j.WriteObjectField("title")
		s.Title.WriteJSON(j)
	}
	if s.UploadDate.Set {
		more.More()
		j.WriteObjectField("upload_date")
		s.UploadDate.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Book json value to io.Writer.
func (s Book) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Book json value from io.Reader.
func (s *Book) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Book from json stream.
func (s *Book) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "id":
			s.ID.Reset()
			if err := s.ID.ReadJSON(i); err != nil {
				i.ReportError("Field ID", err.Error())
				return false
			}
			return true
		case "images":
			s.Images.Reset()
			if err := s.Images.ReadJSON(i); err != nil {
				i.ReportError("Field Images", err.Error())
				return false
			}
			return true
		case "media_id":
			s.MediaID.Reset()
			if err := s.MediaID.ReadJSON(i); err != nil {
				i.ReportError("Field MediaID", err.Error())
				return false
			}
			return true
		case "num_favorites":
			s.NumFavorites.Reset()
			if err := s.NumFavorites.ReadJSON(i); err != nil {
				i.ReportError("Field NumFavorites", err.Error())
				return false
			}
			return true
		case "num_pages":
			s.NumPages.Reset()
			if err := s.NumPages.ReadJSON(i); err != nil {
				i.ReportError("Field NumPages", err.Error())
				return false
			}
			return true
		case "scanlator":
			s.Scanlator.Reset()
			if err := s.Scanlator.ReadJSON(i); err != nil {
				i.ReportError("Field Scanlator", err.Error())
				return false
			}
			return true
		case "tags":
			// Unsupported kind "array" for field "Tags".
			i.Skip()
			return true
		case "title":
			s.Title.Reset()
			if err := s.Title.ReadJSON(i); err != nil {
				i.ReportError("Field Title", err.Error())
				return false
			}
			return true
		case "upload_date":
			s.UploadDate.Reset()
			if err := s.UploadDate.ReadJSON(i); err != nil {
				i.ReportError("Field UploadDate", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

func (GetBookForbidden) WriteJSON(j *json.Stream)        {}
func (GetBookForbidden) ReadJSON(i *json.Iterator) error { return nil }
func (GetBookForbidden) ReadJSONFrom(r io.Reader) error  { return nil }
func (GetBookForbidden) WriteJSONTo(w io.Writer) error   { return nil }

func (GetPageCoverImageForbidden) WriteJSON(j *json.Stream)        {}
func (GetPageCoverImageForbidden) ReadJSON(i *json.Iterator) error { return nil }
func (GetPageCoverImageForbidden) ReadJSONFrom(r io.Reader) error  { return nil }
func (GetPageCoverImageForbidden) WriteJSONTo(w io.Writer) error   { return nil }

func (GetPageCoverImageOK) WriteJSON(j *json.Stream)        {}
func (GetPageCoverImageOK) ReadJSON(i *json.Iterator) error { return nil }
func (GetPageCoverImageOK) ReadJSONFrom(r io.Reader) error  { return nil }
func (GetPageCoverImageOK) WriteJSONTo(w io.Writer) error   { return nil }

func (GetPageImageForbidden) WriteJSON(j *json.Stream)        {}
func (GetPageImageForbidden) ReadJSON(i *json.Iterator) error { return nil }
func (GetPageImageForbidden) ReadJSONFrom(r io.Reader) error  { return nil }
func (GetPageImageForbidden) WriteJSONTo(w io.Writer) error   { return nil }

func (GetPageImageOK) WriteJSON(j *json.Stream)        {}
func (GetPageImageOK) ReadJSON(i *json.Iterator) error { return nil }
func (GetPageImageOK) ReadJSONFrom(r io.Reader) error  { return nil }
func (GetPageImageOK) WriteJSONTo(w io.Writer) error   { return nil }

func (GetPageThumbnailImageForbidden) WriteJSON(j *json.Stream)        {}
func (GetPageThumbnailImageForbidden) ReadJSON(i *json.Iterator) error { return nil }
func (GetPageThumbnailImageForbidden) ReadJSONFrom(r io.Reader) error  { return nil }
func (GetPageThumbnailImageForbidden) WriteJSONTo(w io.Writer) error   { return nil }

func (GetPageThumbnailImageOK) WriteJSON(j *json.Stream)        {}
func (GetPageThumbnailImageOK) ReadJSON(i *json.Iterator) error { return nil }
func (GetPageThumbnailImageOK) ReadJSONFrom(r io.Reader) error  { return nil }
func (GetPageThumbnailImageOK) WriteJSONTo(w io.Writer) error   { return nil }

// WriteJSON implements json.Marshaler.
func (s Image) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.H.Set {
		more.More()
		j.WriteObjectField("h")
		s.H.WriteJSON(j)
	}
	if s.T.Set {
		more.More()
		j.WriteObjectField("t")
		s.T.WriteJSON(j)
	}
	if s.W.Set {
		more.More()
		j.WriteObjectField("w")
		s.W.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Image json value to io.Writer.
func (s Image) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Image json value from io.Reader.
func (s *Image) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Image from json stream.
func (s *Image) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "h":
			s.H.Reset()
			if err := s.H.ReadJSON(i); err != nil {
				i.ReportError("Field H", err.Error())
				return false
			}
			return true
		case "t":
			s.T.Reset()
			if err := s.T.ReadJSON(i); err != nil {
				i.ReportError("Field T", err.Error())
				return false
			}
			return true
		case "w":
			s.W.Reset()
			if err := s.W.ReadJSON(i); err != nil {
				i.ReportError("Field W", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Images) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.Cover.Set {
		more.More()
		j.WriteObjectField("cover")
		s.Cover.WriteJSON(j)
	}
	if s.Pages != nil {
		more.More()
		j.WriteObjectField("pages")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.Pages {
			more.More()
			elem.WriteJSON(j)
		}
		j.WriteArrayEnd()
		more.Up()
	}
	if s.Thumbnail.Set {
		more.More()
		j.WriteObjectField("thumbnail")
		s.Thumbnail.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Images json value to io.Writer.
func (s Images) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Images json value from io.Reader.
func (s *Images) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Images from json stream.
func (s *Images) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "cover":
			s.Cover.Reset()
			if err := s.Cover.ReadJSON(i); err != nil {
				i.ReportError("Field Cover", err.Error())
				return false
			}
			return true
		case "pages":
			// Unsupported kind "array" for field "Pages".
			i.Skip()
			return true
		case "thumbnail":
			s.Thumbnail.Reset()
			if err := s.Thumbnail.ReadJSON(i); err != nil {
				i.ReportError("Field Thumbnail", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON writes json value of Image to json stream.
func (o OptImage) WriteJSON(j *json.Stream) {
	o.Value.WriteJSON(j)
}

// ReadJSON reads json value of Image from json iterator.
func (o *OptImage) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.ObjectValue:
		o.Set = true
		if err := o.Value.ReadJSON(i); err != nil {
			return err
		}
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptImage", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of Images to json stream.
func (o OptImages) WriteJSON(j *json.Stream) {
	o.Value.WriteJSON(j)
}

// ReadJSON reads json value of Images from json iterator.
func (o *OptImages) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.ObjectValue:
		o.Set = true
		if err := o.Value.ReadJSON(i); err != nil {
			return err
		}
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptImages", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int to json stream.
func (o OptInt) WriteJSON(j *json.Stream) {
	j.WriteInt(int(o.Value))
}

// ReadJSON reads json value of int from json iterator.
func (o *OptInt) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = int(i.ReadInt())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptInt", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of string to json stream.
func (o OptString) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptString) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Value = string(i.ReadString())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptString", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of TagType to json stream.
func (o OptTagType) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of TagType from json iterator.
func (o *OptTagType) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Value = TagType(i.ReadString())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptTagType", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of Title to json stream.
func (o OptTitle) WriteJSON(j *json.Stream) {
	o.Value.WriteJSON(j)
}

// ReadJSON reads json value of Title from json iterator.
func (o *OptTitle) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.ObjectValue:
		o.Set = true
		if err := o.Value.ReadJSON(i); err != nil {
			return err
		}
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptTitle", i.WhatIsNext())
	}
	return nil
}

func (SearchByTagIDForbidden) WriteJSON(j *json.Stream)        {}
func (SearchByTagIDForbidden) ReadJSON(i *json.Iterator) error { return nil }
func (SearchByTagIDForbidden) ReadJSONFrom(r io.Reader) error  { return nil }
func (SearchByTagIDForbidden) WriteJSONTo(w io.Writer) error   { return nil }

func (SearchByTagIDOK) WriteJSON(j *json.Stream)        {}
func (SearchByTagIDOK) ReadJSON(i *json.Iterator) error { return nil }
func (SearchByTagIDOK) ReadJSONFrom(r io.Reader) error  { return nil }
func (SearchByTagIDOK) WriteJSONTo(w io.Writer) error   { return nil }

func (SearchForbidden) WriteJSON(j *json.Stream)        {}
func (SearchForbidden) ReadJSON(i *json.Iterator) error { return nil }
func (SearchForbidden) ReadJSONFrom(r io.Reader) error  { return nil }
func (SearchForbidden) WriteJSONTo(w io.Writer) error   { return nil }

func (SearchOK) WriteJSON(j *json.Stream)        {}
func (SearchOK) ReadJSON(i *json.Iterator) error { return nil }
func (SearchOK) ReadJSONFrom(r io.Reader) error  { return nil }
func (SearchOK) WriteJSONTo(w io.Writer) error   { return nil }

// WriteJSON implements json.Marshaler.
func (s SearchResponse) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.NumPages.Set {
		more.More()
		j.WriteObjectField("num_pages")
		s.NumPages.WriteJSON(j)
	}
	if s.PerPage.Set {
		more.More()
		j.WriteObjectField("per_page")
		s.PerPage.WriteJSON(j)
	}
	if s.Result != nil {
		more.More()
		j.WriteObjectField("result")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.Result {
			more.More()
			elem.WriteJSON(j)
		}
		j.WriteArrayEnd()
		more.Up()
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes SearchResponse json value to io.Writer.
func (s SearchResponse) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads SearchResponse json value from io.Reader.
func (s *SearchResponse) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads SearchResponse from json stream.
func (s *SearchResponse) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "num_pages":
			s.NumPages.Reset()
			if err := s.NumPages.ReadJSON(i); err != nil {
				i.ReportError("Field NumPages", err.Error())
				return false
			}
			return true
		case "per_page":
			s.PerPage.Reset()
			if err := s.PerPage.ReadJSON(i); err != nil {
				i.ReportError("Field PerPage", err.Error())
				return false
			}
			return true
		case "result":
			// Unsupported kind "array" for field "Result".
			i.Skip()
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Tag) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.Count.Set {
		more.More()
		j.WriteObjectField("count")
		s.Count.WriteJSON(j)
	}
	if s.ID.Set {
		more.More()
		j.WriteObjectField("id")
		s.ID.WriteJSON(j)
	}
	if s.Name.Set {
		more.More()
		j.WriteObjectField("name")
		s.Name.WriteJSON(j)
	}
	if s.Type.Set {
		more.More()
		j.WriteObjectField("type")
		s.Type.WriteJSON(j)
	}
	if s.URL.Set {
		more.More()
		j.WriteObjectField("url")
		s.URL.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Tag json value to io.Writer.
func (s Tag) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Tag json value from io.Reader.
func (s *Tag) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Tag from json stream.
func (s *Tag) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "count":
			s.Count.Reset()
			if err := s.Count.ReadJSON(i); err != nil {
				i.ReportError("Field Count", err.Error())
				return false
			}
			return true
		case "id":
			s.ID.Reset()
			if err := s.ID.ReadJSON(i); err != nil {
				i.ReportError("Field ID", err.Error())
				return false
			}
			return true
		case "name":
			s.Name.Reset()
			if err := s.Name.ReadJSON(i); err != nil {
				i.ReportError("Field Name", err.Error())
				return false
			}
			return true
		case "type":
			s.Type.Reset()
			if err := s.Type.ReadJSON(i); err != nil {
				i.ReportError("Field Type", err.Error())
				return false
			}
			return true
		case "url":
			s.URL.Reset()
			if err := s.URL.ReadJSON(i); err != nil {
				i.ReportError("Field URL", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s TagType) WriteJSON(j *json.Stream) {
	j.WriteString(string(s))
}

// ReadJSON reads TagType from json stream.
func (s *TagType) ReadJSON(i *json.Iterator) error {
	*s = TagType(i.ReadString())
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Title) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.English.Set {
		more.More()
		j.WriteObjectField("english")
		s.English.WriteJSON(j)
	}
	if s.Japanese.Set {
		more.More()
		j.WriteObjectField("japanese")
		s.Japanese.WriteJSON(j)
	}
	if s.Pretty.Set {
		more.More()
		j.WriteObjectField("pretty")
		s.Pretty.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Title json value to io.Writer.
func (s Title) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Title json value from io.Reader.
func (s *Title) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Title from json stream.
func (s *Title) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "english":
			s.English.Reset()
			if err := s.English.ReadJSON(i); err != nil {
				i.ReportError("Field English", err.Error())
				return false
			}
			return true
		case "japanese":
			s.Japanese.Reset()
			if err := s.Japanese.ReadJSON(i); err != nil {
				i.ReportError("Field Japanese", err.Error())
				return false
			}
			return true
		case "pretty":
			s.Pretty.Reset()
			if err := s.Pretty.ReadJSON(i); err != nil {
				i.ReportError("Field Pretty", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}
