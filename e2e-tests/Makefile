SHELL := /bin/bash

ifeq ($(OS),Windows_NT)
	# Windows commands
	SHELL = cmd
	START_SERVER = .\.ci\server.start.bat
	PREPARE_SERVER = .\.ci\server.prepare.bat
	RUN_CYPRESS = .\.ci\server.run_cypress.bat
	STOP_SERVER = .\.ci\server.stop.bat
	START_DASHBOARD = .\.ci\dashboard.start.bat
	STOP_DASHBOARD = .\.ci\dashboard.stop.bat
	CLEAN_ENV_PLACEHOLDERS = git reset .ci\.env.{cypress,server,dashboard} && git checkout .ci\.env.{cypress,server,dashboard}
else
	# Linux commands
	SHELL = /bin/bash
	START_SERVER = ./.ci/server.start.sh
	PREPARE_SERVER = ./.ci/server.prepare.sh
	RUN_CYPRESS = ./.ci/server.run_cypress.sh
	STOP_SERVER = ./.ci/server.stop.sh
	START_DASHBOARD = ./.ci/dashboard.start.sh
	STOP_DASHBOARD = ./.ci/dashboard.stop.sh
	CLEAN_ENV_PLACEHOLDERS = git reset .ci/.env.{cypress,server,dashboard} && git checkout .ci/.env.{cypress,server,dashboard}
endif

.PHONY: all run stop
all: run
run: start-server prepare-server run-cypress
stop: stop-server stop-dashboard clean-env-placeholders

.PHONY: start-server prepare-server run-cypress stop-server restart-server
start-server:
	$(START_SERVER)

prepare-server:
	$(PREPARE_SERVER)

run-cypress:
	$(RUN_CYPRESS)

stop-server:
	$(STOP_SERVER)

restart-server: stop-server start-server

.PHONY: start-dashboard stop-dashboard
start-dashboard:
	$(START_DASHBOARD)

stop-dashboard:
	$(STOP_DASHBOARD)

.PHONY: print-env-placeholders clean-env-placeholders
print-env-placeholders:
	find .ci -maxdepth 1 -name '.env.*' | xargs --verbose -l cat

clean-env-placeholders:
	$(CLEAN_ENV_PLACEHOLDERS)