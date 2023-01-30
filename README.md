# isbndb-go
Go Client to use isbndb API 


## ISBNdb API

The REST API allows you to retrieve information about millions of books.
Documentation: https://isbndb.com/apidocs/v2


## How to use it

Please create `.env` file for the environment and add your Isbndb API key as the following:


"""
```bash
APIKey=<insert your api key here>
```

After running the

"""
```bash
go run main.go
```
command, you will have a web server ready to fetch book data for you.

Example request and response:

```bash
request: http://localhost:8080/?isbn=9786052116913

response
{
    "title": "Tahıla Karşı - İlk Devletlerin Derin Tarihi",
    "language": "tr",
    "authors": [
        "James C. Scott"
    ],
    "publisher": "Koç Üniversitesi Yayınları",
    "image": "https://images.isbndb.com/covers/69/13/9786052116913.jpg"
}

```

Feel free to extend the exposed fields via extending BookData struct via checking the documentation.

## Next steps

- Filtering by title and author features will be added (on top of the ISBN)
- Filtering by subject feature will be added
