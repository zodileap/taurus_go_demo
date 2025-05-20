package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/zodileap/taurus_go/cache/redis"
)

func TestSetExample(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("set_key", 0, "value")
	s.Get("set_key")
	s.Add("set_key", 0, "value3")
	s.Add("set_key2", 0, "value2")
	s.Del("set_key", "value")
	s.Get("set_key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("set_key")
	fmt.Println(keyRes)
	if err != nil {
		fmt.Println(err)
	}
	s, err = c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("set_key", 0, "value666")
	c.Save()
}

func TestSetExampleDelAll(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("set_key", 0, "value")
	s.Get("set_key")
	s.Add("set_key", 0, "value3")
	s.Del("set_key", "value")
	s.DelAll("set_key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("set_key")
	fmt.Println(keyRes)
}

func TestSetAdd(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("set_key", 0, "value")
	s.Add("set_key", 0, "value2")
	s.Add("set_key2", 0, "value2")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("set_key")
	fmt.Println(keyRes.AddNum)
}

func TestSetGet(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.Get("set_key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("set_key")
	fmt.Println(keyRes.Value)
}

func TestSetDel(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.Del("set_key", "value")
	s.Del("set_key2", "value2")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("set_key")
	fmt.Println(keyRes.DelNum)
}

func TestSetDelAll(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.DelAll("set_key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("set_key")
	fmt.Println(keyRes.DelNum)
}

func TestSetAddR(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	r, err := s.AddR("set_key", 0, "value")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	r, err = s.AddR("key4", 20*time.Second, "value4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestSetGetR(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	r, err := s.GetR("key3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestSetDelR(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	r, err := s.DelR("key3", "value3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	r, err = s.DelR("set_key", "value")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestSetDelAllR(t *testing.T) {
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
	s, err := c.Set()
	if err != nil {
		fmt.Println(err)
	}
	r, err := s.DelAllR("set_key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
