# Data Model

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
<th style="text-align: left;">SQLite</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>CHAR(n)</p></td>
<td style="text-align: left;"><p>CHAR(n)</p></td>
<td style="text-align: left;"><p>CHAR(n)</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>VARCHAR(n)</p></td>
<td style="text-align: left;"><p>VARCHAR(n)</p></td>
<td style="text-align: left;"><p>VARCHAR(n)</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>CLOB</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>NUMERIC(p,s)</p></td>
<td style="text-align: left;"><p>NUMERIC(p,s)</p></td>
<td style="text-align: left;"><p>DECIMAL(p,s)</p></td>
<td style="text-align: left;"><p>NUMERIC</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DECIMAL(p,s)</p></td>
<td style="text-align: left;"><p>DECIMAL(p,s)</p></td>
<td style="text-align: left;"><p>DECIMAL(p,s)</p></td>
<td style="text-align: left;"><p>NUMERIC</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>REAL</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>REAL</p></td>
<td style="text-align: left;"><p>REAL</p></td>
<td style="text-align: left;"><p>REAL</p></td>
<td style="text-align: left;"><p>REAL</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DOUBLE PRECISION</p></td>
<td style="text-align: left;"><p>DOUBLE PRECISION</p></td>
<td style="text-align: left;"><p>DOUBLE</p></td>
<td style="text-align: left;"><p>MORE PRECISELY REAL</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>DATE</p></td>
<td style="text-align: left;"><p>TEXT (formatted as 'YYYY-MM-DD')</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TIME</p></td>
<td style="text-align: left;"><p>TIME</p></td>
<td style="text-align: left;"><p>TIME</p></td>
<td style="text-align: left;"><p>TEXT (formatted as 'HH:MM:SS')</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>TEXT (formatted as 'YYYY-MM-DD HH:MM:SS')</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BOOLEAN</p></td>
<td style="text-align: left;"><p>BOOLEAN</p></td>
<td style="text-align: left;"><p>TINYINT(1)</p></td>
<td style="text-align: left;"><p>INTEGER (0 as false, 1 as true)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BLOB</p></td>
<td style="text-align: left;"><p>BYTEA</p></td>
<td style="text-align: left;"><p>BLOB</p></td>
<td style="text-align: left;"><p>BLOB</p></td>
</tr>
</tbody>
</table>

## References

-   [SQL-92 - Wikipedia](https://en.wikipedia.org/wiki/SQL-92)

    -   [The SQL-92 standard](https://www.contrib.andrew.cmu.edu/~shadow/sql/sql1992.txt)

-   MySQL\[<https://dev.mysql.com>\]

    -   [Data Types](https://dev.mysql.com/doc/refman/8.0/en/data-types.html)

-   PostgreSQL\[<https://www.postgresql.org>\]

    -   [Data Types](https://www.postgresql.org/docs/current/datatype.html)

-   SQLite

    -   [Datatypes In SQLite](https://sqlite.org/datatype3.html)
