package mssql

const qryTaquillas = `
SELECT
	ENT5801C01 AS idTaquilla,
	ENT5801C02 AS codTaquilla,
	ENT5801C03 AS nomTaquilla,
	ENT5801C06 AS estadoTaquilla,
	S.nIdModulo
FROM ENT5801
CROSS APPLY OPENJSON(ENT5801C04)
WITH (
	nIdModulo INT
) AS S
WHERE
	ENT5801C06 = 0
	AND ENT0021C01 = @p1
	AND S.nIdModulo = @p2
`
