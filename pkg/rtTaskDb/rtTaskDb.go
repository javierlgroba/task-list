package rtTaskDb

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/javierlgroba/task-list/pkg/task"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type BunTask struct {
	bun.BaseModel `bun:"table:task_list"`

	ID           string    `bun:"id,pk,type:uuid"`
	Text         string    `bun:"text,notnull"`
	CreationTime time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type RtTaskDb struct {
	db  *bun.DB
	ctx context.Context
}

func (r *RtTaskDb) dbConnect() {
	connParams := make(map[string]interface{})
	connParams["sslmode"] = "disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(os.Getenv("DATABASE_IP_ADDRESS")+":5432"),
		pgdriver.WithUser(os.Getenv("API_POSTGRES_USER")),
		pgdriver.WithPassword(os.Getenv("API_POSTGRES_PASSWORD")),
		pgdriver.WithDatabase("task_list"),
		pgdriver.WithInsecure(true),
	))
	r.db = bun.NewDB(sqldb, pgdialect.New())
	r.db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	r.db.Exec("set schema 'task_list'")
}

func (r *RtTaskDb) dbDisconnect() {
	r.db.Close()
}

func NewRtTaskDb() RtTaskDb {
	var m RtTaskDb
	m.ctx = context.Background()
	m.dbConnect()
	defer m.dbDisconnect()
	_, err := m.db.NewCreateTable().
		Model((*BunTask)(nil)).
		IfNotExists().
		Varchar(100).
		Exec(m.ctx)
	if err != nil {
		panic(err)
	}
	return m
}

func (r RtTaskDb) GetAll() []task.Task {
	r.dbConnect()
	defer r.dbDisconnect()
	value := make([]*BunTask, 0)
	err := r.db.NewSelect().
		Model(&value).
		ExcludeColumn("creation_time").
		Order("creation_time DESC").
		Scan(r.ctx)
	if err != nil {
		fmt.Print(err)
		return make([]task.Task, 0)
	}
	result := make([]task.Task, 0, len(value))
	for _, t := range value {
		result = append(result, task.Task{ID: t.ID, Text: t.Text})
	}
	return result
}

func (r RtTaskDb) GetTask(s string) (task.Task, bool) {
	r.dbConnect()
	defer r.dbDisconnect()
	value := BunTask{ID: s}
	err := r.db.NewSelect().
		Model(&value).
		ExcludeColumn("creation_time").
		Scan(r.ctx)
	if err != nil {
		fmt.Print(err)
		return task.Task{}, false
	}
	result := task.Task{ID: value.ID, Text: value.Text}
	return result, true
}

func (r RtTaskDb) Remove(s string) bool {
	r.dbConnect()
	defer r.dbDisconnect()
	bunTask := BunTask{ID: s}
	res, err := r.db.NewDelete().
		Model(&bunTask).
		WherePK().
		Exec(r.ctx)
	if err != nil {
		fmt.Printf("Error removing a task: %v", err)
		return false
	}
	rows, _ := res.RowsAffected()
	return rows != 0
}

func (r RtTaskDb) RemoveAll() bool {
	r.dbConnect()
	defer r.dbDisconnect()
	err := r.db.ResetModel(r.ctx, (*BunTask)(nil))
	if err != nil {
		fmt.Errorf("Error removing all tasks: %w", err)
	}
	return err == nil
}

func (r RtTaskDb) Add(t task.Task) bool {
	r.dbConnect()
	defer r.dbDisconnect()
	bunTask := BunTask{ID: t.ID, Text: t.Text}
	_, err := r.db.NewInsert().
		Model(&bunTask).
		Exec(r.ctx)
	if err != nil {
		fmt.Errorf("Error adding a tasks: %w", err)
	}
	return err == nil
}
