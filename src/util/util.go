package util

import (
    "strconv"
)

func Atoi32(s string) (*int32, error) {
    if len(s) == 0 {
        return nil, nil
    }
    i, err := strconv.ParseInt(s, 10, 32)
    if err != nil {
        return nil, err
    }

    ret := int32(i)
    return &ret, nil
}

func Atoi64(s string) (*int64, error) {
    if len(s) == 0 {
        return nil, nil
    }
    i, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        return nil, err
    }

    return &i, nil
}
