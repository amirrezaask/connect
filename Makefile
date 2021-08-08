migrate-up:
	psql -h 127.0.0.1 -U connect connect < .scripts/database.sql
