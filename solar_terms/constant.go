package solar_terms

import "time"

type SolarTermsEnum int

const (
	Unknown           SolarTermsEnum = iota
	BeginningOfSpring                // 立春
	RainWater                        // 雨水
	WakingOfInsects                  // 惊蛰
	SpringEquinox                    // 春分
	PureBrightness                   // 清明
	GrainRain                        // 谷雨
	BeginningOfSummer                // 立夏
	GrainFull                        // 小满
	GrainInEar                       // 芒种
	SummerSolstice                   // 夏至
	SlightHeat                       // 小暑
	GreatHeat                        // 大暑
	BeginningOfAutumn                // 立秋
	LimitOfHeat                      // 处暑
	WhiteDew                         // 白露
	AutumnalEquinox                  // 秋分
	ColdDew                          // 寒露
	FrostDescent                     // 霜降
	BeginningOfWinter                // 立冬
	SlightSnow                       // 小雪
	GreatSnow                        // 大雪
	WinterSolstice                   // 冬至
	SlightCold                       // 小寒
	GreatCold                        // 大寒
)

func (v SolarTermsEnum) Label() string {
	switch v {
	case Unknown:
		return ""
	case BeginningOfSpring:
		return "立春"
	case RainWater:
		return "雨水"
	case WakingOfInsects:
		return "惊蛰"
	case SpringEquinox:
		return "春分"
	case PureBrightness:
		return "清明"
	case GrainRain:
		return "谷雨"
	case BeginningOfSummer:
		return "立夏"
	case GrainFull:
		return "小满"
	case GrainInEar:
		return "芒种"
	case SummerSolstice:
		return "夏至"
	case SlightHeat:
		return "小暑"
	case GreatHeat:
		return "大暑"
	case BeginningOfAutumn:
		return "立秋"
	case LimitOfHeat:
		return "处暑"
	case WhiteDew:
		return "白露"
	case AutumnalEquinox:
		return "秋分"
	case ColdDew:
		return "寒露"
	case FrostDescent:
		return "霜降"
	case BeginningOfWinter:
		return "立冬"
	case SlightSnow:
		return "小雪"
	case GreatSnow:
		return "大雪"
	case WinterSolstice:
		return "冬至"
	case SlightCold:
		return "小寒"
	case GreatCold:
		return "大寒"
	default:
		return ""
	}
}

type solarTermDetail struct {
	SolarTermsMonth       time.Month
	SolarTermsSpecialYear []int
	SolarTermsSpecial     int
	SolarTermsCValue      float64
}

var solarTermCValueMap = map[SolarTermsEnum]map[int]solarTermDetail{
	BeginningOfSpring: {21: solarTermDetail{
		SolarTermsMonth:   2,
		SolarTermsSpecial: 0,
		SolarTermsCValue:  3.87,
	}},
	RainWater: {21: solarTermDetail{
		SolarTermsMonth:       2,
		SolarTermsSpecialYear: []int{2026},
		SolarTermsSpecial:     -1,
		SolarTermsCValue:      18.73,
	}},
	WakingOfInsects: {21: solarTermDetail{
		SolarTermsMonth:       3,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      5.63,
	}},
	SpringEquinox: {21: solarTermDetail{
		SolarTermsMonth:       3,
		SolarTermsSpecialYear: []int{2084},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      20.646,
	}},
	PureBrightness: {21: solarTermDetail{
		SolarTermsMonth:       4,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      4.81,
	}},
	GrainRain: {21: solarTermDetail{
		SolarTermsMonth:       4,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      20.1,
	}},
	BeginningOfSummer: {21: solarTermDetail{
		SolarTermsMonth:       5,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      5.52,
	}},
	GrainFull: {21: solarTermDetail{
		SolarTermsMonth:       5,
		SolarTermsSpecialYear: []int{2008},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      21.04,
	}},
	GrainInEar: {21: solarTermDetail{
		SolarTermsMonth:       6,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      5.678,
	}},
	SummerSolstice: {21: solarTermDetail{
		SolarTermsMonth:       6,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      21.37,
	}},
	SlightHeat: {21: solarTermDetail{
		SolarTermsMonth:       7,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      7.108,
	}},
	GreatHeat: {21: solarTermDetail{
		SolarTermsMonth:       7,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      22.83,
	}},
	BeginningOfAutumn: {21: solarTermDetail{
		SolarTermsMonth:       8,
		SolarTermsSpecialYear: []int{2002},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      7.5,
	}},
	LimitOfHeat: {21: solarTermDetail{
		SolarTermsMonth:       8,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      23.13,
	}},
	WhiteDew: {21: solarTermDetail{
		SolarTermsMonth:       9,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      7.646,
	}},
	AutumnalEquinox: {21: solarTermDetail{
		SolarTermsMonth:       9,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      23.042,
	}},
	ColdDew: {21: solarTermDetail{
		SolarTermsMonth:       10,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      8.138,
	}},
	FrostDescent: {21: solarTermDetail{
		SolarTermsMonth:       10,
		SolarTermsSpecialYear: []int{2089},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      23.438,
	}},
	BeginningOfWinter: {21: solarTermDetail{
		SolarTermsMonth:       11,
		SolarTermsSpecialYear: []int{2089},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      7.438,
	}},
	SlightSnow: {21: solarTermDetail{
		SolarTermsMonth:       11,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      22.36,
	}},
	GreatSnow: {21: solarTermDetail{
		SolarTermsMonth:       12,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      7.18,
	}},
	WinterSolstice: {21: solarTermDetail{
		SolarTermsMonth:       12,
		SolarTermsSpecialYear: []int{2021},
		SolarTermsSpecial:     -1,
		SolarTermsCValue:      21.94,
	}},
	SlightCold: {21: solarTermDetail{
		SolarTermsMonth:       1,
		SolarTermsSpecialYear: []int{},
		SolarTermsSpecial:     0,
		SolarTermsCValue:      5.4055,
	}},
	GreatCold: {21: solarTermDetail{
		SolarTermsMonth:       1,
		SolarTermsSpecialYear: []int{2082},
		SolarTermsSpecial:     1,
		SolarTermsCValue:      20.12,
	}},
}
