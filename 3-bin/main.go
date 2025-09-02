package main

import (
	"fmt"

	"github.com/Vadim1t/DZ/3-bin/bins"
	"github.com/Vadim1t/DZ/3-bin/storage"
)

func main() {
	// Создаём новый список bin-объектов
	bl := bins.NewBinList()
	*bl = append(*bl, *bins.NewBin("id1", true, "MyBin"))

	// Сохраняем список в файл bins.json
	err := storage.SaveBins("bins.json", *bl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bins saved to bins.json")

	// Загружаем список из файла
	loadedBins, err := storage.LoadBins("bins.json")
	if err != nil {
		panic(err)
	}

	// Выводим загруженные данные (чтобы использовать loadedBins)
	fmt.Println("Loaded bins:")
	for _, b := range loadedBins {
		fmt.Printf("ID: %s, Private: %t, Name: %s, CreatedAt: %s\n",
			b.ID, b.Private, b.Name, b.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
