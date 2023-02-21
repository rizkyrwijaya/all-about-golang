package main

import "fmt"

func main() {
	listOfData := []Identifiable[int]{
		newDataIdentifier(1, "test"),
		newDataIdentifier(2, "test"),
		newDataIdentifier(3, "test"),
	}

	res := convertListToMap[int](listOfData)
	for i, v := range res {
		fmt.Print(i)
		fmt.Println(v)
	}

	listOfDataString := []Identifiable[string]{
		newDataIdentifierString("123", "test"),
		newDataIdentifierString("1234", "test1"),
		newDataIdentifierString("1235", "test2"),
	}

	res2 := convertListToMap[string](listOfDataString)
	for i, v := range res2 {
		fmt.Print(i)
		fmt.Println(v)
	}
}

type Data struct {
	id   int
	info string
}

func (d *Data) GetIdentifier() int {
	return d.id
}

type DataString struct {
	id   string
	info string
}

func (d *DataString) GetIdentifier() string {
	return d.id
}

func newDataIdentifier(id int, data string) Identifiable[int] {
	return &Data{
		id:   id,
		info: data,
	}
}

func newDataIdentifierString(id string, data string) Identifiable[string] {
	return &DataString{
		id:   id,
		info: data,
	}
}

// Creating a Map converter from List with Identifier
func convertListToMap[K comparable, V Identifiable[K]](dataToChange []V) map[K]V {
	res := map[K]V{}
	for _, data := range dataToChange {
		res[data.GetIdentifier()] = data
	}
	return res
}

type Identifiable[K any] interface {
	GetIdentifier() K
}
