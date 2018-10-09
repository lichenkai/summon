package summon

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ContextHandler interface {
	ServeHTTP(context.Context, ResponseWriter, *Request)
}

type ContextHandlerFunc func(context.Context, ResponseWriter, *Request)

func (h ContextHandlerFunc) ServeHTTP(ctx context.Context, w ResponseWriter, r *Request) {
	h(ctx, w, r)
}

type BytesReadCloser struct {
	*bytes.Reader
}

func (rc *BytesReadCloser) Close() error {
	return nil
}

func Handle(ctx context.Context, h ContextHandler, maxBodyBytes int64) httprouter.Handle {
	f := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var (
			request        *Request
			response       *responseWriter
			err            error
			newctx, cancel = context.WithCancel(ctx)
		)

		defer cancel()

		request = &Request{
			Request: r,
			Path:    params,
		}

		response = &responseWriter{
			ResponseWriter: w,
			wroteHeader:    false,
		}

		if maxBodyBytes > 0 {
			r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
		}

		request.RawBody, err = ioutil.ReadAll(r.Body)
		r.Body.Close()

		r.Body = &BytesReadCloser{
			Reader: bytes.NewReader(request.RawBody),
		}

		err = r.ParseMultipartForm(maxBodyBytes)
		switch {
		case err == http.ErrNotMultipart:
			err = nil
		case err != nil:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		h.ServeHTTP(newctx, response, request)
	}
	return httprouter.Handle(f)
}

func StdHandler(ctx context.Context, h ContextHandler, maxBodyBytes int64) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		var (
			request        *Request
			response       *responseWriter
			err            error
			newctx, cancel = context.WithCancel(ctx)
		)

		defer cancel()

		request = &Request{
			Request: r,
		}

		response = &responseWriter{
			ResponseWriter: w,
			wroteHeader:    false,
		}

		if maxBodyBytes > 0 {
			r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
		}

		request.RawBody, err = ioutil.ReadAll(r.Body)
		r.Body.Close()

		r.Body = &BytesReadCloser{
			Reader: bytes.NewReader(request.RawBody),
		}

		err = r.ParseMultipartForm(maxBodyBytes)
		switch {
		case err == http.ErrNotMultipart:
			err = nil
		case err != nil:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		h.ServeHTTP(newctx, response, request)
	}
	return http.HandlerFunc(f)
}
