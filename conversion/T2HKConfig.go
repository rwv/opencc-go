package conversion

import "github.com/rwv/opencc-go/dictionary"

var T2HKConfig = ConfigType{
	Name:         "Traditional Chinese to Traditional Chinese (Hong Kong standard)",
	Segmentation: dictionary.HKVariantsType,
	ConversionChain: []ConversionChainItem{
		{dictionary.HKVariantsType},
	},
}
