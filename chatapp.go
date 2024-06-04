package main

import (
	"fmt"
)

type User struct {
	Username string
	Password string
	Approved bool
}

type Chat struct {
	Sender    string
	Recipient string
	Message   string
}

type Group struct {
	Name     string
	Admin    string
	Members  []User
	Messages []GroupMessage
}

type GroupMessage struct {
	Sender  string
	Message string
}

const (
	MaxUsers  = 5
	MaxChats  = 100
	MaxGroups = 10
)

var Users [MaxUsers]User
var UserCount int
var Chats [MaxChats]Chat
var ChatCount int
var Groups [MaxGroups]Group
var GroupCount int

func main() {
	initializeDummyData()

	var choice int
	for {
		fmt.Println("----------------------------------")
		fmt.Println("<<  Welcome to Algolia ChatApp! >>")
		fmt.Println("----------------------------------")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Admin Login")
		fmt.Println("4. Exit")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			RegisterUser()
		case 2:
			LoginUser()
		case 3:
			AdminLogin()
		case 4:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func initializeDummyData() {
	Users[0] = User{Username: "user1", Password: "pass1", Approved: true}
	Users[1] = User{Username: "user2", Password: "pass2", Approved: true}
	UserCount = 2

	Groups[0] = Group{Name: "group1", Admin: "user1", Members: []User{Users[0], Users[1]}}
	GroupCount = 1
}

// register user
func RegisterUser() {
	if UserCount >= MaxUsers {
		fmt.Println("Tidak dapat mendaftarkan pengguna baru, batas maksimum tercapai.")
		return
	}

	var username, password string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	Users[UserCount] = User{Username: username, Password: password, Approved: false}
	UserCount++
	fmt.Println("Registrasi berhasil, menunggu persetujuan admin.\n")
}

// login user
func LoginUser() {
	var username, password string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for i := 0; i < UserCount; i++ {
		if Users[i].Username == username && Users[i].Password == password {
			if Users[i].Approved {
				fmt.Println("Login berhasil.\n")
				userMenu(&Users[i])
			} else {
				fmt.Println("Akun belum disetujui oleh admin.")
			}
			return
		}
	}
	fmt.Println("Username atau password salah.")
}

// menu user
func userMenu(user *User) {
	var choice int
	for {
		fmt.Println("----------------------------------")
		fmt.Println("<<           Menu Pengguna      >>")
		fmt.Println("----------------------------------")
		fmt.Println("1. Kirim Pesan Pribadi")
		fmt.Println("2. Baca Pesan Masuk")
		fmt.Println("3. Balas Pesan Masuk")
		fmt.Println("4. Hapus Pesan Masuk")
		fmt.Println("5. Buat Grup Chatting")
		fmt.Println("6. Kirim Pesan ke Grup")
		fmt.Println("7. Lihat Peserta Grup")
		fmt.Println("8. Tambah Peserta Grup")
		fmt.Println("9. Baca Pesan Grup")
		fmt.Println("10. Balas Pesan Grup")
		fmt.Println("11. Logout")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			sendPrivateMessage(user)
		case 2:
			readPrivateMessages(user)
		case 3:
			replyToPrivateMessage(user)
		case 4:
			deletePrivateMessage(user)
		case 5:
			createGroup(user)
		case 6:
			sendGroupMessage(user)
		case 7:
			viewGroupMembers(user)
		case 8:
			addMemberToGroup()
		case 9:
			readGroupMessages(user)
		case 10:
			replyToGroupMessage(user)
		case 11:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// kirim pesan pribadi
func sendPrivateMessage(user *User) {
	var recipientUsername, message string
	fmt.Print("Masukkan username penerima: ")
	fmt.Scan(&recipientUsername)
	fmt.Print("Masukkan pesan: ")
	fmt.Scan(&message)

	for i := 0; i < UserCount; i++ {
		if Users[i].Username == recipientUsername {
			Chats[ChatCount] = Chat{Sender: user.Username, Recipient: recipientUsername, Message: message}
			ChatCount++
			fmt.Println("Pesan terkirim.")
			fmt.Printf("Private message from %s to %s: %s\n", user.Username, recipientUsername, message)
			fmt.Println()
			return
		}
	}
	fmt.Println("Pengguna tidak ditemukan.\n")
}

// baca pesan pribadi
func readPrivateMessages(user *User) {
	found := false
	fmt.Println("Pesan masuk:")
	for i := 0; i < ChatCount; i++ {
		if Chats[i].Recipient == user.Username {
			fmt.Printf("From %s: %s\n", Chats[i].Sender, Chats[i].Message)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada pesan yang tersedia.\n")
	}
	fmt.Println()
}

// balas pesan pribadi
func replyToPrivateMessage(user *User) {
	found := false
	fmt.Print("Masukkan username pengirim pesan yang ingin dibalas: ")
	var senderUsername, message string
	fmt.Scan(&senderUsername)
	fmt.Print("Masukkan balasan: ")
	fmt.Scan(&message)

	for i := 0; i < ChatCount; i++ {
		if Chats[i].Recipient == user.Username {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Tidak ada pesan yang tersedia untuk dibalas.\n")
		return
	}

	for i := 0; i < UserCount; i++ {
		if Users[i].Username == senderUsername {
			Chats[ChatCount] = Chat{Sender: user.Username, Recipient: senderUsername, Message: message}
			ChatCount++
			fmt.Println("Balasan terkirim.")
			fmt.Printf("Reply from %s to %s: %s\n", user.Username, senderUsername, message)
			fmt.Println()
			return
		}
	}
	fmt.Println("Pengguna tidak ditemukan.\n")
}

// hapus pesan pribadi
func deletePrivateMessage(user *User) {
	found := false
	fmt.Println("Pesan masuk:")
	for i := 0; i < ChatCount; i++ {
		if Chats[i].Recipient == user.Username {
			fmt.Printf("From %s: %s\n", Chats[i].Sender, Chats[i].Message)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada pesan yang tersedia untuk dihapus.\n")
		return
	}

	var senderUsername, message string
	fmt.Print("Masukkan username pengirim pesan yang ingin dihapus: ")
	fmt.Scan(&senderUsername)
	fmt.Print("Masukkan pesan yang ingin dihapus: ")
	fmt.Scan(&message)

	for i := 0; i < ChatCount; i++ {
		if Chats[i].Recipient == user.Username && Chats[i].Sender == senderUsername && Chats[i].Message == message {
			// Hapus pesan dengan menggeser elemen setelahnya ke kiri
			copy(Chats[i:], Chats[i+1:ChatCount])
			ChatCount--
			fmt.Println("Pesan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pesan tidak ditemukan.\n")
}

// buat grup
func createGroup(user *User) {
	if GroupCount >= MaxGroups {
		fmt.Println("Tidak dapat membuat grup baru, batas maksimum tercapai.")
		return
	}

	var groupName string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scan(&groupName)

	Groups[GroupCount] = Group{
		Name:    groupName,
		Admin:   user.Username,
		Members: []User{*user}, // Menambahkan user sebagai anggota pertama
	}
	GroupCount++
	fmt.Println("Grup berhasil dibuat.")
	fmt.Println("---------------   Data Grup   ------------------")
	fmt.Printf("Nama Grup: %s, Admin: %s, Members: %s\n", groupName, user.Username, user.Username)
	fmt.Println()
}

// tambah peserta dalam grup
func addMemberToGroup() {
	var groupName string
	fmt.Print("Masukkan nama grup yang akan ditambahkan anggota: ")
	fmt.Scan(&groupName)

	// Cari grup berdasarkan nama
	var groupIndex int = -1
	for i, group := range Groups {
		if group.Name == groupName {
			groupIndex = i
			break
		}
	}

	if groupIndex == -1 {
		fmt.Println("Nama grup tidak ditemukan.\n")
		return
	}

	if len(Groups[groupIndex].Members) >= MaxUsers {
		fmt.Println("Tidak dapat menambahkan anggota baru, batas maksimum tercapai.")
		return
	}

	var username string
	fmt.Print("Masukkan nama pengguna baru: ")
	fmt.Scan(&username)

	newUser := User{Username: username}
	Groups[groupIndex].Members = append(Groups[groupIndex].Members, newUser)
	fmt.Println("Anggota berhasil ditambahkan.")
	fmt.Println("---------------   Data Grup   ------------------")
	fmt.Printf("Nama Grup: %s, Admin: %s, Jumlah Anggota: %d\n",
		Groups[groupIndex].Name, Groups[groupIndex].Admin, len(Groups[groupIndex].Members))
	fmt.Println()
}

// Kirim pesan ke grup
func sendGroupMessage(user *User) {
	var groupName, message string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scan(&groupName)
	fmt.Print("Masukkan pesan: ")
	fmt.Scan(&message)

	for i := 0; i < GroupCount; i++ {
		if Groups[i].Name == groupName {
			Groups[i].Messages = append(Groups[i].Messages, GroupMessage{Sender: user.Username, Message: message})
			fmt.Println("Pesan terkirim ke grup.")
			fmt.Printf("Pesan terkirim ke Grup: %s, Pengirim: %s, Pesan: %s\n", groupName, user.Username, message)
			fmt.Println()
			return
		}
	}
	fmt.Println("Grup tidak ditemukan.")
}

// baca pesan grup
func readGroupMessages(user *User) {
	var groupName string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scan(&groupName)

	group := findGroupByName(groupName)
	if group != nil {
		if isUserInGroup(user, group) {
			if len(group.Messages) > 0 {
				fmt.Println("Pesan dalam grup:")
				for i := 0; i < len(group.Messages); i++ {
					msg := group.Messages[i]
					fmt.Printf("From %s: %s\n", msg.Sender, msg.Message)
				}
			} else {
				fmt.Println("Tidak ada pesan yang tersedia dalam grup ini.")
			}
		} else {
			fmt.Println("Anda bukan anggota dari grup ini.")
		}
	} else {
		fmt.Println("Grup tidak ditemukan.")
	}
}

// balas pesan grup
func replyToGroupMessage(user *User) {
	var groupName, message string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scan(&groupName)
	fmt.Print("Masukkan pesan balasan: ")
	fmt.Scan(&message)

	group := findGroupByName(groupName)
	if group != nil {
		if isUserInGroup(user, group) {
			group.Messages = append(group.Messages, GroupMessage{Sender: user.Username, Message: message})
			fmt.Println("Balasan terkirim ke grup.")
			fmt.Printf("Balasan terkirim ke Grup: %s, Pengirim: %s, Pesan: %s\n", groupName, user.Username, message)
		} else {
			fmt.Println("Anda bukan anggota dari grup ini.")
		}
	} else {
		fmt.Println("Grup tidak ditemukan.")
	}
}

// cari grup berdasarkan nama
func findGroupByName(name string) *Group {
	for i := 0; i < GroupCount; i++ {
		if Groups[i].Name == name {
			return &Groups[i]
		}
	}
	return nil
}

// cek apakah pengguna adalah anggota dari grup
func isUserInGroup(user *User, group *Group) bool {
	for i := 0; i < len(group.Members); i++ {
		if group.Members[i].Username == user.Username {
			return true
		}
	}
	return false
}

// lihat Peserta dalam Group
func viewGroupMembers(user *User) {
	var groupName string
	fmt.Print("Masukkan nama grup: ")
	fmt.Scan(&groupName)

	for i := 0; i < GroupCount; i++ {
		if Groups[i].Name == groupName {
			fmt.Println("Peserta dalam grup : ")
			for j := 0; j < len(Groups[i].Members); j++ {
				member := Groups[i].Members[j]
				fmt.Println("-", member.Username)
			}
			return
		}
	}
	fmt.Println("Grup tidak ditemukan.")
	fmt.Println()
}

func AdminLogin() {
	var password string
	fmt.Print("Masukkan password admin: ")
	fmt.Scan(&password)

	if password == "admin" {
		adminMenu()
	} else {
		fmt.Println("Password salah.")
	}
}

func adminMenu() {
	var choice int
	for {
		fmt.Println("------------------------------------")
		fmt.Println("<<           Menu Admin           >>")
		fmt.Println("------------------------------------")
		fmt.Println("1. Setujui/Penolakan Registrasi Akun")
		fmt.Println("2. Cetak Daftar Akun")
		fmt.Println("3. Hapus Akun Pengguna")
		fmt.Println("4. Logout")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			approveRejectUsers()
		case 2:
			printUserList()
		case 3:
			deleteUserAccount()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func approveRejectUsers() {
	for i := 0; i < UserCount; i++ {
		if !Users[i].Approved {
			fmt.Printf("Setujui akun %s? (y/n): ", Users[i].Username)
			var response string
			fmt.Scan(&response)
			if response == "y" {
				Users[i].Approved = true
				fmt.Println("Akun disetujui.\n")
			} else {
				fmt.Println("Akun ditolak.\n")
			}
		}
	}
}

// func printUserList() {
// 	fmt.Println("=== Daftar Akun Pengguna ===")
// 	for i := 0; i < UserCount; i++ {
// 		fmt.Printf("Username: %s, Disetujui: %t\n", Users[i].Username, Users[i].Approved)
// 	}
// 	fmt.Println()
// }

func printUserList() {
	var sortBy string
	fmt.Println("Pilihan urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&sortBy)

	switch sortBy {
	case "1":
		SelectionSort(Users[:UserCount], "username")
	case "2":
		// Untuk descending, Anda dapat mengubah arah urutan hasil sorting
		SelectionSortDesc(Users[:UserCount], "username")
	default:
		fmt.Println("Pilihan tidak valid. Pengurutan dibatalkan.")
	}

	j := 1
	fmt.Println("Daftar Akun Pengguna :")
	fmt.Println("-------------------------------------")
	fmt.Printf("%-3s %10s %20s \n", "No", "Username", "Disetujui")
	fmt.Println("-------------------------------------")
	for i := 0; i < UserCount; i++ {
		// fmt.Printf("Username: %s, Disetujui: %t\n", Users[i].Username, Users[i].Approved)
		fmt.Printf("%-3d %10s %20t \n", j, Users[i].Username, Users[i].Approved)
		j++
	}
	fmt.Println()
}

// delete akun user
func deleteUserAccount() {
	var username string
	fmt.Print("Masukkan username yang akan dihapus: ")
	fmt.Scan(&username)

	for i := 0; i < UserCount; i++ {
		if Users[i].Username == username {
			// Hapus pengguna dengan menggeser elemen setelahnya ke kiri
			copy(Users[i:], Users[i+1:UserCount])
			UserCount--
			fmt.Println("Pengguna berhasil dihapus.\n")
			return
		}
	}

	fmt.Println("Pengguna tidak ditemukan. \n")
}

// SelectionSortDesc mengurutkan slice dari User berdasarkan kategori tertentu secara descending
func SelectionSortDesc(users []User, sortBy string) {
	n := len(users)
	for i := 0; i < n-1; i++ {
		maxIndex := i
		for j := i + 1; j < n; j++ {
			switch sortBy {
			case "username":
				if users[j].Username > users[maxIndex].Username {
					maxIndex = j
				}
			}
		}
		// Tukar posisi elemen
		users[i], users[maxIndex] = users[maxIndex], users[i]
	}
}

// SelectionSort mengurutkan slice dari User berdasarkan kategori tertentu
func SelectionSort(users []User, sortBy string) {
	n := len(users)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			switch sortBy {
			case "username":
				if users[j].Username < users[minIndex].Username {
					minIndex = j
				}
			}
		}
		// Tukar posisi elemen
		users[i], users[minIndex] = users[minIndex], users[i]
	}
}
