package prialgorithm

import (
	"log"
	"net/http"
	"net/url"
)

const (
	QINIUAPI = "http://ai.qiniuapi.com/v1/image/censor"
	one = iota
	two

)

/*
AK: i0eRBeSz0BTQzQqbxKc9Jm1ymvGkGeUJJ3BZ18N5

SK: 8evFLp1_KiNndNEWRiNszdLClr6aotlXQ1x1KdMB

 */
func ReviewPorn(imgUrl string) bool {
	req := MakeRequest(imgUrl)
	hc := http.Client{}
	hc.Do(req)

	return true
}

func MakeRequest(imgUrl string) *http.Request {
	r, err := url.Parse(QINIUAPI)
	if err != nil {
		log.Fatalln(err)
	}
	q := r.Query()
	r.RawQuery = q.Encode()
	req, err := http.NewRequest(
		"POST",
		r.String(),
		nil,
	)
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Qiniu i0eRBeSz0BTQzQqbxKc9Jm1ymvGkGeUJJ3BZ18N5:8evFLp1_KiNndNEWRiNszdLClr6aotlXQ1x1KdMB")

	return req

}
