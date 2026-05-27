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

const qryPrioridadesPorTaquilla = `
	SELECT  
	J.IdTipoTurno,
	J.nPrioridad
FROM lab5821 AS P
CROSS APPLY OPENJSON(P.lab5821c2)
WITH (
	IdTipoTurno INT '$.IdTipoTurno',
	nPrioridad INT '$.nPrioridad'
) AS J
WHERE P.lab5801c1 = @p1
`

const qryTiemposPorSede = `
	SELECT 
		J.IdTipoTurno,
		J.nCantidad
	FROM LAB5822 AS E
	CROSS APPLY OPENJSON(E.LAB5822C2)
	WITH (
		IdTipoTurno INT '$.IdTipoTurno',
		nCantidad INT '$.nCantidad'
	) AS J
	WHERE E.LAB05C1 = @p1
`

const qryTurnosDisponibles = `
	SELECT 
	lab05c1 AS idSede,
	lab05c4 AS nombreSede,
	lab5800c1 AS idServicio,
	lab5800c2 AS codServicio,
	lab5800c3 AS nomServicio,
	lab5810c1 AS idTipoTurno,
	lab5810c2 AS codTipoTurno,
	lab5810c3 AS nomTipoTurno,
	lab5824c3 AS fechaHoraTurno,
	lab21c2 AS idTipoDocumento,
	lab5802c1 AS idEmpresa,
	lab5802c2 AS codEmpresa,
	lab5802c3 AS nomEmpresa,
	lab5824c2 AS numeroTurno,
	LAB22C1 AS numeroOrden,
	lab5824C20 AS jsonPaciente
FROM lab5824
WHERE lab5800c1 = @p1
	AND lab5824c5 = 0
	AND lab05c1 = @p2
	AND LEFT(lab5824c3, 8) = CONVERT(CHAR(8), GETDATE(), 112)
ORDER BY lab5824c2 ASC
`

const qrySedesUsuario = `
	SELECT
		U.LAB05C1 AS idSede,
		S.LAB05C10 AS codSede,
		S.LAB05C4 AS nomSede
	FROM LAB93 U
	INNER JOIN LAB05 S ON U.LAB05C1 = S.LAB05C1
	INNER JOIN LAB04 L ON U.LAB04C1 = L.LAB04C1
	WHERE L.LAB04C4 = @p1
`

const qryConsumirCredenciales = `
	SELECT lab04c1 as idUsuario 
	FROM lab04 
	WHERE lab04c4 = @p1 AND lab04c5 = @p2
`

const qryTurnosDisponiblesConOrden = `
	SELECT 
	lab05c1 AS idSede,
	lab05c4 AS nombreSede,
	lab5800c1 AS idServicio,
	lab5800c2 AS codServicio,
	lab5800c3 AS nomServicio,
	lab5810c1 AS idTipoTurno,
	lab5810c2 AS codTipoTurno,
	lab5810c3 AS nomTipoTurno,
	lab5824c3 AS fechaHoraTurno,
	lab21c2 AS idTipoDocumento,
	lab5802c1 AS idEmpresa,
	lab5802c2 AS codEmpresa,
	lab5802c3 AS nomEmpresa,
	lab5824c2 AS numeroTurno,
	LAB22C1 AS numeroOrden,
	lab5824C20 AS jsonPaciente
FROM lab5824
WHERE lab5800c1 = @p1
	AND lab5824c5 = 0
	AND lab05c1 = @p2
	AND LEFT(lab5824c3, 8) = CONVERT(CHAR(8), GETDATE(), 112)
	AND lab22c1 <> ''
ORDER BY lab5824c2 ASC
`
