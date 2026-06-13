package rothfusz

const (
	CategoryIdeal          Category = "Ideal"
	CategoryHumid          Category = "Humid"
	CategoryCaution        Category = "Caution"
	CategoryExtremeCaution Category = "Extreme Caution"
	CategoryDanger         Category = "Danger"
	CategoryExtremeDanger  Category = "Extreme Danger"
)

type (
	Category string

	Result struct {
		Data            Data            `json:"data"`
		HeatIndexResult HeatIndexResult `json:"heatIndexResult"`
	}

	Data struct {
		Temp float64 `json:"temp"`
		RH   float64 `json:"rh"`
	}

	HeatIndexResult struct {
		HeatIndexC  float64  `json:"heatIndexC"`
		Category    Category `json:"category"`
		Comfortable bool     `json:"comfortable"`
	}
)

type repo struct {
	minValidTempC float64
	humidRH       float64
}

type RepositoryI interface {
	CalculateHeatIndex(tempC float64, rh float64) *Result
}

func NewRothfusz(minValidTempC float64, humidRH float64) RepositoryI {
	return &repo{
		minValidTempC: minValidTempC,
		humidRH:       humidRH,
	}
}

func (r *repo) CalculateHeatIndex(tempC float64, rh float64) *Result {
	var hiC float64
	// Rothfusz valid only >= configured temp
	if tempC < r.minValidTempC {
		hiC = tempC
	} else {
		hiC = r.calculateRothfusz(tempC, rh)
	}

	category := r.classifyHeatIndex(hiC, rh)

	return &Result{
		Data: Data{
			Temp: tempC,
			RH:   rh,
		},
		HeatIndexResult: HeatIndexResult{
			HeatIndexC:  hiC,
			Category:    category,
			Comfortable: category == CategoryIdeal,
		},
	}
}

func (r *repo) calculateRothfusz(tempC float64, rh float64) float64 {
	T := r.cToF(tempC)
	HI := -42.379 +
		2.04901523*T +
		10.14333127*rh -
		0.22475541*T*rh -
		0.00683783*T*T -
		0.05481717*rh*rh +
		0.00122874*T*T*rh +
		0.00085282*T*rh*rh -
		0.00000199*T*T*rh*rh

	return r.fToC(HI)
}

func (r *repo) classifyHeatIndex(hiC float64, rh float64) Category {
	// Tropical humidity rule
	if hiC < 27 && rh >= r.humidRH {
		return CategoryHumid
	}

	switch {

	case hiC < 27:
		return CategoryIdeal

	case hiC < 32:
		return CategoryCaution

	case hiC < 41:
		return CategoryExtremeCaution

	case hiC < 54:
		return CategoryDanger

	default:
		return CategoryExtremeDanger
	}
}

func (r *repo) cToF(c float64) float64 {
	return (c * 9 / 5) + 32
}

func (r *repo) fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
