# Esqueleto para proyectos API REST con Go

Este proyectos es una plantilla para crear un proyecto API REST con Go y la librería vulcano.
para crear un nuevo proyecto es necesario tener instalado el programa [templet](https://gitlab.com/wfrsgo/templet/-/releases),
debe estan incluido en el PATH del sistema.


## Modo de uso

Para crear un nuevo proyecto basta con ejecutar el siguiente comando:

```bash
templet -r 192.168.1.21:3000 -n <mi_proyecto> "git:develop/api-skeleton"
```

> Cambiar `<mi_proyecto>` por el nombre que desee para el proyecto, puede tener caracteres alfanuméricos y/o guiones bajos (_).

El comando anterior inicará un pequeño formulario para que se introduzca los datos necesarios para la creación del proyecto.

La salida del comando será algo así:

```bash
Cloning into 'C:\Users\FABIAN~1\AppData\Local\Temp\templet-120704655'...
remote: Enumerating objects: 41, done.
remote: Counting objects: 100% (41/41), done.
remote: Compressing objects: 100% (34/34), done.
remote: Total 41 (delta 8), reused 0 (delta 0), pack-reused 0
Unpacking objects: 100% (41/41), 9.35 KiB | 68.00 KiB/s, done.
Asistente para la creación del proyecto mi_proyecto:
====================================================

# Cambiar por el nombre del binario del proyecto
Nombre del binario (sin guiones o puntos) [myapp]: proyecto
# Cambiar por el driver de base de datos que desee usar (pgsql: PostgreSQL, mssql: Microsoft SQL Server)
Driver de base de datos (pgsql/mssql) [pgsql]: mssql
# Cambiar por el nombre del módulo inicial del proyecto, en singular
Nombre del módulo inicial (en singular) [myapp]: documentacion 
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/.gitignore` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/LICENSE` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/README.md` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/Taskfile.yml` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/cmd/api/main.go` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/config/config.json.sample` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/go.mod` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/internal/api/documentacion/app/app.go` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/internal/api/documentacion/domain/domain.go` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/internal/api/documentacion/domain/interfaces.go` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/internal/api/documentacion/infra/pgsql/pgsql.go` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/internal/api/documentacion/infra/pgsql/raw.go` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/internal/api/documentacion/infra/handler/handler.go` creado
> Archivo `C:\Users\Fabian Salazar\tmp\mi_proyecto/internal/api/documentacion/infra/handler/routes.go` creado
# Ejecutando comando `rm -rf .git`
# Ejecutando comando `git init`
Initialized empty Git repository in C:/Users/Fabian Salazar/tmp/mi_proyecto/.git/
# Ejecutando comando `git add .`
# Ejecutando comando `git commit -m 'Commit inicial'`
[master (root-commit) d8d67ce] Commit inicial
 14 files changed, 594 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 LICENSE
 create mode 100644 README.md
 create mode 100644 Taskfile.yml
 create mode 100644 cmd/api/main.go
 create mode 100644 config/config.json.sample
 create mode 100644 go.mod
 create mode 100644 internal/api/documentacion/app/app.go
 create mode 100644 internal/api/documentacion/domain/domain.go
 create mode 100644 internal/api/documentacion/domain/interfaces.go
 create mode 100644 internal/api/documentacion/infra/handler/handler.go
 create mode 100644 internal/api/documentacion/infra/handler/routes.go
 create mode 100644 internal/api/documentacion/infra/pgsql/pgsql.go
 create mode 100644 internal/api/documentacion/infra/pgsql/raw.go
- Proyecto mi_proyecto creado con éxito a las 03/05/2026 09:58:54 PM
```


Al finalizar se creara un proyecto con el siguiente estructura:

```
mi_proyecto
├── cmd
│   └── api
│       └── main.go
├── config
│   └── config.json.sample
├── go.mod
├── internal
│   └── api
│       └── documentacion
│           ├── app
│           │   └── app.go
│           ├── domain
│           │   ├── domain.go
│           │   └── interfaces.go
│           └── infra
│               ├── handler
│               │   ├── handler.go
│               │   └── routes.go
│               └── pgsql
│                   ├── pgsql.go
│                   └── raw.go
├── LICENSE
├── README.md
└── Taskfile.yml
```

Los siguientes pasos son necesarios para que el proyecto funcione:

- El proyecto ya tiene el repositorio git inicializado, el archivo `.gitignore` configurado y el primer commit realizado.
- Reemplazar este archivo `README.md` por la plantilla definida en `README_template.md` del proyecto, editarlo segun las
  instrucciones de esa plantilla.
- Copiar el archivo `config/config.json.sample` a `config/config.json` y editarlo para que se adapte a la configuración
  correspondiente.


Tener en cuenta el flujo de trabajo, gestion de commits y solicitud de pull request postulada en:

- [Commits convencionales](http://192.168.1.21:3000/CLTech/documentacion/src/master/standards/git/02-conventional-commits.md)
- [Solicitudes PR - Pull Requests](http://192.168.1.21:3000/CLTech/documentacion/src/master/standards/git/03-pull-requests.md)
- [Flujo de trabajo Git Flow](http://192.168.1.21:3000/CLTech/documentacion/src/master/standards/git/04-git-flow.md)



