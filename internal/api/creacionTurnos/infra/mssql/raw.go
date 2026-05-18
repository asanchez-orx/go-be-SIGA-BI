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
		ENT5810C01 AS idTipoTurno,
		ENT5810C02 AS codTipoTurno,
		ENT5810C03 AS nomTipoTurno
	FROM ENT5810
`

const qryConfirmarConfigSedes = `
	SELECT
		LAB5803C2 AS datConfig
	FROM LAB5803  
	WHERE LAB5803C1 = 'gen_ManejaSedes'
`

const qrySedeDefault = `
	SELECT 
		LAB05C1 AS idSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C2)), ''), 'N/A') AS codSede,
		ISNULL(NULLIF(LTRIM(RTRIM(LAB05C4)), ''), 'N/A') AS nomSede 
	FROM LAB05
--	WHERE LAB05C22 = 1
`

const qryValidarSedeTurno = `
	SELECT 
		LAB05C4 AS nomSede,
		LAB05C2 AS codSede 
	FROM LAB05
	WHERE LAB05C1 = @p1
`

const qryValidarCompaniaTurno = `
	SELECT 
		LAB5802C1 AS idCompania,
		LAB5802C2 AS codCompania,
		LAB5802C3 AS nomCompania 
	FROM LAB5802
	WHERE LAB5802C1 = @p1
`

const qryValidarTipoTurnoTurno = `
	SELECT 
		LAB5810C1 AS id,
		LAB5810C2 AS cod,
		LAB5810C3 AS nom 
	FROM LAB5810
	WHERE LAB5810C1 = @p1
`

const qryValidarServicioTurno = `
	SELECT 
		LAB5800C1 AS idServicio,
		LAB5800C2 AS codServicio,
		LAB5800C3 AS nomServiciom 
	FROM LAB5800
	WHERE LAB5800C1 = @p1
`

const qryValidarPacienteTurno = `
	SELECT 
		LAB0090C1 AS idPaciente,
		LAB0090C3 AS numeroDocumento,
		LAB0090C10 AS apellido1,
		LAB0090C11 AS apellido2,
		LAB0090C12 AS nombre1,
		LAB0090C13 AS nombre2,
		LAB0090C15 AS sexo,
		LAB0090C14 AS fechaNacimiento,
		LAB0090C2 AS idTipoDocumento,
		JSON_VALUE(LAB0090C81, '$.LAB0090C09') AS nomTipoDocumento
	FROM LAB0090
	WHERE LAB0090C3 = @p1
`

const qryConfigCantidadesTurno = `
	SELECT LAB5818C2 AS jsonCantidad
	FROM LAB5818
	WHERE LAB05C1 = @p1
`

const qryUltimoTurnoModulo = `
	SELECT TOP 1
		LAB5843C2 AS UltimoTurno
	FROM LAB5843
	WHERE LEFT(LAB5843C3, 8) = FORMAT(GETDATE(), 'yyyyMMdd')
	AND LAB5843C14 = @p1
	ORDER BY LAB5843C2 DESC
`

const qryTurnosHoyModulo = `
	SELECT COUNT(DISTINCT LAB5843C2) AS TurnosHoy
	FROM LAB5843
	WHERE LAB5843C14 = @p1
	AND LEFT(LAB5843C3, 8) = FORMAT(GETDATE(), 'yyyyMMdd')
`

const qryInsertarTurno = `
	INSERT INTO LAB5843 (
		LAB05C1, LAB05C4, LAB05C2, 
		LAB5800C1, LAB5800C2, LAB5800C3,
		LAB5810C1, LAB5810C2, LAB5810C3,
		LAB5802C1, LAB5802C2, LAB5802C3,
		LAB5843C2, LAB5843C3, LAB5843C5, LAB5843C6, LAB5843C20,
		LAB0090C2, LAB5843C16, LAB5843C14, LAB5843C15
	) VALUES (
		@p1, @p2, @p3,
		@p4, @p5, @p6,
		@p7, @p8, @p9,
		@p10, @p11, @p12,
		@p13, @p14, 0, 0, @p15,
		@p16, @p17, @p18, @p19
	)
`

const qryObtenerTurnoInsertado = `
	SELECT TOP 1
		LAB5843C2 AS numeroTurno,
		LAB5843C3 AS fechaTurno
	FROM LAB5843 WITH (NOLOCK)
	WHERE LAB5843C14 = @p1
	AND LEFT(LAB5843C3, 8) = FORMAT(GETDATE(), 'yyyyMMdd')
	ORDER BY LAB5843C3 DESC
`

const qryValidarYConfigurarTurno = `
	SELECT 
		(SELECT LAB05C4 FROM LAB05 WHERE LAB05C1 = @p1) AS nomSede,
		(SELECT LAB05C2 FROM LAB05 WHERE LAB05C1 = @p1) AS codSede,
		(SELECT LAB5802C2 FROM LAB5802 WHERE LAB5802C1 = @p2) AS codCompania,
		(SELECT LAB5802C3 FROM LAB5802 WHERE LAB5802C1 = @p2) AS nomCompania,
		(SELECT LAB5810C2 FROM LAB5810 WHERE LAB5810C1 = @p3) AS codTipoTurno,
		(SELECT LAB5810C3 FROM LAB5810 WHERE LAB5810C1 = @p3) AS nomTipoTurno,
		(SELECT LAB5800C2 FROM LAB5800 WHERE LAB5800C1 = @p4) AS codServicio,
		(SELECT LAB5800C3 FROM LAB5800 WHERE LAB5800C1 = @p4) AS nomServicio,
		(SELECT LAB5818C2 FROM LAB5818 WHERE LAB05C1 = @p1) AS jsonCantidad
`

const qryInsertarTurnoConsolidado = `
	BEGIN TRANSACTION;

	DECLARE @ultimoTurno INT;
	DECLARE @turnosHoy INT;
	DECLARE @nuevoTurno INT;

	-- 1. Obtener el ultimo numero de turno de hoy con bloqueo de actualizacion para concurrencia segura
	SELECT TOP 1 @ultimoTurno = LAB5843C2
	FROM LAB5843 WITH (UPDLOCK, HOLDLOCK)
	WHERE LEFT(LAB5843C3, 8) = FORMAT(GETDATE(), 'yyyyMMdd')
	AND LAB5843C14 = @p18
	ORDER BY LAB5843C2 DESC;

	SET @ultimoTurno = ISNULL(@ultimoTurno, 0);
	SET @nuevoTurno = @ultimoTurno + 1;

	-- 2. Contar los turnos de hoy para este modulo
	SELECT @turnosHoy = COUNT(DISTINCT LAB5843C2)
	FROM LAB5843
	WHERE LAB5843C14 = @p18
	AND LEFT(LAB5843C3, 8) = FORMAT(GETDATE(), 'yyyyMMdd');

	SET @turnosHoy = ISNULL(@turnosHoy, 0);

	-- 3. Verificar el limite
	IF @p20 > 0 AND @turnosHoy >= @p20
	BEGIN
		ROLLBACK TRANSACTION;
		THROW 51000, 'LIMIT_EXCEEDED', 1;
	END
	ELSE
	BEGIN
		INSERT INTO LAB5843 (
			LAB05C1, LAB05C4, LAB05C2, 
			LAB5800C1, LAB5800C2, LAB5800C3,
			LAB5810C1, LAB5810C2, LAB5810C3,
			LAB5802C1, LAB5802C2, LAB5802C3,
			LAB5843C2, LAB5843C3, LAB5843C5, LAB5843C6, LAB5843C20,
			LAB0090C2, LAB5843C16, LAB5843C14, LAB5843C15
		) 
		OUTPUT inserted.LAB5843C2 AS numeroTurno, inserted.LAB5843C3 AS fechaTurno
		VALUES (
			@p1, @p2, @p3,
			@p4, @p5, @p6,
			@p7, @p8, @p9,
			@p10, @p11, @p12,
			@nuevoTurno, @p14, 0, 0, @p15,
			@p16, @p17, @p18, @p19
		);

		COMMIT TRANSACTION;
	END
`

