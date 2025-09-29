# Dev process definition

* you might have received a PRD

* from the PRD you need to split FE and BE work

* almost always BE work comes first. FE considerations can inform BE decisions slightly but not too much

* the platform is divided into:

  * core - ruthless business logic, that needs to have its own identity and definition
  * canopy(optional) - can be the bff (backend for frontend) which can listen to how FE is set up or how other people might use our app

* you will either write protos or you will update protos -- this will use c

## JENKINS PIPELINE

Triggers:

* on pull request
* on commit

## ARGO PIPELINE
