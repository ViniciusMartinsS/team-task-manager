package database

import "github.com/ViniciusMartinsS/manager/internal/domain/model"

var users = []model.User{
	{
		Name:     "Vinicius",
		Surname:  "Simone Martins",
		Email:    "visimonemartins@hotmail.com",
		Password: "$2a$12$bvFmmw3X4ctLBf39rF.DqOtp98WPFhm/tztUAXvWpLEgtCpohWzQW", // teste
		Age:      24,
		RoleID:   2,
	},
	{
		Name:     "Fernada",
		Surname:  "Simone Martins",
		Email:    "fesimonemartins@hotmail.com",
		Password: "$2a$12$bvFmmw3X4ctLBf39rF.DqOtp98WPFhm/tztUAXvWpLEgtCpohWzQW", // teste
		Age:      26,
		RoleID:   1,
	},
}

var roles = []model.Role{
	{
		Name: "Manager",
	},
	{
		Name: "Technician",
	},
}

var tasks = []model.Task{
	{
		Name:      "Task Hello World",
		Summary:   "4ea5815b0d1d0d28f446d123f8e751cbca74c5c63a3cde51375c0fad8c947ff5fabc4885656cba0399d86155208d375844c7d3811bf544f99bc5335a",
		UserId:    1,
		Performed: nil,
	},
}
