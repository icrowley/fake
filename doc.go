// Package fake is the fake data generatror for go (Golang), heavily inspired by forgery and ffaker Ruby gems

// Most data and methods are ported from forgery/ffaker Ruby gems.
// Currently english and russian languages are available.
// Feel free to add other languages, but don't forget to regenerate data.go file using `github.com/mjibson/esc` tool and `esc -o data.go -pkg fake data` command because Fake embeds files unless you call `UseExternalData(true)` in order to be able to work file without dependencies when compiled.

// Examples
//  name := fake.FirstName()
//  fullname = := fake.FullName()
//  product := fake.Product()

// Changing language:
//  err := fake.SetLang("ru")
//  if err != nil {
//    panic(err)
//  }
//  password := fake.SimplePassword()

// Using english fallback:
//  err := fake.SetLang("ru")
//  if err != nil {
//    panic(err)
//  }
//  fake.EnFallback(true)
//  password := fake.Paragraph()

// Using external data:
//   fake.UseExternalData(true)
//   password := fake.Paragraph()
package fake
