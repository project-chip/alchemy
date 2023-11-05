package db

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/iancoleman/strcase"
)

type Host struct {
	db *memory.Database

	lock sync.RWMutex
	docs []*ascii.Doc

	tableNames []string
	tables     map[string]*memory.Table

	base *sectionInfo

	ids map[string]int32
}

var (
	dbName = "MatterSpec"
)

type dbRowSet struct {
	rows []*dbRow
}

type dbRow struct {
	values map[matter.TableColumn]interface{}
	extras map[string]interface{}
}

func newDBRow() *dbRow {
	return &dbRow{values: make(map[matter.TableColumn]interface{})}
}

type extraInfo struct {
	name  string
	value string
}

func New() *Host {

	h := &Host{
		db:     memory.NewDatabase(dbName),
		tables: make(map[string]*memory.Table),
		ids:    make(map[string]int32),
	}

	h.db.EnablePrimaryKeyIndexes()
	return h
}

func (h *Host) Load(doc *ascii.Doc) error {
	h.lock.Lock()
	h.docs = append(h.docs, doc)
	h.lock.Unlock()
	return nil
}

func (h *Host) Run(address string, port int) error {
	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
	}
	engine := sqle.NewDefault(
		memory.NewDBProvider(
			h.db,
		))
	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Starting DB at %s:%d...", address, port)
	if err = s.Start(); err != nil {
		return err
	}
	return nil
}

func (h *Host) nextId(s string) int32 {
	id, ok := h.ids[s]
	if !ok {
		h.tableNames = append(h.tableNames, s)
	}
	id++
	h.ids[s] = id
	return id
}

func parseHex(s string) (int64, error) {
	s = strings.TrimPrefix(s, "0x")
	return strconv.ParseInt(s, 16, 64)
}

func toDBName(s string) string {
	return strcase.ToSnake(s)
}

func parseNumber(s string) (int64, error) {
	u, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		u, err = parseHex(s)
	}
	return u, err
}
