package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	"todolist/utl"
)

// ANSI escape codes untuk warna dan gaya
const (
	reset     = "\033[0m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	bold      = "\033[1m"
	boldCyan  = "\033[1;36m"
	boldGreen = "\033[1;32m"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pause() {
	fmt.Print("\nTekan Enter untuk kembali ke menu...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Menu() {
	fmt.Println(boldCyan + "\n=====================================")
	fmt.Println("        APLIKASI TO-DO LIST")
	fmt.Println("=====================================" + reset)
	fmt.Println(blue + "(1)" + reset + " Lihat Daftar Tugas")
	fmt.Println(blue + "(2)" + reset + " Tambah Tugas Baru")
	fmt.Println(blue + "(3)" + reset + " Hapus Tugas")
	fmt.Println(blue + "(4)" + reset + " Tandai Tugas Selesai")
	fmt.Println(blue + "(5)" + reset + " Ubah Deskripsi Tugas")
	fmt.Println(blue + "(6)" + reset + " Simpan Tugas ke File")
	fmt.Println(blue + "(7)" + reset + " Buka Tugas dari File")

	fmt.Println(blue + "(0)" + reset + " Keluar")
	fmt.Println("-------------------------------------")
	fmt.Print("\n" + bold + "Pilih menu: " + reset)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		Menu()

		menuStr, _ := reader.ReadString('\n')
		inputMenu, errMenu := strconv.Atoi(strings.TrimSpace(menuStr))

		if errMenu != nil {
			fmt.Println(yellow + "⚠️  Input tidak valid. Masukkan angka!" + reset)
			pause()
			continue
		}

		switch inputMenu {
		case 1: // Lihat Daftar Tugas
			clearScreen()
			fmt.Println(boldGreen + "📋 Lihat Daftar Tugas:" + reset)
			utl.ViewTask(&utl.TaskList)
			pause()

		case 2: // Tambah Tugas Baru
			clearScreen()
			fmt.Println(cyan + bold + "➕ Tambah Tugas Baru:" + reset)
			utl.ViewTask(&utl.TaskList)
			fmt.Print("\nMasukkan deskripsi tugas: ")
			desc, _ := reader.ReadString('\n')
			utl.AddTask(utl.IdNext, strings.TrimSpace(desc), &utl.TaskList)
			pause()

		case 3: // Hapus Tugas
			clearScreen()
			fmt.Println(bold + "🗑️ Hapus Tugas:" + reset)

			if len(utl.TaskList) == 0 {
				fmt.Println(yellow + "📭 Belum ada tugas yang ditambahkan..." + reset)
				pause()
				continue
			}

			utl.ViewTask(&utl.TaskList)
			fmt.Print("\nMasukkan ID tugas yang ingin dihapus: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))

			if err == nil {
				utl.RemoveTask(id, &utl.TaskList)
			} else {
				fmt.Println(yellow + "⚠️  ID tidak valid!" + reset)
			}
			pause()

		case 4: // Tandai Tugas Selesai
			clearScreen()
			fmt.Println(boldGreen + "✅ Tandai Tugas Selesai:" + reset)

			if len(utl.TaskList) == 0 {
				fmt.Println(yellow + "📭 Belum ada tugas yang ditambahkan..." + reset)
				pause()
				continue
			}

			utl.ViewTask(&utl.TaskList)
			fmt.Print("\nMasukkan ID tugas yang ingin ditandai: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))

			if err == nil {
				utl.MarkTask(id, &utl.TaskList)
			} else {
				fmt.Println(yellow + "⚠️  ID tidak valid!" + reset)
			}
			pause()

		case 5: // Ubah Deskripsi Tugas
			clearScreen()
			fmt.Println(bold + "✏️ Ubah Deskripsi Tugas:" + reset)

			if len(utl.TaskList) == 0 {
				fmt.Println(yellow + "📭 Belum ada tugas yang ditambahkan..." + reset)
				pause()
				continue
			}

			utl.ViewTask(&utl.TaskList)
			fmt.Print("\nMasukkan ID tugas yang ingin diubah: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))

			if err != nil {
				fmt.Println(yellow + "⚠️  ID tidak valid!" + reset)
				pause()
				continue
			}

			fmt.Print("\nMasukkan deskripsi tugas yang baru: ")
			textStr, _ := reader.ReadString('\n')
			text := strings.TrimSpace(textStr)
			utl.ChangeTask(id, text, &utl.TaskList)
			pause()

		case 6: // Simpan file
			clearScreen()
			fmt.Println(bold + "💾 Simpan Tugas ke File" + reset)
			err := utl.SaveTasks("data.json", &utl.TaskList)
			if err != nil {
				fmt.Println("Gagal menyimpan:", err)
			}
			pause()

		case 7: // Load file
			clearScreen()
			fmt.Println(bold + "📂 Buka Tugas dari File" + reset)
			err := utl.LoadTasks("data.json", &utl.TaskList)
			if err != nil {
				fmt.Println("Gagal memuat data:", err)
			}
			pause()

		case 0: // Keluar dari aplikasi
			clearScreen()
			fmt.Println(cyan + "👋 Terima kasih telah menggunakan aplikasi ini!" + reset)
			time.Sleep(1 * time.Second)
			return

		default:
			fmt.Println(yellow + "⚠️  Menu tidak ditemukan!" + reset)
			pause()
		}
	}
}
