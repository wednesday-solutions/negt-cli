# CLI-App

For running this application 
    
    go build gqlgen.go
    
 If your machine already have gqlgen command please delete that.
 
 After that run,
 
    cp gqlgen /usr/local/bin
    
 Then for initializing GraphQL models
 
    gqlgen init
    
  It will create directories for gql and models.
  
  Then for creating new GraphQL model 
  
    gqlgen model
    
  After this, please provide the details for the questions.
  It will create GraphQL Model files in gql/models
