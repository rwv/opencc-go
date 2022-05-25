import os
import glob

files = glob.glob("raw-data/dictionary/*.txt")


for file in files:
    basename = os.path.basename(file)
    dictionaryName = basename.split(".")[0]

    outputFile = "package dictionary\n\n"
    outputFile += "var "+ dictionaryName + " = Dictionary{\n"

    with open(file, "r") as f:
        lines = f.readlines()
        for line in lines:
            mapping = line.split("\t")
            assert len(mapping) == 2
            if len(mapping) == 2:
                outputFile += f'	"{mapping[0]}": "{mapping[1].strip()}",\n'

    outputFile += "}\n"

    with open(f"dictionary/{dictionaryName}.go", "w") as f:
        f.write(outputFile)

# Write types to file

outputFile = "package dictionary\n\n"
outputFile += "const (\n"
outputFile += "\t _ DictionaryType = iota\n"

for file in files:
    basename = os.path.basename(file)
    dictionaryName = basename.split(".")[0]
    outputFile += f"\t{dictionaryName}Type\n"

outputFile += ")\n"

with open(f"dictionary/DictionaryTypes.go", "w") as f:
    f.write(outputFile)

os.system("go fmt ./dictionary/*")