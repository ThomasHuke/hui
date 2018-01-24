package hui

import (
	"database/sql"
)

type DB struct {
	proto string
	raddr string
}

func (d DB) Open(name string) (Conn, error) {

}

func init() {
	sql.Register("hui", &DB{proto: "TCP", raddr: "127.0.0.1:3306"})
}
