package model

import (
	"fmt"
	"time"

	"github.com/mickep76/qry"
	"github.com/pborman/uuid"
)

type Pool struct {
	UUID    string     `json:"uuid"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated,omitempty"`
	Name    string     `json:"name"`
}

type Pools []*Pool

func NewPool(name string) *Pool {
	return &Pool{
		UUID:    uuid.New(),
		Created: time.Now(),
		Name:    name,
	}
}

func (ds *Datastore) AllPools() (Pools, error) {
	kvs, err := ds.Values("pools")
	if err != nil {
		return nil, err
	}

	pools := Pools{}
	if err := kvs.Decode(&pools); err != nil {
		return nil, err
	}

	return pools, nil
}

func (ds *Datastore) QueryPools(q *qry.Query) (Pools, error) {
	pools, err := ds.AllPools()
	if err != nil {
		return nil, err
	}

	filtered, err := q.Query(pools)
	if err != nil {
		return nil, err
	}

	return filtered.(Pools), nil
}

func (ds *Datastore) OnePool(uuid string) (*Pool, error) {
	kvs, err := ds.Values(fmt.Sprintf("pools/%s", uuid))
	if err != nil {
		return nil, err
	}

	pools := Pools{}
	if err := kvs.Decode(&pools); err != nil {
		return nil, err
	}

	if len(pools) > 0 {
		return pools[0], nil
	}

	return nil, nil
}

func (ds *Datastore) CreatePool(pool *Pool) error {
	return ds.Set(fmt.Sprintf("pools/%s", pool.UUID), pool)
}

func (ds *Datastore) UpdatePool(pool *Pool) error {
	now := time.Now()
	pool.Updated = &now
	return ds.Set(fmt.Sprintf("pools/%s", pool.UUID), pool)
}

func (ds *Datastore) DeletePool(uuid string) error {
	if err := ds.Delete(fmt.Sprintf("pools/%s", uuid)); err != nil {
		return err
	}
	return nil
}