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

const queryTaquillasxSedeServicio = `
SELECT 
LAB5801C1 AS id,
LAB5801C3 AS name
FROM
LAB5801
WHERE LAB05C1 = @p1
`

const queryMotivosDescanso = `
SELECT 
    ISNULL(LAB5808C1, 0) AS id,
    ISNULL(LAB5808C3, '') AS name,
    ISNULL(LAB5808C4, '') AS description,
    ISNULL(LAB07C2, '') AS registerDate,
    LAB5808C5 AS type,
    LAB07C1 AS state
FROM
    LAB5808
    WHERE LAB08C5 = 1
`

const qryUpdateEstadoTaquilla = `
UPDATE LAB5801 SET LAB5801C6 = @p1 WHERE LAB05C1 = @p2 AND LAB5801C1 = @p3
`

const qryUpdateEstadoAtencion = `
UPDATE LAB5824 SET LAB5824c5 = @p1 WHERE LAB5824c1 = @p2 AND LAB5800c1 = @p3 AND LAB5801c1 = @p4
`
const qryMotivosCancelacion = `
SELECT 
    ISNULL(LAB5808C1, 0) AS id,
    ISNULL(LAB5808C3, '') AS name,
    ISNULL(LAB5808C5, '') AS type,
    ISNULL(LAB07C1, '') AS state
FROM
    LAB5808
    WHERE LAB5808C5 = 2
`

const qryServicioOrigen = `
SELECT 
LAB5800C1 AS id
FROM
LAB5824
WHERE LAB5824C1 = @p1
`

const qryServiciosDisponiblesParaTransferencia = `
SELECT
    t.lab05c1 AS id,
    pa.LAB5800C1 AS idServicioDes,
	pa.LAB5800C2 as code,
	pa.LAB5800C3 as name,
	nTiempoTransferencia AS waitTime
FROM LAB5801 t
CROSS APPLY OPENJSON(t.lab5801c7)
WITH (
    nIdPuntoAtencion INT '$.nIdPuntoAtencion',
	nTiempoTransferencia INT '$.nTiempoTransferencia'
) j
INNER JOIN LAB5800 pa ON pa.LAB5800C1 = j.nIdPuntoAtencion
WHERE 
t.lab5801c1 = @p1
AND t.lab05c1 = @p2
`

const qryLlamadoTurno = `
SELECT
    lab5824c2 AS turno
FROM lab5824
WHERE lab5824c1 = @p1
  AND lab5800c1 = @p2
  AND LEFT(lab5824c3, 8) = CONVERT(VARCHAR(8), GETDATE(), 112)
`

const qryUpdateTurnoPost = `
UPDATE lab5824 
SET lab5824c5 = 1 
WHERE lab5824c1 = @p1 
  AND lab05c1 = @p2 
  AND lab5800c1 = @p3
`

const qrySelectTurnoPost = `
SELECT
    t.lab5824c1 AS turn_id,
    t.lab5824c2 AS turn_number,
    t.lab5810c1 AS turnType_id,
	t.lab5810c2 AS turnType_code,
	t.lab5810c3 AS turnType_name,
	JSON_VALUE(t.lab5824c20, '$.idPaciente') AS patient_id,
	JSON_VALUE(t.lab5824c20, '$.idPaciente') AS patient_patientId,
	t.lab5800c1 AS service_id,
	t.lab5800c3 AS service_name
FROM lab5824 t
INNER JOIN LAB5821 pa
    ON pa.LAB5801C1 = t.LAB5801C1
WHERE t.lab5824c1 = @p1
  AND t.lab05c1 = @p2
  AND t.lab5800c1 = @p3
`
