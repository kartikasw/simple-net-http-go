package main

import "strings"

type User struct {
	ID          int
	Email       string
	Address     string
	Job         string
	ReasonHappy string
}

var users = []User{
	{
		ID:          1,
		Email:       generateEmail("John Doe"),
		Address:     "123 Main St, City A",
		Job:         "Software Engineer",
		ReasonHappy: "Building innovative software",
	},
	{
		ID:          2,
		Email:       generateEmail("Jane Smith"),
		Address:     "456 Elm St, City B",
		Job:         "Frontend Developer",
		ReasonHappy: "Creating user-friendly interfaces",
	},
	{
		ID:          3,
		Email:       generateEmail("Michael Johnson"),
		Address:     "789 Oak St, City C",
		Job:         "Backend Developer",
		ReasonHappy: "Optimizing server-side performance",
	},
	{
		ID:          4,
		Email:       generateEmail("Emily Brown"),
		Address:     "101 Pine St, City D",
		Job:         "DevOps Engineer",
		ReasonHappy: "Automating deployment processes",
	},
	{
		ID:          5,
		Email:       generateEmail("David Wilson"),
		Address:     "202 Cedar St, City E",
		Job:         "Software Architect",
		ReasonHappy: "Designing scalable systems",
	},
	{
		ID:          6,
		Email:       generateEmail("Sarah Martinez"),
		Address:     "303 Maple St, City F",
		Job:         "QA Engineer",
		ReasonHappy: "Ensuring software quality",
	},
	{
		ID:          7,
		Email:       generateEmail("Christopher Taylor"),
		Address:     "404 Birch St, City G",
		Job:         "Full Stack Developer",
		ReasonHappy: "Working on end-to-end solutions",
	},
}

func generateEmail(name string) string {
	name = strings.ToLower(strings.ReplaceAll(name, " ", ""))

	email := name + "@example.com"
	return email
}
