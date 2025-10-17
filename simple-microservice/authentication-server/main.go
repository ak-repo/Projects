package main

import "os"




var SigningKey = []byte(os.Getenv("SECRET_KEY"))