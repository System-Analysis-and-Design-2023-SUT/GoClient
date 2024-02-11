package sadqueue

import "errors"

var ErrLiveHostNotFound = errors.New("no available host found")
var ErrHostNotAvailable = errors.New("can not connect to host")
var ErrPushFailed = errors.New("can not push message to server")
var ErrPullFailed = errors.New("can not pull message from server")
var ErrSubscribeFailed = errors.New("can not subscribe to server")
