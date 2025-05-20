package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/zodileap/taurus_go/cache/redis"
)

func TestExample(t *testing.T) {
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
		fmt.Println(err.Error())
	}
	s, err := c.Set()
	if err != nil {
		fmt.Println(err.Error())
	}
	s.Add("key", 0, "value")
	s.Add("key2", 0, "value2")
	s.Add("key", 0, "value3")
	result, err := c.Save()
	if err != nil {
		fmt.Println(err.Error())
	}
	keyRes := result.GetSet("key")
	fmt.Println(keyRes.AddNum)
	key2Res := result.GetSet("key2")
	fmt.Println(key2Res.AddNum)
}

func TestSetClient(t *testing.T) {
	redis.SetClient("test", &redis.Options{
		Addr: "localhost:6379",
		// 如果有用户名,否则空字符串
		Username: "",
		// 如果有密码,否则空字符串
		Password: "",
		DB:       1,
	})
}

func TestGetClient(t *testing.T) {
	redis.SetClient("test", &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	})
	c, err := redis.GetClient("test")
	// 请注意关闭连接
	defer c.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestClearClient(t *testing.T) {
	redis.SetClient("test", &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       15,
	})
	defer redis.ClearClient()
}

func TestClose(t *testing.T) {
	redis.SetClient("test", &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       1,
	})
	c, err := redis.GetClient("test")
	// 请注意关闭连接
	defer c.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestDel(t *testing.T) {
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
		fmt.Println(err.Error())
	}
	// 这里可以传任意数量的key
	l, err := c.Del("key1", "key2")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(l)
}

func TestSave(t *testing.T) {
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
		fmt.Println(err.Error())
	}
	s, err := c.Set()
	if err != nil {
		fmt.Println(err.Error())
	}
	s.Add("key", 0, "value")
	s.Add("key2", 0, "value2")
	s.Add("key", 0, "value3")
	result, err := c.Save()
	if err != nil {
		fmt.Println(err.Error())
	}
	keyRes := result.GetSet("key")
	fmt.Println(keyRes.AddNum)
	key2Res := result.GetSet("key2")
	fmt.Println(key2Res.AddNum)
	// 请注意每次调用`Save`后，都需要重新调用`Set`、`Hash`、`String`等方法。
	s, err = c.Set()
	s.Add("key", 0, "value3")
	c.Save()
}

func TestSet(t *testing.T) {
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
		fmt.Println(err.Error())
	}
	s, err := c.Set()
	if err != nil {
		fmt.Println(err.Error())
	}

	s.Add("key", 0, "value")
	s.Add("key", 0, "value3")
	s.Add("key2", 10*time.Second, "value2")
}

func TestSetExpire(t *testing.T) {
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
		fmt.Println(err.Error())
	}
	c.SetExpire("key", 10*time.Second)
}
