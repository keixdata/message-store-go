package eventstore

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type postgresEventStore struct {
	conn *pgx.Conn
}

func (ev *postgresEventStore) Write(msg *Message) (uint64, error) {
	log.Println("Writing", msg.ID)

	if msg.ID == "" {
		return 0, fmt.Errorf("Cannot write a message without an ID")
	}

	if msg.StreamName == "" {
		return 0, fmt.Errorf("Cannot write a message without a stream name")
	}

	sql := "SELECT write_message($1,$2,$3,$4,$5)"
	rows, err := ev.conn.Query(context.Background(), sql, msg.ID, msg.StreamName, msg.Type, msg.Data, msg.Metadata)

	if !rows.Next() {
		return 0, fmt.Errorf("No valid response from the postrgres sql db")
	}
	var position int64
	err = rows.Scan(&position)
	if err != nil {
		return 0, fmt.Errorf("No valid response from the postgres sql fb: %v", err)
	}

	fmt.Print(position)
	return 0, err
}

// WithPgConnString is.
func WithPgConnString(connString string) (EventStore, error) {
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
		return nil, err
	}
	fmt.Println(config)

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	//defer conn.Close(context.Background())

	ev := &postgresEventStore{conn: conn}
	return ev, nil

}

type WriteRes struct {
	Position int64 `db:"position"`
}
