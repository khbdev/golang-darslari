package main

import (
	"errors"
	"fmt"
	"time"
)

type ValidaionError struct
{
	Field string
	Msg string
}

func (v *ValidaionError) Error() string {
    return fmt.Sprintf("Validatsiya xatosi (%s): %s", v.Field, v.Msg)
}


func validateAge(age int) error{
	if age < 0 || age > 125 {
		return &ValidaionError{
			Field: "age",
			Msg: "yosh natogri",
		}
	}
	return nil
}

func divide(a, b  int) (int, error){
	if b ==0  {
		return 0, fmt.Errorf("son 0 bolish mumkun emas a=%d, b=%d", a, b)
	}

	return a / b, nil
}


func great(name string) error{
	if name == "" {
		return errors.New("Isim kiritlmagan")
	} 

		fmt.Println("Salom", name)
		return nil
}

func main(){

	
var surname string



fmt.Print("ismingizni kiriting: ")
fmt.Scanln(&surname)

err := great(surname)

if err != nil {
	fmt.Println("Xato", err)
}

result, ss := divide(10, 5)

if ss != nil {
	fmt.Println("Xato: ", ss)
} else {
	fmt.Println("Natija: ", result)
}


errw := validateAge(100)

  if errw != nil {
        fmt.Println("Xato:", errw)
    } else {
        fmt.Println("Yosh to‘g‘ri")
    }

now := time.Now()

fmt.Println(now.Format("2008-07-01")) // 👉 2025-07-01
fmt.Println(now.Format("15:04:05"))


fmt.Println("Boshladi")
time.Sleep(2 * time.Second)
fmt.Println("2 soniyadan so‘ng...")

}