# verbose-spork


https://github.com/khang00/verbose-spork/assets/40449174/97c27b4b-b9a9-492b-8f1e-3d8214e1757e


## Run
Front end can be run by first `cd` to the `fe` folder then run `npm dev`

Back end can be run by `go build cmd/main.go && ./main`

## Project structure
the front end is in `fe` folder. I chose to make this sub folder for the front-end, because there is not enough time to split front-end to another repo.

the backend is in this folder, and follows the structure outlined in [golang project layout](https://github.com/golang-standards/project-layout)

## Specifications
1. User authentication Sign in & sign up
2. Upload a CSV file contains 1 to 100 keywords
3. For each keywords search it on Google, then stores the followings information of the first result page.
    1. Total number of AdWords advertisers on the page
    2. Total number of links on the page
    3. HTMLM code of the first page
    4. Numbers of search results for this keyword
4. See list of uploaded keyword, and search result information.
5. Search across reports

### Web UI
1. Sign in & sign up
2. Upload file
3. View list keywords
4. View search results for each keyword
5. Search across reports

## APIs
1. Authentication
    1. `POST /api/user/signin` user sign in
    2. `POST /api/user/signup` user sign up
2. `POST /api/keyword/upload` upload a list of keywords for searching
3. `GET /api/keyword?id=[id]` get the search result for this keywords

### Modules
1. Upload file
    1. DB Storage (Postgres)
    2. Mock in memory
2.  Scrape (Google)
    1. Keyword source (DB Storage)
    2. Scraping strategy (Rate limits)
    3. Keyword scraper
        1. Search handler
        2. Result parser
            1. Result page parser
            2. First page parser
        3. Result storage (DB Storage)
3. CURD stuffs for keywords
