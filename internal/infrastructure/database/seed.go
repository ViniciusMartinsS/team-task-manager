package database

import "github.com/ViniciusMartinsS/manager/internal/domain/model"

var users = []model.User{
	{
		Name:     "Jane",
		Surname:  "Doe",
		Email:    "jane.doe@example.com",
		Password: "$2a$12$bvFmmw3X4ctLBf39rF.DqOtp98WPFhm/tztUAXvWpLEgtCpohWzQW", // teste
		Age:      26,
		RoleID:   1,
	},
	{
		Name:     "John",
		Surname:  "Doe",
		Email:    "john.doe@example.com",
		Password: "$2a$12$bvFmmw3X4ctLBf39rF.DqOtp98WPFhm/tztUAXvWpLEgtCpohWzQW", // teste
		Age:      24,
		RoleID:   2,
	},
	{
		Name:     "Julius",
		Surname:  "Rock",
		Email:    "julius.rock@example.com",
		Password: "$2a$12$bvFmmw3X4ctLBf39rF.DqOtp98WPFhm/tztUAXvWpLEgtCpohWzQW", // teste
		Age:      37,
		RoleID:   2,
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
		UserId:    2,
		Performed: nil,
	},
	{
		Name:      "Task Hello World",
		Summary:   "4ea5815b0d1d0d28f446d123f8e751cbca74c5c63a3cde51375c0fad8c947ff5fabc4885656cba0399d86155208d375844c7d3811bf544f99bc5335a",
		UserId:    3,
		Performed: nil,
	},
}
