import os
import glob

files = glob.glob("raw-data/dictionary/*.txt")


for file in files:
    basename = os.path.basename(file)
    dictionaryName = basename.split(".")[0]

    outputFile = "package dictionary\n"
    outputFile += "var "+ dictionaryName + " = map[string]string{\n"

    with open(file, "r") as f:
        lines = f.readlines()
        for line in lines:
            mapping = line.split("\t")
            if len(mapping) == 2:
                outputFile += f'	"{mapping[0]}": "{mapping[1].strip()}",\n'

    outputFile += "}\n"

    with open(f"dictionary/{dictionaryName}.go", "w") as f:
        f.write(outputFile)
