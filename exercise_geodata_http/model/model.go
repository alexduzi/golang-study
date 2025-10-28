package model

type Country struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName struct {
			Por struct {
				Official string `json:"official"`
				Common   string `json:"common"`
			} `json:"por"`
		} `json:"nativeName"`
	} `json:"name"`
	Tld         []string `json:"tld"`
	Cca2        string   `json:"cca2"`
	Ccn3        string   `json:"ccn3"`
	Cioc        string   `json:"cioc"`
	Independent bool     `json:"independent"`
	Status      string   `json:"status"`
	UnMember    bool     `json:"unMember"`
	Currencies  struct {
		Eur struct {
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
		} `json:"EUR"`
	} `json:"currencies"`
	Idd struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
	Capital      []string `json:"capital"`
	AltSpellings []string `json:"altSpellings"`
	Region       string   `json:"region"`
	Subregion    string   `json:"subregion"`
	Languages    struct {
		Por string `json:"por"`
	} `json:"languages"`
	Latlng     []float64 `json:"latlng"`
	Landlocked bool      `json:"landlocked"`
	Borders    []string  `json:"borders"`
	Area       float64   `json:"area"`
	Demonyms   struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
		Fra struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"fra"`
	} `json:"demonyms"`
	Cca3         string `json:"cca3"`
	Translations struct {
		Ara struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ara"`
		Bre struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"bre"`
		Ces struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ces"`
		Cym struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"cym"`
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Est struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"est"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Hun struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hun"`
		Ind struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ind"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Kor struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"kor"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Per struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"per"`
		Pol struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"pol"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Slk struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"slk"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Srp struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"srp"`
		Swe struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"swe"`
		Tur struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"tur"`
		Urd struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"urd"`
		Zho struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"zho"`
	} `json:"translations"`
	Flag string `json:"flag"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	Population int `json:"population"`
	Gini       struct {
		Num2018 float64 `json:"2018"`
	} `json:"gini"`
	Fifa string `json:"fifa"`
	Car  struct {
		Signs []string `json:"signs"`
		Side  string   `json:"side"`
	} `json:"car"`
	Timezones  []string `json:"timezones"`
	Continents []string `json:"continents"`
	Flags      struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
		Alt string `json:"alt"`
	} `json:"flags"`
	CoatOfArms struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"coatOfArms"`
	StartOfWeek string `json:"startOfWeek"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode"`
}
