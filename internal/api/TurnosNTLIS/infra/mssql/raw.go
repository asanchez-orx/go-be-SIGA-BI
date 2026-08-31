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
    ISNULL(LAB5808C5, 0) AS type,
    ISNULL(LAB07C1, 0) AS state
FROM LAB5808
WHERE LAB5808C5 = 1
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
UPDATE LAB5824 SET lab5824c5 = 1 WHERE lab5824c1 = @p1
  AND lab5800c1 = @p2
  AND LEFT(lab5824c3, 8) = CONVERT(VARCHAR(8), GETDATE(), 112)
  
`

const qryUpdateTurnoPost = `
UPDATE T
SET 
    T.lab5824c5 = 1,
    T.lab5801c1 = P.lab5801c1,
    T.lab5801c2 = P.lab5801c2,
    T.lab5801c3 = P.lab5801c3
FROM lab5824 T
INNER JOIN lab5801 P 
    ON P.lab5801c1 = 1
WHERE 
    T.lab5824c1 = @p1
    AND T.lab05c1 = @p2
    AND T.lab5800c1 = @p3;
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

const qryTurnosDisponibles = `
SELECT 

    t.lab5824c1 AS id,
    t.lab5824c2 AS number,
    t.lab5810c1 AS turnType_id,
    t.lab5810c2 AS turnType_code,
    t.lab5810c3 AS turnType_name,
    ISNULL(tt.lab5810c9, '') AS turnType_color,

    ISNULL(JSON_VALUE(t.lab5824c20, '$.idPaciente'), '') AS patient_id,
    ISNULL(JSON_VALUE(t.lab5824c20, '$.idPaciente'), '') AS patient_patientId,

    CONCAT_WS(' ',
        NULLIF(LTRIM(RTRIM(JSON_VALUE(t.lab5824c20, '$.apellido1'))), ''),
        NULLIF(LTRIM(RTRIM(JSON_VALUE(t.lab5824c20, '$.apellido2'))), '')
    ) AS patient_lastName,

    CONCAT_WS(' ',
        NULLIF(LTRIM(RTRIM(JSON_VALUE(t.lab5824c20, '$.nombre1'))), ''),
        NULLIF(LTRIM(RTRIM(JSON_VALUE(t.lab5824c20, '$.nombre2'))), '')
    ) AS patient_name,

    ISNULL(t.lab5800c1, 0) AS service_id,
    ISNULL(t.lab5800c3, '') AS service_name,
    ISNULL(t.lab05c1, 0) AS branch_id,
    ISNULL(pa.lab05c4, '') AS branch_name
FROM lab5824 t
INNER JOIN lab5810 tt ON t.lab5810c1 = tt.lab5810c1
INNER JOIN lab05 pa ON pa.lab05c1 = t.lab05c1
WHERE t.lab05c1 = @p1
  AND t.lab5800c1 = @p2
  AND LEFT(lab5824c3, 8) = CONVERT(VARCHAR(8), GETDATE(), 112)
`

const qryUpdateOldTurn = `
UPDATE lab5824 SET lab5824c5 = 3 WHERE lab5824c1 = @p1 AND lab05c1 = @p2 AND lab5801c1 = @p3
`

const qryInsertNewTurn = `
INSERT INTO lab5824 (
    lab05c1, lab05c4, 
    lab5802c1, lab5802c2, lab5802c3, 
    lab5800c1, lab5800c2, lab5800c3, 
    lab5810c1, lab5810c2, lab5810c3, 
    lab5824c2, lab22c1, lab5824c20, 
    lab5824c14, lab5824c15, lab5824c16, 
    lab5824c3, lab5824c5
)
SELECT 
    t.lab05c1, t.lab05c4, 
    t.lab5802c1, t.lab5802c2, t.lab5802c3, 
    s.lab5800c1, s.lab5800c2, s.lab5800c3, 
    t.lab5810c1, t.lab5810c2, t.lab5810c3, 
    t.lab5824c2, t.lab22c1, t.lab5824c20, 
    t.lab5824c14, t.lab5824c15, t.lab5824c16, 
    FORMAT(GETDATE(), 'yyyyMMddHHmmss'), 0
FROM lab5824 t
INNER JOIN lab5800 s ON s.lab5800c1 = @p2
WHERE t.lab5824c1 = @p1
`

const qryGetLogUserInPoint = `
SELECT TOP 1
    ISNULL(l.lab5821c1, 0)                    AS id,
    ISNULL(DATEDIFF(MILLISECOND, '19700101', l.lab5821c3), 0) AS registerDate,
    ISNULL(l.lab5821c4, 0)                    AS action,
    ISNULL(l.lab05c1, 0)                      AS branch_id,
    ISNULL(l.lab05c6, 0)                      AS branch_state,
    ISNULL(l.lab5801c1, 0)                    AS point_id,
    ISNULL(p.lab5801c8, 0)                    AS point_managePriority,
    ISNULL(l.lab5802c1, 0)                    AS user_id,
    ISNULL(l.lab5808c1, 0)                    AS reason_id,
    ISNULL(r.lab5808c5, 0)                    AS reason_type,
    ISNULL(r.lab07c1, 0)                      AS reason_state,
    ISNULL(l.lab5821c5, 0)                    AS difference
FROM lab5821 l
LEFT JOIN lab5801 p ON p.lab5801c1 = l.lab5801c1
LEFT JOIN lab5808 r ON r.lab5808c1 = l.lab5808c1
WHERE l.lab05c1 = @p1
  AND l.lab5801c1 = @p2
  AND l.lab5802c1 = @p3
ORDER BY l.lab5821c3 DESC
`

const qryGetTurnInPoint = `
SELECT TOP 1
    t.lab5824c1                                                AS id,
    ISNULL(t.lab5810c1, 0)                                    AS turnType_id,
    ISNULL(t.lab5810c2, '')                                   AS turnType_code,
    ISNULL(t.lab5810c3, '')                                   AS turnType_name,
    ISNULL(tt.lab07c1, 0)                                     AS turnType_state,
    ISNULL(CAST(JSON_VALUE(t.lab5824c20, '$.idPaciente') AS INT), 0) AS patient_id,
    ISNULL(JSON_VALUE(t.lab5824c20, '$.idPaciente'), '')      AS patient_patientId,
    ISNULL(JSON_VALUE(t.lab5824c20, '$.apellido1'), '')       AS patient_lastName,
    ISNULL(JSON_VALUE(t.lab5824c20, '$.nombre1'), '')         AS patient_name,
    ISNULL(JSON_VALUE(t.lab5824c20, '$.tipoDocumento'), '')   AS patient_documentType,
    ISNULL(t.lab5824c2, 0)                                    AS number,
    ISNULL(t.lab5824c5, 0)                                    AS state,
    ISNULL(p.lab5801c1, 0)                                    AS point_id,
    ISNULL(p.lab5801c2, '')                                   AS point_code,
    ISNULL(p.lab5801c3, '')                                   AS point_name,
    ISNULL(p.lab5801c8, 0)                                    AS point_managePriority,
    ISNULL(s.lab5800c1, 0)                                    AS service_id,
    ISNULL(s.lab5800c2, '')                                   AS service_code,
    ISNULL(s.lab5800c3, '')                                   AS service_name,
    ISNULL(b.lab05c1, 0)                                      AS branch_id,
    ISNULL(b.lab05c10, '')                                    AS branch_code,
    ISNULL(b.lab05c4, '')                                     AS branch_name,
    ISNULL(b.lab07c1, 0)                                      AS branch_state,
    ISNULL(p.lab07c1, 0)                                      AS point_state
FROM lab5824 t
INNER JOIN lab5801 p ON p.lab5801c1 = t.lab5801c1
INNER JOIN lab5800 s ON s.lab5800c1 = t.lab5800c1
INNER JOIN lab05  b ON b.lab05c1  = t.lab05c1
LEFT  JOIN lab5810 tt ON tt.lab5810c1 = t.lab5810c1
WHERE t.lab5801c1 = @p1
  AND t.lab5824c5 IN (0, 1)
ORDER BY t.lab5824c3 DESC
`

const qryValidateTurnInPoint = `
SELECT 
    --logUserInPoint
    COALESCE(t.lab5824c7,0) as id,
	COALESCE(t.lab5824c9,0) as registerDate,
	COALESCE(t.lab05c1, 0) AS id,
	COALESCE(t.lab5801c1, 0) AS id,
	COALESCE(t.lab5824c7,0) as id,

    -- turnInPoint
    COALESCE(t.LAB5824C1, 0) AS id,

    -- turnType
    COALESCE(t.lab5810c1, 0) AS id,
    COALESCE(t.lab5810c2, '') AS code,
    COALESCE(t.lab5810c3, '') AS name,
    COALESCE(tip.lab07c1, 0) AS state,

    -- patient
    COALESCE(JSON_VALUE(t.lab5824c20, '$.idPaciente'), '') AS id,
    COALESCE(JSON_VALUE(t.lab5824c20, '$.idPaciente'), '') AS patientId,
    JSON_VALUE(t.lab5824c20, '$.apellido1') AS apellido1,
    JSON_VALUE(t.lab5824c20, '$.apellido2') AS apellido2,
    JSON_VALUE(t.lab5824c20, '$.nombre1') AS nombre1,
    JSON_VALUE(t.lab5824c20, '$.nombre2') AS nombre2,
    COALESCE(t.lab21c2, '') AS documentType,

    COALESCE(t.lab5824c2, 0) AS number,
    COALESCE(t.lab5824c5, 0) AS state,

    -- point
    COALESCE(t.lab5801c1, 0) AS id,
    COALESCE(t.lab5801c2, '') AS code,
    COALESCE(t.lab5801c3, '') AS name,

    -- service
    COALESCE(t.lab5800c1, 0) AS id,
    COALESCE(t.lab5800c2, '') AS code,
    COALESCE(t.lab5800c3, '') AS name,

    -- branch
    COALESCE(t.lab05c1, 0) AS id,
    COALESCE(t.lab05c10, '') AS code,
    COALESCE(t.lab05c4, '') AS name

FROM LAB5824 t
INNER JOIN lab5810 tip 
    ON tip.lab5810c1 = t.lab5810c1
WHERE t.lab5824C5 = 1
  AND t.lab5824C8 = @p1
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

const qryTurnosDisponiblesAutomatico = `
SELECT  
    t.lab5824c1 AS id, 
    t.lab5824c2 AS number, 
    t.lab5810c1 AS turnType_id, 
    t.lab5810c2 AS turnType_code, 
    t.lab5810c3 AS turnType_name, 
    ISNULL(tt.lab5810c9, '') AS turnType_color, 

    ISNULL(JSON_VALUE(t.lab5824c20, '$.idPaciente'), '') AS patient_id, 
    ISNULL(JSON_VALUE(t.lab5824c20, '$.idPaciente'), '') AS patient_patientId, 

    ISNULL(t.lab5800c1, 0) AS service_id, 
    ISNULL(t.lab5800c3, '') AS service_name, 
    ISNULL(t.lab05c1, 0) AS branch_id, 
    ISNULL(pa.lab05c4, '') AS branch_name, 

    t.lab5824c5 AS state, 
    t.lab5824c3 AS date 

FROM lab5824 t 
INNER JOIN lab5810 tt 
    ON t.lab5810c1 = tt.lab5810c1 
INNER JOIN lab05 pa 
    ON pa.lab05c1 = t.lab05c1 

WHERE t.lab5800c1 = @p1 
  AND (
      t.lab5824c5 = 0 
      OR (
          t.lab5824c5 = 6 
          AND t.lab5824c18 <= CONVERT(VARCHAR(8), GETDATE(), 112) 
              + REPLACE(CONVERT(VARCHAR(8), GETDATE(), 108), ':', '') 
      ) 
  ) 
  AND t.lab05c1 = @p2 
  AND CONVERT(INT, LEFT(t.lab5824C3, 8)) =
      CONVERT(INT, CONVERT(CHAR(8), GETDATE(), 112))
`

const qryMovimientoTurno = `
UPDATE LAB5824
SET 
    lab5824c5 = @p1,
    lab5824c7 = @p2,
    lab5824c8 = @p3
WHERE 
    lab5824c1 = @p4
    AND lab5800c1 = @p5
    AND lab5801c1 = @p6;
`
