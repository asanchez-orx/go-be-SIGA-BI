package mssql

const qryTipoDocumento = `
	SELECT 
		ENT0024C01 AS idTipoDoc,
		ENT0024C02 AS codTipoDoc,
		ENT0024C03 AS nomTipoDoc
	FROM ENT0024
	WHERE ENT0024C05 = 1
`

const qryCompanias = `
	SELECT 
		LAB5802C1 AS idCompania,
		LAB5802C2 AS codigoCompania,
		LAB5802C3 AS nombreCompania
	FROM LAB5802
`

const qryCompaniasPorSede = `
SELECT 
    C.lab5802c1 AS idCompania,
    C.lab5802c2 AS codigoCompania,
    C.lab5802c3 AS nombreCompania
FROM lab5814 E
INNER JOIN lab5802 C 
    ON E.lab5802c1 = C.lab5802c1
WHERE E.lab05c1 = @p1;
`

const qryVerificarConfigCompanias = `
	SELECT
		LAB5803C2 AS datConfig
	FROM LAB5803  
	WHERE LAB5803C1 = 'gen_ManejaEmpresas'
`

const qryTipoServicio = `
	SELECT 
		LAB5800C1 AS idServicio,
		LAB5800C2 AS codigoServicio,
		LAB5800C3 AS nombreServicio
	FROM LAB5800 
	WHERE LAB85C1 = @p1
	AND LAB05C1 = @p2
`

const qryTipoServicioPorCompania = `
	SELECT 
		C.ENT5800C01 AS idServicio,
		C.ENT5800C02 AS codigoServicio,
		C.ENT5800C03 AS nombreServicio
	FROM ENT5815 E
	INNER JOIN ENT5800 C ON E.ENT5800C01 = C.ENT5800C01
	WHERE E.ENT5802C01 = @p1
	AND C.ENT0008C01 = @p2
	AND C.ENT0021C01 = @p3
`

const qryModulos = `
	SELECT
		LAB5818C1 as id,
		LAB05C1 as idSede,
		LAB5818C2 as jsonModulos
	FROM LAB5818
	WHERE LAB05C1 = @p1
`

const qrySedes = `
	SELECT 
		LAB05C1 AS idSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C2)), ''), 'N/A') AS codSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C4)), ''), 'N/A') AS nomSede
	FROM LAB05
`

const qryTipoTurno = `
	SELECT 
    T.lab5810c1 AS idTipoTurno,
    T.lab5810c2 AS codTipoTurno,
    T.lab5810c3 AS nomTipoTurno
FROM lab5815 REL
INNER JOIN lab5810 T 
    ON T.lab5810c1 = REL.lab5810c1
INNER JOIN lab5800 S 
    ON S.lab5800c1 = REL.lab5800c1
	WHERE REL.lab5800c1 = @p1
	AND S.lab05c1 = @p2
`

const qryTipoTurnoTodos = `
	SELECT 
		ENT5810C01 AS idTipoTurno,
		ENT5810C02 AS codTipoTurno,
		ENT5810C03 AS nomTipoTurno
	FROM ENT5810
`
