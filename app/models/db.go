package models

import (
	"strings"

	"time"

	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"github.com/runner-mei/orm"
)

type DB struct {
	Engine *xorm.Engine
}

type BackgroundTask struct {
	ID        int64     `json:"id" xorm:"id pk autoincr"`
	Name      string    `json:"name,omitempty" xorm:"name"`
	StartedAt time.Time `json:"started_at,omitempty" xorm:"started_at"`
	EndedAt   time.Time `json:"ended_at,omitempty" xorm:"ended_at"`
}

func (backgroundTask *BackgroundTask) TableName() string {
	return "tpt_background_tasks"
}

type BackgroundTaskProgress struct {
	ID               int64     `json:"id" xorm:"id pk autoincr"`
	BackgroundTaskId int64     `json:"background_task_id,omitempty" xorm:"background_task_id"`
	Description      string    `json:"description,omitempty" xorm:"description"`
	CreatedAt        time.Time `json:"created_at,omitempty" xorm:"created_at created"`
}

func (backgroundTaskProgress *BackgroundTaskProgress) TableName() string {
	return "tpt_background_task_progresses"
}

func (db *DB) BackgroundTasks() *orm.Collection {
	return orm.New(func() interface{} {
		return &BackgroundTask{}
	})(db.Engine)
}
func (db *DB) BackgroundTaskProgresses() *orm.Collection {
	return orm.New(func() interface{} {
		return &BackgroundTaskProgress{}
	})(db.Engine)
}

func InitTables(engine *xorm.Engine) error {
	beans := []interface{}{
		&BackgroundTask{},
		&BackgroundTaskProgress{},
	}

	if err := engine.CreateTables(beans...); err != nil {
		return err
	}

	for _, bean := range beans {
		if err := engine.CreateIndexes(bean); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return err
			}
			revel.WARN.Println(err)
		}

		if err := engine.CreateUniques(bean); err != nil {
			if !(strings.Contains(err.Error(), "already exists") ||
				strings.Contains(err.Error(), "已经存在")) &&
				!(strings.Contains(err.Error(), "create unique index") ||
					strings.Contains(err.Error(), "UQE_")) {
				return err
			}
			revel.WARN.Println(err)
		}
	}
	return nil
}
