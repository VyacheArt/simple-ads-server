package ads

import "sort"

type (
	User struct {
		Country string
		Browser string
	}

	Campaign struct {
		Id        uint32
		ClickUrl  string
		Price     float64
		Targeting Targeting
	}

	Targeting struct {
		Country string
		Browser string
	}

	filterFunc func(in []*Campaign, u *User) []*Campaign
)

var (
	filters = []filterFunc{
		filterByCountry,
		filterByBrowser,
	}
)

func MakeAuction(in []*Campaign, u *User) (winner *Campaign) {
	campaigns := make([]*Campaign, len(in))
	copy(campaigns, in)

	for _, f := range filters {
		campaigns = f(campaigns, u)
	}

	if len(campaigns) == 0 {
		return nil
	}

	sort.Slice(campaigns, func(i, j int) bool {
		return campaigns[i].Price < campaigns[j].Price
	})

	return campaigns[0]
}

func filterByBrowser(in []*Campaign, u *User) []*Campaign {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Browser) == 0 {
			//пустой браузер - значит нет ограничений, не удаляем кампанию
			continue
		}

		if in[i].Targeting.Browser == u.Browser {
			//браузер совпадает - всё ок, кампания проходит
			continue
		}

		//а вот тут мы знаем что фильтр по браузеру есть, и браузер пользователя не соответствует ему
		//удаляем кампанию
		in[i] = in[0]
		in = in[1:]
	}

	return in
}

func filterByCountry(in []*Campaign, u *User) []*Campaign {
	for i := len(in) - 1; i >= 0; i-- {
		if len(in[i].Targeting.Country) == 0 {
			//пустой браузер - значит нет ограничений, не удаляем кампанию
			continue
		}

		if in[i].Targeting.Country == u.Country {
			//браузер совпадает - всё ок, кампания проходит
			continue
		}

		//а вот тут мы знаем что фильтр по браузеру есть, и браузер пользователя не соответствует ему
		//удаляем кампанию
		in[i] = in[0]
		in = in[1:]
	}

	return in
}

func GetStaticCampaigns() []*Campaign {
	return []*Campaign{
		{
			Id:    1,
			Price: 1,
			Targeting: Targeting{
				Country: "RU",
				Browser: "Chrome",
			},
			ClickUrl: "https://yandex.ru",
		},
		{
			Id:    2,
			Price: 1,
			Targeting: Targeting{
				Country: "DE",
				Browser: "Chrome",
			},
			ClickUrl: "https://google.com",
		},
		{
			Id:    3,
			Price: 1,
			Targeting: Targeting{
				Browser: "Firefox",
			},
			ClickUrl: "https://duckduckgo.com",
		},
	}
}
