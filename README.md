# go-wunderkammer-www

Go package for web-related operations involving "wunderkammer" databases.

## Important

Work in progress

## Tools

To build binary versions of these tools run the `cli` Makefile target. For example:

```
$> make cli
go build -mod vendor -o bin/wunderkammer-server cmd/wunderkammer-server/main.go
```

### wunderkammer-server

```
$> ./bin/wunderkammer-server -h
Usage of ./bin/wunderkammer-server:
  -database-dsn string
    	A valid wunderkammer database DSN string. (default "sql://sqlite3/oembed.db")
  -path-templates string
    	The path to valid wunderkammer-www HTML templates. (default "static/templates/html/*")
  -server-uri string
    	A valid aaronland/go-http-server URI. (default "http://localhost:8080")
```

For example:

```
$> sqlite3 /usr/local/go-wunderkammer/nasm.db < /usr/local/go-wunderkammer/schema/sqlite/oembed.sql

$> /usr/local/go-smithsonian-openaccess/bin/emit \
	-oembed \
	-bucket-uri file:///Users/asc/code/OpenAccess metadata/objects/NASM \

   | /usr/local/go-wunderkammer-image/bin/append-dataurl \
	-format gif \

   | /usr/local/go-wunderkammer/bin/wunderkammer-db \
	-database-dsn 'sql://sqlite3/usr/local/go-wunderkammer/nasm.db'

...time passes

$> /usr/local/go-wunderkammer-www/bin/wunderkammer-server \
	-database-dsn 'sql://sqlite3/usr/local/go-wunderkammer/nasm.db'
```

And then if you went to `http://localhost:8080/object?url=si://nasm/o/A19480187000` you would see this:

![](docs/images/nasm.jpg)

### Endpoints

#### /image?url={OEMBED_URL}

Display the OEmbed record whose `url` property matches `{OEMBED_URL}`.

#### /object?url={OEMBED_OBJECT_URI}

Display the OEmbed records whose `object_uri` property matches `{OEMBED_OBJECT_URI}`.

#### /oembed?url={OEMBED_URL}

Output OEmbed (as JSON) for the record whose `url` property matches `{OEMBED_URL}`.

For example:

```
$> curl 'http://localhost:8080/oembed?url=https://ids.si.edu/ids/download?id=NASM-NASM2017-03151_screen'

{
  "version": "1.0",
  "type": "photo",
  "width": -1,
  "height": -1,
  "title": "UAV, General Atomics MQ-1L Predator A (Transferred from the United States Air Force)",
  "url": "https://ids.si.edu/ids/download?id=NASM-NASM2017-03151_screen",
  "author_name": "General Atomics Aeronautical Systems, Inc",
  "author_url": "https://airandspace.si.edu/collection/id/nasm_A20040180000",
  "provider_name": "National Air and Space Museum",
  "provider_url": "https://airandspace.si.edu",
  "object_uri": "si://nasm/o/A20040180000",
  "data_url": "data:image/jpeg;base64,R0lGODlhIAPgAYcAAAAAAAAARAAAiAAAzABEAABERABEiABEzACIAACIRACIiACIzADMAADMRADMiADMzADd3REREQAAVQAAmQAA3QBVAABVVQBMmQBJ3QCZAACZTACZmQCT3QDdAADdSQDdkwDungDu7iIiIgAAZgAAqgAA7gBmAABmZgBVqgBP7gCqAACqVQCqqgCe7gDuAADuTwD/VQD/qgD//zMzMwAAdwAAuwAA/wB3AAB3dwBduwBV/wC7AAC7XQC7uwCq ..."
}
```

#### /random

Retrieve a random OEmbed record and redirect to `/object?url={OEMBED_OBJECT_URI}`.

## See also

* https://github.com/search?q=topic%3Awunderkammer+org%3Aaaronland&type=Repositories
* https://github.com/aaronland/go-http-server
* https://github.com/aaronland/go-http-bootstrap