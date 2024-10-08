package cache

import (
	"github.com/sirupsen/logrus"
	"sync"
)

// DataService 数据服务缓存
var DataService = cacheDataService{
	rwMutex:     &sync.RWMutex{},
	dataService: make(map[string]map[string]*DataDataService),
}

type DataDataService struct {
	Method  string
	Source  string
	Timeout int
}

type cacheDataService struct {
	rwMutex *sync.RWMutex

	dataService map[string]map[string]*DataDataService
}

func (cache *cacheDataService) Set(table string, service string, data *DataDataService) {
	cache.rwMutex.Lock()
	defer cache.rwMutex.Unlock()

	mapDataService, ok := cache.dataService[table]
	if !ok {
		mapDataService = make(map[string]*DataDataService)
	}

	mapDataService[service] = data
	cache.dataService[table] = mapDataService
	logrus.Infof("更新缓存用户的数据服务【%s.%s】...", table, service)
}

func (cache *cacheDataService) Get(table string, service string) (*DataDataService, bool) {
	cache.rwMutex.RLock()
	defer cache.rwMutex.RUnlock()

	mapDataService, ok := cache.dataService[table]
	if !ok {
		return nil, false
	}

	data, ok := mapDataService[service]
	logrus.Infof("通过缓存读取数据服务【%s.%s】>>>", table, service)

	return data, ok
}

func (cache *cacheDataService) Delete(table string, service string) {
	cache.rwMutex.Lock()
	defer cache.rwMutex.Unlock()

	mapDataService, ok := cache.dataService[table]
	if !ok {
		return
	}

	delete(mapDataService, service)
	logrus.Infof("删除缓存用户的数据服务【%s.%s】...", table, service)
}

func (cache *cacheDataService) DeleteByTable(table string) {
	cache.rwMutex.Lock()
	defer cache.rwMutex.Unlock()

	delete(cache.dataService, table)
	logrus.Infof("删除缓存用户的数据服务【%s】...", table)
}
