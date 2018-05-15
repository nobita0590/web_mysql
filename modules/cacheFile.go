package modules

import (
	"sync"
	"io/ioutil"
	"github.com/nobita0590/web_mysql/config"
	"fmt"
)

type CacheContent struct {
	files map[string][]byte
	mux sync.Mutex
}

func (c CacheContent) Get(path string) (content []byte,ok bool)  {
	c.mux.Lock()
	defer c.mux.Unlock()
	content,ok = c.files[path]
	return
}

func (c *CacheContent) Set(path string,content []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.files[path] = content
}

func (c *CacheContent) TakeContent(path string) (content []byte) {
	var ok bool
	c.mux.Lock()
	defer c.mux.Unlock()
	if content,ok = c.files[path];ok {
		fmt.Println(path)
		return
	}else{
		var err error
		content,err = ioutil.ReadFile(config.FilePath + path)
		fmt.Println("error:", err)
		c.files[path] = content
		return
	}
}
