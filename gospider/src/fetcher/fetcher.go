/**
根据传入的url，
返回
其中：不需要对网页字符编码进行处理。网页的字符编码就是UTF-8
 */
package fetcher

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Fetch(url string) ([]byte, error) {
	resp,err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()  // ??
	if resp.StatusCode != http.StatusOK {
		//return nil, errors.New("Wrong status code")  //两种写法都是正确的
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
