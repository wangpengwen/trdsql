package trdsql

// SliceImporter is a structure that includes SliceReader.
// SliceImporter can be used as a library from another program.
// It is not used from the command.
// SliceImporter is an importer that reads one slice data.
type SliceImporter struct {
	*SliceReader
}

// NewSliceImporter returns trdsql SliceImporter.
func NewSliceImporter(tableName string, data interface{}) *SliceImporter {
	return &SliceImporter{
		SliceReader: NewSliceReader(tableName, data),
	}
}

// Import is a method to import from SliceReader in SliceImporter.
func (i *SliceImporter) Import(db *DB, query string) (string, error) {
	names, err := i.Names()
	if err != nil {
		return query, err
	}
	types, err := i.Types()
	if err != nil {
		return query, err
	}
	err = db.CreateTable(i.tableName, names, types, true)
	if err != nil {
		return query, err
	}
	err = db.Import(i.tableName, names, i.SliceReader)
	return query, err
}
