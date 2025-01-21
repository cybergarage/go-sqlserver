# Query Model

**go-sqlserver** uses a SQLite database to execute queries. The following table shows the correspondence between the queries supported by **go-sqlserver** and their equivalents in SQL-92, PostgreSQL, and MySQL.

## DDL

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">SQL92</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">go-sqlserver (SQLite)</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>CREATE DATABASE db_name</p></td>
<td style="text-align: left;"><p>CREATE DATABASE db_name</p></td>
<td style="text-align: left;"><p>Not Supported (uses file-based storage)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>DROP DATABASE db_name</p></td>
<td style="text-align: left;"><p>DROP DATABASE db_name</p></td>
<td style="text-align: left;"><p>Not Supported (uses file-based storage)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DROP INDEX idx_name</p></td>
<td style="text-align: left;"><p>DROP INDEX idx_name</p></td>
<td style="text-align: left;"><p>DROP INDEX idx_name</p></td>
<td style="text-align: left;"><p>DROP INDEX idx_name</p></td>
</tr>
</tbody>
</table>

## DML

<table>
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">MERGE INTO table_name USING …​ ON condition WHEN MATCHED THEN …​;</th>
<th style="text-align: left;">Not Supported</th>
<th style="text-align: left;">Not Supported</th>
<th style="text-align: left;">MERGE</th>
<th style="text-align: left;">Not Supported</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>SELECT col_name FROM table_name WHERE condition;</p></td>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>INSERT INTO table_name (col1, col2) VALUES (val1, val2);</p></td>
<td style="text-align: left;"><p>INSERT</p></td>
<td style="text-align: left;"><p>INSERT</p></td>
<td style="text-align: left;"><p>INSERT</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>UPDATE table_name SET col_name = value WHERE condition;</p></td>
<td style="text-align: left;"><p>UPDATE</p></td>
<td style="text-align: left;"><p>UPDATE</p></td>
<td style="text-align: left;"><p>UPDATE</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DELETE FROM table_name WHERE condition;</p></td>
<td style="text-align: left;"><p>DELETE</p></td>
<td style="text-align: left;"><p>DELETE</p></td>
<td style="text-align: left;"><p>DELETE</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>SELECT …​ LIMIT n;</p></td>
<td style="text-align: left;"><p>SELECT …​ LIMIT n;</p></td>
<td style="text-align: left;"><p>LIMIT</p></td>
<td style="text-align: left;"><p>SELECT …​ LIMIT n;</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>SQL92</p></td>
<td style="text-align: left;"><p>PostgreSQL</p></td>
<td style="text-align: left;"><p>MySQL</p></td>
<td style="text-align: left;"><p>go-sqlserver (SQLite)</p></td>
<td style="text-align: left;"><p>SQLite</p></td>
</tr>
</tbody>
</table>

## References

-   [SQL-92 - Wikipedia](https://en.wikipedia.org/wiki/SQL-92)

    -   [The SQL-92 standard](https://www.contrib.andrew.cmu.edu/~shadow/sql/sql1992.txt)

-   [MySQL](https://dev.mysql.com)

    -   [SQL Statements](https://dev.mysql.com/doc/refman/8.0/en/sql-statements.html)

-   [PostgreSQL](https://www.postgresql.org)

    -   [The SQL Language](https://www.postgresql.org/docs/current/sql.html)

-   [SQLite](https://sqlite.org)

    -   [SQL As Understood By SQLite](https://www.sqlite.org/lang.html)
