package mssql

// /////====================//
// /////   Creacion Turnos   //
// ///////////////////////////
// Consultas SQL usadas por el servicio CrearTurno:
// - qrySedeById
// - qryCompaniaById
// - qryPacienteByDocumento
// - qryConfigCantidadPorSede
// - qryUltimoTurnoPorModuloFecha
// - qryTurnosHoyPorModuloFecha
// - qryInsertTurno
// Mantener el orden y los alias de columnas para que los `Scan` coincidan.

const qryTipoDocumento = `
	SELECT 
		LAB54C1 AS idTipoDoc,
		LAB54C2 AS codTipoDoc,
		LAB54C3 AS nomTipoDoc
	FROM LAB54
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
	WHERE LAB07C1 = 1
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
		lab5810C1 AS idTipoTurno,
		lab5810C2 AS codTipoTurno,
		lab5810C3 AS nomTipoTurno
	FROM LAB5810
`

const qryConfirmarConfigSedes = `
	SELECT
		LAB5803C2 AS datConfig
	FROM LAB5803  
	WHERE LAB5803C1 = 'gen_ManejaSedes'
`

const qryCargarConfigLIS = `
	SELECT
		ENT4028C02 AS separadorMuestra
	FROM ENT4028
	WHERE ENT4028C01 = 'ord_separadorMuestras'
`

const qrySedeDefault = `
	SELECT 
		LAB05C1 AS idSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C2)), ''), 'N/A') AS codSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C4)), ''), 'N/A') AS nomSede 
	FROM LAB05
--	WHERE LAB05C22 = 1
`

const qrySedeById = `
	SELECT 
		LAB05C1 AS idSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C2)), ''), 'N/A') AS codSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C4)), ''), 'N/A') AS nomSede
	FROM LAB05
	WHERE LAB05C1 = @p1
`

const qryCompaniaById = `
	SELECT
		LAB5802C1 AS idCompania,
		LAB5802C2 AS codigoCompania,
		LAB5802C3 AS nombreCompania
	FROM LAB5802
	WHERE LAB5802C1 = @p1
`

const qryTipoTurnoById = `
	SELECT
		LAB5810C1 AS idTipoTurno,
		LAB5810C2 AS codTipoTurno,
		LAB5810C3 AS nomTipoTurno
	FROM LAB5810
	WHERE LAB5810C1 = @p1
`

const qryServicioById = `
	SELECT
		LAB5800C1 AS idServicio,
		LAB5800C2 AS codigoServicio,
		LAB5800C3 AS nombreServicio
	FROM LAB5800
	WHERE LAB5800C1 = @p1
`

const qryPacienteByDocumento = `
	SELECT
		LAB21C1 AS idPaciente,
		LAB21C2 AS numeroDocumento,
		LAB21C5 AS apellido1,
		LAB21C6 AS apellido2,
		LAB21C3 AS nombre1,
		LAB21C4 AS nombre2,
		LAB80c1 AS sexo,
		LAB21C7 AS fechaNacimiento,
		LAB54C1 AS idTipoDocumento
	FROM LAB21
	WHERE LAB21C2 = @p1
`

const qryConfigCantidadPorSede = `
	SELECT
		LAB5818C2 AS jsonCantidad
	FROM LAB5818
	WHERE LAB05C1 = @p1
`

const qryUltimoTurnoPorModuloFecha = `
	SELECT TOP 1
		LAB5843C2 AS ultimoTurno
	FROM LAB5843
	WHERE LEFT(LAB5843C3, 8) = @p1
		AND LAB5843C14 = @p2
	ORDER BY LAB5843C2 DESC
`

const qryTurnosHoyPorModuloFecha = `
	SELECT
		COUNT(DISTINCT LAB5843C2) AS turnosHoy
	FROM LAB5843
	WHERE LAB5843C14 = @p1
		AND LEFT(LAB5843C3, 8) = @p2
`

const qryInsertTurno = `
	INSERT INTO LAB5843 (
		LAB05C1, LAB05C2, LAB05C4,
		LAB5802C1, LAB5802C2, LAB5802C3,
		LAB5810C1, LAB5810C2, LAB5810C3,
		LAB5800C1, LAB5800C2, LAB5800C3,
		LAB5843C2, LAB5843C3, LAB5843C4, LAB5843C5, LAB5843C20,
		LAB54C1, LAB5843C16, LAB5843C14, LAB5843C15
	) VALUES (
		@p1, @p2, @p3,
		@p4, @p5, @p6,
		@p7, @p8, @p9,
		@p10, @p11, @p12,
		@p13, @p14, @p15, @p16, @p17,
		@p18, @p19, @p20, @p21
	)
`
