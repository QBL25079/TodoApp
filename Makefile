SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

include .env
export

export PROJECT_ROOT?=$(CURDIR)

env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down todoapp-postgres

env-cleanup:
	$$ans = Read-Host 'Очистить все volume файлы окружения? Опасность утери данных. [Y/N]'; if ($$ans -eq 'Y') { docker compose down todoapp-postgres port-forwarder; docker volume rm todo_pgdata; Write-Host 'Файлы окружения очищены' } else { Write-Host 'Очистка окружения отменена' }

env-port-forward:
	@docker compose up -d port-forwarder 

env-port-close:
	@docker compose down -d port-forwarder 

migrate-create:
	if (-not '$(seq)') { Write-Host 'seq parameter is required'; exit 1 }; docker compose run --rm todoapp-postgres-migrate create -ext sql -dir /migrations -seq '$(seq)'

migrate-up:
	$(MAKE) migrate-action action=up

migrate-down:
	$(MAKE) migrate-action action=down

migrate-action:
	@if (-not '$(action)') { Write-Host 'Action parameter is required'; exit 1 }; docker compose run --rm todoapp-postgres-migrate -path /migrations "-database" "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" $(action)

todoapp-run:
	$$env:POSTGRES_HOST = "localhost"; $$env:POSTGRES_PORT = "5433"; $$env:POSTGRES_DATABASE = "${POSTGRES_DB}"; $$env:POSTGRES_USER = "${POSTGRES_USER}"; $$env:POSTGRES_PASSWORD = "${POSTGRES_PASSWORD}"; $$env:POSTGRES_TIMEOUT = "${POSTGRES_TIMEOUT}"; go mod tidy; go run cmd/todoapp/main.go


todoapp-deploy:
	@docker compose up -d --build todoapp

ps: 
	@docker compose ps