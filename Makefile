
.PHONY: gen
gen:
	@echo "Generating SDK"
	curl -sSL https://insulator.conductor.one/api/v1/openapi.yaml > ./openapi.yaml
	speakeasy generate sdk -s openapi.yaml -o . -d
	rm openapi.yaml
	@echo "Fixing Permissions"
	find * -type f -perm '+rwx' | xargs chmod -x