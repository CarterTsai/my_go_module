package config

// Enums
const (
	VIPER = iota
	DEFAULT

	PATH = "./config/"
)

// Config Struct
type Config interface {
	SetConfig(string) error
	Get(key string) interface{}
}

// CreateConfig new Config
func CreateConfig(t int) (Config, error) {
	switch t {
	case VIPER:
		return new(Viper), nil
	case DEFAULT:
	default:
		return new(Default), nil
	}
	return new(Default), nil
}
