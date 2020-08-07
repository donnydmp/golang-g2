package main

import (
  "fmt"
  "strconv"
  "time"
  "math/rand"
)
var timePark = make([]int64, 0)  //untuk mwnyimpan data integer pada variable timepark
var park = make([]string, 0)	// untukk menyimpan data string pada variabel park

func main()  {
  var menu string				

  for menu != "4"{
    makeMenu(
      "Parkir Masuk",
      "Parkir Keluar",
      "Menampilkan ID kendaran yang parkir",
      "EXIT",
    )
    fmt.Scan(&menu)
    switch menu {
    case "1":
      fmt.Println("ID Parkir anda : ",parkirMasuk())
    case "2":
      var tipe string
      var plat string
      var id_parkir string
      fmt.Print("Masukan Tipe kendaraan : ")
      fmt.Scan(&tipe)
      fmt.Print("Masukan plat kendaraan : ")
      fmt.Scan(&plat)
      fmt.Print("Masukan ID parkir kendaraan : ")
      fmt.Scan(&id_parkir)
      fmt.Print(servParkOut(tipe, plat, id_parkir))
    case "3":
      fmt.Println("ID kendaraan yang terparkir : ")
      fmt.Println(park)
    case "4":
      fmt.Println("====  Terimakasih  ====")
    default:
      fmt.Println("Opsi yang anda masukan salah")
    }
  }
}
func makeMenu(menu ...string)  {
  fmt.Println("")
  fmt.Println("===== Menu =====")
  for menu,val := range menu{
    a:= strconv.Itoa(menu + 1)
    fmt.Println(a + ". " + val)
  }
  fmt.Print("Silahkan Pilih : ")
}

func parkirMasuk() string {
  lenBef := len(park)
  timePark = append(timePark, time.Now().Unix())
  park = append(park, strconv.Itoa(rand.Intn(10000)))
  return park[lenBef]
}
func servParkOut(tipe, plat, id_parkir string) string {
  parkir := false
  var timeParkir int64

  for menu, val := range park{
    if id_parkir == val{
      parkir = true
      timeParkir = timePark[menu]
      RemoveIndexInt64(&timePark, menu)
      RemoveIndexStr(&park, menu)
      break
    }
  }
  if !parkir {
    return "ID Parkir tidak ditemukan"
  }
  total := 0
  currentTime := time.Now().Unix() - timeParkir
  switch tipe {
  case "mobil":
    total = 5000 + ((int(currentTime)) - 1)*3000
    
  case "motor":
    total = 3000 + ((int(currentTime)) - 1)*2000
    
  default:
    return "Tipe kendaraan yang anda masukan salah"
  }
  return "Waktu parkir anda : " + strconv.Itoa(int(currentTime)) + " detik, Jumlah uang yang harus anda bayar adalah : " + strconv.Itoa(total)
}

func RemoveIndexInt64(arr *[]int64, index int) {
    newArr := *arr
    *arr = append(newArr[:index], newArr[index+1:]...)
}
func RemoveIndexStr(arr *[]string, index int) {
  newArr := *arr
  *arr = append(newArr[:index], newArr[index+1:]...)
}

 

