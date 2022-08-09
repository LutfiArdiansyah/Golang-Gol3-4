# Digitask

## Tech

Digitask uses several tools to work properly:

- [Go Language] - open source programming language!
- [Buffalo] - Rapid Web Development in Go
- [Soda CLI] - Soda Core is a command-line interface (CLI) tool that enables you to scan the data in your data source to surface invalid, missing, or unexpected data
- [Bootstrap] - Bootstrap is a powerful, feature-packed frontend toolkit
- [jQuery] - jQuery is a fast, small, and feature-rich JavaScript library
- [Datatable] - DataTables is a plug-in for the jQuery Javascript library

## Installation

Dillinger requires [Go Language](https://go.dev/dl/) v1.19, [Soda CLI](https://gobuffalo.io/documentation/database/soda/), [Buffalo](https://gobuffalo.io/documentation/getting_started/installation/) to run.

Setup .env file to set the database configuration, If you change the ```PORT```, make sure to update the port on index.html.

After the configuration is complete, **there is no need to create a database on the db**, run the following command to create the database automatically and migrate the database:
```sh
soda create -a
soda migrate up
```

## RUN

Run the following command, to run this backend:
```sh
buffalo dev
```

After the backend runs successfully, **open the index.html** file in the browser


[Go Language]: <https://go.dev/>
[Buffalo]: <https://gobuffalo.io/>
[Soda CLI]: <https://docs.soda.io/>
[Bootstrap]: <https://getbootstrap.com/>
[jQuery]: <https://jquery.com/>
[Datatable]: <https://datatables.net/>