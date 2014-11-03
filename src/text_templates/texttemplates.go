package main 

import (
        "text/template"
        "os"
        "log"
        "strings"
)

func main() {
	
		//Exmaple of text template
        type Inventory struct {
        Material string
        Count    uint
        }
        sweaters := Inventory{"wool", 17}
        tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}") //Creating the template
        if err != nil { panic(err) }
        err = tmpl.Execute(os.Stdout, sweaters) //Using the template
        if err != nil { panic(err) }
        
        //Example 1 - Simple Template
       // Define a template.
        const letter = `
        Dear {{.Name}},
        {{if .Attended}}
        It was a pleasure to see you at the wedding.{{else}}
        It is a shame you couldn't make it to the wedding.{{end}}
        {{with .Gift}}Thank you for the lovely {{.}}. {{/* If the value of the pipeline is empty, no output is generated; otherwise, dot is set to the value of the pipeline and T1 is 	executed.*/}}
        {{end}}
        Best wishes,
        Josie
        `

        // Prepare some data to insert into the template.
        type Recipient struct {
                Name, Gift string
                Attended   bool
        }
        var recipients = []Recipient{
                {"Aunt Mildred", "bone china tea set", true},
                {"Uncle John", "moleskin pants", false},
                {"Cousin Rodney", "", false},
        }

        // Create a new template and parse the letter into it.
        t := template.Must(template.New("letter").Parse(letter)) //Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations

        // Execute the template for each recipient.
        for _, r := range recipients {
                err := t.Execute(os.Stdout, r)
                if err != nil {
                        log.Println("executing template:", err)
                }
        }
       
        
        //Example 2 - Template with a function
        // First we create a FuncMap with which to register the function.
        funcMap := template.FuncMap{
                // The name "title" is what the function will be called in the template text.
                "title": strings.Title,
        }

        // A simple template definition to test our function.
        // We print the input text several ways:
        // - the original
        // - title-cased
        // - title-cased and then printed with %q
        // - printed with %q and then title-cased.
        // It looks like the . represents the text itself that is used in the .Execute function
        const templateText = `
        Input: {{printf "%q" .}}
        Output 0: {{title .}}
        Output 1: {{title . | printf "%q"}}
        Output 2: {{printf "%q" . | title}}
        `

        // Create a template, add the function map, and parse the text.
        tmpl, err2 := template.New("titleTest").Funcs(funcMap).Parse(templateText)
        if err2 != nil {
                log.Fatalf("parsing: %s", err2)
        }

        // Run the template to verify the output.
        err2 = tmpl.Execute(os.Stdout, "the go programming language")
        if err2 != nil {
                log.Fatalf("execution: %s", err2)
        } 
        
}

