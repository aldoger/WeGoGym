# WeGoGym

Final Project untuk mata kuliah Konstruksi dan Arsitektur Perangkat Lunak.

## Tim Pengembang 👥

* **Made Satya Dhananjaya** - NRP : 5053231001
* **Nuhito Atmoko** - NRP: 5053231002
* **Nur Muhammad Faiz** - NRP: 5053231008
* **Geraldo Benjamin Nainggolan** - NRP: 5053231014

**Dosen Pembimbing:** Rizky Januar Akbar, S.Kom., M.Eng

---

## Ringkasan Proyek 🏋️‍♀️

WeGoGym adalah sistem manajemen gym komprehensif yang dikembangkan menggunakan bahasa pemrograman Go. Aplikasi ini dirancang untuk menyederhanakan berbagai aspek operasional gym, mulai dari manajemen pengguna dan keanggotaan hingga pemrosesan transaksi dan pencatatan riwayat masuk.

## Fitur ✨

* **Manajemen Pengguna:** Fasilitas lengkap untuk pendaftaran, login, logout, dan akses profil pengguna.
* **Manajemen Keanggotaan yang Efisien:**
    * **Admin-only:** Kelola keanggotaan gym dengan kemampuan membuat, melihat, dan memperbarui detail keanggotaan.
    * **Pengguna:** Proses berlangganan keanggotaan dengan mudah.
    * **Cek Status Keanggotaan:** Verifikasi status keanggotaan pengguna secara instan.
* **Manajemen Sesi Personal Trainer (PT):**
    * Tambahkan dan kelola sesi personal trainer untuk pengguna.
* **Integrasi Pembayaran Midtrans:** Memproses transaksi keanggotaan dan sesi PT dengan mulus melalui Midtrans.
* **Riwayat Masuk Gym:** Catat dan akses riwayat waktu masuk gym setiap pengguna.
* **Generasi Kode QR:** Hasilkan kode QR unik untuk identifikasi pengguna yang cepat.
* **Kontrol Akses Berbasis Peran (RBAC):** Memastikan fungsionalitas yang tepat diakses oleh peran `member` dan `admin`.

---

## Teknologi yang Digunakan 💻

* **Go (Golang):** Bahasa pemrograman utama yang digunakan untuk backend.
* **Gin Web Framework:** Digunakan untuk membangun API RESTful yang cepat dan efisien.
* **GORM:** ORM (Object-Relational Mapping) yang mempermudah interaksi dengan database.
* **PostgreSQL:** Database relasional yang kuat dan andal.
* **Midtrans:** Platform payment gateway untuk memproses transaksi.
* **Logrus:** Pustaka logging yang fleksibel dan terstruktur.
* **go-qrcode:** Pustaka untuk menghasilkan kode QR.
* **godotenv:** Memuat variabel lingkungan dari file `.env`.
* **stretchr/testify:** Pustaka Go untuk pengujian unit, termasuk mock.
