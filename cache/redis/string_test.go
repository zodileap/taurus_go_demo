package redis

import (
	"fmt"
	"testing"

	"github.com/yohobala/taurus_go/cache/redis"
)

func TestStringExample(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	}
	redis.SetClient(name, options)
	c, err := redis.GetClient("test")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	s, err := c.String()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("string_key", 0, "value")
	s.Get("string_key")
	s.Add("string_key2", 0, "value2")
	s.Del("string_key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetString("string_key")
	fmt.Println(keyRes)
	keyRes2 := r.GetString("string_key2")
	fmt.Println(keyRes2)
	if err != nil {
		fmt.Println(err)
	}
	s, err = c.String()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("string_key", 0, "value666")
	c.Save()
}

func TestStringAdd(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	}
	redis.SetClient(name, options)
	c, err := redis.GetClient("test")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	s, err := c.String()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("string_key", 0, "value")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetString("string_key")
	fmt.Println(keyRes)
}

func TestStringGet(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	}
	redis.SetClient(name, options)
	c, err := redis.GetClient("test")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	s, err := c.String()
	if err != nil {
		fmt.Println(err)
	}
	s.Get("string_key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetString("string_key")
	fmt.Println(keyRes.Value)
}

func TestStringDel(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	}
	redis.SetClient(name, options)
	c, err := redis.GetClient("test")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	s, err := c.String()
	if err != nil {
		fmt.Println(err)
	}
	s.Del("string_key")
	_, err = c.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func TestStringAddR(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	}
	redis.SetClient(name, options)
	c, err := redis.GetClient("test")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	s, err := c.String()
	if err != nil {
		fmt.Println(err)
	}
	err = s.AddR("string_key", 0, "value")
	if err != nil {
		fmt.Println(err)
	}
}

func TestStringGetR(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	}
	redis.SetClient(name, options)
	c, err := redis.GetClient("test")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	s, err := c.String()
	if err != nil {
		fmt.Println(err)
	}
	r, err := s.GetR("string_key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestStringDelR(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	}
	redis.SetClient(name, options)
	c, err := redis.GetClient("test")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
	}
	s, err := c.String()
	if err != nil {
		fmt.Println(err)
	}
	err = s.DelR("string_key")
	if err != nil {
		fmt.Println(err)
	}
}
