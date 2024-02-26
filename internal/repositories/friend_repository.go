package repositories

import (
	"math/rand"
	"strconv"

	"kominfo-go/internal/entities"
)

var friends = map[int]entities.Friend{
	1:   {Absen: 1, Nama: "Calman", Alamat: "Jalan Kehidupan No. " + strconv.Itoa(rand.Intn(100)), Pekerjaan: "Software Engineer", Alasan: "Ingin belajar bahasa pemrograman Go"},
	280: {Absen: 280, Nama: "1957356840-818 ROSSARIO CATHERINE ELFRIDA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 51", Pekerjaan: "Golang", Alasan: "006"},
	295: {Absen: 295, Nama: "1957356840-863 AFIF AMIRULLAH Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 42", Pekerjaan: "Golang", Alasan: "006"},
	251: {Absen: 251, Nama: "1957356840-743 MULIA HARTAWAN NEGARA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 75", Pekerjaan: "Golang", Alasan: "006"},
	253: {Absen: 253, Nama: "1957356840-751 RIEFKY ARIF IBRAHIM Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 17", Pekerjaan: "Golang", Alasan: "006"},
	259: {Absen: 259, Nama: "1957356840-77 ALFUL LAILA S Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 33", Pekerjaan: "Golang", Alasan: "006"},
	263: {Absen: 263, Nama: "1957356840-778 KURNIA RAIHAN ARDIAN Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 10", Pekerjaan: "Golang", Alasan: "006"},
	276: {Absen: 276, Nama: "1957356840-809 Rafi Dimas Ariyanto Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 84", Pekerjaan: "Golang", Alasan: "006"},
	278: {Absen: 278, Nama: "1957356840-812 Izhar Kadir Saputra Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 31", Pekerjaan: "Golang", Alasan: "006"},
	283: {Absen: 283, Nama: "1957356840-833 STEVAN GEOVANUS PURBA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 26", Pekerjaan: "Golang", Alasan: "006"},
	285: {Absen: 285, Nama: "1957356840-835 SAHRUL FAZRI UDIN Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 18", Pekerjaan: "Golang", Alasan: "006"},
	252: {Absen: 252, Nama: "1957356840-748 DZULFIQAR AKMALUDDIN Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 57", Pekerjaan: "Golang", Alasan: "006"},
	254: {Absen: 254, Nama: "1957356840-752 Permadi Hidayat Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 68", Pekerjaan: "Golang", Alasan: "006"},
	256: {Absen: 256, Nama: "1957356840-761 ZIKRI KURNIA AIZET Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 26", Pekerjaan: "Golang", Alasan: "006"},
	261: {Absen: 261, Nama: "1957356840-775 SAMUEL IMMANUEL HERLINTON SIBUEA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 10", Pekerjaan: "Golang", Alasan: "006"},
	290: {Absen: 290, Nama: "1957356840-854 DIANDRA AURELLIEA AZRA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 87", Pekerjaan: "Golang", Alasan: "006"},
	287: {Absen: 287, Nama: "1957356840-850 HIDAYATULLAH A Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 86", Pekerjaan: "Golang", Alasan: "006"},
	288: {Absen: 288, Nama: "1957356840-852 NURUL ASMI AMALIA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 68", Pekerjaan: "Golang", Alasan: "006"},
	294: {Absen: 294, Nama: "1957356840-862 RIDWAN SATRIA WICAKSANA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 72", Pekerjaan: "Golang", Alasan: "006"},
	296: {Absen: 296, Nama: "1957356840-867 RIFQI DAMAR ALI Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 6", Pekerjaan: "Golang", Alasan: "006"},
	267: {Absen: 267, Nama: "1957356840-789 Muhammad Dicky Athalla Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 67", Pekerjaan: "Golang", Alasan: "006"},
	269: {Absen: 269, Nama: "1957356840-790 VINCENTIUS CLARISHNA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 78", Pekerjaan: "Golang", Alasan: "006"},
	270: {Absen: 270, Nama: "1957356840-795 DAFFA ASLAM MULTAZAM Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 79", Pekerjaan: "Golang", Alasan: "006"},
	284: {Absen: 284, Nama: "1957356840-834 ILHAM AKBAR SYAHPUTRA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 24", Pekerjaan: "Golang", Alasan: "006"},
	297: {Absen: 297, Nama: "1957356840-872 NAJMUDDIN HABIB Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 73", Pekerjaan: "Golang", Alasan: "006"},
	274: {Absen: 274, Nama: "1957356840-805 Tata Redha Al Fath Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 53", Pekerjaan: "Golang", Alasan: "006"},
	279: {Absen: 279, Nama: "1957356840-813 MIFTAHUL JANNAH Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 13", Pekerjaan: "Golang", Alasan: "006"},
	282: {Absen: 282, Nama: "1957356840-831 SULFAN AIDID Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 98", Pekerjaan: "Golang", Alasan: "006"},
	257: {Absen: 257, Nama: "1957356840-763 Fungki Prasetya Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 18", Pekerjaan: "Golang", Alasan: "006"},
	266: {Absen: 266, Nama: "1957356840-787 Ervany Septa Prawara Arisanto Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 24", Pekerjaan: "Golang", Alasan: "006"},
	268: {Absen: 268, Nama: "1957356840-79 RIZKI KUSNADI Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 78", Pekerjaan: "Golang", Alasan: "006"},
	272: {Absen: 272, Nama: "1957356840-80 ACHSANI TAQWIM Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 99", Pekerjaan: "Golang", Alasan: "006"},
	258: {Absen: 258, Nama: "1957356840-765 SYAMSUL TRI ANDIKA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 12", Pekerjaan: "Golang", Alasan: "006"},
	260: {Absen: 260, Nama: "1957356840-771 SALMIA RAHMAWATI Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 37", Pekerjaan: "Golang", Alasan: "006"},
	264: {Absen: 264, Nama: "1957356840-785 THUFEIL MUHAMMAD NUR AZIZ Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 17", Pekerjaan: "Golang", Alasan: "006"},
	262: {Absen: 262, Nama: "1957356840-777 ROBI AURIO UREKA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 75", Pekerjaan: "Golang", Alasan: "006"},
	273: {Absen: 273, Nama: "1957356840-800 MUH FURQAN SUPRIADI Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 98", Pekerjaan: "Golang", Alasan: "006"},
	289: {Absen: 289, Nama: "1957356840-853 BAYU ADITYA TRIWIBOWO Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 92", Pekerjaan: "Golang", Alasan: "006"},
	291: {Absen: 291, Nama: "1957356840-858 ZACKY FATAYAN Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 4", Pekerjaan: "Golang", Alasan: "006"},
	286: {Absen: 286, Nama: "1957356840-839 M.DIO GEOVANI Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 89", Pekerjaan: "Golang", Alasan: "006"},
	298: {Absen: 298, Nama: "1957356840-90 GHIFFARI AHMADIJAYA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 12", Pekerjaan: "Golang", Alasan: "006"},
	300: {Absen: 300, Nama: "1957356840-96 AZIZ SETYA NURROHMAN Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 91", Pekerjaan: "Golang", Alasan: "006"},
	255: {Absen: 255, Nama: "1957356840-758 Talitha Alda Zafirah Dewi Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 1", Pekerjaan: "Golang", Alasan: "006"},
	265: {Absen: 265, Nama: "1957356840-786 ARMAN FADHILLA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 0", Pekerjaan: "Golang", Alasan: "006"},
	277: {Absen: 277, Nama: "1957356840-81 Rizki Syawaldi Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 34", Pekerjaan: "Golang", Alasan: "006"},
	281: {Absen: 281, Nama: "1957356840-825 MOHAMMAD FIRMAN FARDIANSYAH Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 63", Pekerjaan: "Golang", Alasan: "006"},
	271: {Absen: 271, Nama: "1957356840-797 ARDI PERMANA Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 18", Pekerjaan: "Golang", Alasan: "006"},
	275: {Absen: 275, Nama: "1957356840-808 RAHMAT ADLIN PASARIBU Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 91", Pekerjaan: "Golang", Alasan: "006"},
	292: {Absen: 292, Nama: "1957356840-860 ELIZA FARAHDIBA SALEH Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 58", Pekerjaan: "Golang", Alasan: "006"},
	293: {Absen: 293, Nama: "1957356840-861 Rizqy Ghaniyyu Fadhilah Scalable Web Service with Golang (Batch1)", Alamat: "Jalan Kehidupan No. 46", Pekerjaan: "Golang", Alasan: "006"},
}

type FriendRepository interface {
	GetFriendByAbsen(absen int) entities.Friend
}

type friendRepository struct{}

func NewFriendRepository() FriendRepository {
	return &friendRepository{}
}

func (r *friendRepository) GetFriendByAbsen(absen int) entities.Friend {
	return friends[absen]
}
