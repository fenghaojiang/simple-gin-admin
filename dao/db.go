package dao

import (
	"context"
	"database/sql"
	"simple-gin-admin/model"
	"time"
)

const (
	getSomeData = "Select * from douban_book limit ?, 10"
)

func (d *Dao) GetSomeData(ctx context.Context, id int) (bookInfo []*model.DoubanBook, err error) {
	timeCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if err := d.db.SelectContext(timeCtx, &bookInfo, getSomeData, id); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	return bookInfo, nil
}
