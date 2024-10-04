package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



type templateData struct{
  Data map[string]any
}


func(app *application) render(w http.ResponseWriter, t string, td *templateData){
  var tmpl *template.Template

  if app.config.useCache {
    if templateFromMap, ok := app.templateMap[t]; ok {
      tmpl = templateFromMap
    }
  }  


  if tmpl == nil{
    log.Println("building from cache")

    var newTemplate *template.Template
    var templErr error
    if newTemplate, templErr =  app.buildTemplateFromDisk(t); templErr  != nil{
      log.Println("error building template", templErr)
      return 
    }

    tmpl = newTemplate
    
    if td == nil{
      td = &templateData{}
    }


    if err := tmpl.ExecuteTemplate(w, t, td); err != nil{
      log.Println("error gen template", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}


func(app *application) buildTemplateFromDisk(t string) (*template.Template, error){

  templateSlice := []string{
    "./templates/base.layout.gohtml",
    "./templates/partials/header.partial.gohtml",
    "./templates/partials/footer.partial.gohtml",
    "./templates/partials/navbar.partial.gohtml",
    fmt.Sprintf("./templates/%s", t),
  }

  if tmpl, err := template.ParseFiles(templateSlice...); err != nil{
    return nil,  err
  }else{
    app.templateMap[t] = tmpl
    return tmpl, nil
  }



}
