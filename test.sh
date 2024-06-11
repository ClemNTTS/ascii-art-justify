#!/bin/bash

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align right something standard"
go run main.go --align right something standard

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=right left standard"
go run main.go --align=right left standard

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=left right standard 
"
go run main.go --align=left right standard 

echo "/////////////////////////////////////////////////////////////////////////"
go run main.go --align=center hello shadow

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=justify "1 Two 4" shadow"
go run main.go --align=justify "1 Two 4" shadow

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=right 23/32 standard"
go run main.go --align=right 23/32 standard

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=right ABCabc123 thinkertoy"
go run main.go --align=right ABCabc123 thinkertoy

echo "/////////////////////////////////////////////////////////////////////////"
go run main.go --align=center "#$%&\"" thinkertoy

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=left '23Hello World!' standard "
go run main.go --align=left '23Hello World!' standard 

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=justify 'HELLO there HOW are YOU?!' thinkertoy"
go run main.go --align=justify 'HELLO there HOW are YOU?!' thinkertoy

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=right "a -> A b -> B c -> C" shadow "
go run main.go --align=right "a -> A b -> B c -> C" shadow 

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=right abcd shadow"
go run main.go --align=right abcd shadow 

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=center ola standard "
go run main.go --align=center ola standard 

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=center HeLlO standard"
go run main.go --align=center HeLlO standard

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=justify "hello 2024""
go run main.go --align=justify "hello 2024"

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=right "#?.:;" "
go run main.go --align=right "#?.:;" 

echo "/////////////////////////////////////////////////////////////////////////"
echo "go run main.go --align=left "Hello 2024 ?@" "
go run main.go --align=left "Hello 2024 ?@" 