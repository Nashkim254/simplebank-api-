createdb:
	docker exec -t postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -t postgres14 dropdb  simple_bank	

.PHONY: createdb dropdb