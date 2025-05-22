Berikut adalah catatan penjelasan untuk kode program yang dapat dimasukkan ke dalam README proyek GitHub:

---

# 📌 **To-Do List CLI dengan Go**
Ini adalah program **To-Do List berbasis CLI** yang ditulis dalam bahasa **Go** untuk mengelola daftar tugas dengan berbagai fitur, seperti menambah, menghapus, dan menyimpan ke file JSON.

## 🚀 **Fitur Utama**
✅ Melihat daftar tugas  
➕ Menambah tugas baru  
🗑️ Menghapus tugas  
✅ Menandai tugas sebagai selesai  
✏️ Mengedit deskripsi tugas  
💾 Menyimpan tugas ke file JSON  
📂 Memuat tugas dari file JSON  

## 🏗 **Struktur Kode**
Program ini menggunakan **Go** dengan beberapa package bawaan:
- `bufio`: Membaca input pengguna dari terminal.
- `fmt`: Menampilkan teks ke terminal.
- `os` & `os/exec`: Menjalankan perintah untuk membersihkan layar.
- `strconv`: Konversi string ke integer untuk ID tugas.
- `strings`: Manipulasi teks seperti pemangkasan spasi.
- `time`: Menunda eksekusi saat keluar dari program.
- `runtime`: Mendeteksi sistem operasi untuk membersihkan layar secara otomatis.
- `todolist/utl`: Modul eksternal untuk manajemen tugas.

## 📜 **Penjelasan Kode**
- **`clearScreen()`** → Membersihkan terminal sesuai OS.
- **`pause()`** → Menunda eksekusi sampai pengguna menekan `Enter`.
- **`Menu()`** → Menampilkan menu interaktif dengan pilihan angka.
- **`main()`** → Fungsi utama, mengelola loop interaktif dan menangani input pengguna dengan `switch-case`.

**Setiap opsi menu memanggil fungsi dari `utl` untuk manipulasi daftar tugas:**
- `ViewTask()` → Menampilkan daftar tugas.
- `AddTask()` → Menambahkan tugas baru.
- `RemoveTask()` → Menghapus tugas berdasarkan ID.
- `MarkTask()` → Menandai tugas sebagai selesai.
- `ChangeTask()` → Mengubah deskripsi tugas.
- `SaveTasks()` → Menyimpan daftar tugas ke file JSON.
- `LoadTasks()` → Memuat daftar tugas dari file JSON.

## 🔧 **Cara Menjalankan**
Pastikan Go sudah terinstall, lalu jalankan:
```bash
go run main.go
```
Atau build terlebih dahulu dengan:
```bash
go build -o todolist main.go
./todolist
```

---
