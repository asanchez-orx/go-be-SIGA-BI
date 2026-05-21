package mssql

const qryTaquillas = `
SELECT
	LAB5801C1 AS idTaquilla,
	LAB5801C2 AS codTaquilla,
	LAB5801C3 AS nomTaquilla,
	LAB5801C6 AS estadoTaquilla,
	S.nIdModulo
FROM LAB5801
CROSS APPLY OPENJSON(LAB5801C4)
WITH (
	nIdModulo INT
) AS S
WHERE
	LAB5801C6 = 0
	AND LAB05C1 = @p1
	AND S.nIdModulo = @p2
`

const qryServiciosSiga = `
	SELECT 
	LAB5800C1 AS idServicio,
	LAB5800C2 AS codServicio,
	LAB5800C3 AS nomServicio
	FROM 
	LAB5800
	WHERE LAB05C1 = @p1  AND LAB85C1 = @p2
`
