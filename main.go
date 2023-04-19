package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Food struct {
	Id    uint
	Name  string
	Price uint
}

type Order struct {
	Id    uint
	Foods []string
	Summ  uint
}

func main() {

	GetOrders()

}

func WelcomeClient() bool {

	foods := ReadingFileForFoods()

	var resp bool
	var response string

	fmt.Println("\nWelcome to our 'Fast Food' cafe\n")
	fmt.Println("Here, You can see foods we have and order by food's id:\n")

	for _, food := range foods {
		fmt.Println(food.Id, food.Name, food.Price)
	}

	fmt.Println("\nDo you want to order? (y/n)")
	fmt.Scanln(&response)
	if string(response[0]) == "y" {
		resp = true
	}

	return resp

}

func GetOrders() {

	response := WelcomeClient()

	if response {
		orders := ReadingFileForOrders()
		foods := ReadingFileForFoods()

		newOrder := Order{
			Id: uint(len(orders) + 1),
		}

		var id uint
		var numberOfOrders uint
		fmt.Println("\nHow many foods do you want to get:\n")
		fmt.Scanf("%d\n", &numberOfOrders)

		for i := uint(0); i < numberOfOrders; i++ {
			fmt.Println("\nEnter the food's id: ")
			fmt.Scanln(&id)

			for _, food := range foods {
				if id == food.Id {
					newOrder.Foods = append(newOrder.Foods, food.Name)
					newOrder.Summ = newOrder.Summ + food.Price

				}
			}
		}
		
		fmt.Printf("\nYour Id:%v\n", newOrder.Id)
		fmt.Printf("Foods:%v\n", newOrder.Foods)
		fmt.Printf("Total price:%v\n\n", newOrder.Summ)
	
		fmt.Println("\nWhen your order is ready, you can pay and take it.\n")
		fmt.Println("_______________________________________________________________________________________________________________________________________________\n")

		orders = append(orders, newOrder)

		fmt.Println("Admin, do you want to see all orders? (y/n)")
		var adminRes string
		var adminPass string
		password := "admin10125"
		fmt.Scanln(&adminRes)
		if string(adminRes[0]) == "y" {

			fmt.Println("\nEnter Admin's Password(Only one chance):")
			fmt.Scanln(&adminPass)
			fmt.Println("\n")
			if adminPass == password {
				for _, order := range orders {
					fmt.Println(order.Id, order.Foods, order.Summ)
				}
			} else {
				fmt.Println("Wrong password :(")
			}
		}

		defer WritingFile(orders)

	} else {
		fmt.Println("\nWelcome anytime if you want to experience the difference :)\n")
	}

}

func ReadingFileForFoods() []Food {

	var Foods []Food

	data, err := ioutil.ReadFile("foods.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Foods
	}
	err = json.Unmarshal(data, &Foods)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return Foods
	}
	return Foods
}

func ReadingFileForOrders() []Order {

	var Orders []Order

	data, err := ioutil.ReadFile("admin.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Orders
	}
	err = json.Unmarshal(data, &Orders)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return Orders
	}
	return Orders
}

func WritingFile(order []Order) []Order {
	data, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return order
	}

	err = ioutil.WriteFile("admin.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return order
	}
	return order
}
