package ads

import (
	realip "github.com/Ferluci/fast-realip"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/valyala/fasthttp"
	"log"
	"net"
	"net/http"
)

type Server struct {
	geoip *geoip2.Reader
}

func NewServer(geoip *geoip2.Reader) *Server {
	return &Server{geoip: geoip}
}

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080", s.handler)
}

func (s *Server) handler(ctx *fasthttp.RequestCtx) {
	remoteIp := realip.FromRequest(ctx)
	ua := string(ctx.Request.Header.UserAgent())

	parsed := user_agent.New(ua)
	browserName, _ := parsed.Browser()

	country, err := s.geoip.Country(net.ParseIP(remoteIp))
	if err != nil {
		log.Printf("Failed to parse country: %v", err)
		return
	}

	user := &User{
		Country: country.Country.IsoCode,
		Browser: browserName,
	}
	campaigns := GetStaticCampaigns()

	winner := MakeAuction(campaigns, user)
	if winner == nil {
		ctx.Redirect("https://example.com", http.StatusSeeOther)
		return
	}

	ctx.Redirect(winner.ClickUrl, http.StatusSeeOther)
}
