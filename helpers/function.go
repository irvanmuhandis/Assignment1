package helpers

// Fungsi untuk mendapatkan data teman berdasarkan nomor absen
func GetFriendData(no int) *Friend {

	friends := Getdata()

	// Cek apakah nomor absen valid
	friend, exists := friends[no]
	if !exists {
		return nil
	}

	return &friend
}
