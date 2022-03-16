package crawler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ifty123/crawl_web/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber"
)

// TODO: Write your function here
func CrawlLink(url string) []models.Link {
	webPage := url
	resp, err := http.Get(webPage)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var links []models.Link

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		var link models.Link
		linkTag := item
		if linkTag.Text() == "\n                Dropdown\n              " {
			link.Title = "Dropdown"
		} else {
			link.Title = linkTag.Text()
		}
		link.Url, _ = item.Attr("href")
		link.Title = link.Title

		links = append(links, link)
	})

	bts, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("from json", string(bts))

	return links
}

func GetCraw(c *fiber.Ctx) {
	getCraw := CrawlLink("https://csr-assessment.miqdadyyy.vercel.app")
	//return
	respSend := models.Response{
		OK:   true,
		Msg:  "success",
		Data: getCraw,
	}
	c.JSON(respSend)
}
