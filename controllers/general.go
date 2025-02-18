package controllers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func akses1() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 1")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses2() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 2")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses3() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 3")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses4() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 4")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses5() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 5")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses6() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 6")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses7() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 7")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses8() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 8")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses9() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }

		fmt.Printf("Akses 9")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses10() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 10")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses11() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 11")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses12() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 12")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses13() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 13")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses14() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 14")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses15() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 15")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses16() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 16")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses17() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 17")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses18() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 18")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses19() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 19")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func akses20() {
	url := "https://goldentsunami.id"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Permintaan ke-%d gagal: %v\n", i+1, err)
			continue
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Printf("Membaca respons ke-%d gagal: %v\n", i+1, err)
		// 	continue
		// }
		fmt.Printf("Akses 20")
		fmt.Printf("Permintaan ke-%d berhasil dengan status: %s\n", i+1, resp.Status)
		// fmt.Printf("Respons ke-%d:\n%s\n", i+1, body)
	}
}

func TestAkses(c *gin.Context) {
	var wg sync.WaitGroup

	// Menjalankan setiap fungsi dalam goroutine
	functions := []func(){
		akses1, akses2, akses3, akses4, akses5,
		akses6, akses7, akses8, akses9, akses10,
		akses11, akses12, akses13, akses14, akses15,
		akses16, akses17, akses18, akses19, akses20,
	}

	for _, fn := range functions {
		wg.Add(1) // Menambahkan counter
		go func(f func()) {
			defer wg.Done() // Mengurangi counter saat selesai
			f()             // Menjalankan fungsi
		}(fn)
	}

	wg.Wait() // Menunggu semua goroutine selesai
	c.JSON(200, gin.H{"message": "Semua akses selesai!"})
}
