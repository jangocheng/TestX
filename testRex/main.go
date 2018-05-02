package main

import (
	"fmt"
	"regexp"

	"unicode/utf8"
)

type ss struct {
	ee []interface{}
}

func main() {
	r, _ := regexp.Compile("^[\u4E00-\u9FA5a-zA-Z0-9,]{0,300}$")
	fmt.Println(r.MatchString("张3,fdfs,"))

	a := make([]byte, 10)
	var b byte = 1
	a = append(a, b)
	fmt.Println(a)
	fmt.Print("fadsfadf\n")
	fmt.Println(len("MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCfU8v4BUr81SKm"))

	var str = "张三=213"
	var runes = []rune(str)
	fmt.Println(runes[0] == '张')

	bu := make([]byte, utf8.UTFMax)

	n1 := utf8.EncodeRune(bu, '好')
	fmt.Printf("%v：%v\n", bu, n1)
	n2 := utf8.EncodeRune(bu, runes[0])
	fmt.Printf("%v：%v\n", bu, n2)

	fmt.Println(formatKey(`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqNxzebovJ6R+LF0jFyJD4vgdvj+Apmb5h+pW3T0EtDzWZAr7tyiSAtNedYvRjJCqN5cYw0rIwGMZFbD3lQHbJGC+IvpqXwPB8AWqRAwItI82fo2+AyHkq11yE27IgOjSrKofgg3GWJ6SSQonYuXZ0c09chXXiZPKYe0zRbvq83kAVsYDu1sMwi8mfiVff6CIALsehs1MOjmdLW40N1CicVmJaWuh2yee+sj1/0xMOlV1LyJq63hShBD7T93qpGbHoNkpdz+BFc2byrhv1idbB4DRbUiKynzj3FX2Nz8Dv9TFQv8p2Z8dIOst890atv3P8DO7a9FI8I1reLvFDdyPawIDAQAB`, 1, true))
		fmt.Println(len(`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqNxzebovJ6R+LF0jFyJD4vgdvj+Apmb5h+pW3T0EtDzWZAr7tyiSAtNedYvRjJCqN5cYw0rIwGMZFbD3lQHbJGC+IvpqXwPB8AWqRAwItI82fo2+AyHkq11yE27IgOjSrKofgg3GWJ6SSQonYuXZ0c09chXXiZPKYe0zRbvq83kAVsYDu1sMwi8mfiVff6CIALsehs1MOjmdLW40N1CicVmJaWuh2yee+sj1/0xMOlV1LyJq63hShBD7T93qpGbHoNkpdz+BFc2byrhv1idbB4DRbUiKynzj3FX2Nz8Dv9TFQv8p2Z8dIOst890atv3P8DO7a9FI8I1reLvFDdyPawIDAQAB`))
	fmt.Println(len("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqNxzebovJ6R+LF0jFyJD")*6+8)
		}

func formatKey(key string, keyType int, ifPublic bool) string {
	if keyType == 0 {
		return key
	}
	if ifPublic {
		var publicHeader = "\n-----BEGIN PUBLIC KEY-----\n"
		var publicTail = "-----END PUBLIC KEY-----\n"
		var temp string
		split(key, &temp)
		return publicHeader + temp + publicTail
	} else {
		var publicHeader = "\n-----BEGIN RSA PRIVATE KEY-----\n"
		var publicTail = "-----END RSA PRIVATE KEY-----\n"
		var temp string
		split(key, &temp)
		return publicHeader + temp + publicTail
	}
}

func split(key string, temp *string) {
	if len(key) <= 64 {
		*temp = *temp + key + "\n"
	}
	for i := 0; i < len(key); i++ {
		if (i+1)%64 == 0 {
			*temp = *temp + key[:i+1] + "\n"
			fmt.Println(len(*temp) - 1)
			key = key[i+1:]
			split(key, temp)
			break
		}
	}
}
