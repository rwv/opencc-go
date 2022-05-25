package conversion

import "github.com/rwv/opencc-go/dictionary"

var T2TWConfig = ConfigType{
	Name:         "Traditional Chinese to Traditional Chinese (Taiwan standard)",
	Segmentation: dictionary.TWVariantsType,
	ConversionChain: []ConversionChainItem{
		{dictionary.TWVariantsType},
	},
}
