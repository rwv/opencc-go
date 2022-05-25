import json
import os
import glob

files = glob.glob("raw-data/config/*.json")


for file in files:
    basename = os.path.basename(file)
    configName = basename.split(".")[0]

    upperConfigName = configName.upper()

    upperConfigNameVariable = upperConfigName +"Config"

    outputFile = "package conversion\n\n"

    outputFile += 'import "github.com/rwv/opencc-go/dictionary"\n\n'

    outputFile += "var "+ upperConfigNameVariable + " = ConfigType{\n"

    with open(file, "r") as f:
        data = json.load(f)

        outputFile += f'\tName: "{data["name"]}",\n'

        segmentation_file = data["segmentation"]["dict"]["file"]
        segmentation_type = segmentation_file.split(".")[0]+"Type"

        outputFile += f'\tSegmentation: dictionary.{segmentation_type},\n'

        outputFile += "	ConversionChain: []ConversionChainItem{\n"

        conversion_chain = data["conversion_chain"]

        for conversion_chain_item in conversion_chain:
            if conversion_chain_item["dict"]["type"] == "txt":
                conversionFile = conversion_chain_item["dict"]["file"]
                conversionFileType = conversionFile.split(".")[0]+"Type"
                outputFile += f'\t\t{{dictionary.{conversionFileType},}},\n'
            if conversion_chain_item["dict"]["type"] == "group":
                dicts = conversion_chain_item["dict"]["dicts"]
                outputFile += "\t\t{"
                for dict in dicts:
                    conversionFile = dict["file"]
                    conversionFileType = conversionFile.split(".")[0]+"Type"
                    outputFile += f'dictionary.{conversionFileType}, '
                outputFile += "},\n"

        outputFile += "	},\n"
        outputFile += "}\n"

        with open(f"conversion/{upperConfigNameVariable}.go", "w") as f:
            f.write(outputFile)
