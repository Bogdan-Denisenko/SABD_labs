SRC = main.go
KEY = 42

all: obfuscate deobfuscate

obfuscate:
	go run $(SRC) -mode=obfuscate -input=input.xml -output=obfuscated.xml -key=$(KEY)

deobfuscate:
	go run $(SRC) -mode=deobfuscate -input=obfuscated.xml -output=deobfuscated.xml -key=$(KEY)

