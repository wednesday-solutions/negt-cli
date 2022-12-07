<img align="left" src="https://github.com/wednesday-solutions/negt/blob/develop/negt-cli-preview.png" width="500" height="450" />

<div>
  <a href="https://www.wednesday.is?utm_source=gthb&utm_medium=repo&utm_campaign=serverless" align="left" style="margin-left: 0;">
    <img src="https://uploads-ssl.webflow.com/5ee36ce1473112550f1e1739/5f5879492fafecdb3e5b0e75_wednesday_logo.svg">
  </a>
  <p>
    <h1 align="left">NEGT CLI
    </h1>
  </p>

  <p>
A CLI tool that works with the Node Express GraphQL Template and allows you to create queries, mutations and its scaffold tests for graphql models and stitches them all together preventing wastage of time in setup and boilerplate code.
  </p>

---

  <p>
    <h4>
      Expert teams of digital product strategists, developers, and designers.
    </h4>
  </p>

  <div>
    <a href="https://www.wednesday.is/contact-us?utm_source=gthb&utm_medium=repo&utm_campaign=serverless" target="_blank">
      <img src="https://uploads-ssl.webflow.com/5ee36ce1473112550f1e1739/5f6ae88b9005f9ed382fb2a5_button_get_in_touch.svg" width="121" height="34">
    </a>
    <a href="https://github.com/wednesday-solutions/" target="_blank">
      <img src="https://uploads-ssl.webflow.com/5ee36ce1473112550f1e1739/5f6ae88bb1958c3253756c39_button_follow_on_github.svg" width="168" height="34">
    </a>
  </div>

---

  <p>
    <h3 align="left">Built using <a href="https://github.com/wednesday-solutions/" target="_blank">Node Express GraphQL Template</a>
    </h3>
  </p>

</div>

---


## Installation

    go install github.com/wednesday-solutions/negt@latest
    
## Examples of generated files

<ul>
  <li><a href=gitub.com/wednesday-solutions/negt/blob/develop/generated-files/models></a>models</li>
  <li><a href=gitub.com/wednesday-solutions/negt/blob/develop/generated-files/models/tests></a>tests</li>
  <li><a href=gitub.com/wednesday-solutions/negt/blob/develop/generated-files/mockData></a>mock data</li>
</ul>

## Generating GraphQL models and tests

    image

## Documentation

<h3>Help</h3>

<p>To get a list of commands and usage hints use</p>

    negt --help
    
    negt gqlgen
    
## Creating a new NEGT Application

     negt gqlgen init
     
demo

     image

<p>Hint: If you are using Node-Express-GraphQL-Template, it will create directories `server/gql/models` otherwise `gql/models`.</p>

## Creating GraphQL models

    negt gqlgen model
    
demo
    
    video

## Projects using it

<h3><ul><li><a href="https://github.com/wednesday-solutions/node-express-graphql-template">Node Express GraphQL Template</a></li></ul></h3>

## Generated files contain the following

    Files :-
        1. index.js                 - It contains all relations of the gql model.
        2. models.js                - It contains fields and GraphQLObjectType of the model.
        3. query.js                 - It contains query of the model.
        4. list.js                  - It contains connection and query list of the model.
        5. mutation.js              - It contains the mutation fields and mutation object.
        6. customCreateMutation.js  - If you need custom resolvers for the GraphQL model, you can say yes for the question in the CLI.
           If you said yes, then it will create 3 custom resolver files. This file contains the custom create mutation object.
        7. customUpdateMutation.js  - It contains the custom update mutation object.
        8. customDeleteMutation.js  - It contains the custom delete mutation object.

It will create its test files also in `gql/models/<modelName>/tests` directory.

    Test files :-
        1. index.test.js                - It contains the test cases of index.js file. It checks the importing is working properly.
        2. models.test.js               - It contains the test case for type of the GraphQL model.
        3. query.test.js                - It contains the test case of GraphQL query.
        4. list.test.js                 - It contains the test case of GraphQL query list.
        5. mutation.test.js             - It contains the test cases of GraphQL mutations like creating, updating and deleting.
        6. pagination.test.js           - It contains the test case for checking the pagination of GraphQL queries is working properly.
           Most of the case, we don't need to change this file anymore. So it will create this test case in a different file.
        7. customCreateMutation.test.js - It contains the test case of custom create mutation resolver.
        8. customUpdateMutation.test.js - It contains the test case of custom update mutation resolver.
        9. customDeleteMutation.test.js - It contains the test case of custom delete mutation resolver.

Also, it will create mockData file for the model in `utils/testUtils` directory.

##

<p>This project is released under the <a href="https://github.com/wednesday-solutions/negt/blob/develop/LICENCE">MIT License</a>.</p>

