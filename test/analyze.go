// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

package main

import (
	"fmt"
	"github.com/tencentyun/go-sdk"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ListDir(dirPath string, suffix string) (files []string, err error) {
	files = make([]string, 0, 32)
	suffix = strings.ToUpper(suffix)

	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	pathSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPath+pathSep+fi.Name())
		}
	}

	return files, nil
}

func WalkDir(dirPath string, suffix string) (files []string, err error) {
	files = make([]string, 0, 256)
	suffix = strings.ToUpper(suffix)

	err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
		//遍历目录
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

const appid uint = 200943
const sid = "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
const skey = "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage : analyze [dirpath] [suffix]")
		return
	}

	fmt.Println("start...")
	fmt.Println("====================================")
	dir := os.Args[1]
	suffix := os.Args[2]
	pics, err := ListDir(dir, suffix)
	if err != nil {
		fmt.Printf("err = %s\n", err.Error())
		return
	}

	fuzzy := 0
	notfuzzy := 0
	total := 0

	cloud := qcloud.PicCloud{appid, sid, skey, ""}
	analyze := qcloud.PicAnalyze{1, 0}
	for _, pic := range pics {
		fmt.Printf("analyze pic = %s\r\n", pic)
		total += 1

		fi, err := os.Open(pic)
		if nil != err {
			return
		}
		defer fi.Close()
		picData, err := ioutil.ReadAll(fi)
		if nil != err {
			return
		}

		for i := 0; i < 3; i++ {
			info, err := cloud.UploadBase(picData, "", analyze)
			if err != nil {
				fmt.Printf("err = %s\r\n", err.Error())
				continue
			}
			if info.Analyze.Fuzzy == 1 {
				fmt.Println("this pic is fuzzy")
				fuzzy += 1
				break
			} else {
				fmt.Println("this pic is not fuzzy")
				notfuzzy += 1
				break
			}
		}
	}

	fmt.Println("====================================")
	fmt.Printf("total=%d fuzzy=%d notfuzzy=%d\r\n", total, fuzzy, notfuzzy)
}
