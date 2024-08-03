package dao

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type Relation struct {
	ID        int64
	CreatedAt time.Time

	DistributeID string
	FollowID     int64
	FansID       int64
}

var (
	RelationCount int64 = 0
)

const createRelationTable = `create table if not exists ` + "`%s`" + `
(
    id bigint      auto_increment primary key,
    created_at     timestamp      default CURRENT_TIMESTAMP not null,
    follow_id      int                                      not null,
    fans_id        int                                      not null,
	unique key (fans_id, follow_id)
)`

func (v RelationDao) SepRelation(dbname string) (string, error) {
	atomic.AddInt64(&RelationCount, 1)
	// 方便测试，使用一个小的分表指标
	tablename := fmt.Sprintf("relation_%d", atomic.LoadInt64(&RelationCount)/10)
	if VideoCount%10 == 0 {
		database := DatabasePool[dbname].DB
		if err := database.Exec(fmt.Sprintf(createRelationTable, tablename)); err != nil {
			return "", nil
		}
	}
	return tablename, nil
}

type RelationDao struct{}

func (r RelationDao) Create(dbname string, followID, fansID int64) error {
	var (
		err      error
		database = DatabasePool[dbname].DB
	)
	tablename, err := r.SepRelation(dbname)
	if err != nil {
		return err
	}
	relation := &Relation{
		DistributeID: fmt.Sprintf("%v-%v-%v", uuid1.String(), dbname, tablename),
		FollowID:     followID,
		FansID:       fansID,
	}
	if err = database.Table(tablename).Create(relation).Error; err != nil {
		return err
	}
	go func() {
		err = CacheDB.SAdd(fmt.Sprintf("relation_%v", fansID), followID).Err()
		if err != nil {
			log.Println(err)
		}
	}()
	return nil
}

func (r RelationDao) Delete(dbname string, followID, fansID int64) error {
	var (
		err      error
		database = DatabasePool[dbname].DB
	)
	relation := &Relation{
		DistributeID: fmt.Sprintf("%v-%v", uuid1.String(), dbname),
		FollowID:     followID,
		FansID:       fansID,
	}

	if err = database.Where("fans_id = ? AND follow_id = ?", fansID, followID).Delete(relation).Error; err != nil {
		log.Println(err.Error())
	}
	go func() {
		err = CacheDB.SRem(fmt.Sprintf("relation_%v", fansID), followID).Err()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}
