package helper

import (
	"web-desa/config"
	"web-desa/model"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes),err
}

func SeederRefresh(cfg config.Config) {

	cfg.Database().Migrator().DropTable(&model.User{})
	cfg.Database().Migrator().DropTable(&model.Desa{})
	cfg.Database().Migrator().DropTable(&model.InfoKegiatan{})
	cfg.Database().Migrator().DropTable(&model.Umkm{})
	cfg.Database().Migrator().DropTable(&model.Wisata{})

	hashAdmin1, _ := HashPassword("Admin1**.")
	hashAdmin2, _ := HashPassword("Admin2**.")
	
	users := []model.User{
		{Username: "Admin1", Password: hashAdmin1},
		{Username: "Admin2", Password: hashAdmin2},
	}

	desa := model.Desa{
		ID: 1,
		TentangDesa: "Dilem merupakan sebuah desa yang terletak di wilayah Kecamatan Kepanjen, Kabupaten Malang. Jarak Desa Dilem ke ibu kota kecamatan Kepanjen dapat ditempuh dengan waktu sekitar 7 menit. Secara topopografis, Desa Dilem berada di ketinggian 2000 meter di atas permukaan air laut dengan curah hujan 29 mm/th dan suhu rata-rata 20-35 Celcius. Dengan luas wilayah sekitar 216.133 Ha, Dilem terbagi ke dalam empat hak guna lahan, diantarannya permukiman, pertanian, perikanan, dan tanah bersumber air. Perbatasan Desa Dilem berbatasan langsung dengan beberapa desa dan wilayah sekitarnya. Di sebelah barat, desa ini berbatasan dengan Desa Talangagung dan desa Ngadilangkung. Di sebelah utara, terdapat desa Ngadilangkung dan di sebelah selatan, desa ini berbatasan dengan kelurahan kepanjen. Sedangkan di sebelah timur terdapat Desa Ngadilangkung dan Keluarahan Ardirejo.",
		Visi: "Terwujudnya Desa Dilem yang mandiri, religius, sejahtera dan berkeadilan.",
		Misi: 	"1. Mewujudkan dan meningkatkan tata kelola pemerintahan desa dengan baik "+
				"2. Meningkatkan sarana dan prasarana desa "+
				"3. Meningkatkan kesejahteraan masyarakat melalui program pemerintah dengan pemberdayaan masyarakat"+
				"4. Meningkatkan kehidupan yang harmonis, dan toleran dalam berbudaya dan beragama, mewujudkan keamanan dan kenyamanan serta ketertiban Desa Dilem",
	}

	infoKegiatans := []model.InfoKegiatan{
		{Judul: "Pelatihan Teknis Pemeliharaan Ikan di Kolam Pemancingan Bapak Jumain oleh Mahasiswa MMD", Gambar: "https://qbekdjxviuehumdvbstm.supabase.co/storage/v1/object/public/api-service-desa/informasi-kegiatan/IMG_3846.jpg", Tanggal: "14-07-2023", Deskripsi: "Pelatihan ini berupa sosialiasi dan praktik pembuatan pakan ikan dengan campuran probiotik EM4 dan molase dalam 1 liter air yang dicampurkan pada 5 kg pakan pelet ikan. Setelah itu pakan tersebut ditebar ke kolam-kolam ikan. Pelatihan ini dilakukan oleh Mahasiswa Universitas Brawijaya saat pelaksanaan kegiatan Mahasiswa Membangun Desa (MMD)."},
	}

	umkms := []model.Umkm{
		{Nama: "Lufas Gallery", Alamat: "Jl. Sidoluhur No.15, Lemah Duwur, Dilem, Kec. Kepanjen, Kabupaten Malang, Jawa Timur 65163", Kontak: "082132679994", Gambar: "", Deskripsi: "Sejak didirikan pada tahun 1996, Lufas gallery terus berinovasi dalam mengembangkakn produk-produk unggulan agar tetap selaras dengan perkembangan zaman dan berhasil menarik minat masyarakat. Produk Lufas Gallery selalu menampilkan desain khas yang menarik perhatian para konsumen. Lufas gallery terus berusaha meningkatkan kualitas produknya untuk memastikan kepuasan para pengguna kerajinan kulit. "},
		{Nama: "Trios Collection", Alamat: "Jl.Sidoluhur No.26 RT:06/RW:01, Dilem,Kepanjen - Malang 65163", Kontak: "081359702244", Gambar: "", Deskripsi: "Trios Collection menjadi lambang pusat kerajinan kulit tradisional di Desa Dilem selama dua dekade terakhir. Produk dari Trios Collection menggunakan bahan kulit sapi dan domba dengan kualitas tinggi. Produk Trios Collection tidak hanya tersebar di Desa Dilem saja namun, berhasil menjangkau wilayah Tulungagung, Surabaya, Jakarta, dan bahkan luar Jawa. "},
		{Nama: "syami jaya ", Alamat: "Jl. Sidoluhur, Lemah Duwur, Dilem, Kec. Kepanjen, Kabupaten Malang, Jawa Timur 65163", Kontak: "08885524419", Gambar: "", Deskripsi: "Syami jaya adalah sebuah usaha yang mengkomersialkan produk olahan kacang, yakni kacang oven. Sudah beroperasi sejak 2010, syami jaya telah berhasil menarik minat pelanggan berkat strategi pemasaran baik offline maupun secara online. Selain itu, kacang oven juga menerima pesanan secara offline."},
		{Nama: "Sang Kriyuk", Alamat: "Jl. Sidoluhur No.94, Lemah Duwur, Dilem, Kec. Kepanjen, Kabupaten Malang, Jawa Timur 65163", Kontak: "087809606403", Gambar: "", Deskripsi: "Sangkriyuk merupakan bisnis yang menjual produk olahan pisang. Dimulai sejak tahun 2016, Sangkriyuk berhasil meraih pelanggan di berbagai penjuru wilayah melalui pemasaran online yang efektif. Tidak hanya mengandalkan pemasaran online, Sangkriyuk juga berhasil memperluas jangkauan konsumen melalui jaringan reseller yang kuat. "},
	}

	wisatas := []model.Wisata{
		{Nama: "Eko Wisata Lembah Dilem", HargaTiket: 3000, Alamat: "Lemah Duwur, Dilem, Kec. Kepanjen, Kabupaten Malang, Jawa Timur 65163", Kontak: "08113259311", Gambar: "https://qbekdjxviuehumdvbstm.supabase.co/storage/v1/object/public/api-service-desa/wisata/IMG_1153.jpg", JamBuka: "09:00", JamTutup: "16:00", Deskripsi: "Desa Dilem yang berlokasi di Kecamatan Kepanjen, Kabupaten Malang, merupakan salah satu desa yang mengembangkan objek wisata untuk meningkatkan kesejahteraan masyarakat-nya. Salah satu upaya untuk mengembangkan objek wisata adalah pembangunan Eko Wisata Lembah Dilem. Lembah Dilem yang dibangun oleh BUMDes Dilem Makmur ini menyuguhkan panorama alam dan potensi sumber mata air."},
		{Nama: "Dilem Edu Park", HargaTiket: 0, Alamat: "Jl. Semeru RT. 02 RW. 04, Lemah Duwur, Dilem, Kec. Kepanjen, Kabupaten Malang, Jawa Timur 65163", Kontak: "081333234025", Gambar: "https://qbekdjxviuehumdvbstm.supabase.co/storage/v1/object/public/api-service-desa/wisata/IMG_4606.jpg", JamBuka: "08:00", JamTutup: "22:00", Deskripsi: "Dilem Edu Park merupakan salah satu daya tarik wisata Desa Dilem.Dilem Edu Park didirikan pada tanggal 02 November 2021 yang berlokasi di Jl. Semeru RT. 02 RW 04, Desa Dilem, Kecamatan Kepanjen, Kabupaten Malang. Dilem Edu Park merupakan tempat wisata alam yang didukung oleh fasilitas untuk pengunjung, seperti Wi-Fi, kolam pemancingan ikan lele, dan gazebo. "},
	}

	cfg.Database().Create(&desa)

	for _, user := range users {
		cfg.Database().Create(&user)
	}

	for _, infoKegiatan := range infoKegiatans {
		cfg.Database().Create(&infoKegiatan)
	}

	for _, umkm := range umkms {
		cfg.Database().Create(&umkm)
	}

	for _, wisata := range wisatas {
		cfg.Database().Create(&wisata)
	}

}