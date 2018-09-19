package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/micro/go-log"
	. "github.com/opensds/go-panda/s3/pkg/exception"
)

func (ad *adapter) DeleteBucket(bucketName string) S3Error {
	//Check if the connctor exist or not
	ss := ad.s.Copy()
	defer ss.Close()

	//Delete it from database
	c := ss.DB(DataBaseName).C(BucketMD)
	err := c.Remove(bson.M{"name": bucketName})
	if err == mgo.ErrNotFound {
		log.Log("Delete bucket failed, err:the specified bucket does not exist.")
		return NoSuchBucket
	} else if err != nil {
		log.Log("Delete bucket from database failed,err:%v.\n", err)
		return InternalError
	}
	cc := ss.DB(DataBaseName).C(bucketName)
	deleteErr := cc.DropCollection()
	if deleteErr != nil && deleteErr != mgo.ErrNotFound {
		log.Log("Delete bucket collection from database failed,err:%v.\n", deleteErr)
		return InternalError
	}

	return NoError
}
