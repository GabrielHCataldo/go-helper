package helper

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
	"io"
)

func CompressWithGzip(a any) ([]byte, error) {
	bs, err := ConvertToBytes(a)
	if IsNotNil(err) {
		return nil, err
	}

	var gzipBuffer bytes.Buffer
	gz := gzip.NewWriter(&gzipBuffer)
	defer gz.Close()

	_, err = gz.Write(bs)
	if IsNotNil(err) {
		return nil, err
	}

	return gzipBuffer.Bytes(), nil
}

func CompressWithGzipToBase64(a any) (string, error) {
	bs, err := CompressWithGzip(a)
	if IsNotNil(err) {
		return "", err
	}
	return ConvertToBase64(bs)
}

func CompressWithDeflate(a any) ([]byte, error) {
	bs, err := ConvertToBytes(a)
	if IsNotNil(err) {
		return nil, err
	}

	var buffer bytes.Buffer
	writer, err := flate.NewWriter(&buffer, 9)
	if IsNotNil(err) {
		return nil, err
	}

	_, err = writer.Write(bs)
	if IsNotNil(err) {
		return nil, err
	}
	defer writer.Close()

	return buffer.Bytes(), nil
}

func DecompressWithGzip(a any) ([]byte, error) {
	bs, err := ConvertToBytes(a)
	if IsNotNil(err) {
		return nil, err
	}

	reader, err := gzip.NewReader(bytes.NewReader(bs))
	if IsNotNil(err) {
		return nil, err
	}
	defer reader.Close()

	return io.ReadAll(reader)
}

func DecompressWithDeflate(a any) ([]byte, error) {
	bs, err := ConvertToBytes(a)
	if IsNotNil(err) {
		return nil, err
	}

	reader := flate.NewReader(bytes.NewReader(bs))
	defer reader.Close()

	return io.ReadAll(reader)
}

func DecompressFromBase64WithGzip(a any) ([]byte, error) {
	s64, err := ConvertToString(a)
	if IsNotNil(err) {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(s64)
	if IsNotNil(err) {
		return nil, err
	}

	return DecompressWithGzip(decoded)
}

func DecompressFromBase64WithDeflate(a any) ([]byte, error) {
	s64, err := ConvertToString(a)
	if IsNotNil(err) {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(s64)
	if IsNotNil(err) {
		return nil, err
	}

	return DecompressWithDeflate(decoded)
}

func DecompressWithGzipToDest(a, dest any) error {
	bs, err := DecompressWithGzip(a)
	if err != nil {
		return err
	}
	return ConvertToDest(bs, dest)
}

func DecompressWithDeflateToDest(a, dest any) error {
	bs, err := DecompressWithDeflate(a)
	if err != nil {
		return err
	}
	return ConvertToDest(bs, dest)
}

func DecompressFromBase64WithGzipToDest(a, dest any) error {
	bs, err := DecompressFromBase64WithGzip(a)
	if err != nil {
		return err
	}
	return ConvertToDest(bs, dest)
}

func DecompressFromBase64WithDeflateToDest(a, dest any) error {
	bs, err := DecompressFromBase64WithGzip(a)
	if err != nil {
		return err
	}
	return ConvertToDest(bs, dest)
}
