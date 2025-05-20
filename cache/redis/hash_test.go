package redis

import (
	"fmt"
	"testing"

	"github.com/zodileap/taurus_go/cache/redis"
)

func TestHashExample(t *testing.T) {
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
	h, err := c.Hash()
	h.Del("hash_key", "filed_1")
	h.Del("hash_key", "filed_2")
	h.AddM("hash_key", 0, map[string]string{"filed_1": "value1", "filed_2": "value2"})
	h.Add("hash_key", 0, "filed_3", "value3")
	h.Get("hash_key", "filed_1")
	h.Add("hash_key", 0, "filed_2", "value2")
	h.GetAll("hash_key")
	h.Add("hash_key2", 0, "filed_1", "value1")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes)
	h, err = c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.Add("hash_key3", 0, "filed_1", "value1")
	c.Save()
}

func TestHashAdd(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.Add("hash_key", 0, "filed_1", "value1")
	h.Add("hash_key", 0, "filed_2", "value2")
	h.Add("hash_key2", 0, "filed_1", "value1")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.AddNum)
}

func TestHashAddM(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.AddM("hash_key", 0, map[string]string{"filed_3": "value3", "filed_4": "value4"})
	h.AddM("hash_key2", 0, map[string]string{"filed_2": "value2"})
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.AddNum)
}

func TestHashGet(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.Get("hash_key", "filed_1")
	h.Get("hash_key", "filed_2")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.Value)
}

func TestHashGetM(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.GetM("hash_key", []string{"filed_1", "filed_2"})
	if err != nil {
		fmt.Println(err)
	}
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.Value)
}

func TestHashGetVals(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.GetVals("hash_key")
	if err != nil {
		fmt.Println(err)
	}
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.Value)
}

func TestHashGetKeys(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.GetAll("hash_key")
	if err != nil {
		fmt.Println(err)
	}
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.MapValue)
}

func TestHashDel(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.Del("hash_key", "filed_1")
	h.Del("hash_key2", "filed_2")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.DelNum)
}

func TestHashDelAll(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	h.DelAll("hash_key")
	r, err := c.Save()
	if err != nil {
		fmt.Println(err)
	}
	keyRes := r.GetHash("hash_key")
	fmt.Println(keyRes.DelNum)
}

func TestHashAddR(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	r1, err := h.AddR("hash_key", 0, "filed_1", "value1")
	if err != nil {
		fmt.Println(err)
	}
	r2, err := h.AddR("hash_key", 0, "filed_2", "value2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r1)
	fmt.Println(r2)
}

func TestHashAddMR(t *testing.T) {
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
	h, err := c.Hash()
	if err != nil {
		fmt.Println(err)
	}
	r1, err := h.AddMR("hash_key", 0, map[string]string{"filed_3": "value3", "filed_4": "value4"})
	if err != nil {
		fmt.Println(err)
	}
	r2, err := h.AddMR("hash_key2", 0, map[string]string{"filed_2": "value2"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r1)
	fmt.Println(r2)
}

func TestHashGetR(t *testing.T) {
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
	h, err := c.Hash()
	r1, err := h.GetR("hash_key", "filed_1")
	if err != nil {
		fmt.Println(err)
	}
	r2, err := h.GetR("hash_key", "filed_2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r1)
	fmt.Println(r2)
}

func TestHashGetMR(t *testing.T) {
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
	h, err := c.Hash()
	r, err := h.GetMR("hash_key", "filed_1", "filed_5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestHashGetValsR(t *testing.T) {
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
	h, err := c.Hash()
	r, err := h.GetValsR("hash_key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestHashGetAllR(t *testing.T) {
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
	h, err := c.Hash()
	r, err := h.GetAllR("hash_key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestDelR(t *testing.T) {
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
	h, err := c.Hash()
	r, err := h.DelR("hash_key", "filed_1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestDelAllR(t *testing.T) {
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
	h, err := c.Hash()
	r, err := h.DelAllR("hash_key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
