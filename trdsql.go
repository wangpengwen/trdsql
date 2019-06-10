// Package trdsql implements execute SQL queries on tabular data.
//
// trdsql imports tabular data into a database,
// executes SQL queries, and executes exports.
package trdsql

import (
	"fmt"
	"log"
)

// TRDSQL structure is a structure that defines the whole operation.
type TRDSQL struct {
	Driver   string
	Dsn      string
	Importer Importer
	Exporter Exporter
}

// NewTRDSQL returns a new TRDSQL structure.
func NewTRDSQL(im Importer, ex Exporter) *TRDSQL {
	return &TRDSQL{
		Driver:   "sqlite3",
		Dsn:      "",
		Importer: im,
		Exporter: ex,
	}
}

// Format represents the import/export format
type Format int

// Represents Format
const (
	// import (guesses for import format)
	GUESS Format = iota
	// import/export
	CSV
	// import/export
	LTSV
	// import/export
	JSON
	// import/export
	TBLN
	// export
	RAW
	// export
	MD
	// export
	AT
	// export
	VF
)

func (f Format) String() string {
	switch f {
	case GUESS:
		return "GUESS"
	case CSV:
		return "CSV"
	case LTSV:
		return "LTSV"
	case JSON:
		return "JSON"
	case TBLN:
		return "TBLN"
	case RAW:
		return "RAW"
	case MD:
		return "MD"
	case AT:
		return "AT"
	case VF:
		return "VF"
	default:
		return "Unknown"
	}
}

// Exec is actually executed.
func (trd *TRDSQL) Exec(sql string) error {
	db, err := Connect(trd.Driver, trd.Dsn)
	if err != nil {
		return fmt.Errorf("ERROR(CONNECT):%s", err)
	}
	defer func() {
		err = db.Disconnect()
		if err != nil {
			log.Printf("ERROR(DISCONNECT):%s", err)
		}
	}()

	db.Tx, err = db.Begin()
	if err != nil {
		return fmt.Errorf("ERROR(BEGIN):%s", err)
	}

	if trd.Importer != nil {
		sql, err = trd.Importer.Import(db, sql)
		if err != nil {
			return fmt.Errorf("ERROR(IMPORT):%s", err)
		}
	}

	if trd.Exporter != nil {
		err = trd.Exporter.Export(db, sql)
		if err != nil {
			return fmt.Errorf("ERROR(EXPORT):%s", err)
		}
	}

	err = db.Tx.Commit()
	if err != nil {
		return fmt.Errorf("ERROR(COMMIT):%s", err)
	}

	return nil
}
