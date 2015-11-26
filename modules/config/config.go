package config

// Enums
const (
	VIPER = iota
	DEFAULT

	PATH    = "./config/"
	ExtName = ".json"
)

// Config Struct
type Config interface {
	SetConfig(string) error
	Get(key string) interface{}
	Set(key string, value interface{}) bool
}

// CreateConfig new Config
func CreateConfig() (Config, error) {
	t := VIPER

	switch t {
	case VIPER:
		return new(Viper), nil
	case DEFAULT:
	default:
		return new(Default), nil
	}
	return new(Default), nil
}
