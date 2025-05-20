package utl

import "fmt"

// Mengecek apakah ID ada dalam daftar, mengembalikan indeks jika ditemukan
func idCheck(id int, list *[]Task) (int, error) {
	for i, task := range *list {
		if task.Id == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("⚠️  ID tugas %d tidak ditemukan", id)
}
