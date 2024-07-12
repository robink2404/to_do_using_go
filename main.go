package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"log"
)

type to_dolist struct {
	wk_id       int
	work_name   string
	work_status bool
}

func add_to_list(work_list *[]to_dolist) {
	var wk_id int
	var wk_nm string
	var work_status bool
	var x int
	fmt.Println("Enter Your wk_id: ")
	fmt.Scan(&wk_id)
	fmt.Println("Enter Your Work_Name: ")
	fmt.Scan(&wk_nm)
	fmt.Println("Is Your work done?")
	fmt.Println("1: True")
	fmt.Println("2: False")
	fmt.Scan(&x)
	if x == 1 {
		work_status = true
	} else if x == 2 {
		work_status = false
	} else {
		work_status = true
		fmt.Println("Please Choose Correct Option Next Time: ")
		return
	}
	var new_item = to_dolist{
		wk_id:       wk_id,
		work_name:   wk_nm,
		work_status: work_status,
	}

	*work_list = append(*work_list, new_item)
	fmt.Println("So, Item added to Your list successfully")

}

func search_in_list(work_list *[]to_dolist, x int) int {

	for i, to_do := range *work_list {
		if to_do.wk_id == x {
			return i
		}
	}
	return -1
}
func update_ur_list(work_list *[]to_dolist) {
	var x int
	fmt.Print("Enter Your work_id you want to update: ")
	fmt.Scan(&x)
	var flag int = search_in_list(work_list, x)
	if flag != -1 {
		fmt.Println("So what u want to update, choose one option: ")
		fmt.Println("1:wk_id")
		fmt.Println("2:wk_name")
		fmt.Println("3:Work_statue")
		var y int
		fmt.Scan(&y)
		if y == 1 {
			for _, to_do := range *work_list {
				if to_do.wk_id == x {
					fmt.Println("Enter new wk_id: ")
					var z int
					fmt.Scan(&z)
					to_do.wk_id = z
					fmt.Println("Your wk_id Updated succesfully")
					return
				}
			}
		} else if y == 2 {
			for _, to_do := range *work_list {
				if to_do.wk_id == x {
					fmt.Println("Enter new work_name: ")
					var z string
					fmt.Scan(&z)
					to_do.work_name = z
					fmt.Println("Your wk_name Updated succesfully")
					return
				}
			}

		} else if y == 3 {
			for _, to_do := range *work_list {
				if to_do.wk_id == x {
					to_do.work_status = !to_do.work_status
					fmt.Println("Your wORK status Updated successfully")
					return
				}
			}

		} else {
			fmt.Println("Please Choose correct Option next time")
			return
		}

	} else {
		fmt.Println("Sorry, You Gave either wrong wk_id or that wk_id deleted")
		return
	}

}
func delete_item_from_ur_list(work_list *[]to_dolist) {
	var x int
	fmt.Println("Give the wk_id you want to delete")
	fmt.Scan(&x)

	var flag int = search_in_list(work_list, x)
	if flag != -1 {
		*work_list = append((*work_list)[:flag], (*work_list)[flag+1:]...)
		fmt.Println("so WK_ID %v DELETED successfully ", x)

	} else {
		fmt.Println("Sorry the wk_id You gave is not found in list")
	}
}
func show_list(work_list *[]to_dolist) {
	fmt.Printf("Work_id  Work_Name  Work_done?  \n")

	for _, to_do := range *work_list {
		fmt.Printf("%v  %v  %v  \n", to_do.wk_id, to_do.work_name, to_do.work_status)

	}
}

func main() {
	var work_list = make([]to_dolist, 0)
	fmt.Println("Welcome to To_list Application")

	for {

		fmt.Println("What You want to do, choose from list: ")
		fmt.Println("1:Add to Your List")
		fmt.Println("2:Update to Your List")
		fmt.Println("3:Delete to Your List")
		fmt.Println("4:Show the List")
		fmt.Println("5: Exit")
		var x int
		fmt.Println("Choose Your Option")
		fmt.Scan(&x)
		if x == 1 {
			add_to_list(&work_list)
			fmt.Println("Your work added successfully")
		} else if x == 2 {
			update_ur_list(&work_list)
		} else if x == 3 {
			delete_item_from_ur_list(&work_list)
		} else if x == 4 {
			show_list(&work_list)
		} else {
			fmt.Println("So, Thanku You existed successfully")
			break
		}

	}
	f, err := os.Create("users.csv")
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	header := []string{"Work_id", "Work_Name", "Work_done"}
	if err := w.Write(header); err != nil {
		log.Fatalln("error writing header to CSV:", err)
	}


	for _, record := range work_list {
		stringRecord := []string{
			fmt.Sprint(record.wk_id),
			record.work_name,
			fmt.Sprint(record.work_status),
		}
		if err := w.Write(stringRecord); err != nil {
			log.Fatalln("failed to write record to file", err)
		}
	}

}
