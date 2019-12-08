package master

import (
	"1.imooc_crontab/mycrontab/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// mongodb日志管理
type LogMgr struct {
	mongo.Client
	client *mongo.Client
	logCollection *mongo.Collection
}

var (
	G_logMgr *LogMgr
)

func InitLogMgr() (err error) {
	var (
		client *mongo.Client
	)

	// 建立mongodb连接
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(G_config.MongodbUri))
	if err != nil {
		return
	}

	G_logMgr = &LogMgr{
		client: client,
		logCollection: client.Database("cron").Collection("log"),
	}
	return
}

// 查看任务日志
func (logMgr *LogMgr) ListLog(name string, skip int, limit int) (logArr []*common.JobLog, err error){
	var (
		filter *common.JobLogFilter
		//logSort *common.SortLogByStartTime
		//cursor mongo.Cursor
		jobLog *common.JobLog
	)

	// len(logArr)
	logArr = make([]*common.JobLog, 0)

	// 过滤条件
	filter = &common.JobLogFilter{JobName: name}

	// 按照任务开始时间倒排
	//logSort = &common.SortLogByStartTime{SortOrder: -1}

	// 查询
	options := options.Find()

	// Sort by `_id` field descending
	options.SetSort(bson.D{{"SortOrder", -1}})

	options.SetSkip(int64(skip))

	// Limit by 10 documents only
	options.SetLimit(int64(limit))


/*
		cursor, err = logMgr.logCollection.Find(context.TODO(),
		filter, findopt.Sort(logSort), findopt.Skip(int64(skip)), findopt.Limit(int64(limit)))
*/
	cursor, err := logMgr.logCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return
	}
	// 延迟释放游标
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		jobLog = &common.JobLog{}

		// 反序列化BSON
		if err = cursor.Decode(jobLog); err != nil {
			continue // 有日志不合法
		}

		logArr = append(logArr, jobLog)
	}
	return
}