package utl

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Menampilkan daftar tugas
func ViewTask(list *[]Task) {
	fmt.Println("\n" + bold + cyan + "-------------------------------------------------------------------")
	fmt.Println("                           Daftar Tugas                            ")
	fmt.Println("-------------------------------------------------------------------" + reset)

	if len(*list) == 0 {
		fmt.Println(yellow + "📭 Belum ada tugas yang ditambahkan..." + reset)
		return
	}

	fmt.Printf(bold+"%-4s %-5s %-40s %-15s\n"+reset, "No", "ID", "Deskripsi", "Status")
	fmt.Println("-------------------------------------------------------------------")

	for i, v := range *list {
		statusMark := "❌ Belum Selesai"
		if v.Status {
			statusMark = green + "✅ Selesai" + reset
		}
		fmt.Printf("%-4d %-5d %-39s %-15s\n", i+1, v.Id, v.Deskripsi, statusMark)
	}
}

// Menambahkan tugas baru
func AddTask(id int, deskripsi string, list *[]Task) {
	if deskripsi == "" {
		fmt.Println("⚠️ Deskripsi tugas tidak boleh kosong...")
		return
	}

	newTask := Task{
		Id:        id,
		Deskripsi: deskripsi,
		Status:    false,
	}

	*list = append(*list, newTask)
	IdNext++
	fmt.Println(green + "✅ Tugas baru berhasil ditambahkan!" + reset)
}

// Menghapus tugas berdasarkan ID
func RemoveTask(id int, list *[]Task) {
	index, err := idCheck(id, list)
	if err != nil {
		fmt.Println(err)
		return
	}

	*list = append((*list)[:index], (*list)[index+1:]...)
	fmt.Printf(green+"🗑️  Tugas dengan ID: %d berhasil dihapus.\n"+reset, id)
}

// Menandai tugas sebagai selesai
func MarkTask(id int, list *[]Task) {
	index, err := idCheck(id, list)
	if err != nil {
		fmt.Println(err)
		return
	}

	(*list)[index].Status = true
	fmt.Printf(green+"✅ Tugas dengan ID: %d telah ditandai selesai.\n"+reset, id)
}

// Mengubah deskripsi tugas
func ChangeTask(id int, text string, list *[]Task) {
	if text == "" {
		fmt.Println("⚠️ Deskripsi tugas tidak boleh kosong...")
		return
	}

	index, err := idCheck(id, list)
	if err != nil {
		fmt.Println(err)
		return
	}

	(*list)[index].Deskripsi = text
	fmt.Printf(green+"✏️ Tugas dengan ID: %d telah diubah.\n"+reset, id)
}

func SaveTasks(filename string, list *[]Task) error {
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Println(green + "💾 Data tugas berhasil disimpan!" + reset)
	return nil
}

func LoadTasks(filename string, list *[]Task) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var loaded []*Task
	err = json.Unmarshal(byteValue, &loaded)
	if err != nil {
		return err
	}

	// Reset TaskList dan IdNext
	*list = []Task{}
	maxId := 0

	for _, t := range loaded {
		*list = append(*list, *t)
		if t.Id > maxId {
			maxId = t.Id
		}
	}
	IdNext = maxId + 1

	fmt.Println(green + "📂 Data tugas berhasil dimuat!" + reset)
	return nil
}
