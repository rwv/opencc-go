package config

type SegmentationDict struct {
	Type string `json:"type"`
	File string `json:"file"`
}

type Segmentation struct {
	Type string           `json:"type"`
	Dict SegmentationDict `json:"dict"`
}

type ConversionChainDictDict struct {
	Type string `json:"type"`
	File string `json:"file"`
}

type ConversionChainDict struct {
	Type  string                    `json:"type"`
	Dicts []ConversionChainDictDict `json:"dicts"`
}

type ConversionChainItem struct {
	Dict ConversionChainDict `json:"dict"`
}

type Config struct {
	Name            string                `json:"name"`
	Segmentation    Segmentation          `json:"segmentation"`
	ConversionChain []ConversionChainItem `json:"conversion_chain"`
}
