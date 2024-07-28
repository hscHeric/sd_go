package main

import (
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/hscHeric/sd_go/hello_proto/protos"
	"google.golang.org/protobuf/proto"
)

func main() {
	book := &pb.AddressBook{}

	println("Livro vazio")
	for _, value := range book.People {
		fmt.Println(value)
	}

	println("")

	p := pb.Person{
		Id:    1234,
		Name:  "Heric da Silva Cruz",
		Email: "hericsilvaho@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
		},
	}

	pp := pb.Person{
		Id:    1234,
		Name:  "Maria Anaelica Martins da Silva",
		Email: "aninhaMartins@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
		},
	}

	/*
	 * Em go não usamos funções implementadas diretamente nos objetos como
	 * ListPeople e AddPerson.
	 *
	 * Em GO para esse caso ao modificarmos book.People devemos fazer um append no mesmo
	 * pois estamos trabalhando com um slice. slices em go são interfaces com funções como
	 * a append citada para facilitar a manipulação de vetores.
	 */

	book.People = append(book.People, &p)
	book.People = append(book.People, &pp)

	println("Livro adicionando duas pessoas")
	for _, value := range book.People {
		fmt.Println(value)
	}
	println("")

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	filename := "AddressBook.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}

	_, err = file.Write(out)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	file.Close()

	file, err = os.Open(filename)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	book = &pb.AddressBook{}
	if err = proto.Unmarshal(content, book); err != nil {
		fmt.Println("Falha ao fazer Unmarshal")
		return
	}

	println("Listando livro após fazer marshal, escrever em arquivo e fazer Unmarshal")
	for _, value := range book.People {
		fmt.Println(value)
	}
	fmt.Println("")
}
