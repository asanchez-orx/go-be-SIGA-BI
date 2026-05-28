package mssql

const qrySedes = `
SELECT 
	ISNULL(LAB05C1, 0) AS id,
	ISNULL(LAB05C10, '') AS code,
	ISNULL(LAB05C4, '') AS name,
	ISNULL(LAB05C11, '') AS description,
	LAB05C9 AS registerDate,
	ISNULL(LAB07C1, 0) AS state
FROM LAB05

`

const qryServiciosNTXSede = `
SELECT
     LAB5800C1 AS id,
    ISNULL(LAB5800C2, '') AS code,
    ISNULL(LAB5800C3, '') AS name,
    ISNULL(LAB5800C4, '') AS description,
    ISNULL(lab07c2, '') AS registerDate
FROM
    LAB5800
WHERE
    LAB05C1 = @p1

`

const queryTaquillas = `
SELECT 
LAB5801C1 AS id,
LAB5801C2 AS code,
LAB5801C3 AS name,
LAB07C1 AS state
FROM
LAB5801

`
