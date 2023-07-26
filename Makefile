
.PHONY: gen
gen:
	@echo "Generating SDK"
	curl -sSL -o openapi.yaml https://insulator.conductor.one/api/v1/openapi.yaml
	speakeasy generate sdk -s openapi.yaml -o . -d
	rm openapi.yaml
	@echo "Fixing Permissions"
	find * -type f -perm '+rwx' | xargs chmod -x