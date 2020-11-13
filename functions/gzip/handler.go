package gzip

import (
	"bytes"
	"compress/gzip"
	f "github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("gzip.compress", compress, f.DefParams()),
		f.DefFunc("gzip.decompress", decompress, f.DefParams()),
	)
}

// gzip.compress()
func compress(_ f.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	w := gzip.NewWriter(out)
	if _, err = w.Write(input); err != nil {
		return err
	}
	return w.Close()
}

// gzip.decompress()
func decompress(_ f.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	r, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return err
	}
	uncompressed, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	_, err = out.Write(uncompressed)
	return err
}
