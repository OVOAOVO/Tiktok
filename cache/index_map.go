package cache

import (
	"context"
	"fmt"

	"github.com/config"
	"github.com/redis/go-redis"
)

// 用户id->被点赞的视频id集合->是否含有该视频id

var ctx = context.Background() //返回一个空的context
var rdb *redis.Client          // redis的链接

// const #define     代替
const (
	favor    = "favor"
	relation = "relation"
)

func init() { //redis 新建一个链接
	rdb = redis.NewClient(
		&redis.Options{ //链接设置
			Addr:     fmt.Sprintf("%s:%d", config.Info.RDB.IP, config.Info.RDB.Port),
			Password: "",                       //没有设置密码
			DB:       config.Info.RDB.Database, //选择设置的数据库   这里默认0
		})
}

var (
	proxyIndexOperation ProxyIndexMap //  名  类
)

type ProxyIndexMap struct {
}

func NewProxyIndexMap() *ProxyIndexMap { //指针 返回类型
	return &proxyIndexOperation
}

// 实例一个i   用户id   视频 id   点赞状态
// UpdateVideoFavorState 更新点赞状态，state:true为点赞，false为取消点赞
func (i *ProxyIndexMap) UpdateVideoFavorState(userId int64, videoId int64, state bool) {
	//格式化一个 喜欢的：id（参）
	key := fmt.Sprintf("%s:%d", favor, userId)
	if state { //更新状态
		rdb.SAdd(ctx, key, videoId) //传入context  内容key  视频id
		return
	}
	rdb.SRem(ctx, key, videoId) //删除喜欢
}

// GetVideoFavorState 得到点赞状态
func (i *ProxyIndexMap) GetVideoFavorState(userId int64, videoId int64) bool {
	key := fmt.Sprintf("%s:%d", favor, userId)
	ret := rdb.SIsMember(ctx, key, videoId) //返回并判断 是否已经存在
	return ret.Val()
}

// UpdateUserRelation 更新点赞状态，state:true为点关注，false为取消关注
func (i *ProxyIndexMap) UpdateUserRelation(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", relation, userId)
	if state {
		rdb.SAdd(ctx, key, followId)
		return
	}
	rdb.SRem(ctx, key, followId)
}

// GetUserRelation 得到关注状态
func (i *ProxyIndexMap) GetUserRelation(userId int64, followId int64) bool {
	key := fmt.Sprintf("%s:%d", relation, userId)
	ret := rdb.SIsMember(ctx, key, followId)
	return ret.Val()
}
