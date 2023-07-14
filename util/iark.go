package util

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func URLCPUName(name string) string {
	name = strings.ReplaceAll(name, "(R)", "")
	name = strings.ReplaceAll(name, "(TM)", "")
	name = strings.ReplaceAll(name, "(C)", "")
	name = strings.ReplaceAll(name, "(P)", "")
	name = strings.ReplaceAll(name, "(G)", "")
	name = strings.ReplaceAll(name, "CPU", "Processor")
	name = strings.ReplaceAll(name, "@", "")
	name = strings.ReplaceAll(name, " ", "+")
	if strings.HasSuffix(name, "Hz") {
		name = strings.Split(name, "++")[0]
	}
	return name
}

func isCorrect(query string, url string) bool {
	sp := strings.Split(url, "/")
	u := sp[len(sp)-1]
	query = strings.ToLower(query)
	query = strings.ReplaceAll(query, "-", "")
	query = strings.ReplaceAll(query, "+", "-")
	return strings.HasPrefix(u, query)
}

func IntelArkGetGeneration(data string) string {
	data = strings.Split(data, ">")[1]
	data = strings.TrimSuffix(data, "</a")
	return strings.TrimPrefix(data, "Products formerly ")
}

func intelArkGetCPU(url string) map[string]string {
	res, _ := http.Get("https://ark.intel.com" + url)
	d, _ := io.ReadAll(res.Body)
	sp := []string{}
	data := make(map[string]string)
	for _, l := range strings.Split(string(d), "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		sp = append(sp, l)
	}
	for i, l := range sp {
		if !strings.HasPrefix(l, "<span class=\"value\"") {
			continue
		}
		key := strings.TrimPrefix(strings.TrimSuffix(l, "\">"), "<span class=\"value\" data-key=\"")
		value := sp[i+1]
		value = strings.TrimSuffix(value, "</span>")
		if strings.HasPrefix(key, "<") {
			continue
		}
		data[key] = value
	}
	return data
}

func IntelArkGetCPU(query string) map[string]string {
	res, _ := http.Get(fmt.Sprintf("https://ark.intel.com/content/www/us/en/ark/search.html?_charset_=UTF-8&q=%s", query))
	d, _ := io.ReadAll(res.Body)
	for _, l := range strings.Split(string(d), "\n") {
		if strings.Contains(l, "href=\"/content/www/us/en/ark/products/") && strings.Contains(l, "Processor") && !strings.Contains(l, "Family") {
			url := strings.Split(strings.Split(l, "href=\"")[1], "\"")[0]
			if isCorrect(query, url) {
				return intelArkGetCPU(url)
			}
		}
	}
	return make(map[string]string)
}
