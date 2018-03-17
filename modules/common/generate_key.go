package common

import "time"

func NewProvider() (p Provider) {
	p.source = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	p.length = len(p.source)
	p.previous = 0
	p.previousLength = 0
	return
}

type Provider struct {
	source 		string
	length 		int
	previous	int
	previousLength	int
	cacheStr	[]string
}

func (p *Provider) buildRaw(num int) {
	if num < p.length {
		p.cacheStr = append(p.cacheStr,p.source[num:num+1])
	}else{
		division := num/p.length
		balance := num - division * p.length
		p.cacheStr = append(p.cacheStr,p.source[balance:balance+1])
		p.buildRaw(division)
	}
}
func (p *Provider) GenerateFromInt(input int,_len int) string {
	p.previous = input
	p.previousLength = _len
	p.cacheStr = []string{}
	p.buildRaw(input)
	if _len > 0 {
		for len(p.cacheStr) < _len {
			p.cacheStr = append(p.cacheStr,p.source[0:1])
		}
	}
	s := ""
	for _,str := range p.cacheStr {
		s = str + s
	}
	p.cacheStr = []string{}
	if s == "" {
		return p.source[0:1]
	}
	return s
}
func (p *Provider) Generate(_len int) string {
	t:= time.Now()
	input := int(t.UnixNano())
	return p.GenerateFromInt(input,_len)
}
func (p *Provider) Next() string {
	return p.GenerateFromInt(p.previous + 1,p.previousLength)
}