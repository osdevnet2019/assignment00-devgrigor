
all:
		go build makewords.go
		go build lowercase.go
		cat text_to_spell | ./makewords | ./lowercase | sort | uniq | comm -13 dictionary -
#		cc -o lowercase lowercase.c
#		cc -o makewords makewords.c
