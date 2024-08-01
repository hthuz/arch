package grades

import (
	"errors"
)

type Student struct {
	ID     int
	Name   string
	Grades []Grade
}

type Students []Student

var students Students

type Grade struct {
	Course string
	Credit int
	Score  int
}

func (ss Students) GetStudentByID(ID int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == ID {
			return &ss[i], nil
		}
	}
	return nil, errors.New("student not found")
}

func (s Student) GetGPA() float64 {
	totalCredit := 0.0
	totalGradePoint := 0.0
	for _, g := range s.Grades {
		totalCredit += float64(g.Credit)
		totalGradePoint += g.getGradePoint() * float64(g.Credit)
	}
	return totalGradePoint / totalCredit
}

func (g Grade) getGradePoint() float64 {
	grade := g.Score
	if grade >= 94 {
		return 4.0
	}
	if grade >= 90 {
		return 3.67
	}
	if grade >= 87 {
		return 3.33
	}
	if grade >= 83 {
		return 3.00
	}
	if grade >= 80 {
		return 2.67
	}
	if grade >= 77 {
		return 2.33
	}
	if grade >= 73 {
		return 2.00
	}
	if grade >= 70 {
		return 1.67
	}
	if grade >= 67 {
		return 1.33
	}
	if grade >= 63 {
		return 1.00
	}
	if grade >= 60 {
		return 0.67
	}
	return 0.33
}

func init() {
	students = Students{
		{
			ID:   1,
			Name: "Steve",
			Grades: []Grade{
				{
					Course: "MATH",
					Credit: 4,
					Score:  97,
				},
				{
					Course: "ENGL",
					Credit: 3,
					Score:  85,
				},
			},
		},
		{
			ID:   2,
			Name: "Mark",
			Grades: []Grade{
				{
					Course: "PHIL",
					Credit: 3,
					Score:  98,
				},
				{
					Course: "PSYC",
					Credit: 4,
					Score:  80,
				},
			},
		},
		{
			ID:   3,
			Name: "Amy",
			Grades: []Grade{
				{
					Course: "CHEM",
					Credit: 3,
					Score:  87,
				},
				{
					Course: "PHYS",
					Credit: 4,
					Score:  85,
				},
			},
		},
	}

}
