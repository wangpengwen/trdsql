package trdsql

import (
	"encoding/csv"
)

// CSVWriter provides methods of the Writer interface.
type CSVWriter struct {
	writer    *csv.Writer
	results   []string
	outHeader bool
}

// NewCSVWriter returns CSVWriter.
func NewCSVWriter(writeOpts *WriteOpts) *CSVWriter {
	var err error
	w := &CSVWriter{}
	w.writer = csv.NewWriter(writeOpts.OutStream)
	w.writer.Comma, err = delimiter(writeOpts.OutDelimiter)
	if err != nil {
		debug.Printf("%s\n", err)
	}
	w.outHeader = writeOpts.OutHeader
	return w
}

// PreWrite is output of header and preparation.
func (w *CSVWriter) PreWrite(columns []string, types []string) error {
	w.results = make([]string, len(columns))
	if w.outHeader {
		return w.writer.Write(columns)
	}
	return nil
}

// WriteRow is row write.
func (w *CSVWriter) WriteRow(values []interface{}, columns []string) error {
	for i, col := range values {
		w.results[i] = ValString(col)
	}
	return w.writer.Write(w.results)
}

// PostWrite is flush.
func (w *CSVWriter) PostWrite() error {
	w.writer.Flush()
	return nil
}
