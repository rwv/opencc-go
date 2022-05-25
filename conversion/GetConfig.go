package conversion

func GetConfig(conversion ConversionType) ConfigType {
	switch conversion {
	case S2HK:
		return S2HKConfig
	case TW2S:
		return TW2SConfig
	case T2TW:
		return T2TWConfig
	case S2TWP:
		return S2TWPConfig
	case S2TW:
		return S2TWConfig
	case T2S:
		return T2SConfig
	case T2HK:
		return T2HKConfig
	case HK2S:
		return HK2SConfig
	case TW2SP:
		return TW2SPConfig
	case S2T:
		return S2TConfig
	default:
		return ConfigType{}
	}
}
