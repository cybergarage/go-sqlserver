# Query Model

(Being prepared)

## DDL

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
<th style="text-align: left;">Command</th>
<th style="text-align: left;">SQL92</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">SQLite</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>CREATE DATABASE</p></td>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>CREATE DATABASE db_name</p></td>
<td style="text-align: left;"><p>CREATE DATABASE db_name</p></td>
<td style="text-align: left;"><p>Not Supported (uses file-based storage)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>CREATE TABLE</p></td>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
<td style="text-align: left;"><p>CREATE TABLE table_name (col_name data_type, …​)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>CREATE INDEX</p></td>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
<td style="text-align: left;"><p>CREATE INDEX idx_name ON table_name (col_name)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>ALTER TABLE</p></td>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
<td style="text-align: left;"><p>ALTER TABLE table_name ADD COLUMN col_name data_type</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DROP DATABASE</p></td>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>DROP DATABASE db_name</p></td>
<td style="text-align: left;"><p>DROP DATABASE db_name</p></td>
<td style="text-align: left;"><p>Not Supported (uses file-based storage)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DROP TABLE</p></td>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
<td style="text-align: left;"><p>DROP TABLE table_name</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DROP INDEX</p></td>
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
<th style="text-align: left;">Command</th>
<th style="text-align: left;">SQL92</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">SQLite</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"><p>SELECT col_name FROM table_name WHERE condition;</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>INSERT</p></td>
<td style="text-align: left;"><p>INSERT INTO table_name (col1, col2) VALUES (val1, val2);</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>UPDATE</p></td>
<td style="text-align: left;"><p>UPDATE table_name SET col_name = value WHERE condition;</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DELETE</p></td>
<td style="text-align: left;"><p>DELETE FROM table_name WHERE condition;</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
<td style="text-align: left;"><p>Same</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>LIMIT</p></td>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>SELECT …​ LIMIT n;</p></td>
<td style="text-align: left;"><p>SELECT …​ LIMIT n;</p></td>
<td style="text-align: left;"><p>SELECT …​ LIMIT n;</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>MERGE</p></td>
<td style="text-align: left;"><p>MERGE INTO table_name USING …​ ON condition WHEN MATCHED THEN …​;</p></td>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>Not Supported</p></td>
<td style="text-align: left;"><p>Not Supported</p></td>
</tr>
</tbody>
</table>

## References
