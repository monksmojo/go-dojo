package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

/**
env variables are used to store sensitive information that is used by your production based application
they cannot be stored inside the repo
they are mostly stored in the OS (operating system as the variable) and is fetched from there
**/

/*
following commands work in linux Os.
for other OS pls refer the web
*/

// Using OS method
// used when we have less environment variable

// we use `export` variable to store / update environment variables in the the operating system
// we can also update the environment using `export` command
// to print the OS env variable we use `echo` command

//  to remove the environment variables we use `unset` command unset followed by the var_name

//  list the env variables in OS using `printenv`

// using a third party package github.com/joho/godotenv
// godotenv.load()  method  to load env variables from .env file

//  using a third party package called github.com/caarlos0/env
// documentation link github.com/caarlos0/env/v11
//  now when we talk about the docker and kubernetes cluster where we load env variables directly into the os
//  we create a struct having attributes and applying decorates to assign env variables to these
//  make sure you define the attributes with starting upperCase character

type envConfig struct {
	LoginSecret string `env:"LOGIN_SECRET"`
	AtvKey      int    `env:"ATV_KEY"`
	UserKey     string `env:"USER_KEY"`
	AtvPort     int    `env:"ATV_PORT" envDefault:"9090"`
}

// While `required` tag demands the environment variable to be set, it doesn't check its value. If you want to make sure the environment is set and not empty, you need to use the `notEmpty` tag option instead (`env:"SOME_ENV,notEmpty"`).

type secretConfig struct {
	AtvPassword string `env:"ATV_PASSWORD,notEmpty"`
}

// caarlos0/env also has the feature where we can use env variable having same prefix to be loaded generic struct of environment variable  attributes
// in caarlos0/env it is under ComplexSlices
type devEnvConfig struct {
	DevEnvConfigs envConfig `envPrefix:"DEV_"`
}

func main() {
	// directly using the OS package
	// in go we can use default os package to set,get and unset the env variables
	os.Setenv("DB_HOST", "/localhost") // sets the environment
	os.Setenv("DB_URL", "8090")        // sets the environment
	fmt.Println(os.Getenv("DB_HOST"))  // gets the environment
	fmt.Println(os.Getenv("DB_HOST"))  // gets the environment
	os.Unsetenv("DB_HOST")             // unset the environment
	os.Unsetenv("DB_URL")              // unset the environment

	// using a third party package github.com/joho/godotenv
	// godotenv.load()  method  to load env variables from .env file
	godotenv.Load(".env")
	fmt.Println(os.Getenv("HOST"))        // gets HOST variable from the  env file
	fmt.Println(os.Getenv("PORT"))        // gets PORT from the env file
	fmt.Println(os.Getenv("TOPIC_NAME"))  // gets TOPIC_NAME from the env file
	fmt.Println(os.Getenv("MODULE_NAME")) // gets MODULE_NAME from the env file

	// using a third party package called github.com/caarlos0/env
	os.Setenv("LOGIN_SECRET", "antavo_gogin_908796754")
	os.Setenv("ATV_KEY", "12345678")
	os.Setenv("USER_KEY", "peter_parker_908796754")
	os.Setenv("HOME", "peter_parker_908796754")

	//  picking the special prefix env variables using the same struct
	os.Setenv("DEV_LOGIN_SECRET", "dev_antavo_gogin_908796754")
	os.Setenv("DEV_ATV_KEY", "12345678")
	os.Setenv("DEV_USER_KEY", "dev_peter_parker_908796754")
	os.Setenv("DEV_HOME", "dev_peter_parker_908796754")

	var envConfig envConfig
	err := env.Parse(&envConfig)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("loginSecret: ", envConfig.LoginSecret, reflect.TypeOf(envConfig.LoginSecret))
	fmt.Println("atvKey: ", envConfig.AtvKey, reflect.TypeOf(envConfig.AtvKey))
	fmt.Println("atvPort: ", envConfig.AtvPort, reflect.TypeOf(envConfig.AtvPort))
	fmt.Println("userKey: ", envConfig.UserKey, reflect.TypeOf(envConfig.UserKey))

	var devEnvConfig devEnvConfig
	err = env.Parse(&devEnvConfig)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(devEnvConfig)

	var secretConfig secretConfig
	err = env.Parse(&secretConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("atvPassword", secretConfig.AtvPassword, reflect.TypeOf(secretConfig.AtvPassword))

	os.Unsetenv("LOGIN_SECRET")
	os.Unsetenv("ATV_KEY")
	os.Unsetenv("USER_KEY")
	os.Unsetenv("HOME")
	os.Unsetenv("DEV_LOGIN_SECRET")
	os.Unsetenv("DEV_ATV_KEY")
	os.Unsetenv("DEV_USER_KEY")
	os.Unsetenv("DEV_HOME")

}
