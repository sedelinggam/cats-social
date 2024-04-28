---
runme:
  id: 01HWK92XYFPATPY1X0YWGHFXRV
  version: v3
---

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# -tags are what driver we want to install on our machine
migrate create -ext sql -dir migrations create_table_users
migrate -database "${POSTGRES_DB_URL}" -path EXAMPLE_PATH up