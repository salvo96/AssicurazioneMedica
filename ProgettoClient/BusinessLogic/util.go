package BusinessLogic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

//funzione di lettura stringa da tastiera e formattazione
func ReadStringFormat(text string) string {
	fmt.Println(text)
	in := bufio.NewReader(os.Stdin)
	data, _ := in.ReadString('\n')          //legge la stringa fino al carattere di ritorno a capo '\n'
	data = strings.TrimSuffix(data, "\r\n") //ripulisce la stringa rimuovendo i caratteri '\r' e '\n'
	return data
}

/*func ReadInt32(text string) int32 {
	var value int32
	fmt.Println(text)
	fmt.Scanf("%d\n", &value)
	return value
}

func ReadFloat32(text string) float32 {
	var value float32
	fmt.Println(text)
	fmt.Scanf("%d\n", &value)
	return value
}*/
