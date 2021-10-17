# Source Code for [nicc.io](https://nicc.io) <3

I'm a programmer, so of course I automated this :)

`buildPortfolio/buildPortfolio.go` is a program I wrote that allows me to quickly update my portfolio or template, without having to do the tedious work that is writing repetitive HTML.

### Here's how it works!
1. Reads project, article, and footer data from `buildPortfolio/data.csv`
2. Inserts data into `buildPortfolio/template.tmpl`
3. Outputs completed template to `index.html`

### Flags

- __csv__ _string_
    - A CSV file containing data to be inserted into the template. (_default "data.csv"_)
- __output__ _string_
	- The .html output file. (_default "../index.html"_)
- __tmpl__ _string_
 	- A .tmpl file to insert data into. (_default "template.tmpl"_)

Flags are optional and for debugging mostly. The default flags are set to work for this project exactly.
```bash
go run buildPortfolio.go
```

If no extensions are given, it will fill them in for you. Or you can use some random extension, it doesn't really matter.
```bash
go run buildPortfolio.go -csv="../../far/away/csvfile" -output="/cool_website.lol" -tmpl="some/nested/templatefile"
```


### CSV File Format

`TYPE,TITLE,DESC,URL`
- `TYPE`: Has to be either P, A, or F
    - P for Project
    - A for Article
    - F for Footer
- `TITLE`: A title (For a footer item, this field should be the fontawesome classes for an icon.)
- `DESC`: A short description (For a footer item, this field should be the text  next to the icon.)
- `URL`: A URL
