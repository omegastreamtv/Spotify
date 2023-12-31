package spotify

// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
//
// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
//
// Note: If neither market or user country are provided, the content is considered unavailable for the client.
// Users can view the country that is associated with their account in the account settings.
type Market string

const (
	MarketAndorra                      Market = "AD"
	MarketUnitedArabEmirates           Market = "AE"
	MarketAntiguaAndBarbuda            Market = "AG"
	MarketAlbania                      Market = "AL"
	MarketArmenia                      Market = "AM"
	MarketAngola                       Market = "AO"
	MarketArgentina                    Market = "AR"
	MarketAustria                      Market = "AT"
	MarketAustralia                    Market = "AU"
	MarketAzerbaijan                   Market = "AZ"
	MarketBosniaAndHerzegovina         Market = "BA"
	MarketBarbados                     Market = "BB"
	MarketBangladesh                   Market = "BD"
	MarketBelgium                      Market = "BE"
	MarketBurkinaFaso                  Market = "BF"
	MarketBulgaria                     Market = "BG"
	MarketBahrain                      Market = "BH"
	MarketBurundi                      Market = "BI"
	MarketBenin                        Market = "BJ"
	MarketBrunei                       Market = "BN"
	MarketBolivia                      Market = "BO"
	MarketBrazil                       Market = "BR"
	MarketBahamas                      Market = "BS"
	MarketBhutan                       Market = "BT"
	MarketBotswana                     Market = "BW"
	MarketBelarus                      Market = "BY"
	MarketBelize                       Market = "BZ"
	MarketCanada                       Market = "CA"
	MarketDemocraticRepublicOfCongo    Market = "CD"
	MarketRepublicOfCongo              Market = "CG"
	MarketSwitzerland                  Market = "CH"
	MarketIvoryCoast                   Market = "CI"
	MarketChile                        Market = "CL"
	MarketCameroon                     Market = "CM"
	MarketColombia                     Market = "CO"
	MarketCostaRica                    Market = "CR"
	MarketCapeVerde                    Market = "CV"
	MarketCuracao                      Market = "CW"
	MarketCyprus                       Market = "CY"
	MarketCzechRepublic                Market = "CZ"
	MarketGermany                      Market = "DE"
	MarketDjibouti                     Market = "DJ"
	MarketDenmark                      Market = "DK"
	MarketDominica                     Market = "DM"
	MarketDominicanRepublic            Market = "DO"
	MarketAlgeria                      Market = "DZ"
	MarketEcuador                      Market = "EC"
	MarketEstonia                      Market = "EE"
	MarketEgypt                        Market = "EG"
	MarketSpain                        Market = "ES"
	MarketEthiopia                     Market = "ET"
	MarketFinland                      Market = "FI"
	MarketFiji                         Market = "FJ"
	MarketMicronesia                   Market = "FM"
	MarketFrance                       Market = "FR"
	MarketGabon                        Market = "GA"
	MarketUnitedKingdom                Market = "GB"
	MarketGrenada                      Market = "GD"
	MarketGeorgia                      Market = "GE"
	MarketGhana                        Market = "GH"
	MarketGambia                       Market = "GM"
	MarketGuinea                       Market = "GN"
	MarketEquatorialGuinea             Market = "GQ"
	MarketGreece                       Market = "GR"
	MarketGuatemala                    Market = "GT"
	MarketGuineaBissau                 Market = "GW"
	MarketGuyana                       Market = "GY"
	MarketHongKong                     Market = "HK"
	MarketHonduras                     Market = "HN"
	MarketCroatia                      Market = "HR"
	MarketHaiti                        Market = "HT"
	MarketHungary                      Market = "HU"
	MarketIndonesia                    Market = "ID"
	MarketIreland                      Market = "IE"
	MarketIsrael                       Market = "IL"
	MarketIndia                        Market = "IN"
	MarketIraq                         Market = "IQ"
	MarketIceland                      Market = "IS"
	MarketItaly                        Market = "IT"
	MarketJamaica                      Market = "JM"
	MarketJordan                       Market = "JO"
	MarketJapan                        Market = "JP"
	MarketKenya                        Market = "KE"
	MarketKyrgyzstan                   Market = "KG"
	MarketCambodia                     Market = "KH"
	MarketKiribati                     Market = "KI"
	MarketComoros                      Market = "KM"
	MarketSaintKittsAndNevis           Market = "KN"
	MarketSouthKorea                   Market = "KR"
	MarketKuwait                       Market = "KW"
	MarketKazakhstan                   Market = "KZ"
	MarketLaos                         Market = "LA"
	MarketLebanon                      Market = "LB"
	MarketSaintLucia                   Market = "LC"
	MarketLiechtenstein                Market = "LI"
	MarketSriLanka                     Market = "LK"
	MarketLiberia                      Market = "LR"
	MarketLesotho                      Market = "LS"
	MarketLithuania                    Market = "LT"
	MarketLuxembourg                   Market = "LU"
	MarketLatvia                       Market = "LV"
	MarketLibya                        Market = "LY"
	MarketMorocco                      Market = "MA"
	MarketMonaco                       Market = "MC"
	MarketMoldova                      Market = "MD"
	MarketMontenegro                   Market = "ME"
	MarketMadagascar                   Market = "MG"
	MarketMarshallIslands              Market = "MH"
	MarketMacedonia                    Market = "MK"
	MarketMali                         Market = "ML"
	MarketMongolia                     Market = "MN"
	MarketMacau                        Market = "MO"
	MarketMauritania                   Market = "MR"
	MarketMalta                        Market = "MT"
	MarketMauritius                    Market = "MU"
	MarketMaldives                     Market = "MV"
	MarketMalawi                       Market = "MW"
	MarketMexico                       Market = "MX"
	MarketMalaysia                     Market = "MY"
	MarketMozambique                   Market = "MZ"
	MarketNamibia                      Market = "NA"
	MarketNiger                        Market = "NE"
	MarketNigeria                      Market = "NG"
	MarketNicaragua                    Market = "NI"
	MarketNetherlands                  Market = "NL"
	MarketNorway                       Market = "NO"
	MarketNepal                        Market = "NP"
	MarketNauru                        Market = "NR"
	MarketNewZealand                   Market = "NZ"
	MarketOman                         Market = "OM"
	MarketPanama                       Market = "PA"
	MarketPeru                         Market = "PE"
	MarketPapuaNewGuinea               Market = "PG"
	MarketPhilippines                  Market = "PH"
	MarketPakistan                     Market = "PK"
	MarketPoland                       Market = "PL"
	MarketPalestinianTerritories       Market = "PS"
	MarketPortugal                     Market = "PT"
	MarketPalau                        Market = "PW"
	MarketParaguay                     Market = "PY"
	MarketQatar                        Market = "QA"
	MarketRomania                      Market = "RO"
	MarketSerbia                       Market = "RS"
	MarketRwanda                       Market = "RW"
	MarketSaudiArabia                  Market = "SA"
	MarketSolomonIslands               Market = "SB"
	MarketSeychelles                   Market = "SC"
	MarketSweden                       Market = "SE"
	MarketSingapore                    Market = "SG"
	MarketSlovenia                     Market = "SI"
	MarketSlovakia                     Market = "SK"
	MarketSierraLeone                  Market = "SL"
	MarketSanMarino                    Market = "SM"
	MarketSenegal                      Market = "SN"
	MarketSuriname                     Market = "SR"
	MarketSaoTomeAndPrincipe           Market = "ST"
	MarketElSalvador                   Market = "SV"
	MarketSwaziland                    Market = "SZ"
	MarketChad                         Market = "TD"
	MarketTogo                         Market = "TG"
	MarketThailand                     Market = "TH"
	MarketTajikistan                   Market = "TJ"
	MarketEastTimor                    Market = "TL"
	MarketTunisia                      Market = "TN"
	MarketTonga                        Market = "TO"
	MarketTurkey                       Market = "TR"
	MarketTrinidadAndTobago            Market = "TT"
	MarketTuvalu                       Market = "TV"
	MarketTaiwan                       Market = "TW"
	MarketTanzania                     Market = "TZ"
	MarketUkraine                      Market = "UA"
	MarketUganda                       Market = "UG"
	MarketUnitedStates                 Market = "US"
	MarketUruguay                      Market = "UY"
	MarketUzbekistan                   Market = "UZ"
	MarketSaintVincentAndTheGrenadines Market = "VC"
	MarketVenezuela                    Market = "VE"
	MarketVietnam                      Market = "VN"
	MarketVanuatu                      Market = "VU"
	MarketSamoa                        Market = "WS"
	MarketKosovo                       Market = "XK"
	MarketSouthAfrica                  Market = "ZA"
	MarketZambia                       Market = "ZM"
	MarketZimbabwe                     Market = "ZW"
)

type GetAvailableMarketsResponse struct {
	Markets []string `json:"markets"`
}

// Get the list of markets where Spotify is available.
func (c *Client) GetAvailableMarkets() (*GetAvailableMarketsResponse, error) {
	markets := GetAvailableMarketsResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get("/markets").Receive(&markets, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &markets, nil
}
