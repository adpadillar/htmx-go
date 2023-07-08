# htmx-go

## Development

By using htmx, we only really need to run our go server and go to `localhost:8080` to see our changes.

To run the server, run:

```bash
go run main.go
```

This however, does not watch for changes. To watch for changes, we can use [gin](https://github.com/codegangsta/gin) with a command like this:

```bash
gin -p 8080 -a 8081 run main.go
```

We also need to generate our styles with tailwind. To do this, run:

```bash
./tailwindcss -i src/styles/globals.css -o src/public/tailwind.css
```
