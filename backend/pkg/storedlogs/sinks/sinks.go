package sinks

import . "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/model"

//This file serves as a sink interface for the storedlog package
//It is used to store logs in various different formats such as
//Files, Databases, and other configurations

type LogSink interface {
	StoreLog(log Log) error
}
