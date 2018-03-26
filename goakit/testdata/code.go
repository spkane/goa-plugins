package testdata

var SimpleMethodGoakitResponseEncoderCode = `// EncodeSimpleMethodResponse returns a go-kit EncodeResponseFunc suitable for
// encoding SimpleService SimpleMethod responses.
func EncodeSimpleMethodResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSimpleMethodResponse(encoder)
}
`

var WithPayloadMethodGoakitResponseEncoderCode = `// EncodeWithPayloadMethodResponse returns a go-kit EncodeResponseFunc suitable
// for encoding WithPayloadService WithPayloadMethod responses.
func EncodeWithPayloadMethodResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeWithPayloadMethodResponse(encoder)
}
`

var WithErrorMethodGoakitResponseEncoderCode = `// EncodeWithErrorMethodResponse returns a go-kit EncodeResponseFunc suitable
// for encoding WithErrorService WithErrorMethod responses.
func EncodeWithErrorMethodResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeWithErrorMethodResponse(encoder)
}
`

var Endpoint1GoakitResponseEncoderCode = `// EncodeEndpoint1Response returns a go-kit EncodeResponseFunc suitable for
// encoding MultiEndpointService Endpoint1 responses.
func EncodeEndpoint1Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeEndpoint1Response(encoder)
}
`

var Endpoint2GoakitResponseEncoderCode = `// EncodeEndpoint2Response returns a go-kit EncodeResponseFunc suitable for
// encoding MultiEndpointService Endpoint2 responses.
func EncodeEndpoint2Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeEndpoint2Response(encoder)
}
`

var WithPayloadMethodGoakitRequestDecoderCode = `// DecodeWithPayloadMethodRequest returns a go-kit DecodeRequestFunc suitable
// for decoding WithPayloadService WithPayloadMethod requests.
func DecodeWithPayloadMethodRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeWithPayloadMethodRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}
`

var Endpoint1GoakitRequestDecoderCode = `// DecodeEndpoint1Request returns a go-kit DecodeRequestFunc suitable for
// decoding MultiEndpointService Endpoint1 requests.
func DecodeEndpoint1Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeEndpoint1Request(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}
`

var WithErrorMethodGoakitErrorEncoderCode = `// EncodeWithErrorMethodError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the WithErrorService WithErrorMethod endpoint.
func EncodeWithErrorMethodError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.ErrorEncoder {
	enc := server.EncodeWithErrorMethodError(encoder)
	return func(ctx context.Context, err error, w http.ResponseWriter) error {
		enc(ctx, w, err)
		return nil
	}
}
`

var Endpoint1GoakitErrorEncoderCode = `// EncodeEndpoint1Error returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the MultiEndpointService Endpoint1 endpoint.
func EncodeEndpoint1Error(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.ErrorEncoder {
	enc := server.EncodeEndpoint1Error(encoder)
	return func(ctx context.Context, err error, w http.ResponseWriter) error {
		enc(ctx, w, err)
		return nil
	}
}
`

var Endpoint2GoakitErrorEncoderCode = `// EncodeEndpoint2Error returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the MultiEndpointService Endpoint2 endpoint.
func EncodeEndpoint2Error(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.ErrorEncoder {
	enc := server.EncodeEndpoint2Error(encoder)
	return func(ctx context.Context, err error, w http.ResponseWriter) error {
		enc(ctx, w, err)
		return nil
	}
}
`

var SimpleMethodGoakitResponseDecoderCode = `// DecodeSimpleMethodResponse returns a go-kit DecodeResponseFunc suitable for
// decoding SimpleService SimpleMethod responses.
func DecodeSimpleMethodResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeSimpleMethodResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
`

var WithPayloadMethodGoakitResponseDecoderCode = `// DecodeWithPayloadMethodResponse returns a go-kit DecodeResponseFunc suitable
// for decoding WithPayloadService WithPayloadMethod responses.
func DecodeWithPayloadMethodResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeWithPayloadMethodResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
`

var WithErrorMethodGoakitResponseDecoderCode = `// DecodeWithErrorMethodResponse returns a go-kit DecodeResponseFunc suitable
// for decoding WithErrorService WithErrorMethod responses.
func DecodeWithErrorMethodResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeWithErrorMethodResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
`

var Endpoint1GoakitResponseDecoderCode = `// DecodeEndpoint1Response returns a go-kit DecodeResponseFunc suitable for
// decoding MultiEndpointService Endpoint1 responses.
func DecodeEndpoint1Response(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeEndpoint1Response(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
`

var Endpoint2GoakitResponseDecoderCode = `// DecodeEndpoint2Response returns a go-kit DecodeResponseFunc suitable for
// decoding MultiEndpointService Endpoint2 responses.
func DecodeEndpoint2Response(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeEndpoint2Response(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
`

var WithPayloadMethodGoakitRequestEncoderCode = `// EncodeWithPayloadMethodRequest returns a go-kit EncodeRequestFunc suitable
// for encoding WithPayloadService WithPayloadMethod requests.
func EncodeWithPayloadMethodRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeWithPayloadMethodRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}
`

var Endpoint1GoakitRequestEncoderCode = `// EncodeEndpoint1Request returns a go-kit EncodeRequestFunc suitable for
// encoding MultiEndpointService Endpoint1 requests.
func EncodeEndpoint1Request(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeEndpoint1Request(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}
`

var Endpoint1GoakitMountCode = `// MountEndpoint1Handler configures the mux to serve the "MultiEndpointService"
// service "Endpoint1" endpoint.
func MountEndpoint1Handler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/", f)
}
`

var Endpoint2GoakitMountCode = `// MountEndpoint2Handler configures the mux to serve the "MultiEndpointService"
// service "Endpoint2" endpoint.
func MountEndpoint2Handler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/", f)
}
`

var MixedMethodGoakitMountCode = `// MountMixedMethodHandler configures the mux to serve the "MixedService"
// service "MixedMethod" endpoint.
func MountMixedMethodHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/", f)
}
`

var File1GoakitMountCode = `// MountFile1JSON configures the mux to serve GET request made to "/1.json".
func MountFile1JSON(mux goahttp.Muxer) {
	mux.Handle("GET", "/1.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../file1.json")
	}))
}
`

var File2GoakitMountCode = `// MountFile2JSON configures the mux to serve GET request made to "/2.json".
func MountFile2JSON(mux goahttp.Muxer) {
	mux.Handle("GET", "/2.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../file2.json")
	}))
}
`

var MixedFileGoakitMountCode = `// MountMixedFileJSON configures the mux to serve GET request made to "/1.json".
func MountMixedFileJSON(mux goahttp.Muxer) {
	mux.Handle("GET", "/1.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../mixed_file.json")
	}))
}
`
