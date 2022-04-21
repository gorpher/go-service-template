go run entgo.io/ent/cmd/ent init Playground
entimport -dsn "mysql://root:root@tcp(localhost:3306)/tailwindcss" -tables "playgrounds"
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
