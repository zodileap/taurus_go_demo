package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/yohobala/taurus_go/cache/redis"
)

func TestSetExample(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "", // no password set
		DB:       1,  // use default DB
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
	s.Add("key", 0, "value")
	s.Get("key")
	s.Add("key", 0, "value3")
	s.Add("key2", 0, "value2")
	s.Del("key", "value")
	s.Get("key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("key")
	fmt.Println(keyRes)
	if err != nil {
		fmt.Println(err)
	}
	s, err = c.Set()
	if err != nil {
		fmt.Println(err)
	}
	s.Add("key", 0, "value666")
	c.Save()
}

func TestSetExampleDelAll(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "", // no password set
		DB:       1,  // use default DB
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
	s.Add("key", 0, "value")
	s.Get("key")
	s.Add("key", 0, "value3")
	s.Del("key", "value")
	s.DelAll("key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("key")
	fmt.Println(keyRes)
}

func TestSetAdd(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "", // no password set
		DB:       1,  // use default DB
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
	s.Add("key", 0, "value")
	s.Add("key", 0, "value2")
	s.Add("key2", 0, "value2")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("key")
	fmt.Println(keyRes.AddNum)
}

func TestSetGet(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "", // no password set
		DB:       1,  // use default DB
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
	s.Get("key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("key")
	fmt.Println(keyRes.Value)
}

func TestSetDel(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "", // no password set
		DB:       1,  // use default DB
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
	s.Del("key", "value")
	s.Del("key2", "value2")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("key")
	fmt.Println(keyRes.DelNum)
}

func TestSetDelAll(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "", // no password set
		DB:       1,  // use default DB
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
	s.DelAll("key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetSet("key")
	fmt.Println(keyRes.DelNum)
}

func TestSetAddR(t *testing.T) {
	name := "test"
	options := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "", // no password set
		DB:       1,  // use default DB
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
	r, err := s.AddR("key", 0, "value")
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
		Password: "", // no password set
		DB:       1,  // use default DB
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
		Password: "", // no password set
		DB:       1,  // use default DB
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
	r, err = s.DelR("key", "value")
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
		Password: "", // no password set
		DB:       1,  // use default DB
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
	r, err := s.DelAllR("key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
