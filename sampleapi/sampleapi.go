package sampleapi

import (
    "crypto/tls"
    "encoding/xml"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
    "net/http"
)

const (
    SAMPLE_RSS_URL  = "https://techblog.yahoo.co.jp/index.xml"
    SAMPLE_RSS_URL2 = "https://techblog.yahoo.co.jp/index.xml"
)

func RssApi(c *gin.Context) {
    test   := c.Query("test")
    rssUrl := ""

    if test == "test" {
        rssUrl = SAMPLE_RSS_URL2
    } else {
        rssUrl = SAMPLE_RSS_URL
    }
 
    req, _ := http.NewRequest("GET", rssUrl, nil)
    req.Header.Set("User-Agent", "sample")
    req.Header.Set("Content-Type", "application/xml")
    response, err := RequestAPI(req)

    if err != nil {
        log.SetFlags(log.Lshortfile)
        log.Println("[error]", err)
    }

    rssJson := ParseXml(response)

    c.JSON(200, rssJson)
}

func ParseXml(response string) *RssXml {
    var rssXml RssXml
    xml.Unmarshal([]byte(response), &rssXml)
    return &rssXml
}

func RequestAPI(request *http.Request) (string, error) {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{Transport: tr}

    resp, err := client.Do(request)
    if err != nil {
        log.SetFlags(log.Lshortfile)
        log.Println("[error]", err)
        return "", err
    }

    defer resp.Body.Close()

    byteArray, _ := ioutil.ReadAll(resp.Body)

    return string(byteArray), nil
}

type RssXml struct {
    Channel []Channels `xml:"channel" json:"channels"`
}

type Channels struct {
    Title string `xml:"title" json:"title"`
}

