package postgres

// En Postgres las sentencias a veces no necesitan alias tan estrictos
// y los parámetros posicionales usan el formato $1, $2 en lugar de @p1, @p2
// En este caso, qryTipoDocumento no tiene parámetros, así que la consulta
// puede ser idéntica o adaptada según la convención que uses para tu BD en Postgres.

const qryTipoDocumento = `
	SELECT 
		LAB54C1 AS idTipoDoc,
		LAB54C2 AS codTipoDoc,
		LAB54C3 AS nomTipoDoc
	FROM LAB54
	WHERE LAB07C1 = 1	
`
const qryVerificarConfigCompanias = `
	SELECT
		LAB5803C2 AS datConfig
	FROM LAB5803  
	WHERE LAB5803C1 = 'gen_ManejaEmpresas'
`
