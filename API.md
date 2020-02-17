
# FUPP API
API for FUPP app.

Table of Contents

1. [returns professions array](#categories)
1. [returns profession](#category)
1. [returns professions array](#contact)
1. [returns profession](#expertise)
1. [returns professions array](#expertises)
1. [log-in the user and returns a token](#jobs)
1. [returns professions array](#missions)
1. [returns profession](#profession)
1. [returns professions array](#professions)
1. [returns professions array](#services)
1. [returns professions array](#steps)
1. [returns professions array](#subcategories)
1. [returns profession](#subcategory)
1. [Users API](#users)

<a name="categories"></a>

## categories

| Specification | Value |
|-----|-----|
| Resource Path | /categories |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /categories | [GET](#GetProfessions) | returns professions array |



<a name="GetProfessions"></a>

#### API: /categories (GET)


returns professions array



<a name="category"></a>

## category

| Specification | Value |
|-----|-----|
| Resource Path | /category |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /category/\{id\} | [GET](#GetProfession) | returns profession |



<a name="GetProfession"></a>

#### API: /category/\{id\} (GET)


returns profession



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | query | string | id | Yes |


<a name="contact"></a>

## contact

| Specification | Value |
|-----|-----|
| Resource Path | /contact |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /contact | [GET](#GetProfessions) | returns professions array |



<a name="GetProfessions"></a>

#### API: /contact (GET)


returns professions array



<a name="expertise"></a>

## expertise

| Specification | Value |
|-----|-----|
| Resource Path | /expertise |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /expertise | [GET](#GetProfession) | returns profession |



<a name="GetProfession"></a>

#### API: /expertise (GET)


returns profession



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | query | string | id | Yes |


<a name="expertises"></a>

## expertises

| Specification | Value |
|-----|-----|
| Resource Path | /expertises |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /expertises/\{id\} | [GET](#GetProfessions) | returns professions array |



<a name="GetProfessions"></a>

#### API: /expertises/\{id\} (GET)


returns professions array



<a name="jobs"></a>

## jobs

| Specification | Value |
|-----|-----|
| Resource Path | /jobs |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /jobs/register | [POST](#Login) | log-in the user and returns a token |
| /jobs/myjobs | [GET](#Login) | log-in the user and returns a token |
| /jobs/jobs | [GET](#Login) | log-in the user and returns a token |
| /jobs/pitched | [GET](#Login) | log-in the user and returns a token |
| /jobs/interviewing | [GET](#Login) | log-in the user and returns a token |
| /jobs/offered | [GET](#Login) | log-in the user and returns a token |
| /jobs/accepted | [GET](#Login) | log-in the user and returns a token |
| /jobs/jobstatus | [POST](#Login) | log-in the user and returns a token |



<a name="Login"></a>

#### API: /jobs/register (POST)


log-in the user and returns a token



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| description | query | string | Descripttion | Yes |
| location | query | string | Location | Yes |
| postcode | query | string | Postcode | Yes |
| rate | query | string | Rate | Yes |
| subcategoryid | query | string | SubcategoryID | Yes |
| categoryid | query | string | CategoryID | Yes |


<a name="Login"></a>

#### API: /jobs/myjobs (GET)


log-in the user and returns a token



<a name="Login"></a>

#### API: /jobs/jobs (GET)


log-in the user and returns a token



<a name="Login"></a>

#### API: /jobs/pitched (GET)


log-in the user and returns a token



<a name="Login"></a>

#### API: /jobs/interviewing (GET)


log-in the user and returns a token



<a name="Login"></a>

#### API: /jobs/offered (GET)


log-in the user and returns a token



<a name="Login"></a>

#### API: /jobs/accepted (GET)


log-in the user and returns a token



<a name="Login"></a>

#### API: /jobs/jobstatus (POST)


log-in the user and returns a token



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| status | query | string | Status | Yes |


<a name="missions"></a>

## missions

| Specification | Value |
|-----|-----|
| Resource Path | /missions |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /missions | [GET](#GetMissions) | returns professions array |



<a name="GetMissions"></a>

#### API: /missions (GET)


returns professions array



<a name="profession"></a>

## profession

| Specification | Value |
|-----|-----|
| Resource Path | /profession |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /profession/\{id\} | [GET](#GetProfession) | returns profession |



<a name="GetProfession"></a>

#### API: /profession/\{id\} (GET)


returns profession



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | query | string | id | Yes |


<a name="professions"></a>

## professions

| Specification | Value |
|-----|-----|
| Resource Path | /professions |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /professions | [GET](#GetProfessions) | returns professions array |



<a name="GetProfessions"></a>

#### API: /professions (GET)


returns professions array



<a name="services"></a>

## services

| Specification | Value |
|-----|-----|
| Resource Path | /services |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /services | [GET](#GetProfessions) | returns professions array |



<a name="GetProfessions"></a>

#### API: /services (GET)


returns professions array



<a name="steps"></a>

## steps

| Specification | Value |
|-----|-----|
| Resource Path | /steps |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /steps | [GET](#GetSteps) | returns professions array |



<a name="GetSteps"></a>

#### API: /steps (GET)


returns professions array



<a name="subcategories"></a>

## subcategories

| Specification | Value |
|-----|-----|
| Resource Path | /subcategories |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /subcategories | [GET](#GetProfessions) | returns professions array |



<a name="GetProfessions"></a>

#### API: /subcategories (GET)


returns professions array



<a name="subcategory"></a>

## subcategory

| Specification | Value |
|-----|-----|
| Resource Path | /subcategory |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /subcategory/\{id\} | [GET](#GetProfession) | returns profession |



<a name="GetProfession"></a>

#### API: /subcategory/\{id\} (GET)


returns profession



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | query | string | id | Yes |


<a name="users"></a>

## users

| Specification | Value |
|-----|-----|
| Resource Path | /users |
| API Version | 1.0.0 |
| BasePath for the API | http://host:port/api |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /users/login | [POST](#Login) | log-in the user and returns a token |
| /users/register | [POST](#Register) | user register |



<a name="Login"></a>

#### API: /users/login (POST)


log-in the user and returns a token



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| email | query | string | Email | Yes |
| password | query | string | Password | Yes |


<a name="Register"></a>

#### API: /users/register (POST)


user register



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| firstname | query | string | FristName | Yes |
| surname | query | string | Surname | Yes |
| email | query | string | Email | Yes |
| password | query | string | Password | Yes |
| nicname | query | string | NicName | Yes |
| profession | query | string | Profession | Yes |
| expertise | query | string | Expertise | Yes |
| history | query | string | History | Yes |
| address1 | query | string | Address1 | Yes |
| address2 | query | string | Address2 | Yes |
| city | query | string | City | Yes |
| postcode | query | string | PostCode | Yes |
| phone | query | string | Telephone | Yes |
| buddy | query | string | Buddy | Yes |
| professionid | query | uint | ProfessionID | Yes |
| expertiseid | query | uint | ExpertiseID | Yes |


