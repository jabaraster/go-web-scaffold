package model

import (
    "time"
)

type entity interface {
    setCreatedAt(time.Time)
    setUpdatedAt(time.Time)
}

func beforeInsertCore(e entity) error {
    n := time.Now()
    e.setCreatedAt(n)
    e.setUpdatedAt(n)
    return nil
}

func beforeUpdateCore(e entity) error {
    n := time.Now()
    e.setUpdatedAt(n)
    return nil
}
