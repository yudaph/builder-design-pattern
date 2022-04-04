package main

import (
	"builder/builder"
	"builder/builder2"
	"fmt"
)

func main() {
	selectField := []string{"id", "name"}
	sql := builder.NewSelectBuilder(selectField, "person").Where("id", builder.Equals, 101).Build()
	fmt.Println(sql) // SELECT id,name FROM person WHERE id = 101
	builderSQL := builder.NewSelectBuilder(selectField, "person")
	builderSQL.Where("nationality", builder.Equals, "japan")
	builderSQL.AddSelectedField("nationality", "age")
	builderSQL.Where("age", builder.GreaterThanEquals, 20)
	builderSQL.Order("age", builder.Asc)
	builderSQL.Where("marriage_status", builder.Equals, false)
	fmt.Println(builderSQL.Build())
	// SELECT id,name,nationality,age FROM person WHERE nationality = "japan" AND age >= 20 AND marriage_status = false ORDER BY age asc

	personBuilder := builder2.NewPersonBuilder("anto", 20)
	personBuilder.
		Address().
		At("Dragon Street").
		City("Dragon City").
		PostalCode("123123").
		Collage().
		At("Oxford University").
		Major("Computer Science")
	fmt.Println(personBuilder.Build()) // &{anto 20 Dragon Street Dragon City 123123 Oxford University Computer Science}

}
