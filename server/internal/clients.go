
package clients

import (

	"fmt" //good for output in console
	"os" //for accessing the environment variables
	"net/http" //for making https requests
	"io"
	"github.com/joho/godotenv"
)

func GetWeatherData ()(string, error) {
	err := godotenv.Load("../.env");
	if err != nil {
		fmt.Println("error loading .env file")
		return "", err
	}

	//get api key from env to be used later, this is a necesary step
	apiKey := os.Getenv("Golang_API_KEY")
	//if the os is empty, it returns an empty string not nil 
	if apiKey == "" {
		fmt.Println("api key is missing")
		return "", fmt.Errorf("Api Key Is Missing") //what does errorf do?
	}

	resp, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=London&limit=5&appid=" + apiKey)

	if err != nil {
		fmt.Println("Error making HTTP request: ", err)
		return "", err
	}

	defer resp.Body.Close() //makes sure the response body is closed and nothing is leaked


	body, err := io.ReadAll(resp.Body) 
	if err != nil{
		fmt.Println("error reading the response body ", err)
		return "", err
	}

	fmt.Println("Resp status ", resp.Status)
	fmt.Println("Weather data: ", string(body))

	return string(body), nil
}