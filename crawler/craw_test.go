package crawler

import (
	"testing"

	"github.com/ifty123/crawl_web/models"

	"github.com/stretchr/testify/assert"
)

func TestCSR(t *testing.T) {
	assert.Equal(t, getExpectedOutput(), CrawlLink("https://csr-assessment.miqdadyyy.vercel.app/"), "Wrong Answer")
}

func TestSSR(t *testing.T) {
	assert.Equal(t, getExpectedOutput(), CrawlLink("https://ssr-assessment.miqdadyyy.vercel.app/"), "Wrong Answer")
}

func getExpectedOutput() []models.Link {
	return []models.Link{
		{Url: "#", Title: "Navbar"},
		{Url: "#", Title: "Dropdown"},
		{
			Url:   "https://www.example.com/bridge/bait.html",
			Title: "Crawl This Url",
		},
		{Url: "http://www.example.com/", Title: "This Url Too"},
		{Url: "https://www.example.org/", Title: "Ok Ganbatte"},
		{
			Url:   "https://www.example.org/bear/attraction.html",
			Title: "More",
		},
		{Url: "https://example.com/bridge/bell.html", Title: "Again"},
		{Url: "https://www.example.org/", Title: "Chotto"},
		{Url: "http://example.com/", Title: "Good Luck"},
		{
			Url:   "http://www.example.com/adjustment.php?branch=branch&birth=berry",
			Title: "Hufft",
		},
		{
			Url:   "http://example.com/adjustment.php",
			Title: "Ok I'm Getting Bored",
		},
		{Url: "https://www.example.edu/#bear", Title: "hmmmmm"},
		{Url: "https://www.example.com/", Title: "Ok Boomer"},
		{
			Url:   "http://www.example.com/believe/acoustics.html",
			Title: "Demo Sonna Ja Dame",
		},
		{
			Url:   "http://apparel.example.com/blade.html",
			Title: "Mo Sonna ja hora",
		},
		{
			Url:   "http://www.example.com/addition/bomb.php",
			Title: "Kokoro wa shinkai suru yo",
		},
		{
			Url:   "https://www.example.com/?birthday=bird",
			Title: "Motto Motto",
		},
		{
			Url:   "http://www.example.com/bait/bag?breath=animal",
			Title: "Teret tereret",
		},
		{Url: "http://www.example.com/", Title: "Sasageyo Jare"},
		{Url: "https://www.example.com/", Title: "Lorem"},
		{Url: "https://bear.example.com/", Title: "Ipsum"},
		{Url: "https://www.example.edu/", Title: "Dolor"},
		{
			Url:   "http://example.com/?appliance=actor&arithmetic=battle",
			Title: "Sit",
		},
		{Url: "http://bird.example.net/addition/blade.html", Title: "Amet"},
		{
			Url:   "https://www.example.com/brick/animal.html",
			Title: "Constrector",
		},
		{Url: "https://www.example.com/belief/breath", Title: "sapien"},
		{
			Url:   "http://example.com/arch?afternoon=bait&acoustics=boy",
			Title: "molli",
		},
		{Url: "https://www.example.net/afternoon", Title: "auctor"},
		{Url: "https://example.com/", Title: "lacinia"},
		{Url: "https://www.example.com/", Title: "vitae"},
		{Url: "https://www.example.com/#basketball", Title: "Ut magna"},
		{
			Url:   "https://bag.example.com/bed/blade",
			Title: "Bang jual seblak gak",
		},
		{Url: "https://example.com/amount/badge#airport", Title: "sodales"},
		{Url: "https://www.example.edu/", Title: "vel"},
		{
			Url:   "http://example.com/brick?blood=arithmetic&brick=baseball",
			Title: "diam",
		},
		{Url: "https://brick.example.net/anger", Title: "cursus"},
		{Url: "https://birth.example.net/board.php", Title: "ante"},
		{Url: "https://amount.example.com/", Title: "Sed"},
		{
			Url:   "https://arm.example.com/?account=approval&arm=achiever#airport",
			Title: "faucibus",
		},
		{Url: "http://example.com/bath.php", Title: "Nam commodo"},
		{Url: "https://www.example.net/bottle#anger", Title: "Tellus"},
		{Url: "http://www.example.com/#brake", Title: "Alliquam"},
		{Url: "https://example.com/berry", Title: "Iri"},
		{
			Url:   "https://www.example.net/board/aunt.aspx#anger",
			Title: "bilang bos",
		},
		{Url: "http://example.com/", Title: "sodales"},
		{Url: "https://example.edu/berry/badge.html", Title: "tortor"},
		{Url: "http://example.net/", Title: "Nunc"},
		{
			Url:   "http://www.example.net/?anger=agreement&border=boot",
			Title: "elefiend",
		},
		{Url: "http://www.example.net/", Title: "tempor"},
		{Url: "https://example.com/", Title: "nibh"},
		{Url: "https://book.example.com/", Title: "Jual"},
		{Url: "https://example.com/advice#brick", Title: "Tanah Kampus"},
		{Url: "http://www.example.com/", Title: "Stikom banyuwangi"},
		{Url: "https://example.com/arm.html", Title: "Murah"},
		{Url: "https://bath.example.com/", Title: "Minat PM"},
		{Url: "https://example.com/", Title: "Donec"},
		{Url: "http://www.example.com/anger/balance", Title: "Turpis"},
		{Url: "https://aftermath.example.com/bikes", Title: "Condimentum"},
		{Url: "http://example.com/", Title: "Tristique"},
		{Url: "https://www.example.com/aunt", Title: "auctor"},
		{Url: "http://www.example.com/bead#bath", Title: "justo"},
		{Url: "http://bed.example.com/bell.html", Title: "Vehucula"},
		{Url: "https://aunt.example.net/brother", Title: "Arcu"},
		{Url: "https://www.example.com/box/brass", Title: "Integer"},
		{Url: "http://bikes.example.com/", Title: "pal pale"},
		{Url: "https://example.com/bell/adjustment.aspx", Title: "This"},
		{Url: "https://www.example.com/", Title: "Is"},
		{Url: "http://www.example.com/", Title: "Just"},
		{Url: "http://example.com/", Title: "For"},
		{Url: "https://example.edu/advertisement", Title: "Test Case"},
		{Url: "http://www.example.com/arch.aspx", Title: "Crawllo"},
		{Url: "https://www.example.com/books/bridge.php", Title: "Dama"},
		{Url: "http://example.com/", Title: "Genera"},
		{Url: "https://www.example.com/brake.html", Title: "Salamat"},
		{Url: "http://www.example.com/?amusement=back", Title: "Malesuada"},
		{Url: "https://bedroom.example.com/", Title: "Brotha"},
		{Url: "https://www.example.com/angle", Title: "Morbui"},
		{
			Url:   "https://example.com/art?brass=box&bells=bird",
			Title: "England",
		},
		{
			Url:   "http://attraction.example.org/brother.aspx",
			Title: "Brotherrr",
		},
		{Url: "https://example.com/anger.html", Title: "Anger"},
		{Url: "http://example.com/?achiever=bath", Title: "Achievement"},
		{Url: "http://www.example.org/", Title: "ORG"},
		{Url: "https://example.com/art", Title: "ART"},
		{Url: "http://www.example.com/", Title: "Def"},
		{Url: "https://anger.example.com/advertisement.htm", Title: "Ads"},
		{
			Url:   "http://www.example.com/arch/attraction.php?bat=blade",
			Title: "Attraction",
		},
		{Url: "http://example.com/back.php", Title: "Back"},
		{Url: "https://example.com/aftermath/attack", Title: "Attack"},
		{Url: "http://brake.example.edu/", Title: "Edu"},
		{Url: "http://example.org/", Title: "Again"},
	}
}
